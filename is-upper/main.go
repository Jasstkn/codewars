package main

import "strings"

type MyString string

func (s MyString) IsUpperCase() bool {
	if string(s) == strings.ToUpper(string(s)) {
		return true
	}
	return false

}

func main() {

}
