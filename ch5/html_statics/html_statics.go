package html_statics

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML :%s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		return 0, 1
	}
	if n.Type == html.TextNode {
		br := []byte(n.Data)
		for start, step := 0, 0; start < len(br); start += step {
			var err error
			step, _, err = bufio.ScanWords(br[start:], true)
			if err != nil {
				return -1, -1
			}
			words++
		}
		return words, 0
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "style" || c.Data == "script" {
			continue
		}
		cwords, cimages := countWordsAndImages(c)
		words += cwords
		images += cimages
	}
	return words, images
}
