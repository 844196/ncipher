package ncipher

import (
	"strconv"
)

type Options struct {
	Seed      string
	Delimiter string
}

func (opts *Options) seedSize() int {
	return len([]rune(opts.Seed))
}

func (opts *Options) encodeMap() map[string]string {
	m, i := make(map[string]string, opts.seedSize()), 0
	for _, c := range opts.Seed {
		m[strconv.Itoa(i)] = string(c)
		i++
	}

	return m
}

func (opts *Options) decodeMap() map[string]string {
	m := make(map[string]string, opts.seedSize())
	for t, f := range opts.encodeMap() {
		m[f] = t
	}

	return m
}
