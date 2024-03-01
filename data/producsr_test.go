package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "soap",
		Price: 1,
		SKU:   "arst-arst-arst",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
