package main

import (
	"fmt"
)

func foo() {
	for i := 0; i <= 50; i++ {
		fmt.Println(i)
	}
}

func bar() {
	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("%c\n", ch)
	}
}

func main() {
	foo()
	bar()
}