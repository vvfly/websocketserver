package model

import (
	mongoclient "github.com/luckyweiwei/base/cache/mongo-client"
	redisclient "github.com/luckyweiwei/base/cache/redis-client"
	"github.com/luckyweiwei/base/logger"
	mysqlclient "github.com/luckyweiwei/base/orm/mysql-client"
)

// redis key
const (
	OnlineUserKey = "live:onlineAudience:"
)

const (
	TAG_BIZ          = "Biz"         // 处理本节点biz
	TAG_EXTERNAL_FOO = "ExternalFoo" // 非跨分组
	TAG_EXTERNAL_BAR = "ExternalBar" // 跨分组
)

const (
	EntryTimeInterval           int64 = 5   // 入场时间间隔 s
	PingTimeInterval            int64 = 100 // 心跳时间间隔 s
	OnlineInfoStatisticInterval int   = 8   // 在线信息统计时间间隔 s
)

const (
	GroupIDForWsBiz = "ForWsBiz"
)

const (
	RedisDBNameDB1 = "db1"
)

const (
	MongoDBNameDB1 = "db1"
)

type PprofConfig struct {
	PprofPort int
	Disabled  int
}

type ApiServerConfig struct {
	ServerAddr string
	ServerName string
	ServerID   string
	ApiPort    int
}

type KafkaConsumerConfig struct {
	Version         string   `toml:"version"`
	Topics          []string `toml:"topics"`
	GroupID         string   `toml:"group_id"`
	BrokerAddresses []string `toml:"broker_addresses"`
}
type KafkaProducerConfig struct {
	BrokerList []string
	BatchNumer int
}
type ServerConfig struct {
	DebugMode     int
	ConsulAddr    string
	ServerID      string
	WebsocketPort int

	Des3Key4WsMsg               string
	SignatureSecretKey          string
	EntryTimeInterval           int64
	PingTimeInterval            int64
	OnlineInfoStatisticInterval int

	PprofConf     PprofConfig
	ApiServerConf ApiServerConfig
	ProducerConf  KafkaProducerConfig
	ConsumerConf  map[string]KafkaConsumerConfig
	RedisConf     []redisclient.RedisClientConfig
	MongoConf     []mongoclient.MongoClientConfig
	MysqlConf     []mysqlclient.MysqlClientConfig
	LogConf       logger.TLogConfig
}

var config = &ServerConfig{}

func GetServerConfig() *ServerConfig {
	return config
}
