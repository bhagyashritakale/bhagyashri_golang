package main

import "fmt"

func main() {
	var p int = 10
	var q int = 20
	// Simple Assignment
	p = q
	fmt.Println(p)

	//Add Assignment
	p += q
	fmt.Println(p)

	//subtract Assignment
	p -= q
	fmt.Println(p)

	//Multiply Assignment
	p *= q
	fmt.Println(p)

	//Divide Assignment
	p /= q
	fmt.Println(p)

	//Modulus Assignment
	p %= q
	fmt.Println(p)

}
