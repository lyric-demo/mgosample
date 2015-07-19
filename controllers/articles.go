package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"

	"github.com/LyricTian/mgosample/models"
	"github.com/LyricTian/mgosample/modules/util"
)

// 注册文章管理路由
func RegisterArticle(g *echo.Group) {
	article := new(ArticleController)
	art := g.Group("/article")
	{
		art.Get("", article.Index)
		art.Post("/save", article.Save)
		art.Delete("/delete/:id", article.Remove)
		art.Get("/:id", article.GetId)
	}
}

// 文章管理
type ArticleController struct{}

func (this *ArticleController) Index(c *echo.Context) error {
	article := new(models.Article)
	var data []models.Article
	if err := article.GetData(&data); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, util.Success(data))
}

func (this *ArticleController) Save(c *echo.Context) error {
	var article models.Article
	if err := c.Bind(&article); err != nil {
		return err
	}
	if !bson.IsObjectIdHex(article.Id.Hex()) {
		article.Id = bson.NewObjectId()
	}
	if _, err := article.Save(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, util.Success(true, "保存成功！"))
}

func (this *ArticleController) Remove(c *echo.Context) error {
	id := c.Param("id")
	if id == "" || !bson.IsObjectIdHex(id) {
		return errors.New("Id error.")
	}
	article := new(models.Article)
	article.Id = bson.ObjectIdHex(id)
	if err := article.Remove(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, util.Success(true, "删除成功！"))
}

func (this *ArticleController) GetId(c *echo.Context) error {
	id := c.Param("id")
	if id == "" || !bson.IsObjectIdHex(id) {
		return errors.New("Id error.")
	}
	article := new(models.Article)
	article.Id = bson.ObjectIdHex(id)
	data, err := article.GetId()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, util.Success(data))
}
