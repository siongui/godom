package godom

// This file implements EventTarget interface
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (o *Object) AddEventListener(t string, listener func(Event), args ...interface{}) {
	if len(args) == 1 {
		o.Call("addEventListener", t, listener, args[0])
	} else {
		o.Call("addEventListener", t, listener)
	}
}

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener
func (o *Object) RemoveEventListener(t string, listener func(Event), args ...interface{}) {
	if len(args) == 1 {
		o.Call("removeEventListener", t, listener, args[0])
	} else {
		o.Call("removeEventListener", t, listener)
	}
}
