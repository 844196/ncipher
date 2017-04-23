package ncipher

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

type Converter struct {
	opts *Options
}

func (c *Converter) convert(target *string, convertMap map[string]string) {
	for f, t := range convertMap {
		*target = strings.Replace(*target, f, t, -1)
	}
}

func (c *Converter) Encode(origin string) (encoded string) {
	rdx := c.opts.seedSize()
	b, i := make([]string, utf8.RuneCountInString(origin)), 0
	for _, r := range origin {
		b[i] = strconv.FormatInt(int64(r), rdx)
		i++
	}

	encoded = strings.Join(b, string(c.opts.Delimiter))
	c.convert(&encoded, c.opts.encodeMap())

	return
}

func (c *Converter) Decode(origin string) (decoded string) {
	c.convert(&origin, c.opts.decodeMap())

	rdx := c.opts.seedSize()
	b := strings.Split(origin, c.opts.Delimiter)
	for i := 0; i < len(b); i++ {
		cp, _ := strconv.ParseInt(b[i], rdx, 0)
		b[i] = string(rune(cp))
	}

	decoded = strings.Join(b, "")

	return
}
