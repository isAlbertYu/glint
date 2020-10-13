package trie

import (
	"fmt"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	tireObj := NewTrie()
	tireObj.Insert([]string{"lang", "go", "web"})
	fmt.Println(tireObj.root.String())
	//fmt.Println(tireObj.root.children)
	var a []int
	fmt.Println(a==nil)
	fmt.Println(len(a))
}