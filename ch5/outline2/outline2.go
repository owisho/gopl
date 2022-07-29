package outline2

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
)

//forEachNode 遍历所有节点
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s %s/>\n", depth*2, "", n.Data, attrs(n))
		} else {
			fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data, attrs(n))
		}
		depth++
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	} else if n.Type == html.TextNode {
		fmt.Printf("%*s%s", depth*2, "", n.Data)
		depth++
	}
}

func attrs(n *html.Node) string {
	var bf bytes.Buffer
	for _, attr := range n.Attr {
		bf.WriteString(fmt.Sprintf("%s='%s' ", attr.Key, attr.Val))
	}
	return bf.String()
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
