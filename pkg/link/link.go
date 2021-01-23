package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Link is a parsed representation of HTML <a> tag
type Link struct {
	Href string `json:"href,omitempty"`
	Text string `json:"text,omitempty"`
}

// Find reads HTML from provided reader and returns found links
func Find(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	return findLinks(doc), nil
}

func findLinks(n *html.Node) []Link {
	if isLink(n) {
		link := buildLink(n)
		return []Link{link}
	}

	var links []Link

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, findLinks(c)...)
	}

	return links
}

func isLink(n *html.Node) bool {
	return n.Type == html.ElementNode && n.DataAtom == atom.A
}

func buildLink(n *html.Node) Link {
	var sb strings.Builder
	text(n, &sb)
	text := sb.String()
	if len(text) > 0 {
		text = text[:len(text)-1]
	}

	return Link{
		Href: href(n),
		Text: text,
	}
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
		sb.WriteString(text)
		if len(text) > 0 {
			sb.WriteRune(' ')
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
