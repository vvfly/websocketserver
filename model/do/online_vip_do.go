package do

import (
	"strconv"

	"github.com/luckyweiwei/websocketserver/proto"
)

type OnlineVIP struct {
	ID       string       `bson:"_id"`
	ServerID string       `bson:"serverId"` // websocket网关节点id
	LiveID   string       `bson:"liveId"`
	Count    int          `bson:"count"`
	Rank     []OnlineUser `bson:"rank"`
}

type OnlineUser struct {
	SdkVersion         string `bson:"sdkVersion" json:"sdkVersion"`
	AppID              string `bson:"appId" json:"appId"`
	LiveAppID          string `bson:"liveAppId" json:"liveAppId"`
	SessionID          string `bson:"sessionId" json:"sessionId"`
	IsReconnect        string `bson:"isReconnect" json:"isReconnect"`
	UserID             string `bson:"userId" json:"userId"`
	OpenID             string `bson:"openId" json:"openId"`
	Avatar             string `bson:"avatar" json:"avatar"`
	UserName           string `bson:"userName" json:"userName"`
	Sex                string `bson:"sex" json:"sex"`
	Role               string `bson:"role" json:"role"`         // 字段是判断用户角色的，包括主播、用户、房管、超管，主要是用来对用户进行相关操作
	UserRole           string `bson:"userRole" json:"userRole"` // 判断当前用户是主播还是用户，主要是用来区分用户名片的，是显示主播名片还是用户名片
	ExpGrade           int    `bson:"expGrade" json:"expGrade"`
	GuardType          string `bson:"guardType" json:"guardType"`
	CarID              string `bson:"carId" json:"carId"`
	CarName            string `bson:"carName" json:"carName"`
	CarIcon            string `bson:"carIcon" json:"carIcon"`
	CarOnlineURL       string `bson:"carOnlineUrl" json:"carOnlineUrl"`
	CarResURL          string `bson:"carResUrl" json:"carResUrl"`
	IsPlayCarAnim      string `bson:"isPlayCarAnim" json:"isPlayCarAnim"`
	MarkUrlsJoinString string `bson:"markUrlsJoinString" json:"markUrlsJoinString"`
	NobilityType       string `bson:"nobilityType" json:"nobilityType"`
	IsEnterHide        int    `bson:"isEnterHide" json:"isEnterHide"`
	TokenType          string `bson:"tokenType" json:"tokenType"`
}

func (o *OnlineUser) IsNobility() bool {
	return o.NobilityType != strconv.Itoa(proto.NobilityTypeNo)
}

func (o *OnlineUser) IsGuard() bool {
	return o.GuardType != proto.GuardTypeNo
}
