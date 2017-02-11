package main

import (
	. "github.com/siongui/godom"
)

func main() {
	div := Document.GetElementById("info")
	print(div.ClassList().Contains("invisible"))

	p := Document.CreateElement("p")
	p.SetTextContent("content of p")
	div.AppendChild(p)
	print(div.InnerHTML())

	strong := Document.CreateElement("strong")
	strong.SetInnerHTML("bold")
	div.AppendChild(strong)
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
}
