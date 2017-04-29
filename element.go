package godom

// This file implements Element interface
// https://developer.mozilla.org/en-US/docs/Web/API/Element

// Properties

func (o *Object) ClassList() *DOMTokenList {
	return &DOMTokenList{o.Get("classList")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (o *Object) InnerHTML() string {
	return o.Get("innerHTML").String()
}
func (o *Object) SetInnerHTML(html string) {
	o.Set("innerHTML", html)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML
func (o *Object) OuterHTML() string {
	return o.Get("outerHTML").String()
}
func (o *Object) SetOuterHTML(html string) {
	o.Set("outerHTML", html)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/tagName
func (o *Object) TagName() string {
	return o.Get("tagName").String()
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/Element/getAttribute
func (o *Object) GetAttribute(attributeName string) string {
	return o.Call("getAttribute", attributeName).String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
func (o *Object) GetBoundingClientRect() *DOMRect {
	return &DOMRect{o.Call("getBoundingClientRect")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelector
func (o *Object) QuerySelector(selectors string) *Object {
	return &Object{o.Call("querySelector", selectors)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelectorAll
func (o *Object) QuerySelectorAll(selectors string) []*Object {
	nodeList := o.Call("querySelectorAll", selectors)
	length := nodeList.Get("length").Int()
	var nodes []*Object
	for i := 0; i < length; i++ {
		nodes = append(nodes, &Object{nodeList.Call("item", i)})
	}
	return nodes
}
