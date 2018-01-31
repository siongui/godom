package godom

// This file implements Document interface
// https://developer.mozilla.org/en-US/docs/Web/API/Document

// Properties

// Returns the currently focused element
// https://developer.mozilla.org/en-US/docs/Web/API/Document/activeElement
func (o *Object) ActiveElement() *Object {
	return &Object{o.Get("activeElement")}
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
func (o *Object) CreateElement(tag string) *Object {
	return &Object{Document.Call("createElement", tag)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/createTextNode
func (o *Object) CreateTextNode(textContent string) *Object {
	return &Object{Document.Call("createTextNode", textContent)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
func (o *Object) GetElementById(id string) *Object {
	return &Object{o.Call("getElementById", id)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/write
func (o *Object) Write(markup string) {
	Document.Call("write", markup)
}
