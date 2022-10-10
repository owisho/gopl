package url

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	n := m
	n["p"] = []string{"1", "2", "3"}
	fmt.Println(m)

	n = nil
	fmt.Println(m.Get("item"))
	m.Add("item", "3")
}
