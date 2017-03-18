package mynumber_test

import (
	"fmt"

	"github.com/nasa9084/go-mynumber"
)

func ExampleValidate() {
	mynum := "889654457031"
	var isValid string
	if mynumber.Validate(mynum) {
		isValid = "valid"
	} else {
		isValid = "invalid"
	}
	fmt.Printf("%s is %s mynumber.\n", mynum, isValid)
}

func ExampleGenerate() {
	mynum := mynumber.Generate()
	fmt.Printf("%s\n", mynum)
}
