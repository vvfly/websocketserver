package wproto

// ws to kafka
type MsgEntryWsReq struct {
	SdkVersion         string `json:"sdkVersion"`
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
	NobilityType       int `json:"nobilityType"`
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

	IsPlayNobilityEnterAnimation int `json:"isPlayNobilityEnterAnimation"` //是否播放贵族入场动画
	IsWeekStar                   int `json:"isWeekStar"`                   //是否展示周星标识
	IsdOrdinary                  int `json:"isdOrdinary"`                  //是否是普通用户（无任何特权和身份，区别在于只广播等级和昵称，显示在入场消息最先面滚动区域）
}

// ws to client
// 普通用户
type MsgOrdinaryEntryRespData struct {
	UserName string `json:"userName"`
	ExpGrade int    `json:"expGrade"`
}
type MsgOrdinaryEntryResp struct {
	MessageType  string `json:"messageType"`
	BusinessData struct {
		Code       int                      `json:"code"`
		Message    string                   `json:"message"`
		ResultData MsgOrdinaryEntryRespData `json:"resultData"`
	} `json:"businessData"`
}

// 非普通用户
type MsgEntryRespData struct {
	UserName                     string `json:"userName"`
	UserID                       string `json:"userId"`
	Role                         string `json:"role"`     // 字段是判断用户角色的，包括主播、用户、房管、超管，主要是用来对用户进行相关操作
	UserRole                     string `json:"userRole"` // 判断当前用户是主播还是用户，主要是用来区分用户名片的，是显示主播名片还是用户名片
	Sex                          string `json:"sex"`
	Avatar                       string `json:"avatar"`
	ExpGrade                     int    `json:"expGrade"`
	GuardType                    string `json:"guardType"`
	CarID                        string `json:"carId"`
	CarName                      string `json:"carName"`
	CarIcon                      string `json:"carIcon"`
	CarOnlineURL                 string `json:"carOnlineUrl"`
	CarResURL                    string `json:"carResUrl"`
	IsPlayCarAnim                string `json:"isPlayCarAnim"`
	NobilityType                 int    `json:"nobilityType"`
	IsEnterHide                  int
	IsPlayNobilityEnterAnimation string `json:"isPlayNobilityEnterAnimation"` //是否播放贵族入场动画
	IsWeekStar                   int    `json:"isWeekStar"`                   //是否展示周星标识

	//加属性，用于渠道方app,用户详情跳转
	AppID    string   `json:"appId"`
	OpenID   string   `json:"openId"`
	MarkUrls []string `json:"markUrls"`
}
type MsgEntryResp struct {
	MessageType  string `json:"messageType"`
	BusinessData struct {
		Code       int              `json:"code"`
		Message    string           `json:"message"`
		ResultData MsgEntryRespData `json:"resultData"`
	} `json:"businessData"`
}
