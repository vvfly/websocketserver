package dao

import (
	"context"

	mongoclient "github.com/luckyweiwei/base/cache/mongo-client"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/websocketserver/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// 集合名称
const SocketConnectCountInfoCollectionName = "SocketConnectCountInfo"

// SocketConnectCountInfoModel
type MgoSocketConnectCountInfoDao struct {
	collection *mongo.Collection
}

func NewMgoSocketConnectCountInfoDao() *MgoSocketConnectCountInfoDao {
	mgo := new(MgoSocketConnectCountInfoDao)
	mgo.collection = mongoclient.GetMongoClientManager().GetMongoDatabaseClient(model.MongoDBNameDB1).Collection(SocketConnectCountInfoCollectionName)
	return mgo
}

// 查找一个
func (m *MgoSocketConnectCountInfoDao) FindOne(filter interface{}) *mongo.SingleResult {
	return m.collection.FindOne(context.TODO(), filter)

}

// 查找多个
func (m *MgoSocketConnectCountInfoDao) Find(filter interface{}) (*mongo.Cursor, error) {
	result, err := m.collection.Find(context.TODO(), filter)
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	return result, nil
}

// 新增一个
func (m *MgoSocketConnectCountInfoDao) InsertOne(value interface{}) (*mongo.InsertOneResult, error) {
	result, err := m.collection.InsertOne(context.TODO(), value)
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	return result, nil
}

// 新增多个
func (m *MgoSocketConnectCountInfoDao) InsertMany(docs []interface{}) (*mongo.InsertManyResult, error) {
	result, err := m.collection.InsertMany(context.TODO(), docs)
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	return result, nil
}

// 更新一个
func (m *MgoSocketConnectCountInfoDao) UpdateOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	result, err := m.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	return result, nil
}

// 删除多个
func (m *MgoSocketConnectCountInfoDao) DeleteMany(filter interface{}) (*mongo.DeleteResult, error) {
	result, err := m.collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	return result, nil
}
