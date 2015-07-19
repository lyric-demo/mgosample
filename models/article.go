package models

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 文章管理
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

// 集合名称
func (this *Article) CName() string {
	return "articles"
}

// 查询所有数据
func (this *Article) GetData(data *[]Article) error {
	return this.Find(this.CName(), nil, nil).All(data)
}

// 查询单条数据
func (this *Article) GetId() (Article, error) {
	var article Article
	err := this.Collection(this.CName()).FindId(this.Id).One(&article)
	return article, err
}

// 保存数据
func (this *Article) Save() (*mgo.ChangeInfo, error) {
	return this.Collection(this.CName()).Upsert(bson.M{"_id": this.Id}, this)
}

// 移除数据（id）
func (this *Article) Remove() error {
	return this.Collection(this.CName()).RemoveId(this.Id)
}
