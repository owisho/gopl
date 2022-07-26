package ch7

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	w := WordCounter(0)
	n, err := fmt.Fprintf(&w, "this is a test by %s", "owisho")
	if err != nil {
		t.Errorf("err=%v", err)
		return
	}
	fmt.Printf("count=%d,word count=%d\v", n, w)
}

func TestLineCounter(t *testing.T) {
	l := LineCounter(0)
	n, err := fmt.Fprintf(&l, "this is a test by %s\n hello interface\n known something new", "owisho")
	if err != nil {
		t.Errorf("err=%v", err)
		return
	}
	fmt.Printf("count=%d,line count=%d\v", n, l)
}
