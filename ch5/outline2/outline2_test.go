package outline2

import (
	"fmt"
	"net/http"
	"testing"
)
import "golang.org/x/net/html"

func TestForEachNode(t *testing.T) {
	resp, err := http.Get("https://downloader-344303.web.app/")
	if err != nil {
		fmt.Printf("http error=%v", err)
		return
	}
	node, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("html parse error=%v", err)
		return
	}
	forEachNode(node, startElement, endElement)
}
