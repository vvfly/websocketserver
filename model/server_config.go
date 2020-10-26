/*
** 使用 Spring Cloud Config Server 管理配置
 */

package model

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	mongoclient "github.com/luckyweiwei/base/cache/mongo-client"
	redisclient "github.com/luckyweiwei/base/cache/redis-client"
	mysqlclient "github.com/luckyweiwei/base/orm/mysql-client"
	"github.com/luckyweiwei/base/utils"

	. "github.com/luckyweiwei/base/logger"
	"github.com/spf13/viper"
)

const (
	// flag
	ConfigName          = "spring.cloud.config.name"
	ConfigFullName      = "spring.cloud.config.full.name"
	ConfigLabel         = "spring.cloud.config.label"
	ConfigProfile       = "spring.cloud.config.profile"
	ConfigUri           = "spring.cloud.config.uri"
	ConfigServiceAddr   = "service_addr"
	ConfigKafkaGroupUri = "kafka_group_uri"

	ConfigPprofPort       = "pprof.port"
	ConfigPprofDisabled   = "pprof.disabled"
	ConfigDebugLogEnabled = "logging.debug.enabled"

	EnableDebug     = "enableDebug"
	ApplicationName = "spring.application.name"
	ProfilesActive  = "spring.profiles.active"

	// log
	LoggingLevel      = "logging.level.root"
	LoggingMaxSize    = "logging.max-size"
	LoggingMaxBackups = "logging.max-backups"
	LoggingMaxAge     = "logging.max-age"

	// websocket server
	WebsocketServerHost         = "websocket.server.host"
	WebsocketServerPort         = "websocket.netty.port"
	WebsocketEncryptKey         = "websocket.encryptKey"
	WebsocketSignatureSecretKey = "websocket.signatureSecretKey"

	// api server
	ApiServerPort = "server.port"

	// consul
	ConsulHost = "spring.cloud.consul.host"
	ConsulPort = "spring.cloud.consul.port"

	// kafka
	KafkaVersion          = "kafka.version"
	KafkaGroupID          = "kafka.groupId"
	KafkaGroupIDForAcross = "kafka.groupIdForAcross"
	KafkaServer           = "kafka.server"
	KafkaServerForAcross  = "kafka.serverForAcross"

	// redis
	RedisDatabase     = "spring.redis.database"
	RedisHost         = "spring.redis.host"
	RedisPort         = "spring.redis.port"
	RedisPassword     = "spring.redis.password"
	RedisActive       = "spring.redis.active"
	RedisIdle         = "spring.redis.idle"
	RedisDialTimeout  = "spring.redis.dialTimeout"
	RedisReadTimeout  = "spring.redis.readTimeout"
	RedisWriteTimeout = "spring.redis.writeTimeout"
	RedisIdleTimeout  = "spring.redis.idleTimeout"

	// mongo
	MongoDBURL      = "spring.data.mongodb.uri"
	MongoDBDatabase = "spring.data.mongodb.database"

	// mysql
	DatasourceDsn         = "spring.datasource.dsn"
	DatasourceIdleSize    = "spring.datasource.idleSize"
	DatasourceMaxSize     = "spring.datasource.maxSize"
	DatasourceMaxLifeTime = "spring.datasource.maxLifeTime"
	DatasourceSqlDebug    = "spring.datasource.sqlDebug"
)

var (
	ConfigNameValue            *string = nil
	ConfigFullNameValue        *string = nil
	ConfigLabelValue           *string = nil
	ConfigProfileValue         *string = nil
	ConfigUriValue             *string = nil
	ConfigServiceAddrValue     *string = nil
	ConfigPprofPortValue       *string = nil
	ConfigPprofDisabledValue   *string = nil
	ConfigDebugLogEnabledValue *string = nil
	ConfigKafkaGroupUriValue   *string = nil
)

func flagInit() {
	ConfigNameValue = flag.String(ConfigName, "ws", "server config name")
	ConfigFullNameValue = flag.String(ConfigFullName, "tl-ws-server", "server config full name")
	ConfigLabelValue = flag.String(ConfigLabel, "master", "server config label")
	ConfigProfileValue = flag.String(ConfigProfile, "rfbak", "server config profile")
	ConfigUriValue = flag.String(ConfigUri, "http://api.zhangling.link:8084", "server config uri")
	ConfigServiceAddrValue = flag.String(ConfigServiceAddr, "192.168.88.206", "server config service addr")
	ConfigPprofPortValue = flag.String(ConfigPprofPort, "23000", "server config pprof port")
	ConfigPprofDisabledValue = flag.String(ConfigPprofDisabled, "1", "server config pprof disabled")
	ConfigDebugLogEnabledValue = flag.String(ConfigDebugLogEnabled, "0", "server config log debug enabled")
	ConfigKafkaGroupUriValue = flag.String(ConfigKafkaGroupUri, "", "server config kafka server")

	flag.Parse()
}

