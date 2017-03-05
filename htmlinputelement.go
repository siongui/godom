package godom

// This file implements HTMLInputElement interface
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
// https://developer.mozilla.org/en/docs/Web/API/HTMLInputElement

// Properties

func (o *Object) Value() string {
	return o.Get("value").String()
}

func (o *Object) SetValue(s string) {
	o.Set("value", s)
}
