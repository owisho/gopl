package squares

import (
	"fmt"
	"testing"
)

func TestSquares(t *testing.T) {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	f1 := squares()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
}