func loadRemoteConfig() (err error) {
	// http://api.zhangling.link:8084/tl-ws-server/rfbak/master/ws-rfbak.properties
	confAddr := fmt.Sprintf("%v/%v/%v/%v/%v-%v.properties",
		*ConfigUriValue,
		*ConfigFullNameValue,
		*ConfigProfileValue,
		*ConfigLabelValue,
		*ConfigNameValue,
		*ConfigProfileValue,
	)
	Log.Debugf("Spring Cloud Config Server 管理配置地址为:%v", confAddr)

	resp, err := http.Get(confAddr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 设置配置文件格式: properties
	viper.SetConfigType("properties")

	// 载入配置文件
	err = viper.ReadConfig(resp.Body)
	if err != nil {
		return err
	}

	// 载入flag参数
	viper.Set(ConfigServiceAddr, *ConfigServiceAddrValue)
	viper.Set(ConfigPprofPort, *ConfigPprofPortValue)
	viper.Set(ConfigPprofDisabled, *ConfigPprofDisabledValue)
	viper.Set(ConfigDebugLogEnabled, *ConfigDebugLogEnabledValue)

	if *ConfigKafkaGroupUriValue != "" {
		viper.Set(ConfigKafkaGroupUri, *ConfigKafkaGroupUriValue)
	}

	return nil
}

func setServerConfig() {
	// serverID
	serverID := viper.GetString(ConfigServiceAddr) + "_" + viper.GetString(ApiServerPort)

	// consul
	consulAddr := net.JoinHostPort(viper.GetString(ConsulHost), viper.GetString(ConsulPort))

	// api server config
	apiServerConfig := ApiServerConfig{
		ServerAddr: viper.GetString(ConfigServiceAddr),
		ServerName: viper.GetString(ApplicationName),
		ServerID:   viper.GetString(ApplicationName) + "-" + viper.GetString(ConfigServiceAddr),
		ApiPort:    viper.GetInt(ApiServerPort),
	}

	// pprof config
	pprofConfig := PprofConfig{
		PprofPort: viper.GetInt(ConfigPprofPort),
		Disabled:  viper.GetInt(ConfigPprofDisabled),
	}

	// redis
	var redisConfigs = []redisclient.RedisClientConfig{}
	redisConfig := redisclient.RedisClientConfig{
		Name:         RedisDBNameDB1,
		Addr:         net.JoinHostPort(viper.GetString(RedisHost), viper.GetString(RedisPort)),
		Active:       viper.GetInt(RedisActive),
		Idle:         viper.GetInt(RedisIdle),
		DialTimeout:  viper.GetInt(RedisDialTimeout),
		ReadTimeout:  viper.GetInt(RedisReadTimeout),
		WriteTimeout: viper.GetInt(RedisWriteTimeout),
		IdleTimeout:  viper.GetInt(RedisIdleTimeout),
		DBNum:        viper.GetString(RedisDatabase),
		Password:     viper.GetString(RedisPassword),
	}
	redisConfigs = append(redisConfigs, redisConfig)

	// mongo
	var mongoConfigs = []mongoclient.MongoClientConfig{}
	mongoConfig := mongoclient.MongoClientConfig{
		Name:     MongoDBNameDB1,
		URL:      viper.GetString(MongoDBURL),
		Database: viper.GetString(MongoDBDatabase),
	}
	mongoConfigs = append(mongoConfigs, mongoConfig)

	// mysql
	var mysqlConfigs = []mysqlclient.MysqlClientConfig{}
	mysqlConfig := mysqlclient.MysqlClientConfig{
		URL:         viper.GetString(DatasourceDsn),
		IdleSize:    viper.GetInt(DatasourceIdleSize),
		MaxSize:     viper.GetInt(DatasourceMaxSize),
		MaxLifeTime: viper.GetInt64(DatasourceMaxLifeTime),
		SqlDebug:    viper.GetInt(DatasourceSqlDebug),
		Memory:      false,
	}
	mysqlConfigs = append(mysqlConfigs, mysqlConfig)

	// kafka producer
	var producerBrokerList = make([]string, 0)
	producerBrokerList = append(producerBrokerList, viper.GetString(KafkaServer))
	kafkaProducerConfig := KafkaProducerConfig{
		BrokerList: producerBrokerList,
		BatchNumer: 3,
	}

	// kafka consumer
	var (
		kafkaConsumerConfigMap = make(map[string]KafkaConsumerConfig)
		kVersion               = viper.GetString(KafkaVersion)
		kGroupID               = viper.GetString(KafkaGroupID)
		kGroupIDForAcross      = viper.GetString(KafkaGroupIDForAcross)
		kServer                = viper.GetString(KafkaServer)
		kServerForAcross       = viper.GetString(KafkaServerForAcross)
	)
	serviceAddr := viper.GetString(ConfigServiceAddr)

	if !strings.Contains(kGroupID, serviceAddr) {
		kGroupID = serviceAddr + kGroupID
	}
	if !strings.Contains(kGroupIDForAcross, serviceAddr) {
		kGroupIDForAcross = serviceAddr + kGroupIDForAcross
	}

	fooTopics := make([]string, 0)
	fooBrokers := make([]string, 0)
	fooBrokers = append(fooBrokers, kServer)
	kafkaFooConfig := KafkaConsumerConfig{
		Version:         kVersion,
		Topics:          fooTopics,
		GroupID:         kGroupID,
		BrokerAddresses: fooBrokers,
	}
	kafkaConsumerConfigMap[TAG_EXTERNAL_FOO] = kafkaFooConfig

	barTopics := make([]string, 0)
	barBrokers := make([]string, 0)
	barBrokers = append(barBrokers, kServerForAcross)
	kafkaBarConfig := KafkaConsumerConfig{
		Version:         kVersion,
		Topics:          barTopics,
		GroupID:         kGroupIDForAcross,
		BrokerAddresses: barBrokers,
	}
	kafkaConsumerConfigMap[TAG_EXTERNAL_BAR] = kafkaBarConfig

	// log config
	logLevel := viper.GetString(LoggingLevel)
	if logLevel == "" {
		logLevel = "ERROR"
	}
	logDebugEnabled := viper.GetInt(ConfigDebugLogEnabled)
	if logDebugEnabled == 1 {
		logLevel = "DEBUG"
	}
	logMaxSize := viper.GetInt(LoggingMaxSize)
	if logMaxSize == 0 {
		logMaxSize = 200
	}
	logMaxAge := viper.GetInt(LoggingMaxAge)
	if logMaxAge == 0 {
		logMaxAge = 7
	}
	LoggingMaxBackups := viper.GetInt(LoggingMaxBackups)
	if LoggingMaxBackups == 0 {
		LoggingMaxBackups = 10
	}
	logConfig := TLogConfig{
		Level:      logLevel,
		MaxSize:    logMaxSize,
		MaxBackups: LoggingMaxBackups,
		MaxAge:     logMaxAge,
	}

	config = &ServerConfig{
		DebugMode:     viper.GetInt(EnableDebug),
		ConsulAddr:    consulAddr,
		ServerID:      serverID,
		WebsocketPort: viper.GetInt(WebsocketServerPort),

		Des3Key4WsMsg:               viper.GetString(WebsocketEncryptKey),
		SignatureSecretKey:          viper.GetString(WebsocketSignatureSecretKey),
		EntryTimeInterval:           EntryTimeInterval,
		PingTimeInterval:            PingTimeInterval,
		OnlineInfoStatisticInterval: OnlineInfoStatisticInterval,

		PprofConf:     pprofConfig,
		ApiServerConf: apiServerConfig,
		ProducerConf:  kafkaProducerConfig,
		ConsumerConf:  kafkaConsumerConfigMap,
		RedisConf:     redisConfigs,
		MongoConf:     mongoConfigs,
		MysqlConf:     mysqlConfigs,
		LogConf:       logConfig,
	}

	Log.Debugf("server config = %v", utils.FormatStruct(config))
}

func ServerConfigInit() {
	flagInit()

	err := loadRemoteConfig()
	if err != nil {
		Log.Error(err)
		os.Exit(1)
	}

	setServerConfig()
}
