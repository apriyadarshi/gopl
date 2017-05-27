package getelement

import (
	//"fmt"
	"golang.org/x/net/html"
)

func ElementById(doc *html.Node, id string) *html.Node {
	var result = html.Node{}
	result.Data = "INVALID"
	recurse(&result, doc, id, checkId)
	return &result
}

func recurse(out, in *html.Node, id string, pre func(out, in *html.Node, id string) bool) {
	//fmt.Printf("Value of out: %s\n", out.Data)
	var ok bool
	if pre != nil {
		ok = pre(out, in, id)
	}
	//fmt.Println(in.Data)
	if ok {
		for c := in.FirstChild; c != nil; c = c.NextSibling {
			recurse(out, c, id, pre)
		}
	} else {
		return
	}
}

//True value returned => Continue propogation
//True value returned in 2 cases: 1. out already found in a prev call. 2.out found in this call
func checkId(out, in *html.Node, id string) bool {
	//zero type for a pointer is nil
	if out.Data != "INVALID" {
		return false
	}
	if in != nil {
		for _, v := range in.Attr {
			if v.Key == "id" && v.Val == id {
				//fmt.Printf("id value: %s \n", v.Val)
				*out = *in //This is the only way to assign as out=in will only copy in's address in the current call to checkID
				return false
			}
		}
		return true
	}
	return true
}
