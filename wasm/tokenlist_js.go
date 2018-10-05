package wasm

// This file implements DOMTokenList interface
// https://developer.mozilla.org/en/docs/Web/API/DOMTokenList

import (
	"syscall/js"
)

type DOMTokenList struct {
	js.Value
}

func (t DOMTokenList) Length() int {
	return t.Get("length").Int()
}

func (t DOMTokenList) Contains(s string) bool {
	return t.Call("contains", s).Bool()
}

func (t DOMTokenList) Add(s string) {
	t.Call("add", s)
}

func (t DOMTokenList) Remove(s string) {
	t.Call("remove", s)
}

func (t DOMTokenList) Toggle(s string) {
	t.Call("toggle", s)
}
