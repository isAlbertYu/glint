package glint

//所有RouterGroup持有一个共同的单例engine
type RouterGroup struct {
	prefix      string        //该分组所拥有的共同前缀
	middlewares []HandlerFunc //该分组的中间件函数
	parent      *RouterGroup  //该分组的父类
	engine      *engine
}

//创建一个顶级父类RouterGroup
func NewGroup(prefix string) *RouterGroup {
	return &RouterGroup{
		prefix: prefix,
		parent: nil,
		engine: GetSingleEngine(),
	}
}

//在父类RouterGroup father下创建一个子类RouterGroup
func (father *RouterGroup) NewSubGroup(subPrefix string) *RouterGroup {
	return &RouterGroup{
		prefix: father.prefix + subPrefix,
		parent: father,
		engine: GetSingleEngine(),
	}
}

//get请求的路由与处理函数
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	pattern = group.prefix + pattern
	group.engine.addRoute("GET", pattern, handler)
}

//post请求的路由与处理函数
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	pattern = group.prefix + pattern
	group.engine.addRoute("POST", pattern, handler)
}
