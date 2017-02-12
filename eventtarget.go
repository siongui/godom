// This file implements EventTarget interface
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
package godom

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (o *Object) AddEventListener(t string, listener func(Event), args ...interface{}) {
	o.Call("addEventListener", t, listener, args)
}
