package tmp

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		fmt.Println(t)
	}
}
