package godom

// This file implements Document interface
// https://developer.mozilla.org/en-US/docs/Web/API/Document

// Methods

func (o *Object) CreateElement(tag string) *Object {
	return &Object{Document.Call("createElement", tag)}
}

func (o *Object) CreateTextNode(textContent string) *Object {
	return &Object{Document.Call("createTextNode", textContent)}
}

func (o *Object) GetElementById(id string) *Object {
	return &Object{o.Call("getElementById", id)}
}
