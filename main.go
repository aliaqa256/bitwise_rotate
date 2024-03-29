package main

import (
	"errors"
	"fmt"
)

func main() {
	number := gettingInput()
	L := GetlenBinary(number)
	ShowDetail(number, L)

	fmt.Print("for how many times you want to rotate ")
	numberOfRotations := int64(gettingInput())
	// if the number is negative, convert it to positive
	numberOfRotations = if_negative(numberOfRotations)

	var R int64 = Devideable(numberOfRotations, L)

	dirction, err := getRightOrLeftInput()
	if err != nil {
		fmt.Println(err)
	}

	res := Rotate(dirction, number, L, R)

	fmt.Println(res)
	fmt.Printf("%b \n", res)
}

func gettingInput() (number int64) {
	fmt.Printf("Enter a number: ")
	fmt.Scan(&number)
	number = if_negative(number)
	return
}

func if_negative(num int64) int64 {
	if num < 0 {
		num = 0
	}
	return num
}

// NOTE: this funtion returns rune
func getRightOrLeftInput() (rune, error) {
	var direct string
	fmt.Println("select the direction that you want to rotate ?(right ro left/ << or >> )")
	fmt.Scan(&direct)
	switch direct {
	case "right":
		return 'r', nil
	case "left":
		return 'l', nil
	case "<<":
		return 'l', nil
	case "l":
		return 'l', nil
	case ">>":
		return 'r', nil
	case "r":
		return 'r', nil
	default:
		return 'e', errors.New("wrong input")
	}
}

func Devideable(Num, len int64) int64 {
	if Num > len {
		return Num % len
	} else {
		return Num
	}
}

// getRightOrLeftInput is function that gets inputs from user and decides
func Rotate(direction rune, number, len, r int64) int64 {
	switch direction {
	case 'r':
		b := number >> r
		b_masked := b & ((1 << len) - 1)
		c := number << (len - r)
		c_masked := c & ((1 << len) - 1)
		res := b_masked | c_masked
		return res
	case 'l':
		b := number << r
		b_masked := b & ((1 << len) - 1)
		c := number >> (len - r)
		c_masked := c & ((1 << len) - 1)
		res := b_masked | c_masked
		return res

	default:
		return 0
	}
}

func GetlenBinary(num int64) int64 {
	var inlen int64

	for {
		num = num / 2
		inlen++
		if num < 1 {
			break
		}
	}
	return inlen
}

func ShowDetail(number, L int64) {
	fmt.Printf("The number of bits is: [%d]\n", L)
	fmt.Printf("The binary number is: [%b] \n", number)
}
