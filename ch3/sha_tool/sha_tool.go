package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	sumType := flag.String("type", "sha256", "choose type of sum method")
	flag.Parse()
	args := os.Args
	flagArgs := flag.Args()
	fmt.Printf("args = %v, sumType = %s, flag args=%v\n", args[1:], *sumType, flagArgs)

	switch *sumType {
	case "sha256":
		arr := sha256.Sum256([]byte(flagArgs[0]))
		fmt.Println(arr)
	case "sha384":
		arr := sha512.Sum384([]byte(flagArgs[0]))
		fmt.Println(arr)
	case "sha512":
		arr := sha512.Sum512([]byte(flagArgs[0]))
		fmt.Println(arr)
	}
}
