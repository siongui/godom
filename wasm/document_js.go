package wasm

// This file implements Document interface
// https://developer.mozilla.org/en-US/docs/Web/API/Document

// Properties

// Returns the currently focused element
// https://developer.mozilla.org/en-US/docs/Web/API/Document/activeElement
func (v Value) ActiveElement() Value {
	return Value{v.Get("activeElement")}
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
func (v Value) CreateElement(tag string) Value {
	return Value{Document.Call("createElement", tag)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/createTextNode
func (v Value) CreateTextNode(textContent string) Value {
	return Value{Document.Call("createTextNode", textContent)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
func (v Value) GetElementById(id string) Value {
	return Value{v.Call("getElementById", id)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/write
func (v Value) Write(markup string) {
	Document.Call("write", markup)
}
