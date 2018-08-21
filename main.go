package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func init() {
	t := time.Now()
	rand.Seed(t.UnixNano())
}

func main() {

	ssn := generateSSN()
	for !isValid(ssn) {
		ssn = generateSSN()
	}
	fmt.Println(ssn)
}

func generateSSN() string {
	var builder strings.Builder

	for i := 1; i <= 9; i++ {
		if i == 4 || i == 6 {
			builder.WriteString("-")
		}
		d := rand.Intn(9)
		builder.WriteString(strconv.Itoa(d))
	}
	return builder.String()
}

// Just focussing on the first grouping for now:
// https://en.wikipedia.org/wiki/Social_Security_number#Valid_SSNs
func isValid(ssn string) bool {
	startsWith9, _ := regexp.MatchString("^9[0-9][0-9]", ssn)
	if strings.HasPrefix(ssn, "666") || startsWith9 {
		return false
	}
	return true
}
