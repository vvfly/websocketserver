package schedule

import (
	"strconv"

	httpclient "github.com/luckyweiwei/base/http-client"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/model/dao"
	"github.com/luckyweiwei/websocketserver/model/do"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
	"github.com/luckyweiwei/websocketserver/server/consul"
	remoteservice "github.com/luckyweiwei/websocketserver/server/remote-service"
)

// 付费直播间用户在线数据
func RefreshLiveOnlineUser() {
	Log.Debug("enter...")

	var (
		cm       = connectionm.GetConnectionManager()
		serverID = model.GetServerConfig().ServerID
		docs     = []do.LiveOnlineUser{}
	)

	addr := consul.GetBalanceAddr(proto.LiveServerName)
	addr += proto.FindLivingChargeRoom

	liveIDs := cm.LiveMap.Keys()
	postData := utils.SerializeToJson(liveIDs)

	// 调用接口获取所有的在线付费直播间
	resp := FindLivingChargeRoom(addr, postData)
	if resp == nil {
		return
	}

	chargeRooms := *resp

	if len(chargeRooms) > 0 {
		for _, liveID := range chargeRooms {
			roomConnection := cm.GetRommConnection(liveID)
			if roomConnection == nil {
				Log.Warnf("cannot find roomID in map, liveID = %v", liveID)
				continue
			}

			// 遍历直播间用户
			roomConnection.Connections.IterCb(func(key string, v interface{}) {
				connection := v.(*connectionm.Connection)

				wsConnInfo := connection.WsConnectInfo
				onlineUserInfo := &do.OnlineUser{
					SdkVersion:         wsConnInfo.SdkVersion,
					AppID:              wsConnInfo.Sid.AppID,
					LiveAppID:          wsConnInfo.Authen.UsrInfo.AppIDForCurrentLive,
					SessionID:          wsConnInfo.SessionsInfoStr,
					IsReconnect:        wsConnInfo.Authen.UsrInfo.IsReconnect,
					UserID:             wsConnInfo.Sid.UserID,
					OpenID:             wsConnInfo.Authen.UsrInfo.OpenID,
					Avatar:             wsConnInfo.Authen.UsrInfo.Avatar,
					UserName:           wsConnInfo.Authen.UsrInfo.UserName,
					Sex:                wsConnInfo.Authen.UsrInfo.Sex,
					Role:               wsConnInfo.Authen.UsrInfo.Role,
					UserRole:           wsConnInfo.Authen.UsrInfo.UserRole,
					ExpGrade:           wsConnInfo.Authen.UsrInfo.ExpGrade,
					GuardType:          wsConnInfo.Authen.UsrInfo.GuardType,
					CarID:              wsConnInfo.Authen.UsrInfo.CarID,
					CarName:            wsConnInfo.Authen.UsrInfo.CarName,
					CarIcon:            wsConnInfo.Authen.UsrInfo.CarIcon,
					CarOnlineURL:       wsConnInfo.Authen.UsrInfo.CarOnlineURL,
					CarResURL:          wsConnInfo.Authen.UsrInfo.CarResURL,
					IsPlayCarAnim:      wsConnInfo.Authen.UsrInfo.IsPlayCarAnim,
					MarkUrlsJoinString: wsConnInfo.Authen.UsrInfo.MarkUrlsJoinString,
					NobilityType:       strconv.Itoa(wsConnInfo.Authen.UsrInfo.NobilityType),
					IsEnterHide:        wsConnInfo.Authen.UsrInfo.IsEnterHide,
					TokenType:          wsConnInfo.Authen.UsrInfo.TokenType,
				}

				doc := do.LiveOnlineUser{
					ServerID:      serverID,
					LiveID:        liveID,
					UserID:        connection.WsConnectInfo.Sid.UserID,
					EntryLiveTime: connection.WsConnectInfo.Authen.UsrInfo.LastEntryLiveTime,
					OnlineUser:    utils.SerializeToJson(onlineUserInfo),
				}
				docs = append(docs, doc)
			})
		}
	}

	// delete
	doc := &do.LiveOnlineUser{
		ServerID: serverID,
	}
	err := dao.NewLiveOnlineUserDao().Delete(doc)
	if err != nil {
		Log.Error(err)
		return
	}

	// insert
	if len(docs) > 0 {
		err := dao.NewLiveOnlineUserDao().CreateMany(&docs)
		if err != nil {
			Log.Error(err)
			return
		}
	}
}

func FindLivingChargeRoom(addr, postData string) *proto.FindLivingChargeRoomResp {

	resp, body, errs := httpclient.New().
		Post(addr).
		Timeout(remoteservice.TimeOut).
		Set("Content-Type", "application/json").
		SendRawBodyData(postData).
		End()

	if errs != nil {
		Log.Error(errs)
		return nil
	}

	if resp.StatusCode != 200 {
		Log.Error("req status code != 200, resp = %v", resp)
		return nil
	}

	var respData = &proto.FindLivingChargeRoomResp{}
	err := utils.DecodeFromJson(body, respData)
	if err != nil {
		Log.Error(err)
		return nil
	}

	return respData
}
