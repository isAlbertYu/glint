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
//handlers：存储路由映射
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
//如，对于路由URL /lang/go/web 变为 [lang, go, web]
//对于路由URL / 变为 []
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
//向trie树中插入url
//并向handlers中添加路由映射
func (r *router) addRoute(method string, url string, handler HandlerFunc) {
	urlParts := parsePattern(url)
	var key string
	if len(urlParts) == 0 {
		key = method + "-" + ""
	} else {
		key = method + "-" + url
	}
	if _, ok := r.tries[method]; !ok {
		r.tries[method] = trie.NewTrie()
	}
	r.tries[method].Insert(urlParts)
	r.handlers[key] = handler
}

//根据url在trie树中查找路由映射
//返回trie树中对应的url节点 与 以map方式存储trie树中通配符所匹配的值
func (r *router) getRoute(method string, url string) (*trie.TrieNode, map[string]string) {
	trie, ok := r.tries[method]
	if !ok { //不存在该请求方法对应的trie树
		return nil, nil
	}

	paramMap := make(map[string]string)

	urlParts := parsePattern(url) // /hello/:name => [hello, :name]
	node := trie.Search(urlParts) //查找对应于trie树中的节点 [hello, :name] => web节点

	if node != nil {
		//node.GetPattern() : trie树中的路由节点，可能带有通配符/hello/:name
		parts := parsePattern(node.GetPattern()) //parts = [hello, :name]
		for index, part := range parts { // [hello, :name]
			if part[0] == ':' {
				paramMap[part[1:]] = urlParts[index]
			}
			if part[0] == '*' {
				paramMap[part[1:]] = strings.Join(urlParts[index:], "/")
				break
			}
		}
		return node, paramMap
	}
	return nil, nil
}

//执行路由映射的处理函数
func (r *router) handle(c *Context) {
	node, paramMap := r.getRoute(c.Method, c.Path)
	if node != nil {
		c.RouteParamMap = paramMap
		key := c.Method + "-" + node.GetPattern()
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
