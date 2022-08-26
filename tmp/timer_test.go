package tmp

import (
	"fmt"
	"testing"
	"time"
)

func TestAftTrFunc(t *testing.T) {
	after()
	for {
		time.Sleep(1 * time.Second)
	}
}

func after() {
	time.AfterFunc(10*time.Second, func() {
		fmt.Println("after 10s")
	})
}
