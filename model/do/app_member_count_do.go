package do

type AppMemberCount struct {
	ID       string `bson:"_id"`
	ServerID string `bson:"serverId"` // websocket网关节点id
	AppID    string `bson:"appId"`
	Count    int64  `bson:"count"`
}
