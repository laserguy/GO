package main

func main() {
	/*
		type 'deck' is defined in the file 'deck.go', which is also in the package 'main'
		therefore it is not required to include deck.go in thes file
		But while running command should be "go run main.go deck.go"
	*/
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
