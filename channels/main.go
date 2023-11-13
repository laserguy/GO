package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://twitter.com",
		"http://github.com",
		"http://stackoverflow.com",
		"http://go.dev",
	}

	/*
		'chan' is a channel => channels are used do communication between different go routines(threads)
		Every channel is 'typed', data of that specific type can only be sent through that channel
	*/
	c := make(chan string)

	for _, link := range links {
		/* 
			'go' keyword is only used in front of the function, it used for the creating a child routine
			"routine" is same as thread
			'go' creates a child routine
		*/
		go checkLink(link, c)
	}

	for { // This is a infinite loop
		fmt.Println(<-c,c)  // => recieving info from the channel
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil{
		fmt.Println(link, "might be down")
		c <- link // passing info through the channel
		return
	}
	fmt.Println(link, "is up!!!")
	c <- link
}