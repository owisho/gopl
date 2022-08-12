package errorf

import "testing"

func TestErrorf(t *testing.T) {
	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name)
}
