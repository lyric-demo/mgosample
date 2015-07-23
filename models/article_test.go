package models

import (
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var (
	TestId = bson.NewObjectId()
)

func TestArticleInsert(t *testing.T) {
	art := new(Article)
	art.Id = TestId
	art.Title = "测试标题"
	art.Body = "测试标题内容"
	art.Author = "Lyric"
	art.PublishDate = time.Now()
	art.Tags = []string{"测试"}
	art.Type = "测试分类"
	if err := art.Insert(); err != nil {
		t.Error(err)
		return
	}
	t.Log("Insert success.")
}

func TestArticleGetSingleData(t *testing.T) {
	art := new(Article)
	art.Id = TestId
	data, err := art.GetSingleData()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}

func TestArticleUpdate(t *testing.T) {
	art := new(Article)
	art.Id = TestId
	art.Title = "Test Title"
	if err := art.Update(); err != nil {
		t.Error(err)
		return
	}
	t.Log("Update success.")
}

func TestArticleGetData(t *testing.T) {
	var data []Article
	art := new(Article)
	if err := art.GetData(&data); err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}

func TestArticleDelete(t *testing.T) {
	art := new(Article)
	art.Id = TestId
	if err := art.Delete(); err != nil {
		t.Error(err)
		return
	}
	t.Log("Delete success.")
}
