//findlinks1 prints the links in an HTML document read from standard input
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin) //doc is a pointer to html node
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visitNode(nil, doc) {
		//for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func appendLink(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "link" || n.Data == "script") {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				links = append(links, a.Val)
				break
			}
		}
	}
	return links
}

/*Traverses HTML node tree;
 *For each node if its an anchor element, extracts the href attr
 *Appends the extracted link to the links slice*/
func visit(links []string, n *html.Node) []string {

	if n != nil {

		links = appendLink(links, n)

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			links = visit(links, c)
		}
	}

	return links
}

/*
	Full recursive Version
 	Programming for recursion:
 	1. Write the actual calls that need to be made
 	2. Check how the same sequence of calls can be attained via recursion.
 	Ex: Here, we need the calls visitNode(x.FirstChild), visitNode(x.FirstChild.NextSibling), vN(x.FC.NS.NS)..vN(x.lastChild)
 	This sequence can easily be attained via a func b(x) { vN(x); if x.NS != nil { b(x.NS) }; }
 	Thus vN(x) { doWhateverWith(x); if x.FC !=nil { b(x.FC) } }
*/
func visitNode(links []string, n *html.Node) []string {

	if n != nil {

		links = appendLink(links, n)
		if n.FirstChild != nil {
			links = visitBreadth(links, n.FirstChild)
		}
	}

	return links
}

func visitBreadth(links []string, n *html.Node) []string {

	links = visitNode(links, n)
	if n.NextSibling != nil {
		links = visitBreadth(links, n.NextSibling)
	}
	return links
}
