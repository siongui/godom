// Package wasm aims to make DOM manipulation in Go as similar to JavaScript as
// possible. The Go code is compiled to WebAssembly via official Go compiler
// 1.11 or later. This package is used for front-end programming.
//
// You can import wasm as follows:
//
//   import (
//   	. "github.com/siongui/godom/wasm"
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
// js.Value in "syscall/js" to manipulate DOM API directly.
//
// For example, if **textContent** property of foo is not implemented, you can
// use Get method of js.Value in "syscall/js" to get the **textContent**
// property as follows:
//
//   t := foo.Get("textContent").String()
//
// If some DOM method is not implemented, you can use Call method provided by
// js.Value in "syscall/js". For example, if foo is an audio element and
// **play** method of foo is not implemented. You can use Call method of
// js.Value in "syscall/js" to call **play** method of DOM audio element
// directly.
//
//   foo.Call("play")
//
// You can also read the godoc of "syscall/js":
//
//   https://tip.golang.org/pkg/syscall/js/
//   https://godoc.org/syscall/js
//
package wasm

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
