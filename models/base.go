package models

import "gopkg.in/mgo.v2"

// DbBase 提供操作数据库的基类
type DbBase struct{}

// Session 获取当前操作数据库Session
func (d *DbBase) Session() *mgo.Session {
	return _Session
}

// Database 获取当前操作数据库
func (d *DbBase) Database() *mgo.Database {
	return _Database
}

// Collection 获取操作集合
func (d *DbBase) Collection(cName string) *mgo.Collection {
	return d.Database().C(cName)
}

// Find 提供简单的Find查询
func (d *DbBase) Find(cName string, query, selector interface{}) *mgo.Query {
	return d.Database().C(cName).Find(query).Select(selector)
}
