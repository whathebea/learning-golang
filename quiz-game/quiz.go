package main

import "fmt"

func main() {
	fmt.Println("Hey! What's your name?")

	var name string
	
	fmt.Scan(&name)

	fmt.Printf("Hey %v, let's test your knowledge with a quiz!", name)
}
