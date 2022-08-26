package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestGetFeatures(t *testing.T) {
	f, err := os.Open("/Users/wangyang/Desktop/opera/demo_data/features.txt")
	if err != nil {
		t.Errorf("open file error: %v", err)
	}

	s := bufio.NewScanner(f)
	for i := 0; s.Scan(); i++ {
		line := s.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			continue
		}
		arr := strings.Split(line, "=")
		featureVal := arr[1][1 : len(arr[1])-1]
		feature, version := strings.Split(featureVal, "$")[0], strings.Split(featureVal, "$")[1]
		fmt.Printf("A:%s:%s:userId\n", feature, version)
	}
}
