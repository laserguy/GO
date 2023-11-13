package main

import "fmt"

type bot interface{
	/*
		This function declaration inside the interface means that any function within the program
		which has a same declaration will become an honoary member of type 'bot'.
		Manual link is not required: That's why it is implicit, Golang automatically figures out what
		variables become member of 'bot'
	*/
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}


func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "English"
}

/*
	The difference between the two function declaration, is in the way the reciever is written
	(englishBot) => It can be written in this way, when the variable name is not used within the function
	(sb spanishBot)
*/

func (sb spanishBot) getGreeting() string {
	return "Spanish"
}