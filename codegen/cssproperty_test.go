package codegen

import (
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
}
