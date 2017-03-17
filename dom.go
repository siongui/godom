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

type window struct {
	*js.Object
}

var Window = &window{js.Global}
var Document = &Object{js.Global.Get("document")}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/location
func (w *window) Location() *Location {
	return &Location{w.Get("location")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/alert
func (w *window) Alert(s string) {
	w.Call("alert", s)
}
func Alert(s string) {
	Window.Alert(s)
}
