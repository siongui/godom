package wasm

// This file implements Element interface
// https://developer.mozilla.org/en-US/docs/Web/API/Element

// Properties

func (v Value) ClassList() DOMTokenList {
	return DOMTokenList{v.Get("classList")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (v Value) InnerHTML() string {
	return v.Get("innerHTML").String()
}
func (v Value) SetInnerHTML(html string) {
	v.Set("innerHTML", html)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML
func (v Value) OuterHTML() string {
	return v.Get("outerHTML").String()
}
func (v Value) SetOuterHTML(html string) {
	v.Set("outerHTML", html)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/tagName
func (v Value) TagName() string {
	return v.Get("tagName").String()
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/Element/getAttribute
func (v Value) GetAttribute(attributeName string) string {
	return v.Call("getAttribute", attributeName).String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
func (v Value) GetBoundingClientRect() DOMRect {
	return DOMRect{v.Call("getBoundingClientRect")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelector
func (v Value) QuerySelector(selectors string) Value {
	return Value{v.Call("querySelector", selectors)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelectorAll
func (v Value) QuerySelectorAll(selectors string) []Value {
	nodeList := v.Call("querySelectorAll", selectors)
	length := nodeList.Get("length").Int()
	var nodes []Value
	for i := 0; i < length; i++ {
		nodes = append(nodes, Value{nodeList.Call("item", i)})
	}
	return nodes
}
