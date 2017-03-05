package codegen

import (
	"os"
	"testing"
)

func TestGetAllStyleProperty(t *testing.T) {
	cssprops, err := GetAllStyleProperty()
	if err != nil {
		t.Error(err)
	}

	for _, prop := range cssprops {
		println(prop)
	}

	fo, err := os.Create("../cssstyledeclaration.go")
	if err != nil {
		t.Error(err)
	}
	defer fo.Close()

	err = GenCssProp(cssprops, fo)
	if err != nil {
		t.Error(err)
	}
}
