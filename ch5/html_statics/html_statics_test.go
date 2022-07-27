package html_statics

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestCountWordsAndImages2(t *testing.T) {
	tmp := `<ol>
      <li>
		<p>
          (You may need to run the command as root or through <code>sudo</code>).
        </p>
		<p>
		  <strong>Do not</strong> untar the archive into an existing /usr/local/go tree. This is known to
		  produce broken Go installations.
		 </p>
      </li>
      <li>Confirm that the command prints the installed version of Go.</li>
    </ol>`
	bf := bytes.Buffer{}
	bf.WriteString(tmp)
	doc, err := html.Parse(&bf)
	if err != nil {
		fmt.Printf("parse err=%v\n", err)
		return
	}
	words, images := countWordsAndImages(doc)
	fmt.Printf("words=%d,images=%d\n", words, images)
}

func TestCountWordsAndImages(t *testing.T) {
	words, images, err := CountWordsAndImages("https://go.dev/doc/install")
	if err != nil {
		fmt.Printf("CountWordsAndImages err = %v", err)
		return
	}
	fmt.Printf("words=%d,images=%d", words, images)
}

func TestRune(t *testing.T) {
	s := "国"
	for i, c := range s {
		fmt.Printf("i=%d,b=%d\n", i, c)
	}
	fmt.Printf("bytes=%v", []byte(s))
}

func TestStringsMapping(t *testing.T) {
	s := strings.Map(func(r rune) rune {
		return r + 1
	}, "abc")
	fmt.Printf("s=%s\n", s)
}
