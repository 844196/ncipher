package ncipher

import (
	"strconv"
	"unicode/utf8"
)

type Options struct {
	Seed      string
	Delimiter string
}

func (opts *Options) seedSize() int {
	return utf8.RuneCountInString(opts.Seed)
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
	m, i := make(map[string]string, opts.seedSize()), 0
	for _, c := range opts.Seed {
		m[string(c)] = strconv.Itoa(i)
		i++
	}

	return m
}
