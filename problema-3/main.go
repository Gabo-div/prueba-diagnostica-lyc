package main

import (
	"fmt"
	"regexp"
)

func main() {
	testStrings := []string{
		"\"this is a string\"",
		"1.23e-5",
		"192.168.1.1",
		"test@example.com",
		"not a valid string",
		"1.2.3.4.5",
		"invalid-email",
	}

	for _, s := range testStrings {
		fmt.Printf("Testing: %s\n", s)

		if isString(s) {
			fmt.Println("  - Recognized as a string")
		}

		if isScientific(s) {
			fmt.Println("  - Recognized as scientific notation")
		}

		if isIP(s) {
			fmt.Println("  - Recognized as an IP address")
		}

		if isEmail(s) {
			fmt.Println("  - Recognized as an email address")
		}
	}
}

func isString(s string) bool {
	re := regexp.MustCompile(`^".*"$`)
	return re.MatchString(s)
}

func isScientific(s string) bool {
	re := regexp.MustCompile(`^[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?$`)
	return re.MatchString(s)
}

func isIP(s string) bool {
	re := regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	return re.MatchString(s)
}

func isEmail(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(s)
}
