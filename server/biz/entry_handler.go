package biz

import (
	"errors"
	"strconv"

	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/proto"
	"github.com/luckyweiwei/websocketserver/proto/kafkaproto"
	"github.com/luckyweiwei/websocketserver/proto/wproto"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
	remoteservice "github.com/luckyweiwei/websocketserver/server/remote-service"
)

func checkLastEnterTime() bool {
	serverConfig := model.GetServerConfig()
	timeNow := utils.GetTimeOfS()

	cm := connectionm.GetConnectionManager()
	mutex := cm.Mutex

	mutex.Lock()
	defer mutex.Unlock()
	lastEnterTime := cm.LastEnterTime

	ct := utils.GetTimeOfS()
	if (ct - lastEnterTime) > serverConfig.EntryTimeInterval { // 可以发送入场消息
		cm.LastEnterTime = timeNow // 更新最后入场时间
		return true
	}
	return false
}

func produceData(
	topic string,
	wsConnInfo *proto.WsConnectInfo,
	isdOrdinary int,
	isWeekStar int,
	isPlayNobilityEnterAnimation int,
) error {

	key := wsConnInfo.SessionsInfoStr

	value := wproto.MsgEntryWsReq{
		SdkVersion:                   wsConnInfo.SdkVersion,
		AppID:                        wsConnInfo.Sid.AppID,
		LiveAppID:                    wsConnInfo.Authen.UsrInfo.AppIDForCurrentLive,
		SessionID:                    wsConnInfo.SessionsInfoStr,
		IsReconnect:                  wsConnInfo.Authen.UsrInfo.IsReconnect,
		UserID:                       wsConnInfo.Sid.UserID,
		OpenID:                       wsConnInfo.Authen.UsrInfo.OpenID,
		Avatar:                       wsConnInfo.Authen.UsrInfo.Avatar,
		UserName:                     wsConnInfo.Authen.UsrInfo.UserName,
		Sex:                          wsConnInfo.Authen.UsrInfo.Sex,
		Role:                         wsConnInfo.Authen.UsrInfo.Role,
		UserRole:                     wsConnInfo.Authen.UsrInfo.UserRole,
		ExpGrade:                     wsConnInfo.Authen.UsrInfo.ExpGrade,
		GuardType:                    wsConnInfo.Authen.UsrInfo.GuardType,
		CarID:                        wsConnInfo.Authen.UsrInfo.CarID,
		CarName:                      wsConnInfo.Authen.UsrInfo.CarName,
		CarIcon:                      wsConnInfo.Authen.UsrInfo.CarIcon,
		CarOnlineURL:                 wsConnInfo.Authen.UsrInfo.CarOnlineURL,
		CarResURL:                    wsConnInfo.Authen.UsrInfo.CarResURL,
		IsPlayCarAnim:                wsConnInfo.Authen.UsrInfo.IsPlayCarAnim,
		MarkUrlsJoinString:           wsConnInfo.Authen.UsrInfo.MarkUrlsJoinString,
		NobilityType:                 wsConnInfo.Authen.UsrInfo.NobilityType,
		IsEnterHide:                  wsConnInfo.Authen.UsrInfo.IsEnterHide,
		TokenType:                    wsConnInfo.Authen.UsrInfo.TokenType,
		Reconnect:                    wsConnInfo.IsReconnect(),
		Hide:                         wsConnInfo.IsHide(),
		PushSide:                     wsConnInfo.IsPushSide(),
		Nobility:                     wsConnInfo.IsNobility(),
		Guard:                        wsConnInfo.IsGuard(),
		LiveID:                       wsConnInfo.Sid.LiveID,
		EnterType:                    wsConnInfo.Sid.EnterType,
		Login:                        wsConnInfo.IsLogin(),
		PullSide:                     wsConnInfo.IsPullSide(),
		RoomManager:                  wsConnInfo.IsRoomManager(),
		Vip:                          wsConnInfo.IsVip(),
		NotHide:                      !wsConnInfo.IsHide(),
		IsPlayNobilityEnterAnimation: isPlayNobilityEnterAnimation,
		IsWeekStar:                   isWeekStar,
		IsdOrdinary:                  isdOrdinary,
	}

	valueJSONStr := utils.SerializeToJson(&value)

	produceData := kafkaproto.SetKfkProducerJobData(topic, key, []byte(valueJSONStr))
	kafkaproto.KafkaProducerJobChan <- produceData

	return nil
}

