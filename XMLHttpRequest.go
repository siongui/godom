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

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/responseText
func (x *XMLHttpRequest) ResponseText() string {
	return x.Get("responseText").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/responseURL
func (x *XMLHttpRequest) ResponseURL() string {
	return x.Get("responseURL").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/responseXML
func (x *XMLHttpRequest) ResponseXML() *Object {
	return &Object{x.Get("responseXML")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/statusText
func (x *XMLHttpRequest) StatusText() string {
	return x.Get("statusText").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/withCredentials
func (x *XMLHttpRequest) WithCredentials() bool {
	return x.Get("withCredentials").Bool()
}

func (x *XMLHttpRequest) SetWithCredentials(c bool) {
	x.Set("withCredentials", c)
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/abort
func (x *XMLHttpRequest) Abort() {
	x.Call("abort")
}

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

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/overrideMimeType
func (x *XMLHttpRequest) OverrideMimeType(mimeType string) {
	x.Call("overrideMimeType", mimeType)
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/send
func (x *XMLHttpRequest) Send(args ...interface{}) {
	if len(args) == 1 {
		x.Call("send", args[0])
	} else {
		x.Call("send")
	}
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/setRequestHeader
func (x *XMLHttpRequest) SetRequestHeader(header, value string) {
	x.Call("setRequestHeader", header, value)
}
