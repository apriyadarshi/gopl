package prettyprint

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"strings"
)

/*func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Errorf("error while parsing html: ", err)
	}
	visit(doc, startElement, endElement)
}*/
/*
func PrettyPrint(n *html.Node) {
	visit(doc, startEleme)
}*/

func PrettyPrint(w io.Writer, n *html.Node) {
	var depth int

	startElement := func(w io.Writer, n *html.Node) {
		switch n.Type {
		case html.ElementNode:
			fmt.Fprintf(w, "%*s<%s", depth*2, "", n.Data)
			for _, attr := range n.Attr {
				fmt.Fprintf(w, " %s=%q", attr.Key, strings.Replace(attr.Val, "\"", "&quot;", -1))
				//Quotes are html escaped toa avoid injection issues
			}
			if n.FirstChild == nil {
				fmt.Fprint(w, "/>\n")
			} else {
				fmt.Fprintf(w, ">\n")
				depth++
			}
		//Following cases don't increase depth as they are leaf nodes
		case html.CommentNode:
			fmt.Fprintf(w, "%*s<!--%s-->\n", depth*2, "", n.Data)
		case html.TextNode:
			if strings.TrimSpace(n.Data) != "" {
				fmt.Fprintf(w, "%*s%s\n", depth*2, "", n.Data)
			}
		}
	}

	endElement := func(w io.Writer, n *html.Node) {
		switch n.Type {
		case html.ElementNode:
			if n.FirstChild != nil {
				depth--
				fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
			}
		case html.CommentNode:
			//depth--
		case html.TextNode:
			//depth--
		}

	}

	visit(w, n, startElement, endElement)
}

func visit(w io.Writer, n *html.Node, pre, post func(w io.Writer, n *html.Node)) {

	if pre != nil {
		pre(w, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(w, c, pre, post)
	}

	if post != nil {
		post(w, n)
	}
}
