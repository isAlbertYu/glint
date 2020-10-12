package trie

//trie树节点
type trieNode struct {
	//pattern   string      //到达当前节点的路由
	curRouVal string      //当前路由节点值
	children  []*trieNode //子路由节点
	isWild    bool        //是否模糊匹配，含通配符
}

//trie树
type Trie struct {
	root *trieNode
}

//查看某个节点是否有值为routVal的子节点，
//有则返回子节点，无则返回nil
func (n *trieNode) matchChild(routVal string) *trieNode {
	for _, child := range n.children {
		if child.curRouVal == routVal || child.isWild {
			return child
		}
	}
	return nil
}

//向trie树中插入一个路由
//parts为待插入路由的各层节点
//如，对于路由/lang/go/web , parts为[lang, go, web]
func (trie *Trie) insert(parts []string) {
	curNode := trie.root
	for i := 0; i < len(parts); i++ {
		next := curNode.matchChild(parts[i])//找出与之匹配的当前节点的子节点
		if next != nil {//若有与之匹配的子节点
			curNode = next
		} else {//若无与之匹配的子节点，则新建并插入新节点
			t := &trieNode{//新建节点
				curRouVal: parts[i],
				isWild:    parts[i][0] == '*' || parts[i][0] == ':',
			}
			curNode.children = append(curNode.children, t)//插入节点
		}
	}
}















