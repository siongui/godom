package main

import (
	"syscall/js"
)

// This type usually represents a DOM element or node. The type is actually a
// wrapper for js.Value in syscall/js package. Any instance of the type can also
// use the methods provided by syscall/js, such as Call, Get, and Set methods.
// If something is not implemented in this library, you can use the methods
// provided by syscall/js to manipulate JavaScript DOM API directly.
type Value struct {
	js.Value
}

type window struct {
	js.Value
}

// equivalent to window object in JavaScript DOM API.
var Window = &window{js.Global()}

// equivalent to document object in JavaScript DOM API.
var Document = Value{js.Global().Get("document")}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/alert
func (w *window) Alert(s string) {
	w.Call("alert", s)
}
func Alert(s string) {
	Window.Alert(s)
}
