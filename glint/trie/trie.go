package trie

//trie树
type Trie struct {
	root *TrieNode
}

//新建一个trie树，根节点值为空串
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			pattern:    "",
			curNodeVal: "",
			children:   make([]*TrieNode, 0),
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
			}
			curNode.children = append(curNode.children, t) //插入节点
			curNode = t
		}
	}
}

//查找trie树中是否存在指定的路由
//存在则返回该路由节点
//parts为路由URL的各层节点
//如，对于URL /lang/go/web , parts为[lang, go, web]
func (trie *Trie) Search(parts []string) *TrieNode {
	curNode := trie.root
	if len(parts) == 0 {
		return curNode
	}
	for i := 0; i < len(parts); i++ {
		child := curNode.matchChild(parts[i]) //找出与之匹配的当前节点的子节点
		//没有可以匹配的子节点
		if child == nil {
			return nil
		}
		//已到达待匹配的url的最后一层 且 也到达了trie树的叶子节点, 则存在匹配
		if i == len(parts)-1 && len(child.children) == 0 {
			return child
		}

		//匹配到trie树中的*通配符
		if child.curNodeVal[0] == '*' {
			return child
		}
		curNode = child
	}
	return nil
}

//func (trie *Trie) Display() {
//	arr := make([]string, 0)
//	var str strings.Builder
//
//	str.WriteString("└─")
//	str.WriteString(trie.root.curNodeVal)
//	arr = append(arr, str.String())
//
//	for index, child := range trie.root.children {
//		child.show(&str, index == len(trie.root.children)-1, 1)
//	}
//
//	trie.root.show(&str)
//	fmt.Println(str.String())
//
//}
//
//func (node *TrieNode) show(res *strings.Builder, isLastOne bool, level int) {
//	if isLastOne {
//		(*res).WriteString("└─")
//		(*res).WriteString(node.curNodeVal)
//		(*res).WriteString("\n")
//		for i := 0; i < level; i++ {
//			(*res).WriteString("  ")
//		}
//	} else {
//		(*res).WriteString("├─")
//		(*res).WriteString(node.curNodeVal)
//		(*res).WriteString("\n")
//		for i := 0; i < level; i++ {
//			(*res).WriteString("  ")
//		}
//	}
//
//
//	if len(node.children) > 0 && !isLastOne {
//		for index, child := range node.children {
//			child.show(res, index == len(node.children)-1, 1)
//		}
//	}
//
//}
/*
└─
  └─lang
    ├─go
    │ ├─web
    │ └─rpc
    └─py
      ├─nlp
      └─cv

*/