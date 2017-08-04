package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // get value
	fmt.Println(&a) //get address

	var b *int = &a

	fmt.Println(b)  // memory address
	fmt.Println(*b) // dereference... get the actual value
}
