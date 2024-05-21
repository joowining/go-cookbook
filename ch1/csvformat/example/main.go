package main

import (
	"fmt"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter1/csvformat"
)

func main() {
	if err := csvformat.AddMoviesFromText(); err != nil {
		return
	}

	if err := csvformat.WriteCSVOutput(); err != nil {
		return
	}

	buffer, err := csvformat.WriteCSVBuffer()
	if err != nil {
		return
	}

	fmt.Println("Buffer = ", buffer.String())

}
