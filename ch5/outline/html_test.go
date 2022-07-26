package outline

import (
	"fmt"
	"net/http"
	"testing"
)
import "golang.org/x/net/html"

func TestHtmlParse(t *testing.T) {
	resp, err := http.Get("https://go.dev/")
	if err != nil {
		t.Errorf("http get error %v", err)
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		t.Errorf("html parse error %v", err)
		return
	}
	doc = doc.FirstChild
	doc = doc.NextSibling.FirstChild
	fmt.Printf("node type=%v,data=%s,namespace=%s,attrs=%v", doc.Type, doc.Data, doc.Namespace, doc.Attr)
	fmt.Printf("node atom=%v", doc.DataAtom)
}
