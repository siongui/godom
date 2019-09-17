package godom

// This file implements Document interface
// https://developer.mozilla.org/en-US/docs/Web/API/Document

// Properties

// Returns the currently focused element
// https://developer.mozilla.org/en-US/docs/Web/API/Document/activeElement
func (v Value) ActiveElement() Value {
	return Value{v.Get("activeElement")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/cookie
func (v Value) Cookie() Value {
	return Document.Get("cookie").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/title
func (v Value) Title() Value {
	return Document.Get("title").String()
}
func (v Value) SetTitle(title string) Value {
	Document.set("title", title)
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

// https://developer.mozilla.org/en-US/docs/Web/API/Document/writeln
func (v Value) Writeln(markup string) {
	Document.call("writeln", markup)
}
