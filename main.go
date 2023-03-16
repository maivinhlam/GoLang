package main

import (
	"fmt"
	"regexp"
)

func main() {
	//pattern
	r, _ := regexp.Compile(`^(\d{1,5}(\.\d{1,5})?)+%$`)

	fmt.Println(r.MatchString("1.01.0%"))
}
