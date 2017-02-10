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
