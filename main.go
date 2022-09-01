package main

import "log"

func main() {

	myString := "Green"
	log.Println("My string is set to ", myString)
	changeUsingPointer(&myString)
	log.Println("My string is set to ", myString)
}
func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue

}
