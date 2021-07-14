package main

import (
	"fmt"
	"net/url"
	"reflect"
)

func main() {
	m := url.Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")
	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"   (first value)
	fmt.Println(m["item"])     // "[1 2]" (direct map access)
	m = nil
	fmt.Println(reflect.TypeOf(m))
	fmt.Println(m.Get("item")) // ""
	vs := m["item"]
	fmt.Println(vs)
	// m.Add("item", "3") // panic: assignment to entry in nil map
}
