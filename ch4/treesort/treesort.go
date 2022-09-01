package treesort

import (
	"bytes"
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	values = appendValues(values[:0], root)
}

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

func (t *tree) String() string {
	var buf bytes.Buffer
	if t.left != nil {
		buf.WriteString(t.left.String() + " ")
	}
	buf.WriteString(fmt.Sprintf("%d ", t.value))
	if t.right != nil {
		buf.WriteString(t.right.String() + " ")
	}
	s := buf.String()
	if strings.HasSuffix(s, " ") {
		s = s[0 : len(s)-1]
	}
	return s
}
