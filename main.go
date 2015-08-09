package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/LyricTian/mgosample/conf"
	"github.com/LyricTian/mgosample/controllers"
	"github.com/LyricTian/mgosample/modules/util"
)

func main() {
	e := echo.New()

	e.SetDebug(conf.APPCONFIG.Debug)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// 解析绑定请求
	e.SetBinder(func(r *http.Request, v interface{}) error {
		defer r.Body.Close()
		return json.NewDecoder(r.Body).Decode(v)
	})

	// 错误处理
	e.SetHTTPErrorHandler(func(err error, c *echo.Context) {
		code := http.StatusInternalServerError
		msg := http.StatusText(code)
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code()
			msg = he.Error()
		}
		if conf.APPCONFIG.Debug {
			msg = err.Error()
		}
		c.JSON(code, util.Error(errors.New(msg)))
	})

	api := e.Group("/api").Group("/v1")

	// 注册文章管理路由
	controllers.RegisterArticle(api)

	log.Println(fmt.Sprintf("[%s] is running at %d port.", conf.APPCONFIG.Name, conf.APPCONFIG.Port))
	e.Run(fmt.Sprintf(":%d", conf.APPCONFIG.Port))
}
