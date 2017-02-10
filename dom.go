package godom

import (
	"github.com/gopherjs/gopherjs/js"
)

type Object struct {
	*js.Object
}

var Window = &Object{js.Global}
var Document = &Object{js.Global.Get("document")}

func Alert(s string) {
	Window.Call("alert", s)
}

func (o *Object) GetElementById(id string) *Object {
	return &Object{o.Call("getElementById", id)}
}

func (o *Object) QuerySelector(selectors string) *Object {
	return &Object{o.Call("querySelector", selectors)}
}

func (o *Object) QuerySelectorAll(selectors string) []*Object {
	nodeList := o.Call("querySelectorAll", selectors)
	length := nodeList.Get("length").Int()
	var nodes []*Object
	for i := 0; i < length; i++ {
		nodes = append(nodes, &Object{nodeList.Call("item", i)})
	}
	return nodes
}

func (o *Object) CreateElement(tag string) *Object {
	return &Object{Document.Call("createElement", tag)}
}

func (o *Object) CreateTextNode(textContent string) *Object {
	return &Object{Document.Call("createTextNode", textContent)}
}

func (o *Object) SetInnerHTML(html string) {
	o.Set("innerHTML", html)
}

func (o *Object) InnerHTML() string {
	return o.Get("innerHTML").String()
}

func (o *Object) ClassList() *DOMTokenList {
	return &DOMTokenList{o.Get("classList")}
}
