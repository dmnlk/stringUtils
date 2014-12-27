package stringUtils

import (

)

func IsEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}
