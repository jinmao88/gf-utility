package response

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// JsonRes 数据返回通用JSON数据结构
type Resp struct {
	g.Meta  `mime:"json" example:"string"`
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"result"`  // 返回数据(业务接口定义具体数据结构)
}

func (j *Resp) Format(code int, msg string, data ...interface{}) JsonRes {
	j.Code = code
	j.Message = msg
	j.Data = data
	return j
}

func (j *Resp) C(code int) JsonRes {
	j.Code = code
	return j
}

func (j *Resp) M(msg string) JsonRes {
	j.Message = msg
	return j
}

func (j *Resp) D(data interface{}) JsonRes {
	j.Data = data
	return j
}

func (j *Resp) JsonExit(r *ghttp.Request) {
	r.Response.WriteJsonExit(j)
}

func Code(code int) gcode.Code {
	return gcode.New(code, "", nil)
}

func ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}
	res := r.GetHandlerResponse()
	err := r.GetError()
	if err != nil {
		r.Response.WriteJson(Resp{
			Code:    gerror.Code(err).Code(),
			Message: err.Error(),
		})
		return
	}
	r.Response.WriteJson(res)

}

// Json 标准返回结果数据结构封装。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(Resp{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// JsonExit 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	Json(r, err, msg, data...)
	r.Exit()
}

// Resp 定义了一个响应接口，用于构建灵活的API响应
type JsonRes interface {
	// Format 用于格式化响应，包括状态码、消息和数据
	// 参数:
	//   code: 响应状态码
	//   msg: 响应消息
	//   data: 可变参数，用于传递响应数据
	// 返回值:
	//   Resp: 返回Resp接口，支持链式调用
	Format(code int, msg string, data ...interface{}) JsonRes

	// C 用于设置响应状态码
	// 参数:
	//   code: 响应状态码
	// 返回值:
	//   Resp: 返回Resp接口，支持链式调用
	C(code int) JsonRes

	// M 用于设置响应消息
	// 参数:
	//   msg: 响应消息
	// 返回值:
	//   Resp: 返回Resp接口，支持链式调用
	M(msg string) JsonRes

	// D 用于设置响应数据
	// 参数:
	//   data: 响应数据
	// 返回值:
	//   Resp: 返回Resp接口，支持链式调用
	D(data interface{}) JsonRes

	// JsonExit 用于将响应以JSON格式输出并结束请求处理
	// 参数:
	//   r: ghttp.Request指针，用于获取请求相关上下文
	// 返回值: 无
	// 注意: 该方法会直接输出JSON响应并结束请求处理流程
	JsonExit(r *ghttp.Request)
}
