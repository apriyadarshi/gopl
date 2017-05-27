package xmltree

import (
	"encoding/xml"
	"fmt"
	"io"
)

type Node interface {
	String() string
}

type CharData string

func (c CharData) String() string {
	return string(c)
}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e Element) String() string {
	var s string
	s = fmt.Sprintf("<%s>\n", e.Type)
	for _, v := range e.Children {
		s += fmt.Sprintf("%s", v)
	}
	s += fmt.Sprintf("</%s>\n", e.Type)
	return s
}

func Parse(r io.Reader) (*Node, error) {
	dec := xml.NewDecoder(r)
	var stack []*Node
	var root Node
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("parse: decode token: ", err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			curr := Node(Element{Type: tok.Name, Attr: tok.Attr, Children: []Node{}})
			if len(stack) == 0 {
				root = curr
				stack = append(stack, &root)
			} else {
				//peek top element in stack and make it child of current one
				parent := (*(stack[len(stack)-1])).(Element)
				parent.Children = append(parent.Children, curr)
				stack = append(stack, &curr)
			}
		case xml.EndElement:
			//pop
			stack = stack[:len(stack)-1]
		case xml.CharData:
			curr := Node(CharData(tok))
			parent := (*(stack[len(stack)-1])).(Element)
			parent.Children = append(parent.Children, curr)
		}
	}
	return &root, nil
}
