package glint

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

//定义了一次http请求的上下文内容
//包含：
//		Req 来自客户端的http请求
//		Writer 服务端的响应输出流
//		Path 路由匹配
//		Method 请求方法
//		StatusCode 状态码
type Context struct {
	Req    *http.Request
	Writer http.ResponseWriter

	Path   string
	Method string

	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

//获取post表单中key对应的value
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//获取get请求的URL中key对应的value
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

//将格式化字符串写入响应输出流
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Context-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

//将JSON写入响应输出流
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Context-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	err := encoder.Encode(obj)
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

//将HTML写入响应输出流
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Context-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

//将字节数据写入响应输出流
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}
