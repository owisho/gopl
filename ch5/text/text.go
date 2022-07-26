package text

import (
	"fmt"
	"net/http"
)
import "golang.org/x/net/html"

func PrintAllTextContent(url string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	doc, err := html.Parse(resp.Body)
	contents := make([]string, 0)
	contents = GetNodeTextContent(contents, doc)
	//for line := range contents {
	//	fmt.Println(contents[line])
	//}
	return nil
}

func GetNodeTextContent(contents []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		contents = append(contents, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			if c.Type == html.TextNode {
				fmt.Println(c.Data)
			}
			continue
		}
		contents = append(contents, GetNodeTextContent(nil, c)...)
	}
	return contents
}
