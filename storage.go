package godom

// This file implements Storage interface
// https://developer.mozilla.org/en-US/docs/Web/API/Storage

import (
	"github.com/gopherjs/gopherjs/js"
)

type Storage struct {
	*js.Object
}

var LocalStorage = &Storage{Window.Get("localStorage")}
var SessionStorage = &Storage{Window.Get("sessionStorage")}

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/Storage/length
func (s *Storage) Length() int {
	return s.Get("length").Int()
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/Storage/key
func (s *Storage) Key(index int) string {
	return s.Call("key", index).String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Storage/getItem
func (s *Storage) GetItem(keyName string) string {
	return s.Call("getItem", keyName).String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Storage/setItem
func (s *Storage) SetItem(keyName, keyValue string) {
	s.Call("setItem", keyName, keyValue)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Storage/removeItem
func (s *Storage) RemoveItem(keyName string) {
	s.Call("removeItem", keyName)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Storage/clear
func (s *Storage) Clear() {
	s.Call("clear")
}
