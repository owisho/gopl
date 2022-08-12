package max_min

import (
	"bytes"
	"math"
)

func max(vals ...int) int {
	ret := -1
	for _, val := range vals {
		if val > ret {
			ret = val
		}
	}
	return ret
}

func min(vals ...int) int {
	ret := math.MaxInt
	for _, val := range vals {
		if val < ret {
			ret = val
		}
	}
	return ret
}

func multiJoin(sep string, vals ...string) string {
	var bf bytes.Buffer
	for i, val := range vals {
		if i > 0 {
			bf.WriteString(sep)
		}
		bf.WriteString(val)
	}
	return bf.String()
}
