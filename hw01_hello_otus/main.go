package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	greeting := "Hello, OTUS!"
	fmt.Print(stringutil.Reverse(greeting))
}
