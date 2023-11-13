/*
	 There are two kind of package in GO
		1 - executable
		2 - reuseable (something that can be used as a library)

		executable must have a "main" function inside them

		When we write "package main" that means when we compile it (go build main.go), it will result in main.exe file
		There "package main" is a keyword

		If we write anything else like "package apple" that will result into nothing when using command "go build main.go"
*/
package main

/*
	One important distinction to make here between package and import
	package will lead to creattion of "package" whereas import is for using the package
	writing "package main" on top means that the current file is part of the package main

	package can have multiple files in them it is analogus to python folder
*/

import "fmt"

func main() {
	fmt.Println("Hi there!")
}