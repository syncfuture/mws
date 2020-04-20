package xmlutil

import (
	"github.com/subchen/go-xmldom"
)

func GetNodeText(node *xmldom.Node, path string) string {
	x := node.QueryOne(path)
	if x != nil {
		return x.Text
	}
	return ""
}
