package tools

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	f1, _ := os.Open("/Users/wangyang/Desktop/cmp1")
	f2, _ := os.Open("/Users/wangyang/Desktop/cmp2")
	br1, _ := io.ReadAll(f1)
	br2, _ := io.ReadAll(f2)
	lines1 := strings.Split(string(br1), "\n")
	lines2 := strings.Split(string(br2), "\n")
	fMap1 := make(map[string]bool)
	fMap2 := make(map[string]bool)
	for _, line := range lines1 {
		arr := strings.Split(line, "$")
		fMap1[arr[0]+"_"+arr[1]] = true
	}
	for _, line := range lines2 {
		line = strings.ReplaceAll(line, "Davinci_DB_Import_20220912_", "")
		line = strings.ReplaceAll(line, "_af", "")
		fMap2[line] = true
	}
	for key, _ := range fMap2 {
		if _, ok := fMap1[key]; !ok {
			fmt.Println(key)
		}
	}
}
