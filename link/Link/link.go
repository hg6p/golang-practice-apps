package link

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func TraverseHtml(file *os.File) {
	doc, err := html.Parse(file)
	if err != nil {
		fmt.Println("Error parsing html ", err)
		return
	}
	str := ""

	var arr []Link

	f(doc, str, &arr)
	fmt.Printf("%+v\n", arr)
}

func iterateSiblings(n *html.Node, str *string) string {
	if n.Type == html.TextNode {
		*str += n.Data
	}
	if n.FirstChild != nil {

		return iterateSiblings(n.FirstChild, str)
	}
	if n.NextSibling == nil {
		return ""
	}

	return iterateSiblings(n.NextSibling, str)
}

func f(n *html.Node, str string, arr *[]Link) string {
	if n.FirstChild == nil {
		return n.Data
	}
	if n.Type == html.ElementNode && n.Data == "a" {

		str = ""
		iterateSiblings(n.FirstChild, &str)
		*arr = append(*arr, Link{Href: n.Attr[0].Val, Text: str})
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {

		f(c, str, arr)
	}

	return n.Data
}
