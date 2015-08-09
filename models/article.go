package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Article 文章管理结构
type Article struct {
	DbBase      `bson:",omitempty"`
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title       string        `bson:",omitempty" json:"title"`        //标题
	Body        string        `bson:",omitempty" json:"body"`         // 内容
	Author      string        `bson:",omitempty" json:"author"`       // 作者
	PublishDate time.Time     `bson:",omitempty" json:"publish_date"` // 发布日期
	Tags        string        `bson:",omitempty" json:"tags"`         // 标签
	Type        string        `bson:",omitempty" json:"type"`         // 分类
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
	return a.Collection(a.CName()).UpdateId(a.ID, a)
}

// Delete 删除记录
func (a *Article) Delete() error {
	return a.Collection(a.CName()).RemoveId(a.ID)
}

// GetSingleData 获取单条数据
func (a *Article) GetSingleData() (Article, error) {
	var article Article
	err := a.Collection(a.CName()).FindId(a.ID).One(&article)
	return article, err
}

// GetData 获取所有数据
func (a *Article) GetData(data *[]Article) error {
	return a.Find(a.CName(), nil, nil).All(data)
}

// AllCount 获取所有数据条数
func (a *Article) AllCount() (int, error) {
	return a.Find(a.CName(), nil, nil).Count()
}

// GetPageData 获取所有分页数据
func (a *Article) GetPageData(skip, limit int, data *[]Article) error {
	return a.Find(a.CName(), nil, nil).Skip(skip).Limit(limit).All(data)
}
