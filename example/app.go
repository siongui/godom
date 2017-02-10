package main

import (
	. "github.com/siongui/godom"
)

func main() {
	div := Document.GetElementById("info")
	div.SetInnerHTML("abc")
	print(div.ClassList().Contains("invisible"))
}
