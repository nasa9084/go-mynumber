package mynumber_test

import (
	"testing"

	mynumber "github.com/nasa9084/go-mynumber"
)

func TestValidate(t *testing.T) {
	candidates := []struct {
		In  string
		Out bool
	}{
		// True Patterns
		{"889654457031", true}, // 1
		{"660033834012", true}, // 2
		{"840387915883", true}, // 3
		{"974040902044", true}, // 4
		{"621791573895", true}, // 5
		{"815173424586", true}, // 6
		{"788294996377", true}, // 7
		{"732801235078", true}, // 8
		{"515068736919", true}, // 9
		{"618587094570", true}, // 10
		{"517850989500", true}, // 11
		// False Patterns
		{"889654457032", false}, // 1
		{"660033834013", false}, // 2
		{"840387915884", false}, // 3
		{"974040902045", false}, // 4
		{"621791573896", false}, // 5
		{"815173424587", false}, // 6
		{"788294996378", false}, // 7
		{"732801235079", false}, // 8
		{"515068736910", false}, // 9
		{"618587094571", false}, // 10
		{"517850989502", false}, // 11
		// Abnormal Input Patterns
		{"abcdefghijkl", false}, // alphabet, length=12
		{"12345678", false}, // number, length!=12
		{"abcdefgh", false}, // alphabet, length!=12
		{"12a34b56d78e", false}, // alphabet and number
		{"１２３４５６７８９０１１", false}, // full-width digits
	}

	for _, c := range candidates {
		result := mynumber.Validate(c.In)
		if result != c.Out {
			t.Errorf("%t != %t.\n", result, c.Out)
			return
		}
	}
}

func TestGenerate(t *testing.T) {
	for i := 0; i < 100; i++ {
		generated := mynumber.Generate()
		t.Logf("%s\n", generated)
		if !mynumber.Validate(generated) {
			t.Errorf("Invalid mynumber is generated: %s\n", generated)
			return
		}
	}
}
