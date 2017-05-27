package ch5

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"testing"
)

func TestElementsByTagName(t *testing.T) {
	testhtml := "<html><head></head><body><h1>Hi I am h1</h1><h2>Hi I am h2 </h2><h3>Hi! I am h3</h3><h4>h1</h4</body></html>"
	doc, err := html.Parse(bytes.NewBufferString(testhtml))
	if err != nil {
		t.Errorf("parse: ", err)
	}

	r := ElementsByTagName(doc, "h1", "h2", "body")
	if len(r) != 3 {
		t.Errorf("ElementsByTagName: Expected length of result: %d Got: %d", 3, len(r))
	}
	//fmt.Printf("r was of len %d", len(r))
	for _, v := range r {
		fmt.Println(v.Data)
	}

}
