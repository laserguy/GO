package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*	create a new type 'decl'
	which is a slice of strings
	Difference between `slice' and `array`, slice is dynamic whose length can be increased while array is static`
*/
type deck []string

/*
	The below function are called RECIEVER function

	The below function means that any variable of type deck will have access to this function

	func (d deck)  => deck "word" is important here also called "RECIEVER"
*/

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i,card)
	}
}

/*
	This one is not a reciever, this function can be called anywhere in the same package
*/

func newDeck() deck{
	cards := deck{}
	
	cardSuits := []string{"Spades","Diamonds", "Hearts","Clubs"}
	cardValues := []string{"Ace", "Two","Three","Four"}

	/*
		for i, suit   => will give an error that 'i is never used'

		'_' is used to indicate we won't be using that variable
	*/
	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value +" of " + suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck,deck){
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	/*
		writing `[]string(d) => is a type conversion`
	*/
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename,[]byte(d.toString()),0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ",err)
		os.Exit(1)
	}

	s := strings.Split(string(byteSlice),",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}