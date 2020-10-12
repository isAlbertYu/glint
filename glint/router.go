package glint

import (
	"go-glint/glint/trie"
	"net/http"
	"strings"
)

//定义了基于map实现的路由表，可以高效查询
//请求方法-匹配模式 : 处理函数
//GET与POST方法将分别对应两个trie树
//roots的key为GET与POST，值为对应的trie树
type router struct {
	tries    map[string]*trie.Trie
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		tries:    make(map[string]*trie.Trie),
		handlers: make(map[string]HandlerFunc),
	}
}

//将路由URL解析为各节点值组成的切片
//如将路由URL lang/go/web 变为 [lang, go, web]
func parsePattern(url string) []string {
	arr := strings.Split(url, "/")
	parts := make([]string, 0)
	for _, value := range arr {
		if value != "" {
			parts = append(parts, value)
			if value[0] == '*' {
				break
			}
		}
	}
	return parts
}

//添加路由映射
func (r *router) addRoute(method string, url string, handler HandlerFunc) {
	parts := parsePattern(url)
	key := method + "-" + url
	if _, ok := r.tries[method]; !ok {
		r.tries[method] = trie.NewTrie()
	}
	r.tries[method].Insert(parts)
	r.handlers[key] = handler
}

//查找路由映射
func (r *router)getRoute(method string, url string) {
	trie, ok := r.tries[method]
	if !ok {//不存在该请求方法对应的trie树
		return
	}

	routeMap := make(map[string]string)

	urlParts := parsePattern(url)
	node := trie.Search(urlParts)

	if node != nil {
		parts := parsePattern(node.GetPattern())
		for index, part := range parts {
			if part[0] == ':' {
				routeMap[part[1:]] = urlParts[index]
			}
		}
	}

}

//处理路由
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handlerFunc, ok := r.handlers[key]; ok {
		handlerFunc(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
