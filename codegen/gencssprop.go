package codegen

import (
	"io"
	"strings"
	"text/template"
)

const gotmpl = `package godom
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
{{range $prop := .}}
func (s *CSSStyleDeclaration) {{title $prop}}() string {
	return s.Get("{{$prop}}").String()
}

func (s *CSSStyleDeclaration) Set{{title $prop}}(v string) {
	s.Set("{{$prop}}", v)
}
{{end}}
`

func GenCssProp(cssprops []string, w io.Writer) (err error) {
	funcMap := template.FuncMap{
		"title": strings.Title,
	}

	tmpl, err := template.New("").Funcs(funcMap).Parse(gotmpl)
	if err != nil {
		return
	}

	return tmpl.Execute(w, cssprops)
}
