package trie

import (
	"fmt"
	"strings"
)

/*
			lang
			  |
		   -------
		  |      |
		 go    python
		 |        |
	  -----     ------
	 |    |    |     |
	web  rpd  py 	pyc


*/
//trie树节点
//如对于节点 python，
//其pattern为 /lang/python
//curNodeVal为 python
//children为 [py, pyc]
//isWild为 false
type TrieNode struct {
	pattern    string      //当前节点的url
	curNodeVal string      //当前节点值
	children   []*TrieNode //子路由节点
	isWild     bool        //是否模糊匹配，含通配符
}

func (n *TrieNode) GetPattern() string {
	return n.pattern
}

func (n *TrieNode) String() string {
	var childrenStr strings.Builder
	childrenStr.WriteString("[")
	for index, child := range n.children {
		childrenStr.WriteString(child.curNodeVal)
		if index < len(n.children)-1 {
			childrenStr.WriteString(",")
		}
	}
	childrenStr.WriteString("]")

	return fmt.Sprintf("pattern: %s\ncurNodeVal: %s\nchildren: %s\nisWild: %t\n",
		n.pattern, n.curNodeVal, childrenStr.String(), n.isWild)
}

//查找出第一个匹配成功的子节点
//查看某个节点是否有值为routVal的子节点，
//有则返回子节点，无则返回nil
func (n *TrieNode) matchChild(routVal string) *TrieNode {
	for _, child := range n.children {
		if child.curNodeVal == routVal || child.isWild {
			return child
		}
	}
	return nil
}

//trie树
type Trie struct {
	root *TrieNode
}

//新建一个trie树，根节点值为ROOT，根节点不参与路由匹配
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			pattern:    "/ROOT",
			curNodeVal: "ROOT",
			children:   make([]*TrieNode, 0),
			isWild:     false,
		},
	}
}

//向trie树中插入一个路由
//parts为路由URL的各层节点
//如，对于URL /lang/go/web , parts为[lang, go, web]
func (trie *Trie) Insert(parts []string) {
	curNode := trie.root
	for i := 0; i < len(parts); i++ {
		child := curNode.matchChild(parts[i]) //找出与之匹配的当前节点的子节点
		if child != nil { //若有与之匹配的子节点
			curNode = child
		} else { //若无与之匹配的子节点，则新建并插入新节点
			t := &TrieNode{ //新建节点
				pattern:    curNode.pattern + "/" + parts[i],
				curNodeVal: parts[i],
				isWild:     parts[i][0] == '*' || parts[i][0] == ':',
			}
			curNode.children = append(curNode.children, t) //插入节点
		}
	}
}

//查找trie树中是否存在指定的路由
//存在则返回该路由节点
//parts为路由URL的各层节点
//如，对于URL /lang/go/web , parts为[lang, go, web]
func (trie *Trie) Search(parts []string) *TrieNode {
	curNode := trie.root
	for i := 0; i < len(parts); i++ {
		child := curNode.matchChild(parts[i]) //找出与之匹配的当前节点的子节点
		if i == len(parts)-1 && child != nil {
			return child
		}

		if child == nil { //没有与之匹配的子节点
			return nil
		}
		curNode = child
	}
	return nil
}
