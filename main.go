package main

import (
	"fmt"
	"github.com/844196/ncipher/lib"
)

func main() {
	opts := ncipher.Options{
		Seed:      "そぞ",
		Delimiter: "\n",
	}
	conv, _ := ncipher.NewConverter(&opts)

	origin := "にゃんぱす"
	encoded := conv.Encode(origin)
	decoded := conv.Decode(encoded)

	fmt.Printf("平文:\n%s\n\n", origin)
	fmt.Printf("暗号化:\n%s\n\n", encoded)
	fmt.Printf("復号化:\n%s\n\n", decoded)
}
