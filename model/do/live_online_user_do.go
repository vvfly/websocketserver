package do

type LiveOnlineUser struct {
	ID            int64  `gorm:"AUTO_INCREMENT;primary_key"`
	ServerID      string `gorm:"column:serverId;size:128;not null"` // websocket网关节点id
	LiveID        string `gorm:"column:liveId;size:64;not null"`
	UserID        string `gorm:"column:userId;size:128"`
	EntryLiveTime int64  `gorm:"column:entryLiveTime"`
	OnlineUser    string `gorm:"column:onlineUser;type:text"`
}

func (LiveOnlineUser) TableName() string {
	return "LiveOnlineUser"
}
