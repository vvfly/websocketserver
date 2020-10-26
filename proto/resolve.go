package proto

type SessionsInfo struct {
	LiveID    string // 房间ID
	UserID    string // 用户ID
	EnterType string // 进入类型  1为开播端， 2 为观看端
	AppID     string // 渠道ID
}

type UserInfo struct {
	AppIDForCurrentUser string
	AppIDForCurrentLive string
	SessionID           string
	IsReconnect         string
	UserID              string
	OpenID              string // 渠道方用户id
	Avatar              string // 图像
	UserName            string
	Sex                 string
	Role                string // 计算后的值，比如，主播，房管，是根据所进入的直播间动态确定
	UserRole            string // 数据库原始值
	ExpGrade            int    // 等级
	GuardType           string // 守护类型 0，无；1，周守护；2，月守护；3，年守护
	CarID               string // 座驾id -1表示没有座驾
	CarName             string // 座驾名字
	CarIcon             string // 座驾图标 对应于 Car 类的 imgUrl
	CarOnlineURL        string // 座驾在线地址 Car: animalUrl
	CarResURL           string // 座驾资源下载 Car: zipUrl，现在已经废弃
	IsPlayCarAnim       string // 是否播放
	MarkUrlsJoinString  string // 动态标示
	NobilityType        int    //  贵族类型 -1 表未开通 贵族等级1~7
	IsEnterHide         int    // 是否入场隐身，1是，0否
	TokenType           string // 1 登录用户  2是游客
	IsRankHide          int    // 是否排行榜隐身 1是 0否

	ClientIP string // websocket填充

	IsPlayNobilityEnterAnimation int //是否播放贵族入场动画
	IsWeekStar                   int //是否展示周星标识

	LastPingTime      int64 // 最后心跳时间
	LastEntryLiveTime int64 // 最后进入直播间时间
}



type AuthenInfo struct {
	TimeStamp       int64     // 时间戳
	RequestUniqueID string    // 序列号
	Signature       string    // 签名
	UsrInfo         *UserInfo // 用户信息
}



type WsConnectInfo struct {
	Sid             *SessionsInfo
	SessionsInfoStr string
	Token           string
	Authen          *AuthenInfo
	SubAppId        string // 分销渠道ID
	SdkVersion      string // SDK版本号
	SdkType         string // SDK类型 1看播端 2开播端
}



func (uwc *WsConnectInfo) IsRoomManager() bool {
	return uwc.Authen.UsrInfo.Role == ROOM_MANAGER
}

func (uwc *WsConnectInfo) IsClanAdmin() bool {
	return uwc.Authen.UsrInfo.Role == CLAN_ADMIN
}

func (uwc *WsConnectInfo) IsLiveAdmin() bool {
	return uwc.Authen.UsrInfo.Role == LIVE_ADMIN
}

func (uwc *WsConnectInfo) IsLogin() bool {
	return uwc.Authen.UsrInfo.TokenType == TokenTypeLogin
}

func (uwc *WsConnectInfo) IsReconnect() bool {
	return uwc.Authen.UsrInfo.IsReconnect == "1"
}

func (uwc *WsConnectInfo) IsEnterHide() bool {
	return uwc.Authen.UsrInfo.IsEnterHide == 1
}

func (uwc *WsConnectInfo) IsPushSide() bool {
	return uwc.Sid.EnterType == EnterTypeAnchor
}

func (uwc *WsConnectInfo) IsPullSide() bool {
	return uwc.Sid.EnterType == EnterTypeAudience
}

func (uwc *WsConnectInfo) HasCar() bool {
	return uwc.Authen.UsrInfo.CarID != "-1"
}

func (uwc *WsConnectInfo) IsNobility() bool {
	return uwc.Authen.UsrInfo.NobilityType != NobilityTypeNo
}

func (uwc *WsConnectInfo) IsNobilityLevel1() bool {
	return uwc.Authen.UsrInfo.NobilityType == NobilityTypeLevel1
}

func (uwc *WsConnectInfo) IsNobilityLevel2() bool {
	return uwc.Authen.UsrInfo.NobilityType == NobilityTypeLevel2
}

func (uwc *WsConnectInfo) IsNobilityLevel3() bool {
	return uwc.Authen.UsrInfo.NobilityType == NobilityTypeLevel3
}

func (uwc *WsConnectInfo) IsNobilityLevel4() bool {
	return uwc.Authen.UsrInfo.NobilityType == NobilityTypeLevel4
}

func (uwc *WsConnectInfo) IsNobilityLevel5() bool {
	return uwc.Authen.UsrInfo.NobilityType == NobilityTypeLevel5
}

func (uwc *WsConnectInfo) IsNobilityLevel6() bool {
	return uwc.Authen.UsrInfo.NobilityType == NobilityTypeLevel6
}

func (uwc *WsConnectInfo) IsNobilityLevel7() bool {
	return uwc.Authen.UsrInfo.NobilityType == NobilityTypeLevel7
}

func (uwc *WsConnectInfo) IsGuard() bool {
	return uwc.Authen.UsrInfo.GuardType != GuardTypeNo
}

func (uwc *WsConnectInfo) IsGuardWeek() bool {
	return uwc.Authen.UsrInfo.GuardType == GuardTypeWeek
}

func (uwc *WsConnectInfo) IsGuardMonth() bool {
	return uwc.Authen.UsrInfo.GuardType == GuardTypeMonth
}

func (uwc *WsConnectInfo) IsGuardYear() bool {
	return uwc.Authen.UsrInfo.GuardType == GuardTypeYear
}

func (uwc *WsConnectInfo) IsVip() bool {
	return uwc.IsGuard() || uwc.IsNobility()
}

func (uwc *WsConnectInfo) IsHide() bool {
	return uwc.Authen.UsrInfo.IsEnterHide == 1
}

func (uwc *WsConnectInfo) IsOrdinary() bool {
	return !uwc.HasCar() && !uwc.IsNobility() && !uwc.IsGuard()
}
