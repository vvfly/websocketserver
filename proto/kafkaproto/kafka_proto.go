package kafkaproto

const (
	TOPIC_BIZ = "ws_biz"

	TOPIC_ENTRY        = "entry"
	TOPIC_ANCHOR_ENTER = "anchorEnter"

	TOPIC_NOTIFY_ONLINE_USER = "notifyOnlineUser"
)

var (
	FooTopics = []string{
		"entry",
		"notifyOnlineUser",
		"propSend_return",
		"goOut_return",
		"leave_return",
		"userGrade_return",
		"notify_return",
		"liveAdminBanPost_return",
		"liveControl_return",
		"grabGiftBoxBroadcast_return",
		"banPostAll_return",
		"liveAdminGoOut_return",
		"chatReceipt_return",
		"grabGiftBoxNotified_return",
		"postInterval_return",
		"liveSetting_return",
		"chat_return",
		"gamount_return",
		"banPost_return",
		"shield_return",
	}

	BarTopics = []string{
		"entry",
		"notifyOnlineUser",
		"balance_return",
		"liveSetting_return",
		"assistKing_return",
		"warn_return",
		"buyLiveTicket_return",
		"gamount_return",
		"generalFlutterScreen_return",
		"lianmaiInviting_return",
		"lianmaiMatchSuccess_return",
		"consumeNotify_return",
		"pkEnd_return",
		"turntableStatusUpdate_return",
		"giftTrumpet_return",
		"notifyFP_return",
		"propSend_return",
		"lianmaiMatchTimeout_return",
		"userGrade_return",
		"nobilityTrumpetBroadcast_return",
		"lianmaiSuccess_return",
		"openNobilityBroadcast_return",
		"gift_return",
		"forbidLive_return",
		"msgBroadcast_return",
		"leave_return",
		"lianmaiEnd_return",
		"chat_return",
		"universalBroadcast_return",
		"pkStart_return",
		"declineInviting_return",
		"firstKill_return",
		"putGiftBox_return",
		"cancelInviting_return",
		"liveDrawExamineResult_return",
		"liveDrawStart_return",
		"liveDrawFinished_return",
		"userPrivateMessage_return",
		"userPrivateMessageReceipt_return",
	}
)
