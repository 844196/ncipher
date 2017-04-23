package ncipher_test

import (
	"fmt"
	"github.com/844196/ncipher/lib"
)

func Example() {
	opts := ncipher.Options{
		Seed:      "にゃんぱす",
		Delimiter: "〜",
	}
	conv, _ := ncipher.NewConverter(&opts)

	target := "わたしはサーバルキャットのサーバル！"
	encoded := conv.Encode(target)

	fmt.Println(encoded)
	// Output: ぱすすんゃゃ〜ぱすすにゃぱ〜ぱすすににに〜ぱすすにすす〜ぱすすぱぱす〜すににゃぱに〜ぱすすすすゃ〜すにににすぱ〜ぱすすぱんゃ〜すにににぱに〜ぱすすすゃぱ〜ぱすすすんぱ〜ぱすすにすぱ〜ぱすすぱぱす〜すににゃぱに〜ぱすすすすゃ〜すにににすぱ〜すにすんゃゃゃ
}

func ExampleNewConverter() {
	opts := ncipher.Options{
		Seed:      "0123456789ABCDEF",
		Delimiter: ".",
	}
	conv, _ := ncipher.NewConverter(&opts)

	fmt.Printf("%T\n", conv)
	// Output: *ncipher.Converter
}

func ExampleConverter_Encode() {
	opts := ncipher.Options{
		Seed:      "0123456789ABCDEF",
		Delimiter: ".",
	}
	conv, _ := ncipher.NewConverter(&opts)

	target := "0123" // [U+0030 U+0031 U+0032 U+0033]
	result := conv.Encode(target)

	fmt.Println(result)
	// Output: 30.31.32.33
}

func ExampleConverter_Decode() {
	opts := ncipher.Options{
		Seed:      "0123456789ABCDEF",
		Delimiter: ".",
	}
	conv, _ := ncipher.NewConverter(&opts)

	target := "30.31.32.33"
	result := conv.Decode(target)

	fmt.Println(result)
	// Output: 0123
}
