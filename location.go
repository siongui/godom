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

func (l *Location) Host() string {
	return l.Get("host").String()
}

func (l *Location) Hostname() string {
	return l.Get("hostname").String()
}

func (l *Location) Href() string {
	return l.Get("href").String()
}

func (l *Location) Origin() string {
	return l.Get("origin").String()
}

func (l *Location) Pathname() string {
	return l.Get("pathname").String()
}

func (l *Location) Port() string {
	return l.Get("port").String()
}

func (l *Location) Protocol() string {
	return l.Get("protocol").String()
}

func (l *Location) Search() string {
	return l.Get("search").String()
}
