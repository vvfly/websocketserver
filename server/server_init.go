package server

import (
	"os"

	"github.com/BurntSushi/toml"
	mongoclient "github.com/luckyweiwei/base/cache/mongo-client"
	redisclient "github.com/luckyweiwei/base/cache/redis-client"
	. "github.com/luckyweiwei/base/logger"
	mysqlclient "github.com/luckyweiwei/base/orm/mysql-client"
	"github.com/luckyweiwei/base/utils"
	"github.com/luckyweiwei/websocketserver/kafka/consumer"
	"github.com/luckyweiwei/websocketserver/kafka/producer"
	"github.com/luckyweiwei/websocketserver/model"
	"github.com/luckyweiwei/websocketserver/schedule"
	"github.com/luckyweiwei/websocketserver/server/connectionm"
	"github.com/luckyweiwei/websocketserver/server/consul"
	remoteservice "github.com/luckyweiwei/websocketserver/server/remote-service"
)

func LoadConfig() *model.ServerConfig {
	cfile := SERVER_CONFIG
	if len(os.Args) == 2 {
		cfile = os.Args[1]
	}

	_, err := os.Stat(cfile)
	if err != nil {
		Log.Error(err)
		os.Exit(4)
	}

	serverConfig := model.GetServerConfig()
	_, err = toml.DecodeFile(cfile, serverConfig)
	if err != nil {
		Log.Error(err)
		os.Exit(4)
	}

	Log.Debugf("server config = %v", utils.FormatStruct(serverConfig))

	return serverConfig
}

func ServerInit() {
	// server config init
	serverConfig := model.GetServerConfig()

	// init redis
	Log.Info("redis init...")
	err := redisclient.RedisClientManagerInit(serverConfig.RedisConf)
	if err != nil {
		Log.Error(err)
		os.Exit(1)
	}

	// init mongoDB
	Log.Info("mongo init...")
	err = mongoclient.MongoClientManagerInit(serverConfig.MongoConf)
	if err != nil {
		Log.Error(err)
		os.Exit(1)
	}

	// init mysql
	Log.Info("mysql init...")
	mysqlclient.MysqlClientManagerInit(serverConfig.MysqlConf)

	// consul init
	Log.Info("consul init...")
	consul.ConsulManagerInit()

	// connection manager init
	Log.Info("connection manager init...")
	connectionm.ConnectionManagerInit()

	// init kafka producer
	Log.Info("kafka producer init...")
	producer.InitKafkaProducer()

	// init kafka consumer
	Log.Info("kafka consumer init...")
	consumer.InitKafkaConsumer()

	// remote service call manager init
	Log.Info("remote service call manager init")
	remoteservice.RemoteServiceCallManagerInit()

	// schedule init
	Log.Info("schedule init")
	schedule.ScheduleInit()
}
