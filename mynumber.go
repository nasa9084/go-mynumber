package mynumber

import (
	"bytes"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

var numRegexp = regexp.MustCompile(`^[[:digit:]]{12}$`)

func init() {
	rand.Seed(time.Now().Unix())
}

// Validate mynumber
func Validate(number string) bool {
	if !numRegexp.MatchString(number) {
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
		if 1 <= n && n <= 6 {
			qn = n + 1
		} else {
			qn = n - 5
		}
		sum += (pn * qn)
	}
	var result int
	if sum%11 <= 1 {
		result = 0
	} else {
		result = 11 - (sum % 11)
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
		if 1 <= n && n <= 6 {
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
