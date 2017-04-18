package ncipher

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

type Options struct {
	Seed      []rune
	Delimiter []rune
}

type converter struct {
	Options *Options
}

func (c *converter) Encode(origin string) (encoded string) {
	return "a"
}

func (c *converter) Decode(origin string) (decoded string) {
	return "a"
}

func NewConverter(opts *Options) (*converter, error) {
	return &converter{Options: opts}, nil
}
