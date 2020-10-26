package proto

const (
	RankLength = 10
)

const (
	WEBSOCKET_PING = "ping"
	WEBSOCKET_PONG = "pong"
)

const (
	WEBSOCKET_PROTO_REQUEST_CONNECT_INFO = "WEBSOCKET_PROTO_REQUEST_CONNECT_INFO"
)

// NobilityType 贵族类型
const (
	NobilityTypeNo     = -1 // 未开通贵族
	NobilityTypeLevel1 = 1  // 游侠
	NobilityTypeLevel2 = 2  // 骑士
	NobilityTypeLevel3 = 3  // 子爵
	NobilityTypeLevel4 = 4  // 伯爵
	NobilityTypeLevel5 = 5  // 公爵
	NobilityTypeLevel6 = 6  // 国王
	NobilityTypeLevel7 = 7  // 皇帝
)

// GuardType
const (
	GuardTypeNo    = "0" // 无守护
	GuardTypeWeek  = "1" // 无守护
	GuardTypeMonth = "2" // 无守护
	GuardTypeYear  = "3" // 无守护
)

// TokenType
const (
	TokenTypeLogin   = "1" // 登陆用户
	TokenTypeVisitor = "2" // 游客
)

// EnterType
const (
	EnterTypeAnchor   = "1" // 开播端
	EnterTypeAudience = "2" // 观看端
)

// ScopeType
const (
	_             int = iota
	BY_APP_ID         // 1 根据单个渠道ID广播
	BY_LIVE_ID        // 2 根据单个直播间ID广播
	BY_SESSION_ID     // 3 根据单个sessionID广播
	BY_ALL            // 4 向整个平台所有用户广播
	BY_BULK           // 5 批量消息广播

	// 运营后台系统广播
	BY_APP_IDS    = 21 // 根据单个渠道ID, 找到其它的所有在线的直播间，向每个直播间的所有该渠道的用户广播（粒度为用户）
	BY_LIVE_IDS   = 22 // 根据单个直播间ID广播, 给本直播间内当前渠道的所有用户广播
	BY_SUB_APP_ID = 23 // 根据子渠道ID广播
)

// 错误码
const (
	SUCCESS = 1
	FAIL    = -1
)

// Role 直播间用户角色
const (
	ANCHOR       = "1" // 主播
	AUDIENCE     = "2" // 观众
	ROOM_MANAGER = "3" // 房间管理员
	CLAN_ADMIN   = "4" // 家族管理员
	LIVE_ADMIN   = "5" // 直播间超级管理员
)

// client to ws
type MsgReq struct {
	MessageType  string      `json:"messageType"`
	BusinessData interface{} `json:"businessData"`
	R            string      `json:"r"`
	T            string      `json:"t"`
	S            string      `json:"s"`
}

// ws to client
type MsgResp struct {
	MessageType  string `json:"messageType"`
	BusinessData struct {
		Code       int         `json:"code"` // 返回数据code=1表示成功，-1表示失败
		Message    string      `json:"message"`
		ResultData interface{} `json:"resultData"`
	} `json:"businessData"`
}

// kafka to ws
type MsgReturnKey struct {
	ScopeType           int      `json:"scopeType"`
	ScopeIDList         []string `json:"scopeIdList"`        // 向多用户、多房间、多渠道广播，指定的ID列表,兼容单个的ID
	ExcludeScopeIDList  []string `json:"excludeScopeIdList"` // 要排除的ID的列表，预留设计
	Code                int      `json:"code"`               // 成功或异常编码
	CodeMessage         string   `json:"codeMessage"`        // 成功固定为 SUCCESS, 失败为失败(异常)原因
	MessageType         string   `json:"messageType"`
	AppIDForCurrentUser string   `json:"appIdForCurrentUser"` // 当前用户所在渠道
}

type MsgReturnValueBulkData struct {
	MessageType string      `json:"messageType"`
	ScopeType   int         `json:"scopeType"`
	ScopeID     string      `json:"scopeId"`
	Text        interface{} `json:"text"`
}
type MsgReturnValueBulk struct {
	MsgReturnValueBulkData []MsgReturnValueBulkData
}
