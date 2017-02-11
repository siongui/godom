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
