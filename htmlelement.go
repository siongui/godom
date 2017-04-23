package godom

// This file implements HTMLElement interface
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement

import (
	"github.com/gopherjs/gopherjs/js"
)

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/style
func (o *Object) Style() *CSSStyleDeclaration {
	return &CSSStyleDeclaration{o.Get("style")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dataset
func (o *Object) Dataset() *js.Object {
	return o.Get("dataset")
}

// Methods

// Remove keyboard focus from the current element
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/blur
func (o *Object) Blur() {
	o.Call("blur")
}

// Set focus on the specified element, if it can be focused.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus
func (o *Object) Focus() {
	o.Call("focus")
}
