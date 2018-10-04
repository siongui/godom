package wasm

// This file implements Document interface
// https://developer.mozilla.org/en-US/docs/Web/API/Document

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
func (v Value) GetElementById(id string) Value {
	return Value{v.Call("getElementById", id)}
}
