package godom

// This file implements KeyboardEvent interface
// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key
func (e Event) Key() string {
	return e.Get("key").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/keyCode
// MDN says keyCode is deprecated, use key instead.
func (e Event) KeyCode() int {
	return e.Get("keyCode").Int()
}

// Methods
