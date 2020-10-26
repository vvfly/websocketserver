package do

type SocketConnectCountInfo struct {
	ID       string       `bson:"_id"`
	ServerID string       `bson:"serverId"` // websocket网关节点id
	AppID    string       `bson:"appId"`
	LiveID   string       `bson:"liveId"`
	Count    int          `bson:"count"`
	Rank     []OnlineUser `bson:"rank"`
}
