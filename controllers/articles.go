package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"

	"github.com/LyricTian/mgosample/models"
)

// RegisterArticle 注册文章管理路由
func RegisterArticle(g *echo.Group) {
	article := new(ArticleController)
	group := g.Group("/articles")
	{
		group.Get("", article.GetData)
		group.Get("/:id", article.GetSingleData)
		group.Post("", article.Insert)
		group.Put("/:id", article.Update)
		group.Delete("/:id", article.Delete)
	}
}

// ArticleController 文章管理控制器
type ArticleController struct {
	BaseController
}

// GetResModel 获取响应模型
func (a *ArticleController) GetResModel(data interface{}) map[string]interface{} {
	return map[string]interface{}{"articles": data}
}

// GetData 获取所有数据
func (a *ArticleController) GetData(c *echo.Context) error {
	a.Context = c
	article := new(models.Article)
	var data []models.Article
	skip, limit := a.GetPageParams()
	if err := article.GetPageData(skip, limit, &data); err != nil {
		return err
	}
	if data == nil {
		data = make([]models.Article, 0)
	}
	return c.JSON(http.StatusOK, a.GetResModel(data))
}

// GetSingleData 获取单条记录
func (a *ArticleController) GetSingleData(c *echo.Context) error {
	id := c.Param("id")
	if id == "" || !bson.IsObjectIdHex(id) {
		return errors.New("ID error")
	}
	article := new(models.Article)
	article.ID = bson.ObjectIdHex(id)
	data, err := article.GetSingleData()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, a.GetResModel(data))
}

// Insert 插入记录
func (a *ArticleController) Insert(c *echo.Context) error {
	var reqData map[string]*models.Article
	if err := c.Bind(&reqData); err != nil {
		return err
	}
	if err := reqData["article"].Insert(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, reqData)
}

// Update 更新记录
func (a *ArticleController) Update(c *echo.Context) error {
	id := c.Param("id")
	if id == "" || !bson.IsObjectIdHex(id) {
		return errors.New("ID error")
	}
	var reqData map[string]*models.Article
	if err := c.Bind(&reqData); err != nil {
		return err
	}
	article := reqData["article"]
	article.ID = bson.ObjectIdHex(id)
	fmt.Println("===> Request data:", article)
	if err := article.Update(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, reqData)
}

// Delete 删除记录
func (a *ArticleController) Delete(c *echo.Context) error {
	id := c.Param("id")
	if id == "" || !bson.IsObjectIdHex(id) {
		return errors.New("ID error")
	}
	article := new(models.Article)
	article.ID = bson.ObjectIdHex(id)
	if err := article.Delete(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}
