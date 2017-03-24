package main

import (
	. "github.com/siongui/godom"
)

func myListener(e Event) {
	print(e.Target().TextContent())
	e.Target().RemoveEventListener("click", myListener)
}

func main() {
	div := Document.GetElementById("info")
	print(div.ClassList().Contains("invisible"))

	p := Document.CreateElement("p")
	p.SetTextContent("content of p")
	div.AppendChild(p)
	print(div.InnerHTML())

	strong := Document.CreateElement("strong")
	strong.SetInnerHTML("bold")
	p.AppendAfter(strong)
	print(strong.TextContent())

	print(div.GetBoundingClientRect().X())
	print(div.GetBoundingClientRect().Y())
	print(div.GetBoundingClientRect().Width())
	print(div.GetBoundingClientRect().Height())
	print(div.GetBoundingClientRect().Top())
	print(div.GetBoundingClientRect().Right())
	print(div.GetBoundingClientRect().Bottom())
	print(div.GetBoundingClientRect().Left())

	div.RemoveAllChildNodes()

	p = Document.CreateElement("p")
	p.SetTextContent("click me")
	p.AddEventListener("click", myListener)
	div.AppendChild(p)

	input := Document.GetElementById("ip")
	print(input.IsFocused())
	input.Focus()
	print(input.IsFocused())
	print(input.OuterHTML())

	print(Window.Location().Hostname())
	print(Window.Navigator().Language())
	print(Window.Navigator().Languages())
}
