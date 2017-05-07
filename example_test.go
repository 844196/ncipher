package ncipher_test

import (
	"fmt"

	"github.com/844196/ncipher"
)

func Example() {
	in := "われわれはかしこいので"
	out := ncipher.StdEncoding.Encode(in)

	fmt.Println(out)
	// Output: ぱすすんゃゃ〜ぱすすんにぱ〜ぱすすんゃゃ〜ぱすすんにぱ〜ぱすすにすす〜ぱすぱすんぱ〜ぱすすににに〜ぱすぱすすゃ〜ぱすぱすゃゃ〜ぱすすにすぱ〜ぱすすにぱゃ〜
}

func ExampleEncoding_Encode() {
	cnf := ncipher.Config{
		Seed:      "あいうえおかきくけこさしすせそたちつてとなにぬねの",
		Delimiter: "、",
	}
	enc, _ := ncipher.NewEncoding(&cnf)

	in := "ふーbar ほげfuga"
	out := enc.Encode(in)

	fmt.Println(out)
	// Output: とにか、ないた、えね、えぬ、おそ、いく、とにし、ととな、おう、おつ、おえ、えぬ、
}

func ExampleEncoding_Decode() {
	cnf := ncipher.Config{
		Seed:      "あいうえおかきくけこさしすせそたちつてとなにぬねの",
		Delimiter: "、",
	}
	enc, _ := ncipher.NewEncoding(&cnf)

	in := "とにか、ないた、えね、えぬ、おそ、いく、とにし、ととな、おう、おつ、おえ、えぬ、"
	out, _ := enc.Decode(in)

	fmt.Println(out)
	// Output: ふーbar ほげfuga
}
