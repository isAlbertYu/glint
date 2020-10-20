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

func TestShow(t *testing.T) {
	tireObj := NewTrie()
	tireObj.Insert([]string{"lang", "go", "web"})
	tireObj.Insert([]string{"lang", "go", "rpc"})
	tireObj.Insert([]string{"lang", "py", "nlp"})
	tireObj.Insert([]string{"lang", "py", "cv"})
	tireObj.Display()
}