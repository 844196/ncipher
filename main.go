package main

import (
	"fmt"
	"github.com/844196/ncipher/ncipher"
)

func main() {
	opts := ncipher.Options{
		Seed:      []rune("01234"),
		Delimiter: []rune("-"),
	}
	conv, _ := ncipher.NewConverter(&opts)

	fmt.Printf(conv.Encode("foo"))
}
