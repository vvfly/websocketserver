package dao

import (
	mysqlclient "github.com/luckyweiwei/base/orm/mysql-client"
	"github.com/luckyweiwei/websocketserver/model/do"
	"gorm.io/gorm"
)

type LiveOnlineUserDao struct {
	db *gorm.DB
}

var liveOnlineUserDao *LiveOnlineUserDao = nil

func NewLiveOnlineUserDao() *LiveOnlineUserDao {
	db := mysqlclient.MysqlClientInstance().Master()

	liveOnlineUserDao = &LiveOnlineUserDao{
		db: db,
	}
	return liveOnlineUserDao
}

// delete
func (dao *LiveOnlineUserDao) Delete(m *do.LiveOnlineUser) error {
	return dao.db.Where("serverId", m.ServerID).Delete(m).Error
}

// insert
func (dao *LiveOnlineUserDao) CreateOne(m *do.LiveOnlineUser) error {
	return dao.db.Create(m).Error
}
func (dao *LiveOnlineUserDao) CreateMany(m *[]do.LiveOnlineUser) error {
	return dao.db.Create(m).Error
}
