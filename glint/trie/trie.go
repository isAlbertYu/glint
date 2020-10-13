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
		//已到达待匹配的url的最后一层 且 也到达了trie树的叶子节点
		if i == len(parts)-1 && len(child.children)==0{
			return child
		}

		if child.curNodeVal[0] == '*' {
			return child
		}
		curNode = child
	}
	return nil
}
