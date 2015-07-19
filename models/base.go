package models

import "gopkg.in/mgo.v2"

// 提供操作数据库的基类
type DbBase struct{}

func (d *DbBase) Session() *mgo.Session {
	return _Session
}

func (d *DbBase) Database() *mgo.Database {
	return _Database
}

func (d *DbBase) Collection(cName string) *mgo.Collection {
	return d.Database().C(cName)
}

func (d *DbBase) Find(cName string, query, selector interface{}) *mgo.Query {
	return d.Database().C(cName).Find(query).Select(selector)
}
