// This file implements Element interface
// https://developer.mozilla.org/en-US/docs/Web/API/Element
package godom

// Properties

func (o *Object) ClassList() *DOMTokenList {
	return &DOMTokenList{o.Get("classList")}
}

func (o *Object) InnerHTML() string {
	return o.Get("innerHTML").String()
}

func (o *Object) SetInnerHTML(html string) {
	o.Set("innerHTML", html)
}

// Methods

func (o *Object) GetBoundingClientRect() *DOMRect {
	return &DOMRect{o.Call("getBoundingClientRect")}
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
