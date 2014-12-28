//port ApacheCommons lang string utils
package stringUtils

// Check if string is blank
func IsEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

// Check if string is not blank
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// Check if any one of strings are blank
func IsAnyEmpty(strings ...string) bool {
	for _, v := range strings {
		if len(v) == 0 {
			return true
		}
	}
	return false
}
