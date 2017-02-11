// This file implements Node interface
// https://developer.mozilla.org/en-US/docs/Web/API/Node
package godom

func (o *Object) TextContent() string {
	return o.Get("textContent").String()
}

func (o *Object) SetTextContent(s string) {
	o.Set("textContent", s)
}

func (o *Object) AppendChild(c *Object) {
	o.Call("appendChild", c)
}

func (o *Object) HasChildNodes() bool {
	return o.Call("hasChildNodes").Bool()
}

func (o *Object) RemoveChild(c *Object) *Object {
	return &Object{o.Call("removeChild", c)}
}

func (o *Object) LastChild() *Object {
	return &Object{o.Get("lastChild")}
}

func (o *Object) FirstChild() *Object {
	return &Object{o.Get("firstChild")}
}
