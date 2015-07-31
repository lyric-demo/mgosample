package controllers

import (
	"strconv"

	"github.com/labstack/echo"
)

const (
	// PAGESIZE 分页大小
	PAGESIZE = 10
)

// BaseController 提供控制器的基础结构
type BaseController struct {
	Context *echo.Context
}

// GetPageParams 获取分页参数
func (b *BaseController) GetPageParams() (skip, limit int) {
	page := b.Context.Query("page")
	if page == "" {
		page = "0"
	}
	iPage, _ := strconv.Atoi(page)
	if iPage <= 0 {
		iPage = 1
	}

	pageSize := b.Context.Query("pagesize")
	if pageSize == "" {
		pageSize = "0"
	}
	iPageSize, _ := strconv.Atoi(pageSize)
	if iPageSize <= 0 {
		iPageSize = PAGESIZE
	}

	skip = (iPage - 1) * iPageSize
	limit = iPageSize

	return
}
