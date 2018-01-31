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

	print(Window.History().Length())

	ih := Document.QuerySelector("#infoHistory")
	persons := Document.QuerySelectorAll(".person")
	for _, person := range persons {
		person.AddEventListener("click", func(e Event) {
			e.PreventDefault()
			p := e.Target()
			href := p.GetAttribute("href")
			content := p.Dataset().Get("content").String()
			Window.History().PushState(content, "", href)
			ih.SetInnerHTML(content)
		})
	}

	Window.AddEventListener("popstate", func(e Event) {
		if e.Get("state") == nil {
			ih.SetInnerHTML("Entry Page")
		} else {
			ih.SetInnerHTML(e.Get("state").String())
		}
	})

	println(Window.PageXOffset())
	println(Window.PageYOffset())
	println(Window.ScrollX())
	println(Window.ScrollY())

	Document.Write("<br><strong>Hello</strong> <em>World</em>")
}
