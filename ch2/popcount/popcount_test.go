package popcount

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestPopCount(t *testing.T) {
	fmt.Println(pc)
	fmt.Println(256 >> 8)
}

func TestDiffCount(t *testing.T) {
	sum1 := sha256.Sum256([]byte("x"))
	sum2 := sha256.Sum256([]byte("x"))
	fmt.Println(DiffCount(sum2, sum1))
}
