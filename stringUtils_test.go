package stringUtils

import "testing"

func TestIsEmpty(t *testing.T) {
	var notBlank string = "golang"
	actual := IsEmpty(notBlank)
	if actual == true {
		t.Errorf("fail test, not blank string should return false")
	}
	var blank string = ""
	actual = IsEmpty(blank)
	if actual == false {
		t.Errorf("fail test, black should empty")
	}
	var spaceIncludeText string = " golang "
	actual = IsEmpty(spaceIncludeText)
	if actual == true {
		t.Errorf("fail test, not blank string should return false")
	}
}

func TestIsNotEmpty(t *testing.T) {
	var notBlank string = "golang"
	actual := IsNotEmpty(notBlank)
	if actual == false {
		t.Errorf("fail test, not blank string should return true")
	}
	var blank string = ""
	actual = IsNotEmpty(blank)
	if actual == true {
		t.Errorf("fail test, black should return false")
	}
	var spaceIncludeText string = " golang "
	actual = IsNotEmpty(spaceIncludeText)
	if actual == false {
		t.Errorf("fail test, not blank string should return false")
	}
}

func TestIsAnyEmpty(t *testing.T) {
	actual := IsAnyEmpty("")
	if actual == false {
		t.Errorf("fail test, blank string should return true")
	}
	actual = IsAnyEmpty("", "golang")
	if actual == false {
		t.Errorf("fail test, blank string and not blank string should return true")
	}
	actual = IsAnyEmpty("golang", "")
	if actual == false {
		t.Errorf("fail test, not blank and blank string should return true")
	}
	actual = IsAnyEmpty(" golang ", "")
	if actual == false {
		t.Errorf("fail test, not blank and blank string should return true")
	}
	actual = IsAnyEmpty(" ", "golang")
	if actual {
		t.Errorf("fail test, not blank strings should return false")
	}
	actual = IsAnyEmpty("golang", "gophers")
	if actual {
		t.Errorf("fail test, not blank strings should return false")
	}
}

func TestIsNoneEmpty(t *testing.T) {
	actual := IsNoneEmpty("")
	if actual == true {
		t.Errorf("fail test, blank string should return false")
	}
	actual = IsNoneEmpty("", "golang")
	if actual == true {
		t.Errorf("fail test, blank string and string should return false")
	}
	actual = IsNoneEmpty("golang", "")
	if actual == true {
		t.Errorf("fail test, string and blank string should return false")
	}
	actual = IsNoneEmpty(" golang ", "")
	if actual == true {
		t.Errorf("fail test, string and blank string should return false")
	}
	actual = IsNoneEmpty(" ", "golang")
	if actual == false {
		t.Errorf("fail test, strings  should return true")
	}
	actual = IsNoneEmpty("golang", "gophers")
	if actual == false {
		t.Errorf("fail test, strings should return true")
	}
}

func TestIsBlank(t *testing.T) {
	actual := IsBlank("")
	if actual == false {
		t.Errorf("fail test, blank string should return true")
	}
	actual = IsBlank(" ")
	if actual == false {
		t.Errorf("fail test, space string should return true")
	}
	actual = IsBlank("    ")
	if actual == false {
		t.Errorf("fail test, space strings should return true")
	}
	actual = IsBlank("golang")
	if actual == true {
		t.Errorf("fail test, string should return false")
	}
	actual = IsBlank(" golang ")
	if actual == true {
		t.Errorf("fail test, string should return false")
	}
	actual = IsBlank("   golang")
	if actual == true {
		t.Errorf("fail test, string should return false")
	}
}

func TestIsNotBlank(t *testing.T) {
	actual := IsNotBlank("")
	if actual == true {
		t.Errorf("fail test, blank string should return false")
	}
	actual = IsNotBlank(" ")
	if actual == true {
		t.Errorf("fail test, space string should return false")
	}
	actual = IsNotBlank("    ")
	if actual == true {
		t.Errorf("fail test, space strings should return false")
	}
	actual = IsNotBlank("golang")
	if actual == false {
		t.Errorf("fail test, string should return true")
	}
	actual = IsNotBlank(" golang ")
	if actual == false {
		t.Errorf("fail test, string should return true")
	}
}

func TestIsAnyBlank(t *testing.T) {
	actual := IsAnyBlank("")
	if actual == false {
		t.Errorf("fail test, blank string should return true")
	}
	actual = IsAnyBlank("", "golang")
	if actual == false {
		t.Errorf("fail test, blank string should return true")
	}
	actual = IsAnyBlank("golang", "")
	if actual == false {
		t.Errorf("fail test, blank string should return true")
	}
	actual = IsAnyBlank(" golang ", "")
	if actual == false {
		t.Errorf("fail test, blank string should return true")
	}
	actual = IsAnyBlank(" ", "golang")
	if actual == false {
		t.Errorf("fail test, blank string should return true")
	}
	actual = IsAnyBlank("golang", "gophers")
	if actual == true {
		t.Errorf("fail test, blank string should return false")
	}

}
