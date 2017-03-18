package mynumber

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Validate mynumber
func Validate(number string) bool {
	if _, err := strconv.Atoi(number); err != nil || len(number) != 12 {
		return false
	}
	sum := 0
	checkDigit, err := strconv.Atoi(string(number[11]))
	if err != nil {
		panic(err)
	}
	for i, d := range number[:11] {
		pn, err := strconv.Atoi(string(d))
		if err != nil {
			panic(err)
		}
		n := 11 - i
		var qn int
		if n <= 6 { // always 1 <= n
			qn = n + 1
		} else {
			qn = n - 5
		}
		sum += (pn * qn)
	}
	var result int
	sum %= 11
	if sum <= 1 {
		result = 0
	} else {
		result = 11 - sum
	}
	return result == checkDigit
}

// Generate dummy mynumber
func Generate() string {
	buf := bytes.NewBufferString("")
	cd := 0
	for n := 11; n > 0; n-- {
		pn := rand.Intn(10)
		buf.WriteString(strconv.Itoa(pn))
		var qn int
		if n <= 6 { // always 1 <= n
			qn = n + 1
		} else {
			qn = n - 5
		}
		cd += pn * qn
	}
	if cd%11 <= 1 {
		buf.WriteString("0")
	} else {
		buf.WriteString(strconv.Itoa(11 - (cd % 11)))
	}
	return buf.String()
}
