// Package godom aims to make DOM manipulation in Go as similar to JavaScript
// as possible. The Go code is compiled to JavaScript via GopherJS. This package
// is used for front-end programming.
//
// You can import godom as follows:
//
//   import (
//   	. "github.com/siongui/godom"
//   )
//
// Access dom element with id="foo" as follows:
//
//   foo := Document.GetElementById("foo")
//
// Or
//
//   foo := Document.QuerySelector("#foo")
//
// If something is not implemented in this library, you can call methods of
// *js.Object of GopherJS to manipulate JavaScript DOM API directly.
//
// For example, if **textContent** property of foo is not implemented, you can
// use Get method to get the **textContent** property as follows:
//
//   t := foo.Get("textContent").String()
//
// If some method is not implemented, you can use Call method provided by
// GopherJS. For example, if foo is audio element and **play** method of foo is
// not implemented. You can use GopherJS Call method to call **play** method of
// foo directly.
//
//   foo.Call("play")
//
package godom

import (
	"github.com/gopherjs/gopherjs/js"
)

// This type usually represents a DOM element or node. The type is actually a
// wrapper for *js.Object of GopherJS. Any instance of the type can also use the
// methods provided by GopherJS, such as Call, Get, and Set methods. If
// something is not implemented in this library, you can use the methods
// provided by GopherJS to manipulate JavaScript DOM API directly.
type Object struct {
	*js.Object
}

type window struct {
	*js.Object
}

// equivalent to window object in JavaScript DOM API.
var Window = &window{js.Global}

// equivalent to document object in JavaScript DOM API.
var Document = &Object{js.Global.Get("document")}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/location
func (w *window) Location() *Location {
	return &Location{w.Get("location")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/navigator
func (w *window) Navigator() *Navigator {
	return &Navigator{w.Get("navigator")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/alert
func (w *window) Alert(s string) {
	w.Call("alert", s)
}
func Alert(s string) {
	Window.Alert(s)
}
