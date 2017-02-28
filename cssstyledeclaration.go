package godom
// This file implements CSSStyleDeclaration interface
// https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleDeclaration

import (
	"github.com/gopherjs/gopherjs/js"
)

type CSSStyleDeclaration struct {
	*js.Object
}

// Attributes

func (s *CSSStyleDeclaration) CssText() string {
	return s.Get("cssText").String()
}

func (s *CSSStyleDeclaration) Length() int {
	return s.Get("length").Int()
}

// CSS Properties

/* left */
func (s *CSSStyleDeclaration) Left() string {
	return s.Get("cssText").String()
}

func (s *CSSStyleDeclaration) SetLeft(v string) {
	s.Set("left", v)
}

/* max-width */
func (s *CSSStyleDeclaration) MaxWidth() string {
	return s.Get("maxWidth").String()
}

func (s *CSSStyleDeclaration) SetMaxWidth(v string) {
	s.Set("maxWidth", v)
}

/* min-width */
func (s *CSSStyleDeclaration) MinWidth() string {
	return s.Get("minWidth").String()
}

func (s *CSSStyleDeclaration) SetMinWidth(v string) {
	s.Set("minWidth", v)
}
