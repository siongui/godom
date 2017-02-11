// This file implements useful functions not in DOM API
package godom

func (o *Object) RemoveAllChildNodes() {
	for o.HasChildNodes() {
		o.RemoveChild(o.LastChild())
	}
}
