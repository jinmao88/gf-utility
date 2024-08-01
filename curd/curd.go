package curd

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type Curd interface {
	// SetCtx 设置Ctx
	SetCtx(context.Context)
	// List 列表
	List() (*List, error)
	// Add 新增
	Add() error
	// Edit 编辑
	Edit() error
	// Del 删
	Del() error
	// Tree 返回树结构
	Tree() (g.Map, error)
	// Options 返回options
	Options() ([]Option, error)
}

type Pagination struct {
	Page     int    `p:"page"`
	PageSize int    `p:"page_size"`
	Order    string `p:"order"`
}

type List struct {
	Items    interface{} `json:"items"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type Context struct {
	Pagination
	Operator Operator `p:"operator"`
}

type Operator struct {
	User int `p:"uid"`
	Role int `p:"role"`
}

func (ctx *Context) ListResult(m *gdb.Model) (*List, error) {
	count, err := m.Count()
	if err != nil {
		return nil, err
	}
	all, err := m.Page(ctx.Pagination.Page, ctx.Pagination.PageSize).All()
	if err != nil {
		return nil, err
	}
	return &List{
		Items:    all,
		Total:    count,
		Page:     ctx.Pagination.Page,
		PageSize: ctx.Pagination.PageSize,
	}, nil
}
