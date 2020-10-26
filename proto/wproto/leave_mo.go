package wproto

// client to ws
type MsgLeaveReqData struct {
}
type MsgLeaveReq struct {
	MessageType  string          `json:"messageType"`
	BusinessData MsgLeaveReqData `json:"businessData"`
	R            string          `json:"r"`
	T            string          `json:"t"`
	S            string          `json:"s"`
}

// ws to kafka
type MsgLeaveWsReqData struct {
	LiveID    string `json:"liveId"`
	Role      string `json:"role"`
	ExpGrade  int    `json:"expGrade"`
	GuardType string `json:"guardType"`
	UserName  string `json:"userName"`
	Avatar    string `json:"avatar"`
	ClientIP  string `json:"clientIp"`
	KS        int64  `json:"ks"` // 当前时间戳秒数
}

type MsgLeaveWsReq struct {
	R            string            `json:"r"`
	S            string            `json:"s"`
	MessageType  string            `json:"messageType"`
	T            string            `json:"t"`
	BusinessData MsgLeaveWsReqData `json:"businessData"`
}

// kafka to ws
type MsgLeaveReturnValue struct {
	Role         string `json:"role"`
	UserID       string `json:"userId"`
	LastLiveData struct {
		StartTime     int64  `json:"startTime"`
		EndTime       int64  `json:"endTime"`
		Herald        string `json:"herald"`
		PublishTime   string `json:"publishTime"`
		MaxPopularity int    `json:"maxPopularity"`
		NobilityType  int    `json:"nobilityType"`
		OpenID        string `json:"openId"`
		AppID         string `json:"appId"`
	} `json:"lastLiveData"`
}

type MsgLeaveReturnValueBulkData struct {
	MessageType string              `json:"messageType"`
	ScopeType   int                 `json:"scopeType"`
	ScopeID     string              `json:"scopeId"`
	Text        MsgLeaveReturnValue `json:"text"`
}
type MsgLeaveReturnValueBulk struct {
	MsgLeaveReturnValueBulkData []MsgLeaveReturnValueBulkData
}

// ws to client
type MsgLeaveRespData struct {
	UserID string `json:"userId"`
	Role   string `json:"role"`
	// StreamName   string `json:"streamName"`
	LastLiveData struct {
		StartTime     int64  `json:"startTime"`
		EndTime       int64  `json:"endTime"`
		Herald        string `json:"herald"`
		PublishTime   string `json:"publishTime"`
		MaxPopularity string `json:"maxPopularity"`
	} `json:"lastLiveData"`
}
type MsgLeaveResp struct {
	MessageType  string `json:"messageType"`
	BusinessData struct {
		Code       int              `json:"code"`
		Message    string           `json:"message"`
		ResultData MsgLeaveRespData `json:"resultData"`
	} `json:"businessData"`
}
