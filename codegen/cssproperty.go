package codegen

import (
	"github.com/PuerkitoBio/goquery"
)

func GetAllStyleProperty() (cssprops []string, err error) {
	url := "https://www.w3schools.com/jsref/dom_obj_style.asp"

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return
	}

	csspropslinks := doc.Find(".w3-table-all").First().Find("a")
	csspropslinks.Each(func(i int, s *goquery.Selection) {
		cssprops = append(cssprops, s.Text())
	})

	return
}
