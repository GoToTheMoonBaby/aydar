package main

import "fmt"

func StrConcat(a string, b string) (res string) {
	return a + " " + b
}

func main() {
	var a = "Go to the"
	var b = "Moon"
	fmt.Println(StrConcat(a, b))
}