func audience(wsConnInfo *proto.WsConnectInfo) error {
	if wsConnInfo.IsReconnect() { //重连不发入场消息
		return nil
	}
	if wsConnInfo.IsHide() { //设置了入场隐身的用户不广播入场消息
		return nil
	}

	var (
		isOrdinary                   = 0
		isWeekStar                   = 0
		isPlayNobilityEnterAnimation = 0

		appID  = wsConnInfo.Sid.AppID
		liveID = wsConnInfo.Sid.LiveID
		rm     = remoteservice.GetRemoteServiceCallManager()
	)

	// 普通用户
	if wsConnInfo.IsOrdinary() {

		sysParam := rm.GetSysParamByApp(appID)
		if sysParam == nil {
			Log.Errorf("获取系统配置失败. appID=%v", appID)
			return errors.New("获取系统配置失败")
		}

		// 用户等级门槛
		entryNoticeLevelThreshold, err := strconv.Atoi(sysParam.EntryNoticeLevelThreshold)
		if err != nil {
			Log.Error(err)
			return err
		}
		if wsConnInfo.Authen.UsrInfo.ExpGrade < entryNoticeLevelThreshold {
			Log.Warningf("用户等级小于入场消息等级门槛. 用户等级=%v, 等级门槛=%v",
				wsConnInfo.Authen.UsrInfo.ExpGrade, entryNoticeLevelThreshold)

			isOrdinary = 1
		}

		// 入场频率检查，如果满足条件同时更新最后入场时间
		if !checkLastEnterTime() {
			Log.Debugf("不在入场频率范围内")
			return nil
		}
	} else { // 非普通用户
		// 是否周星用户
		user := remoteservice.GetUserWeekStarByApp(appID)
		if user == nil {
			Log.Errorf("获取用户服务失败. appID=%v", appID)
			return errors.New("获取用户服务失败")
		}
		userAttr := *user
		if len(userAttr) > 0 {
			for _, v := range userAttr {
				if v == wsConnInfo.Sid.UserID {
					isWeekStar = 1
					wsConnInfo.Authen.UsrInfo.IsWeekStar = isWeekStar
					break
				}
			}
		}

		// 是否贵族且需播放贵族动画
		if wsConnInfo.IsNobility() {
			playNobilityEnterAnimation := remoteservice.GetPlayNobilityEnterAnimationByLive(liveID)
			if playNobilityEnterAnimation == "" {
				Log.Errorf("获取贵族服务失败. liveID=%v", liveID)
				return errors.New("获取贵族服务失败")
			}

			playNobilityEnterAnimationInt, err := strconv.Atoi(playNobilityEnterAnimation)
			if err != nil {
				Log.Error(err)
				return err
			}
			isPlayNobilityEnterAnimation = playNobilityEnterAnimationInt

			wsConnInfo.Authen.UsrInfo.IsPlayNobilityEnterAnimation = isPlayNobilityEnterAnimation
		}
	}

	// 房管 超管
	if wsConnInfo.IsRoomManager() || wsConnInfo.IsClanAdmin() || wsConnInfo.IsLiveAdmin() {
		isOrdinary = 0
	}

	topic := kafkaproto.TOPIC_ENTRY
	return produceData(topic, wsConnInfo, isOrdinary, isWeekStar, isPlayNobilityEnterAnimation)
}

func anchor(wsConnInfo *proto.WsConnectInfo) error {

	topic := kafkaproto.TOPIC_ANCHOR_ENTER

	return produceData(topic, wsConnInfo, 0, 0, 0)
}

func EntryHandler(wsConnInfo *proto.WsConnectInfo) error {
	Log.Debug("entering ...")

	enterType := wsConnInfo.Sid.EnterType
	if enterType == proto.EnterTypeAnchor {
		return anchor(wsConnInfo)
	} else if enterType == proto.EnterTypeAudience {
		return audience(wsConnInfo)
	} else {
		Log.Errorf("invalid enterType. enterType=%v", enterType)
		return errors.New("invalid enterType:" + enterType)
	}
}
