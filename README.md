<h1>
  ncipher
  <a href="https://circleci.com/gh/844196/ncipher">
    <img src="https://circleci.com/gh/844196/ncipher.svg?style=shield&circle-token=eb43591a571f24a883128f0c4bf7b68776897ac3">
  </a>
  <a href="https://goreportcard.com/report/github.com/844196/ncipher">
    <img src="https://goreportcard.com/badge/github.com/844196/ncipher">
  </a>
  <a href="https://godoc.org/github.com/844196/ncipher">
    <img src="https://godoc.org/github.com/844196/ncipher?status.svg">
  </a>
  <a href="https://github.com/844196/ncp_spec">
    <img src="https://img.shields.io/badge/spec-844196/ncp__spec-blue.svg?style=flat&colorB=5272B4">
  </a>
</h1>

ncipher provides "Nyanpasu Cipher (N-Cipher)" encoder & decoder

## Usage

```go
package main

import (
	"fmt"

	"github.com/844196/ncipher"
)

func main() {
	src := "われわれはかしこいので"
	out := ncipher.StdEncoding.Encode(src)

	fmt.Println(out)
	// Output: ぱすすんゃゃ〜ぱすすんにぱ〜ぱすすんゃゃ〜ぱすすんにぱ〜ぱすすにすす〜ぱすぱすんぱ〜ぱすすににに〜ぱすぱすすゃ〜ぱすぱすゃゃ〜ぱすすにすぱ〜ぱすすにぱゃ〜
}
```

```go
package main

import (
	"fmt"

	"github.com/844196/ncipher"
)

func main() {
	enc, _ := ncipher.NewEncoding(ncipher.Config{
		Seed:      "あいうえおかきくけこさしすせそたちつてとなにぬねの",
		Delimiter: "、",
	})

	src := "ふーbar ほげfuga"

	encoded := enc.Encode(src)
	fmt.Println(encoded)
	// Output: とにか、ないた、えね、えぬ、おそ、いく、とにし、ととな、おう、おつ、おえ、えぬ、

	decoded, _ := enc.Decode(encoded)
	fmt.Println(decoded)
	// Output: ふーbar ほげfuga
}
```

## Installation

```console
$ go get github.com/844196/ncipher
```
