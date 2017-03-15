package godom

// This file implements useful functions not in DOM API

func (o *Object) RemoveAllChildNodes() {
	for o.HasChildNodes() {
		o.RemoveChild(o.LastChild())
	}
}

// element.AppendBefore(newElement) - insert newElement before element.
// Some people call this as *insert before*, but this calling is confusing
// because the meaning of insertBefore() in DOM API is different.
// http://stackoverflow.com/a/32135318
func (o *Object) AppendBefore(n *Object) {
	o.ParentNode().InsertBefore(n, o)
}

// element.AppendAfter(newElement) - insert newElement after element.
// Some people call this as *insert after*.
// http://stackoverflow.com/a/32135318
func (o *Object) AppendAfter(n *Object) {
	o.ParentNode().InsertBefore(n, o.NextSibling())
}

// Check if the element is currently focused
func (o *Object) IsFocused() bool {
	return o.IsEqualNode(Document.ActiveElement())
}
