package proto

type UserOnlineInfo struct {
	SdkVersion         string `json:"sdkVersion"`
	SdkType            string `json:"sdkType"`
	AppID              string `json:"appId"`
	LiveAppID          string `json:"liveAppId"`
	SessionID          string `json:"sessionId"`
	IsReconnect        string `json:"isReconnect"`
	UserID             string `json:"userId"`
	OpenID             string `json:"openId"`
	Avatar             string `json:"avatar"`
	UserName           string `json:"userName"`
	Sex                string `json:"sex"`
	Role               string `json:"role"`     // 字段是判断用户角色的，包括主播、用户、房管、超管，主要是用来对用户进行相关操作
	UserRole           string `json:"userRole"` // 判断当前用户是主播还是用户，主要是用来区分用户名片的，是显示主播名片还是用户名片
	ExpGrade           int    `json:"expGrade"`
	GuardType          string `json:"guardType"`
	CarID              string `json:"carId"`
	CarName            string `json:"carName"`
	CarIcon            string `json:"carIcon"`
	CarOnlineURL       string `json:"carOnlineUrl"`
	CarResURL          string `json:"carResUrl"`
	IsPlayCarAnim      string `json:"isPlayCarAnim"`
	MarkUrlsJoinString string `json:"markUrlsJoinString"`
	NobilityType       int    `json:"nobilityType"`
	IsEnterHide        int    `json:"isEnterHide"`
	TokenType          string `json:"tokenType"`
	Reconnect          bool   `json:"reconnect"`
	Hide               bool   `json:"hide"`
	PushSide           bool   `json:"pushSide"`
	Nobility           bool   `json:"nobility"`
	Guard              bool   `json:"guard"`
	LiveID             string `json:"liveId"`
	EnterType          string `json:"enterType"`
	Login              bool   `json:"login"`
	PullSide           bool   `json:"pullSide"`
	RoomManager        bool   `json:"roomManager"`
	Vip                bool   `json:"vip"`
	NotHide            bool   `json:"notHide"`
	IsRankHide         int    `json:"isRankHide"`
}
