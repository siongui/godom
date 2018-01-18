package godom

// This file implements Event interface
// https://developer.mozilla.org/en-US/docs/Web/API/Event

import (
	"github.com/gopherjs/gopherjs/js"
)

type Event struct {
	*js.Object
}

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/Event/target
func (e Event) Target() *Object {
	return &Object{e.Get("target")}
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/Event/preventDefault
func (e Event) PreventDefault() {
	e.Call("preventDefault")
}

// https://developer.mozilla.org/en-US/docs/Web/API/Event/stopImmediatePropagation
func (e Event) StopImmediatePropagation() {
	e.Call("stopImmediatePropagation")
}

// https://developer.mozilla.org/en-US/docs/Web/API/Event/stopPropagation
func (e Event) StopPropagation() {
	e.Call("stopPropagation")
}
