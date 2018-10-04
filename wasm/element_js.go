package wasm

// This file implements Element interface
// https://developer.mozilla.org/en-US/docs/Web/API/Element

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelector
func (v Value) QuerySelector(selectors string) Value {
	return Value{v.Call("querySelector", selectors)}
}
