package proto

const (
	// 直播服务
	MT_LEAVE      = "leave"      // 主播离开直播间
	MT_DISCONNECT = "disconnect" // 主播主动断开连接

	// 聊天服务
	MT_CHAT         = "chat"        // 聊天
	MT_CHAT_RECEIPT = "chatReceipt" // 聊天回执

	// websocket服务
	MT_ENTRY        = "entry"       // 进入直播间
	MT_ANCHOR_ENTER = "anchorEnter" // 主播入场

	// 礼物服务
	MT_GIFT                     = "gift"                 // 发送礼物
	MT_GIFT_TRUMPET             = "giftTrumpet"          // 礼物喇叭
	MT_PUT_GIFT_BOX             = "putGiftBox"           // 投放空投宝箱
	MT_GRAB_GIFT_BOX            = "grabGiftBox"          // 抢礼物宝箱
	MT_GRAB_GIFT_BOX_NOTIFIED   = "grabGiftBoxNotified"  // 抢礼物宝箱结果通知
	MT_GRAB_GIFT_BOX_BROAD_CAST = "grabGiftBoxBroadcast" // 抢礼物宝箱结果广播
	MT_BALANCE                  = "balance"              //余额

	// 道具服务
	MT_PROP_SEND = "propSend" // 发送道具 道具消息广播

	// 直播服务
	MT_LIVE_SETTING = "liveSetting" // 主播设置了直播间属性的通用的通知消息设置值，仅本场有效

	MT_TOKEN_INVALID_NOTIFY = "tokenInvalidNotify"

	MT_GAMOUNT = "gamount" //礼物总金额

	MT_SHIELD = "shield" //用户屏蔽

	MT_BANPOST = "banPost" //用户禁言

	MT_BANPOST_ALL = "banPostAll" //整个房间禁言

	MT_POST_INTERVAL = "postInterval" //房间发言间隔时间设置

	MT_WARN = "warn" //直播警告

	MT_UNIVERSAL_BROADCAST = "universalBroadcast" //直播间通用广播消息

	MT_MSG_BROADCAST = "msgBroadcast" //组件消息广播

	MT_NOTIFY = "notify" //直播通知

	MT_CONSUME_NOTIFY = "consumeNotify" //消费通知

	MT_GENERAL_FLUTTER_SCREEN = "generalFlutterScreen" //通用飘屏消息

	MT_FORBIDLIVE = "forbidLive" //禁止推流-封停

	MT_LIVE_CONTROL = "liveControl" //设置场控

	MT_GOOUT = "goOut" // 踢出房间

	MT_BROKE_STREAM = "brokeStream"

	MT_PUSH_STREAM = "pushStream"

	MT_USER_GRADE = "userGrade" // 用户|主播等级

	MT_LIVE_ADMIN_BANPOST = "liveAdminBanPost" // 超级管理员禁言

	MT_NOTIFY_ONLINE_USER = "notifyOnlineUser"

	MT_LIVE_ADMIN_GOOUT = "liveAdminGoOut" // 超级管理员全平台踢出

	MT_CACHE_SYN = "cacheSyn" // 广播里消息，通知每个节点同步缓存信息

	MT_RECEIVE_ORDER_QUEUE = "receiveOrderQueue" // 调取收单方加钱待处理队列

	MT_NOBILITY_TRUMPET_BROADCAST = "nobilityTrumpetBroadcast" // 贵族喇叭广播消息

	MT_OPEN_NOBILITY_BROADCAST = "openNobilityBroadcast" // 贵族开通广播消息(横幅，不是消费广播)

	MT_TURNTABLE_STATUS_UPDATE = "turntableStatusUpdate" // 轮盘状态推送

	MT_P2P_MESSAGE = "P2Pmessage"

	/*
	 * 连麦PK消息类型
	 */
	MT_LIANMAI_INVITING = "lianmaiInviting" //指定连麦邀请

	MT_DECLINE_INVITING = "declineInviting" // 拒绝指定连麦邀请

	MT_CANCEL_INVITING = "cancelInviting" // 取消指定连麦

	MT_LIANMAI_MATCH_SUCCESS = "lianmaiMatchSuccess" // 匹配成功（包括指定连麦）

	MT_LIANMAI_MATCH_TIMEOUT = "lianmaiMatchTimeout" // 连麦匹配超时

	MT_LIANMAI_SUCCESS = "lianmaiSuccess" // 连麦成功

	MT_LIANMAI_END = "lianmaiEnd" // 连麦结束

	MT_PK_START = "pkStart" // pk开始

	MT_NOTIFY_FP = "notifyFP" // 通知前端火力值

	MT_PK_END = "pkEnd" // pk结束

	MT_FIRST_KILL = "firstKill" // 首杀通知

	MT_ASSIST_KING = "assistKing" // 助攻王通知

	// V06B01抽奖消息协议
	MT_LIVE_DRAW_EXAMINE_RESULT     = "liveDrawExamineResult"     // 直播抽奖审核结果通知->通知直播间主播
	MT_LIVE_DRAW_START              = "liveDrawStart"             // 直播抽奖开始通知->通知直播主播和观众抽奖开始
	MT_LIVE_DRAW_FINISHED           = "liveDrawFinished"          // 直播抽奖结束通知->通知主播和观众抽奖结束
	MT_USER_PRIVATE_MESSAGE         = "userPrivateMessage"        // 中奖通知->通知主播和中奖观众
	MT_USER_PRIVATE_MESSAGE_RECEIPT = "userPrivateMessageReceipt" // 用户私信回执
)
