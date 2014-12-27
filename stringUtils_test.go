package stringUtils

import "testing"

func TestIsEmpty(t *testing.T) {
	var notBlank string = "golang"
	actual := IsEmpty(notBlank)
	if  actual == true {
		t.Errorf("fail test, not blank string should return false")
	}
	var blank string = ""
	actual = IsEmpty(blank)
	if actual == false {
		t.Errorf("fail test, black should empty")
	}

}

