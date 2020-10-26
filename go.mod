module github.com/luckyweiwei/websocketServer

go 1.14

replace github.com/luckyweiwei/base => ../base

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Shopify/sarama v1.27.2
	github.com/forgoer/openssl v0.0.0-20201023062029-c3112b0c8700
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/gomodule/redigo/redis v0.0.0-20200429221454-e14091dffc1b
	github.com/gorilla/websocket v1.4.2
	github.com/hashicorp/consul/api v1.7.0
	github.com/orcaman/concurrent-map v0.0.0-20190826125027-8c72a8bb44f6
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/viper v1.7.1
	go.mongodb.org/mongo-driver v1.4.2
	gorm.io/gorm v1.20.5
)
