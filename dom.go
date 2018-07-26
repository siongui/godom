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
// If something is not implemented in this library, you can use methods of
// *js.Object of GopherJS to manipulate JavaScript DOM API directly.
//
// For example, if **textContent** property of foo is not implemented, you can
// use Get method of *js.Object in GopherJS to get the **textContent** property
// as follows:
//
//   t := foo.Get("textContent").String()
//
// If some DOM method is not implemented, you can use Call method provided by
// *js.Object in GopherJS. For example, if foo is an audio element and **play**
// method of foo is not implemented. You can use Call method of *js.Object in
// GopherJS to call **play** method of DOM audio element directly.
//
//   foo.Call("play")
//
// You can also read the godoc of GopherJS:
//
//   https://godoc.org/github.com/gopherjs/gopherjs/js
//   https://godoc.org/github.com/gopherjs/gopherjs/js#Object
//   https://godoc.org/github.com/gopherjs/gopherjs/js#Object.Call
//   https://godoc.org/github.com/gopherjs/gopherjs/js#Object.Set
//   https://godoc.org/github.com/gopherjs/gopherjs/js#Object.Get
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

// https://developer.mozilla.org/en-US/docs/Web/API/Window/history
func (w *window) History() *History {
	return &History{w.Get("history")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/alert
func (w *window) Alert(s string) {
	w.Call("alert", s)
}
func Alert(s string) {
	Window.Alert(s)
}

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (w *window) AddEventListener(t string, listener func(Event), args ...interface{}) {
	if len(args) == 1 {
		w.Call("addEventListener", t, listener, args[0])
	} else {
		w.Call("addEventListener", t, listener)
	}
}

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener
func (w *window) RemoveEventListener(t string, listener func(Event), args ...interface{}) {
	if len(args) == 1 {
		w.Call("removeEventListener", t, listener, args[0])
	} else {
		w.Call("removeEventListener", t, listener)
	}
}

// Properties of window object

// https://developer.mozilla.org/en-US/docs/Web/API/Window/pageXOffset
func (w *window) PageXOffset() float64 {
	return w.Get("pageXOffset").Float()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/pageYOffset
func (w *window) PageYOffset() float64 {
	return w.Get("pageYOffset").Float()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollX
func (w *window) ScrollX() float64 {
	return w.Get("scrollX").Float()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollY
func (w *window) ScrollY() float64 {
	return w.Get("scrollY").Float()
}
