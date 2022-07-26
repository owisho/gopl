package text

import (
	"bytes"
	"fmt"
	"testing"
)
import "golang.org/x/net/html"

func TestPrintAllTextContent(t *testing.T) {
	PrintAllTextContent("https://zhuanlan.zhihu.com/p/411315625")
}

func TestGetNodeTextContent(t *testing.T) {
	h := `<dev><p>aaaa</p><p>bbbbb</p></dev>`
	bf := bytes.Buffer{}
	bf.WriteString(h)
	doc, err := html.Parse(&bf)
	if err != nil {
		t.Errorf("err = %v", err)
		return
	}
	contents := make([]string, 0)
	contents = GetNodeTextContent(contents, doc)
	for _, line := range contents {
		fmt.Println(line)
	}
}
