package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Link is a parsed representation of HTML <a> tag
type Link struct {
	Href string
	Text string
}

// Find reads HTML from provided reader and returns found links
func Find(r io.Reader) ([]Link, error) {
	root, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var links []Link

	for node := root; node != nil; node = node.NextSibling {
		findLinks(node, &links)
	}

	return links, nil
}

func findLinks(n *html.Node, links *[]Link) {
	if isLink(n) {
		var sb strings.Builder
		text(n, &sb)
		link := Link{
			Href: href(n),
			Text: sb.String(),
		}
		if len(link.Text) > 0 {
			link.Text = link.Text[:len(link.Text)-1]
		}
		*links = append(*links, link)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findLinks(c, links)
	}
}

func isLink(n *html.Node) bool {
	return n.Type == html.ElementNode && n.DataAtom == atom.A
}

func href(n *html.Node) string {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

func text(n *html.Node, sb *strings.Builder) {
	if isText(n) {
		text := strings.TrimSpace(n.Data)
		(*sb).WriteString(text)
		if len(text) > 0 {
			(*sb).WriteRune(' ')
		}
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text(c, sb)
	}
}

func isText(n *html.Node) bool {
	return n.Type == html.TextNode
}
