package godom

// This file implements History interface
// https://developer.mozilla.org/en-US/docs/Web/API/History

import (
	"github.com/gopherjs/gopherjs/js"
)

type History struct {
	*js.Object
}

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/History/length
func (h *History) Length() int {
	return h.Get("length").Int()
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/History/back
func (h *History) Back() {
	h.Call("back")
}

// https://developer.mozilla.org/en-US/docs/Web/API/History/forward
func (h *History) Forward() {
	h.Call("forward")
}

// https://developer.mozilla.org/en-US/docs/Web/API/History/go
func (h *History) Go(p int) {
	h.Call("go", p)
}

// https://developer.mozilla.org/en-US/docs/Web/API/History_API#The_pushState()_method
func (h *History) PushState(stateObj interface{}, title, url string) {
	h.Call("pushState", stateObj, title, url)
}

// https://developer.mozilla.org/en-US/docs/Web/API/History/replaceState
func (h *History) ReplaceState(stateObj interface{}, title, url string) {
	h.Call("replaceState", stateObj, title, url)
}
