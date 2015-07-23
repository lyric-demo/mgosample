package main

import (
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

	e.SetDebug(conf.APP_CONFIG.Debug)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// 错误处理
	e.SetHTTPErrorHandler(func(err error, c *echo.Context) {
		code := http.StatusInternalServerError
		msg := http.StatusText(code)
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code()
			msg = he.Error()
		}
		if conf.APP_CONFIG.Debug {
			msg = err.Error()
		}
		c.JSON(code, util.Error(errors.New(msg)))
	})

	api := e.Group("/api").Group("/v1")

	// 注册文章管理路由
	controllers.RegisterArticle(api)

	log.Println(fmt.Sprintf("[%s] is running at %d port.", conf.APP_CONFIG.Name, conf.APP_CONFIG.Port))
	e.Run(fmt.Sprintf(":%d", conf.APP_CONFIG.Port))
}
