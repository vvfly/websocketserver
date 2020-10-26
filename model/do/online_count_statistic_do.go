package do

type OnlineCountStatistic struct {
	ID                  string `bson:"_id"`
	ServerID            string `bson:"serverId"` // websocket网关节点id
	LiveID              string `bson:"liveId"`
	TotalCount          int64  `bson:"totalCount"`
	GuardCount          int64  `bson:"guardCount"`
	MonthGuardCount     int64  `bson:"monthGuardCount"`
	YearGuardCount      int64  `bson:"yearGuardCount"`
	NobilityCount       int64  `bson:"nobilityCount"`
	RoomManagerCount    int64  `bson:"roomManagerCount"`
	NobilityLevel1Count int64  `bson:"nobilityLevel1Count"`
	NobilityLevel2Count int64  `bson:"nobilityLevel2Count"`
	NobilityLevel3Count int64  `bson:"nobilityLevel3Count"`
	NobilityLevel4Count int64  `bson:"nobilityLevel4Count"`
	NobilityLevel5Count int64  `bson:"nobilityLevel5Count"`
	NobilityLevel6Count int64  `bson:"nobilityLevel6Count"`
	NobilityLevel7Count int64  `bson:"nobilityLevel7Count"`
}
