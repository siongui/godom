package godom

// This file implements XMLHttpRequest (XHR)
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest

import (
	"github.com/gopherjs/gopherjs/js"
)

type XMLHttpRequest struct {
	*js.Object
}

func NewXMLHttpRequest() *XMLHttpRequest {
	return &XMLHttpRequest{Window.Get("XMLHttpRequest").New()}
}

// Properties

func (x *XMLHttpRequest) ResponseXML() *Object {
	return &Object{x.Get("responseXML")}
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/open
//
// Optional parameters 'user' and 'password' are not supported.
func (x *XMLHttpRequest) Open(method, url string, args ...interface{}) {
	if len(args) == 1 {
		// args[0] should be bool async
		x.Call("open", method, url, args[0])
	} else {
		x.Call("open", method, url)
	}
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/send
func (x *XMLHttpRequest) Send(args ...interface{}) {
	if len(args) == 1 {
		x.Call("send", args[0])
	} else {
		x.Call("send")
	}
}
