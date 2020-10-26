package aproto

import "github.com/luckyweiwei/websocketserver/model/do"

type OnlineCountStatistic struct {
	TotalCount          int64 `json:"totalCount"`
	GuardCount          int64 `json:"guardCount"`
	MonthGuardCount     int64 `json:"monthGuardCount"`
	YearGuardCount      int64 `json:"yearGuardCount"`
	NobilityCount       int64 `json:"nobilityCount"`
	RoomManagerCount    int64 `json:"roomManagerCount"`
	NobilityLevel1Count int64 `json:"nobilityLevel1Count"`
	NobilityLevel2Count int64 `json:"nobilityLevel2Count"`
	NobilityLevel3Count int64 `json:"nobilityLevel3Count"`
	NobilityLevel4Count int64 `json:"nobilityLevel4Count"`
	NobilityLevel5Count int64 `json:"nobilityLevel5Count"`
	NobilityLevel6Count int64 `json:"nobilityLevel6Count"`
	NobilityLevel7Count int64 `json:"nobilityLevel7Count"`
}

type OnlineUserDto struct {
	UserID             string `json:"userId"`
	OpenID             string `json:"openId"`
	Avatar             string `json:"avatar"`
	UserName           string `json:"userName"`
	Sex                string `json:"sex"`
	Role               string `json:"role"`
	ExpGrade           int    `json:"expGrade"`
	GuardType          string `json:"guardType"`
	NobilityType       string `json:"nobilityType"`
	IsEnterHide        int    `json:"isEnterHide"`
	AppID              string `json:"appId"`
	UserRole           string `json:"userRole"`
	MarkUrlsJoinString string `json:"markUrlsJoinString"`
	//直播间停留时间
	StaySeconds int64 `json:"staySeconds"`
}

type OnlineUserListDto struct {
	VipCount   int             `json:"vipCount"`   // 在线贵宾(贵族数加守护数量)用户数
	TotalCount int             `json:"totalCount"` // 总的在线用户数
	List       []do.OnlineUser `json:"list"`
}
