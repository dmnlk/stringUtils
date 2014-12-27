package stringUtils

import "testing"

func TestIsEmpty(t *testing.T) {
	blank := ""
	actual := IsEmpty(blank)
	if actual == false {
		t.Fail()
	}
}
