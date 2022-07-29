package topsort

import (
	"fmt"
	"testing"
)

func TestTopsort(t *testing.T) {
	for i, course := range topsort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func TestTopsort1(t *testing.T) {
	for i, course := range topsort1(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
