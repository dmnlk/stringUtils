//port ApacheCommons lang string utils
package stringUtils

import "regexp"

// Check if string is empty
func IsEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

// Check if string is not empty
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// Check if any one of strings are empty
func IsAnyEmpty(strings ...string) bool {
	for _, s := range strings {
		if len(s) == 0 {
			return true
		}
	}
	return false
}

// Checks if none of the strings are empty
func IsNoneEmpty(strings ...string) bool {
	for _, s := range strings {
		if len(s) == 0 {
			return false
		}
	}
	return true
}

// Checks if a string is whitespace, empty ("")
func IsBlank(s string) bool {
	if len(s) == 0 {
		return true
	}
	reg := regexp.MustCompile(`^\s+$`)
	actual := reg.MatchString(s)
	if actual {
		return true
	}
	return false
}
