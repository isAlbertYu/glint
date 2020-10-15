package glint

import (
	"log"
	"net/http"
	"sync"
)

type HandlerFunc func(c *Context)

//http服务端的处理引擎，内置一个路由表router
//由router来管理路由映射
//全局只有一个engine
type engine struct {
	router *router
}

var (
	once sync.Once
	singleEngine *engine
)

//获取单例engine指针，指向同一个engine实体
func GetSingleEngine() *engine {
	once.Do(func() {
		singleEngine = &engine{
			router: newRouter(),
		}
	})
	return singleEngine
}

func (engine *engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

//get请求的路由与处理函数
func (engine *engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

//post请求的路由与处理函数
func (engine *engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *engine) Run(addr string) error {
	log.Println("golint start running...")
	return http.ListenAndServe(addr, engine)
}

//使得Engine实现Handler接口，可以注册到http服务端
func (engine *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
