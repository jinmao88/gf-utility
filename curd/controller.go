package curd

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/jinmao88/gf-utility/response"
)

type CurdReq struct {
	g.Meta    `path:"/curd" method:"post" summary:"Curd请求" tags:"Curd"`
	Interface string `p:"i" v:"required"`
	Action    string `p:"a" v:"required"`
}

func Controller(ctx context.Context, req *CurdReq, check func(i string) (Curd, error)) (res response.JsonRes, err error) {
	r := new(response.Resp)
	cu, err := check(req.Interface)
	if err != nil {
		return nil, err
	}
	if err = g.RequestFromCtx(ctx).Parse(cu); err != nil {
		return nil, gerror.NewCode(response.Code(2), err.Error())
	}
	cu.SetCtx(ctx)
	switch req.Action {
	case "list":
		r.Data, err = cu.List()
	case "tree":
		r.Data, err = cu.Tree()
	case "options":
		r.Data, err = cu.Options()
	case "add":
		err = cu.Add()
		r.Message = "新增成功"
	case "edit":
		err = cu.Edit()
		r.Message = "修改成功"
	case "del":
		err = cu.Del()
		r.Message = "删除成功"
	default:
		err = gerror.NewCode(response.Code(3), "接口参数错误")
	}
	return r, nil
}
