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
