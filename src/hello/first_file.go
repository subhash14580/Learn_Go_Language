package main

import (
	"fmt"
	"strconv"
)

var I int = 1

const (
	x1 = iota
	x2 = iota
	x3 = iota
)

func main() {

	var i int
	i = 42
	//i = 30

	j := 50
	fmt.Println("This is the first time i am running the code.")
	fmt.Println(i + j)
	fmt.Println(i+j == 90)

	fmt.Printf(" %v, %T \n", I, i)

	var theUrl string = "https://www.google.com"

	fmt.Println(theUrl, " any")
	fmt.Println(theUrl, " 2any")

	var convertedString string
	convertedString = strconv.Itoa(j)

	fmt.Printf("converted string is %v and type is %T \n", convertedString, convertedString)

	// agenda Boolean Values

	n := 1

	var result bool

	result = n == 1

	fmt.Println(result)

	// numeric Values int8 --> -128 to 128
	// int 16, int32, int64
	// uint16, uint32

	a := 10
	b := 12
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b) //  not producing float values
	fmt.Println(a * b)
	fmt.Println(b % a)

	a = 0
	b = 1
	//  Binary Operators
	fmt.Println("Binary Operations")
	fmt.Println(a | b)
	fmt.Println(a & b)
	fmt.Println(a ^ b)
	fmt.Println(b &^ a)

	// Shift Operations

	a = 8

	fmt.Println(a << 1)
	fmt.Println(a >> 1)

	n1 := 3.14
	n1 = 13.7e72
	n1 = 2.1e9
	fmt.Printf("%v %T \n", n1/2, n1)

	// complex numbers

	// var complex complex64 = 2i
	var complex complex64 = complex(2, 1222)
	fmt.Printf("%v %v %T \n", real(complex), imag(complex), complex)

	// Strings of text -- strings are immuatable

	s := "Hello I am Learning Go Programming Language"
	s1 := " <<<This is new String>>>"
	fmt.Printf(" %v , %T \n", s+s1, s)

	// Byte Slices of strings
	s3 := []byte(s)
	fmt.Printf(" %v , %T \n", s3, s3)

	// characters  Rune is integer is 32

	c := 'a'

	fmt.Printf("%v, %T \n", c, c)

	// Constants
	// 1. Naming Convention
	// 2. Typed Constatns
	// 3. UnTyped Constants
	// 4. Enumerated Constants
	// 5. Enumeration Expressions

	// const keyword is used to define the contant.
	// If first letter is capital then the const will be exported.
	// const is like a final keyword in java.
	// once assigned, cant be reassigned again.

	const myConst int = 42
	var integerA int = 32
	fmt.Printf("%v %T", myConst, myConst)

	r1 := myConst + integerA
	fmt.Println(r1)

	// Enumerated Constants

	fmt.Printf("%v %T\n", x1, x1)
	fmt.Printf("%v %T\n", x2, x2)
	fmt.Printf("%v %T\n", x3, x3)

	const (
		x4 = iota
		x5 = iota
		x6
	)

	fmt.Println(x4)
	fmt.Println(x6) // answer 2, values are automatically assigned

	const (
		_ = iota + 5
		catSpecialist
		dogSpecialist
		pigSpecialist
	)

	const cat = 6
	const dog = 7

	fmt.Println(catSpecialist == cat)
	fmt.Println(dogSpecialist != dog)

	// Arrays
	// 1.Creation
	// 2.Built In Functions
	// 3.Working with Arrays
	// Slices
	// 1.Creation
	// 2.Built In Functions
	// 3.Working with Slices

}
