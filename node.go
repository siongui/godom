package godom

// This file implements Node interface
// https://developer.mozilla.org/en-US/docs/Web/API/Node

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/Node/childNodes
func (o *Object) ChildNodes() []*Object {
	nodeList := o.Get("childNodes")
	length := nodeList.Get("length").Int()
	var nodes []*Object
	for i := 0; i < length; i++ {
		nodes = append(nodes, &Object{nodeList.Call("item", i)})
	}
	return nodes
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/firstChild
func (o *Object) FirstChild() *Object {
	return &Object{o.Get("firstChild")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/lastChild
func (o *Object) LastChild() *Object {
	return &Object{o.Get("lastChild")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/nextSibling
func (o *Object) NextSibling() *Object {
	return &Object{o.Get("nextSibling")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeType
func (o *Object) NodeType() int {
	return o.Get("nodeType").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeValue
func (o *Object) NodeValue() string {
	return o.Get("nodeValue").String()
}
func (o *Object) SetNodeValue(s string) {
	o.Set("nodeValue", s)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/parentNode
func (o *Object) ParentNode() *Object {
	return &Object{o.Get("parentNode")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
// Returns the textual content of an element and all its descendants.
func (o *Object) TextContent() string {
	return o.Get("textContent").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
// Sets the textual content of an element and all its descendants.
func (o *Object) SetTextContent(s string) {
	o.Set("textContent", s)
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/Node/appendChild
func (o *Object) AppendChild(c *Object) {
	o.Call("appendChild", c)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/hasChildNodes
func (o *Object) HasChildNodes() bool {
	return o.Call("hasChildNodes").Bool()
}

// Inserts the first Node given in a parameter immediately before the second,
// child of this element, Node.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/insertBefore
func (o *Object) InsertBefore(newNode, referenceNode *Object) *Object {
	return &Object{o.Call("insertBefore", newNode, referenceNode)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/isEqualNode
func (o *Object) IsEqualNode(n *Object) bool {
	return o.Call("isEqualNode", n).Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/isSameNode
func (o *Object) IsSameNode(n *Object) bool {
	return o.Call("isSameNode", n).Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/removeChild
func (o *Object) RemoveChild(c *Object) *Object {
	return &Object{o.Call("removeChild", c)}
}
