package prettyprint

import (
	"bytes"
	"golang.org/x/net/html"
	//"io"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPrettyPrint(t *testing.T) {

	doc := getSampleHTMLDoc(t)
	//Now real error will start
	buf := bytes.NewBuffer([]byte{})
	PrettyPrint(buf, doc)

	var htmlBytes []byte
	buf.Write(htmlBytes)
	fmt.Println(buf.String())
	_, errParse := html.Parse(bytes.NewReader(htmlBytes))
	if errParse != nil {
		t.Fatalf("parse formed html: ", errParse)
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
