package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"

	"github.com/LyricTian/mgosample/models"
)

// 注册文章管理路由
func RegisterArticle(g *echo.Group) {
	article := new(ArticleController)
	groupArt := g.Group("/article")
	{
		groupArt.Get("", article.GetData)
		groupArt.Get("/:id", article.GetSingleData)
		groupArt.Post("", article.Insert)
		groupArt.Put("/:id", article.Update)
		groupArt.Delete("/:id", article.Delete)
	}
}

// ArticleController
// 文章管理控制器
type ArticleController struct{}

// GetResModel 获取响应模型
func (a *ArticleController) GetResModel(data interface{}) map[string]interface{} {
	return map[string]interface{}{"article": data}
}

// GetData 获取所有数据
func (a *ArticleController) GetData(c *echo.Context) error {
	article := new(models.Article)
	var data []models.Article
	if err := article.GetData(&data); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, a.GetResModel(data))
}

// GetSingleData 获取单条记录
func (a *ArticleController) GetSingleData(c *echo.Context) error {
	id := c.Param("id")
	if id == "" || !bson.IsObjectIdHex(id) {
		return errors.New("Id error.")
	}
	article := new(models.Article)
	article.Id = bson.ObjectIdHex(id)
	data, err := article.GetSingleData()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, a.GetResModel(data))
}

// Insert 插入记录
func (a *ArticleController) Insert(c *echo.Context) error {
	var article models.Article
	if err := c.Bind(&article); err != nil {
		return err
	}
	if err := article.Insert(); err != nil {
		return err
	}
	return c.String(http.StatusOK, "新增成功")
}

// Update 更新记录
func (a *ArticleController) Update(c *echo.Context) error {
	id := c.Param("id")
	if id == "" || !bson.IsObjectIdHex(id) {
		return errors.New("Id error")
	}
	var article models.Article
	if err := c.Bind(&article); err != nil {
		return err
	}
	article.Id = bson.ObjectIdHex(id)
	if err := article.Update(); err != nil {
		return err
	}
	return c.String(http.StatusOK, "更新成功")
}

// Delete 删除记录
func (a *ArticleController) Delete(c *echo.Context) error {
	id := c.Param("id")
	if id == "" || !bson.IsObjectIdHex(id) {
		return errors.New("Id error")
	}
	article := new(models.Article)
	article.Id = bson.ObjectIdHex(id)
	if err := article.Delete(); err != nil {
		return err
	}
	return c.String(http.StatusOK, "删除成功")
}
