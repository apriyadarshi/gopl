//Prints the text of selected xml elements
//./xmlselect div div h2 -> will print all headings that are a child of a div of a div
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

var count int32

func main() {
	//Process selector
	selector := strings.Split(os.Args[1], " ")
	fmt.Println(selector)
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement //stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
			if len(selector) == 1 {
				if matchIDSelector(tok, selector[0]) {
					fmt.Printf("1 id match for type %s: %s\n", tok.Name.Local, tok)
				} else {
					if matchClassSelector(tok, selector[0]) {
						count++
						fmt.Printf("%d match for class selector %s: %s\n", count, selector[0], tok.Name.Local)
					}
				}
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(selector) > 1 {
				if containsAll(stack, selector) {
					for _, v := range stack {
						fmt.Printf("%s ", v.Name.Local)
					}
					fmt.Printf(": %s\n", tok)
				}
			}
		}
	}
}

func matchIDSelector(x xml.StartElement, selector string) bool {
	if strings.Contains(selector, "#") {
		id := strings.TrimPrefix(selector, "#")
		if hasAttr(x, "id", id) {
			return true
		}
	}
	return false
}
func matchClassSelector(x xml.StartElement, selector string) bool {
	if strings.Contains(selector, ".") {

		selectors := strings.Split(selector, ".")
		elem, class := "", ""
		if len(selectors) == 2 {
			elem, class = selectors[0], selectors[1]
		} else {
			class = selectors[1]
		}
		if (elem == "" || elem == x.Name.Local) && hasAttr(x, "class", class) {
			return true
		}
	}
	return false
}

func hasAttr(x xml.StartElement, name, value string) bool {
	for _, attr := range x.Attr {
		if attr.Name.Local == name && attr.Value == value {
			return true
		}
	}
	return false
}

//whether x contains all of y in specified order
func containsAll(x []xml.StartElement, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if matchIDSelector(x[0], y[0]) {
			y = y[1:]
		} else if matchClassSelector(x[0], y[0]) {
			y = y[1:]
		} else {
			if x[0].Name.Local == y[0] {
				y = y[1:]
			}
		}
		x = x[1:]
	}
	return false
}
