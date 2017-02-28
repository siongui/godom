package godom
// This file implements HTMLElement interface
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement

// Properties

//https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/style
func (o *Object) Style() *CSSStyleDeclaration {
	return &CSSStyleDeclaration{o.Get("style")}
}
