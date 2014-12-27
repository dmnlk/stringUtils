package stringUtils

import "testing"

func TestIsEmpty(t *testing.T) {
	blank := ""
	actual := IsEmpty(blank)
	if actual == false {
		t.Errorf("fail test, black should empty")
	}
	nilString := nil
	actual = IsEmpty(nilString)
	if actual == false {
		t.Error("fail test, nil should empty")
	}
}
