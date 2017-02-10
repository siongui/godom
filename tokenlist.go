package godom

import (
	"github.com/gopherjs/gopherjs/js"
)

type DOMTokenList struct {
	*js.Object
}

func (t *DOMTokenList) Contains(s string) bool {
	return t.Call("contains", s).Bool()
}
