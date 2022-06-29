package main

import "fmt"

func main() {
	fmt.Println("Hey! What's your name?")

	var name string

	fmt.Scan(&name)

	fmt.Printf("Hey %v, let's test your knowledge with a quiz!", name)

	fmt.Printf("\nInside of which HTML element do we put javascript? ")

	var answer string
	fmt.Scan(&answer)

	if answer == "<script>" {
		fmt.Println("You're right!")
	} else {
		fmt.Println("You're wrong :C")
	}

	fmt.Printf("\n Next! Where is the correct place to insert JavaScript? ")
	var answer2 string
	fmt.Scan(&answer2)

	if answer2 == "<body>" {
		fmt.Println("You're right!")
	} else {
		fmt.Println("You're wrong :C")
	}
}
