package proto

const (
	SysParamAddrSuf                   = "/sys/param/getSysParamListByAppId"      // 系统配置服务
	UserWeekStarAddrSuf               = "/user/getWeekStarUsers"                 // 用户周星服务
	PlayNobilityEnterAnimationAddrSuf = "/nobility/isPlayNobilityEnterAnimation" // 是否播放贵族动画
	GiftSendSuf                       = "/gift/send"                             // 礼物发送
	ScoreGiftSend                     = "/scoreGift/send"                        // 积分发送

	/*
	* 过滤出正在直播的付费房
	*	实时查询
	* @param liveIds
	 */
	FindLivingChargeRoom = "/live/findLivingChargeRoom"
)

type SysParamByAppResp struct {
	EnableAchievement                   string `json:"enableAchievement"`
	MsgBroadcastSwitch                  string `json:"msgBroadcastSwitch"`
	NeteaseStreamDefaultPullFormat      string `json:"neteaseStreamDefaultPullFormat"`
	EnableCar                           string `json:"enableCar"`
	LiveManagerCountLimit               string `json:"liveManagerCountLimit"`
	EnableTurntable                     string `json:"enableTurntable"`
	BanGroupThresholdCount              string `json:"banGroupThresholdCount"`
	EnableGradeUpperLimit               string `json:"enableGradeUpperLimit"`
	EnableTranslationLevel              string `json:"enableTranslationLevel"`
	SpeakLevelFilter4Admin              string `json:"speakLevelFilter4Admin"`
	EnableTaskBox                       string `json:"enableTaskBox"`
	EnableGiftWall                      string `json:"enableGiftWall"`
	AnchorHeartBeatTimeoutInterval      string `json:"anchorHeartBeatTimeoutInterval"`
	NoticeUserGrade                     string `json:"noticeUserGrade"`
	LiveDefaultCoverPictureURL          string `json:"liveDefaultCoverPictureUrl"`
	LiveInitOnlineUserListSize          string `json:"liveInitOnlineUserListSize"`
	EntryNoticeLevelThreshold           string `json:"entryNoticeLevelThreshold"`
	EnableGiftBox                       string `json:"enableGiftBox"`
	EnableAnchorHomepageJump            string `json:"enableAnchorHomepageJump"`
	LiveOnlineUserListLevelFilter       string `json:"liveOnlineUserListLevelFilter"`
	NetEaseStreamCallBackDomain         string `json:"NetEaseStreamCallBackDomain"`
	EnableInteract                      string `json:"enableInteract"`
	EnableTranslate                     string `json:"enableTranslate"`
	FestivalDayTime                     string `json:"festivalDayTime"`
	ChatLenLimit                        string `json:"chatLenLimit"`
	NoticeDayRank                       string `json:"noticeDayRank"`
	OnlineCountSynInterval              string `json:"onlineCountSynInterval"`
	EnableVisitorLive                   string `json:"enableVisitorLive"`
	LiveStatisticHistoryDayCount        string `json:"liveStatisticHistoryDayCount"`
	BoomTime                            string `json:"boomTime"`
	EnableUserHomepageJump              string `json:"enableUserHomepageJump"`
	HomepageJumpURL                     string `json:"homepageJumpUrl"`
	EnableFestival                      string `json:"enableFestival"`
	NoticeAnchorGrade                   string `json:"noticeAnchorGrade"`
	EnableProp                          string `json:"enableProp"`
	EnablePrivateMessage                string `json:"enablePrivateMessage"`
	GradeSet10CharacterLimit            string `json:"gradeSet10CharacterLimit"`
	ExpScore                            string `json:"expScore"`
	AudienceEnterLiveBroadcastInterval  string `json:"audienceEnterLiveBroadcastInterval"`
	EnableLive                          string `json:"enableLive"`
	EnableOffence                       string `json:"enableOffence"`
	BoomMultiple                        string `json:"boomMultiple"`
	EnablePaster                        string `json:"enablePaster"`
	EnableNobility                      string `json:"enableNobility"`
	EnableRecordAPI                     string `json:"enableRecordApi"`
	EnableGuard                         string `json:"enableGuard"`
	EnableScore                         string `json:"enableScore"`
	EnableWeekStar                      string `json:"enableWeekStar"`
	OffenceNotifyThreshold              string `json:"offenceNotifyThreshold"`
	GiftTrumpetPlayPeriod               string `json:"giftTrumpetPlayPeriod"`
	GiftRatio                           string `json:"giftRatio"`
	LiveWatchCountStatisticIntervalMins string `json:"liveWatchCountStatisticIntervalMins"`
	SocketHeartBeatInterval             string `json:"socketHeartBeatInterval"`
	SpeakKeyWordsBlackList              string `json:"speakKeyWordsBlackList"`
}

type UserWeekStarByAppResp []string

type FindLivingChargeRoomResp []string
