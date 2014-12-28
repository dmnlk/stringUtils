package stringUtils

func IsEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

func IsAnyEmpty(strings ...string) bool {
	for _, v := range strings {
		if len(v) == 0 {
			return true
		}
	}
	return false
}
