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
	web  rpc  py 	pyc


*/
//trie树节点
//如对于节点 python，
//其pattern为 /lang/python
//curNodeVal为 python
//children为 [py, pyc]
//isWild为 false
type TrieNode struct {
	pattern    string      //节点的url
	curNodeVal string      //节点值
	children   []*TrieNode //子节点切片
}

//节点值是否含有通配符
func (n *TrieNode) hasWildcard() bool {
	//len(n.curNodeVal) == 0则是根节点，必定不会含通配符
	return len(n.curNodeVal) == 0 || n.curNodeVal[0] == '*' || n.curNodeVal[0] == ':'
}

//获取节点url
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

	return fmt.Sprintf("{pattern:%s, curNodeVal:%s, nchildren:%s, isWild:%t}\n",
		n.GetPattern(), n.curNodeVal, childrenStr.String(), n.hasWildcard())
}

//查找出第一个匹配成功的子节点
//查看某个节点是否有值为routVal的子节点，
//有则返回子节点，无则返回nil
func (n *TrieNode) matchChild(routVal string) *TrieNode {
	for _, child := range n.children {
		if child.curNodeVal == routVal || child.hasWildcard() {
			return child
		}
	}
	return nil
}
