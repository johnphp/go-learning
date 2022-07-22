package main

import "fmt"

func main() {
	fmt.Println("This is about Structs in Go")

	//declaring a struct type
	type person struct {
		name string
		age  int
		pet  string
	}

	//once a struct type is declared, define a variable of that type
	var fred person
	//since no value is assigned, the variable gets zero value.
	//a zero value struct has every field set to the field's zero value
	fmt.Println(fred)

	//assigning a struct literal to a variable
	bob := person{}
	fmt.Println(bob)

	//assigning non-empty struct literal
	//style #1
	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println(julia)

	//style #2
	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println(beth)

}
