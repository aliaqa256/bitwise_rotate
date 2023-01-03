package main

import (
	"fmt"
	"strconv"
)

func main() {
	// get the number and convert it to binary
	var number int64
	
	fmt.Printf("Enter a number: ")
	fmt.Scan(&number)
	if number < 0 {
		number = 0
	}

	init := strconv.FormatInt(number, 2)
	fmt.Printf("The binary number is: [%s]\n", init)
	L := len(init)
	fmt.Printf("The number of bits is: [%d]\n", L)

	var numberOfRotations int
	fmt.Printf("Enter the number of rotations: ")
	fmt.Scan(&numberOfRotations)
	// if the number is negative, convert it to positive
		if numberOfRotations < 0 {
		numberOfRotations = 0
	}


	var R int
	if numberOfRotations > L {
		R = int(numberOfRotations) % int(L)
	} else {
		R = numberOfRotations
	}
	b := number >> R
	c := number << (L - R)
	res := b | c
	fmt.Printf("The result is: [%s]\n", strconv.FormatInt(res, 2)[len(strconv.FormatInt(res, 2))-L:])

}
