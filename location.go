package godom

// This file implements Location interface
// https://developer.mozilla.org/en-US/docs/Web/API/Location

import (
	"github.com/gopherjs/gopherjs/js"
)

type Location struct {
	*js.Object
}

// Properties

func (l *Location) Hostname() string {
	return l.Get("hostname").String()
}
