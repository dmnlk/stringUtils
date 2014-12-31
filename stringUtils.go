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
func IsAnyEmpty(ss ...string) bool {
	for _, s := range ss {
		if len(s) == 0 {
			return true
		}
	}
	return false
}

// Checks if none of the strings are empty
func IsNoneEmpty(ss ...string) bool {
	for _, s := range ss {
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

// Checks if a string is not empty (""), not null and not whitespace only.
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

