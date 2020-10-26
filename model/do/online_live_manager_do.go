package do

type OnlineLiveManager struct {
	ID        string   `bson:"_id"`
	ServerID  string   `bson:"serverId"` // websocket网关节点id
	LiveID    string   `bson:"liveId"`
	UserIDSet []string `bson:"userIdSet"`
}
