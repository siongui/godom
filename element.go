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

// Methods

// Remove keyboard focus from the current element
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/blur
func (o *Object) Blur() {
	o.Call("blur")
}

// Set focus on the specified element, if it can be focused.
// https://developer.mozilla.org/en/docs/Web/API/HTMLElement/focus
func (o *Object) Focus() {
	o.Call("focus")
}

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
