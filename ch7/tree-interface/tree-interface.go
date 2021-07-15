package main

import (
	"bytes"
	"fmt"
)

func main() {
	t := new(tree)
	val := []int{1, 23, 4}
	add(t, 2)
	add(t, 3)
	add(t, -1)
	values := appendValues(val, t)
	fmt.Println(values)
	Sort(values)
	fmt.Println(values)
	newValues := t.String()
	fmt.Println(newValues)

}

type tree struct {
	value       int
	left, right *tree
}

// Exercise 7.3
func (t *tree) String() string {
	order := make([]int, 0)
	order = appendValues(order, t)

	if len(order) == 0 {
		return "[]"
	}
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%v", order)
	return b.String()
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
