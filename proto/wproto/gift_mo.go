package wproto

/*
{
        "giftCostType": "1",
        "markId": "97",
        "giftName": "礼物批量赠送66",
        "sex": "1",
        "boxType": "1",
        "avatar": "http://imgdown.fangqie666.com:9327/img///5/91a263efdbcb74fea94548e66e98762a_s1.jpg",
        "anchorId": "1ed8bfc1-324f-4834-8065-d3fad9943f88",
        "isStarGift": "0",
        "anchorName": "安卓用户123",
        "uuid": "2d4bf89e-f208-4259-bd6b-a8ae0c48238c",
        "effectType": "2",
        "liveCount": "0",
        "createTime": "1593661373572",
        "price": "100",
        "appId": "201",
        "clientIp": "10.0.3.15",
        "giftNum": "1",
        "guardType": "0"
    }

*/
// client to ws
type MsgGiftReqData struct {
	GiftCostType string `json:"giftCostType"` // 礼物消费类型 1人民币，2积分
	MarkID       string `json:"markId"`       // 礼物的markId
	GiftName     string `json:"giftName"`     // 礼物名字
	Sex          string `json:"sex"`
	BoxType      string `json:"boxType"`
	Avatar       string `json:"avatar"`
	AnchorID     string `json:"anchorId"`
	IsStarGift   string `json:"isStarGift"`
	AnchorName   string `json:"anchorName"`
	UUID         string `json:"uuid"`       // 全局唯一标识
	EffectType   string `json:"effectType"` //礼物特效类型 1：静态，2：动态
	LiveCount    string `json:"liveCount"`  // 直播场次
	CreateTime   string `json:"createTime"` // 礼物发送时间
	Price        string `json:"price"`      // 礼物价格
	AppID        string `json:"appId"`
	ClientIP     string `json:"clientIp"`
	GiftNum      string `json:"giftNum"` // 礼物数量
	GuardType    string `json:"guardType"`
	FollowStatus int    `json:"followStatus"` // 1关注 0未关注
}
type MsgGiftReq struct {
	MessageType  string         `json:"messageType"`
	BusinessData MsgGiftReqData `json:"businessData"`
	R            string         `json:"r"`
	T            string         `json:"t"`
	S            string         `json:"s"`
}

// ws to gift
type MsgGiftWsReqData struct {
	GiftCostType string `json:"giftCostType"`
	Role         string `json:"role"`
	MarkID       string `json:"markId"`
	ExpGrade     string `json:"expGrade"`
	GiftName     string `json:"giftName"`
	Sex          string `json:"sex"`
	Ks           string `json:"ks"`
	BoxType      string `json:"boxType"`
	Avatar       string `json:"avatar"`
	AnchorID     string `json:"anchorId"`
	IsStarGift   string `json:"isStarGift"`
	UserName     string `json:"userName"`
	AnchorName   string `json:"anchorName"`
	UUID         string `json:"uuid"`
	UserID       string `json:"userId"`
	LiveID       string `json:"liveId"`
	EffectType   string `json:"effectType"`
	LiveCount    string `json:"liveCount"`
	CreateTime   string `json:"createTime"`
	Price        string `json:"price"`
	AppID        string `json:"appId"`
	ClientIP     string `json:"clientIp"`
	GiftNum      string `json:"giftNum"`
	GuardType    string `json:"guardType"`
	FollowStatus int    `json:"followStatus"`
}

type MsgGiftWsReq struct {
	R            string           `json:"r"`
	S            string           `json:"s"`
	MessageType  string           `json:"messageType"`
	T            string           `json:"t"`
	BusinessData MsgGiftWsReqData `json:"businessData"`
}

// post data to gift
type PostBody struct {
	SessionID string `json:"sessionId"`
	Message   string `json:"message"`
}
