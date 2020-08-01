package godom

// This file implements JavaScript/XSLT Bindings
// https://developer.mozilla.org/en-US/docs/Web/XSLT/XSLT_JS_interface_in_Gecko/JavaScript_XSLT_Bindings

import (
	"github.com/gopherjs/gopherjs/js"
)

type XSLTProcessor struct {
	*js.Object
}

func NewXSLTProcessor() *XSLTProcessor {
	return &XSLTProcessor{Window.Get("XSLTProcessor").New()}
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/XSLTProcessor#Methods
func (x *XSLTProcessor) ImportStylesheet(node *Object) {
	x.Call("importStylesheet", node)
}
