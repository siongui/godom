// Package godom aims to make DOM manipulation in Go as similar to JavaScript
// as possible. The Go code is compiled to JavaScript via GopherJS. This package
// is used for front-end programming.
package godom

import (
	"github.com/gopherjs/gopherjs/js"
)

type Object struct {
	*js.Object
}

var Window = &Object{js.Global}
var Document = &Object{js.Global.Get("document")}

func Alert(s string) {
	Window.Call("alert", s)
}
