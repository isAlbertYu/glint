package glint

import (
	"fmt"
	"reflect"
	"testing"
)

var r *router

func init() {
	r = newRouter()
	//r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/coder/:name", nil)
	r.addRoute("GET", "/student/tom", nil)
	r.addRoute("GET", "/student/tony", nil)
	//r.addRoute("GET", "/hi/:name", nil)
	//r.addRoute("GET", "/assets/*filepath", nil)
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/:word/*name/hhh"), []string{"p", ":word", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	n, ps := r.getRoute("GET", "/coder/albert")
	stu, _ := r.getRoute("GET", "/student")

	if stu == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}
	fmt.Println("n.GetPattern()= ", n.GetPattern())
	if n.GetPattern() != "/coder/:name" {
		t.Fatal("should match /coder/:name")
	}
	fmt.Println("ps[name]= ", ps["name"])

	if ps["name"] != "albert" {
		t.Fatal("name should be equal to 'albert'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.GetPattern(), ps["name"])

}

func TestP(t *testing.T) {
	a := parsePattern("/")
	fmt.Println(len(a))
}