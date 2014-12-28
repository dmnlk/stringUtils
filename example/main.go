package main

import (
	"fmt"
	"github.com/dmnlk/stringUtils"
)


func main() {
	a := "test"
	fmt.Println(stringUtils.IsEmpty(a))
	fmt.Println(stringUtils.IsAnyEmpty("test", "aaaa"))
}
