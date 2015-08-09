package models

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"

	"github.com/LyricTian/mgosample/conf"
)

var (
	_Session  *mgo.Session
	_Database *mgo.Database
)

func init() {
	_Session, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%d", conf.DBCONFIG.Host, conf.DBCONFIG.Port))
	if err != nil {
		panic(fmt.Sprintf("Initialize mongodb error:%v", err))
	}
	_Database = _Session.DB(conf.DBCONFIG.Database)
	if err = _Session.Ping(); err != nil {
		panic(fmt.Sprintf("MongoDB execute ping error:%v", err))
	}
	log.Println("MongoDB initialize success.")
	mgo.SetDebug(conf.APPCONFIG.Debug)
}
