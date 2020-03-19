package main

import "fmt"

func main() {
	fmt.Println("WELCOME!!!")
	fmt.Println("This Application allows you to convert any Integer to it's binary value!")
	fmt.Println("Kindly Input the Number: ")

	var Integer int;

	fmt.Scan(&Integer);

	fmt.Printf("The Binary representation of the value %d is: %b", Integer, Integer)

}