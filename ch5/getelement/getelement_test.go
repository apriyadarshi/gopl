package getelement

import (
	"bytes"
	//"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestElementById(t *testing.T) {
	doc := getSampleHTMLDoc(t)

	node := ElementById(doc, "playground")
	var class string
	if node == nil {
		t.Error("ElementById: element not found")
	} else {
		for _, v := range node.Attr {
			if v.Key == "class" && v.Val == "play" {
				class = v.Val
			}
		}
		if class != "play" {
			t.Error("ElementById: incorrect element not found")
		}
	}

}

func getSampleHTMLDoc(t *testing.T) *html.Node {
	url := "https://golang.org"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("fetch: getting golang.org: ", err)
	}
	b, err2 := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err2 != nil {
		t.Errorf("fetch: reading %s:%v\n", url, err)
	}
	//s := fmt.Sprintf("%s", b)
	doc, err3 := html.Parse(bytes.NewReader(b))
	if err3 != nil {
		t.Errorf("parse fetched html: ", err3)
	}
	return doc
}
