package models

import (
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var (
	TestId = bson.NewObjectId()
)

func TestSave(t *testing.T) {
	art := new(Article)
	art.Id = TestId
	art.Title = "测试标题"
	art.Body = "测试标题内容"
	art.Author = "Lyric"
	art.PublishDate = time.Now()
	art.Tags = []string{"测试"}
	art.Type = "测试分类"
	if _, err := art.Save(); err != nil {
		t.Error(err)
		return
	}
	t.Log("Save success.")
}

func TestGetId(t *testing.T) {
	art := new(Article)
	art.Id = TestId
	data, err := art.GetId()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}

func TestGetData(t *testing.T) {
	var data []Article
	art := new(Article)
	if err := art.GetData(&data); err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}

func TestRemove(t *testing.T) {
	art := new(Article)
	art.Id = TestId
	if err := art.Remove(); err != nil {
		t.Error(err)
		return
	}
	t.Log("Remove success.")
}
