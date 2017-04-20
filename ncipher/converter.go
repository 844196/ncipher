package ncipher

import (
	"strconv"
	"strings"
)

type converter struct {
	opts *Options
}

func (c *converter) convert(target *string, convertMap map[string]string) {
	for f, t := range convertMap {
		*target = strings.Replace(*target, f, t, -1)
	}
}

func (c *converter) Encode(origin string) (encoded string) {
	rdx := c.opts.seedSize()
	b, i := make([]string, len([]rune(origin))), 0
	for _, r := range origin {
		b[i] = strconv.FormatInt(int64(r), rdx)
		i++
	}

	encoded = strings.Join(b, string(c.opts.Delimiter))
	c.convert(&encoded, c.opts.encodeMap())

	return
}

func (c *converter) Decode(origin string) (decoded string) {
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

func NewConverter(opts *Options) (*converter, error) {
	return &converter{opts: opts}, nil
}
