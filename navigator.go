package godom

// This file implements Navigator interface
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator

import (
	"github.com/gopherjs/gopherjs/js"
)

type Navigator struct {
	*js.Object
}

// Properties

func (n *Navigator) Language() string {
	return n.Get("language").String()
}

func (n *Navigator) Languages() string {
	return n.Get("languages").String()
}
