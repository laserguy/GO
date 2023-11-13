package main

import (
	"fmt"
	"io"
	"os"
)

type logwriter struct{}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Usage: go run main.go <filename>.go")
	}

	filename := os.Args[1]
	fp, error := os.Open(filename)
	if error == nil {
		lw := logwriter{}
		io.Copy(lw, fp)
	}else{
		fmt.Println("Error opening file: ", error)
	}

}

func (logwriter) Write(bs []byte) (int, error) {
	fmt.Println((string(bs)))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}