package xmltree

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestParse(t *testing.T) {
	xml, _ := GetXML()

	fmt.Printf(xml)
}

func GetXML() (string, error) {
	resp, err := http.Get("http://www.w3.org/TR/2006/REC-xml11-20060816")
	if err != nil {
		return "", fmt.Errorf("fetch: %v\n", err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", fmt.Errorf("fetch: reading %s:%v\n", "http://www.w3.org/TR/2006/REC-xml11-20060816", err)
	}
	return fmt.Sprintf("%s", b), nil
}
