package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Article 文章管理结构
type Article struct {
	DbBase
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Title       string        `bson:",omitempty"` //标题
	Body        string        `bson:",omitempty"` // 内容
	Author      string        `bson:",omitempty"` // 作者
	PublishDate time.Time     `bson:",omitempty"` // 发布日期
	Tags        []string      `bson:",omitempty"` // 标签
	Type        string        `bson:",omitempty"` // 分类
}

// CName 获取当前集合名称
func (a *Article) CName() string {
	return "articles"
}

// Insert 插入记录
func (a *Article) Insert() error {
	return a.Collection(a.CName()).Insert(a)
}

// Update 更新记录
func (a *Article) Update() error {
	return a.Collection(a.CName()).UpdateId(a.Id, bson.M{"$set": a})
}

// Delete 删除记录
func (a *Article) Delete() error {
	return a.Collection(a.CName()).RemoveId(a.Id)
}

// GetSingleData 获取单条数据
func (a *Article) GetSingleData() (Article, error) {
	var article Article
	err := a.Collection(a.CName()).FindId(a.Id).One(&article)
	return article, err
}

// GetData 获取所有数据
func (a *Article) GetData(data *[]Article) error {
	return a.Find(a.CName(), nil, nil).All(data)
}
