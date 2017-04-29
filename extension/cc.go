package main

import (
	. "github.com/siongui/godom"
	"github.com/siongui/gojianfan"
	"strings"
)

func traverse(elm *Object) {
	nodeType := elm.NodeType()

	if nodeType == 1 || nodeType == 9 {
		// element node or document node
		for _, node := range elm.ChildNodes() {
			// recursively call to traverse
			traverse(node)
		}
		return
	}

	if nodeType == 3 {
		// text node
		v := strings.TrimSpace(elm.NodeValue())
		if v != "" {
			println(gojianfan.S2T(v))
		}
		return
	}
}

func main() {
	traverse(Document)
}
