package ncipher

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type Options struct {
	Seed      string
	Delimiter string
}

func NewOptions(seed, delimiter string) (*Options, error) {
	opts := &Options{
		Seed:      seed,
		Delimiter: delimiter,
	}

	if err := opts.ensureValid(); err != nil {
		return nil, err
	}

	return opts, nil
}

func (opts *Options) ensureValid() error {
	if opts.seedSize() < 2 {
		return fmt.Errorf("seed to short")
	}

	if opts.seedSize() > 36 {
		return fmt.Errorf("seed to long")
	}

	encounter := make(map[rune]int, opts.seedSize())
	for _, r := range opts.Seed {
		encounter[r] += 1

		if encounter[r] != 1 {
			return fmt.Errorf("duplicate seed char: %s", r)
		}
	}

	if utf8.RuneCountInString(opts.Delimiter) < 1 {
		return fmt.Errorf("delimiter too short")
	}

	for _, dr := range opts.Delimiter {
		for _, sr := range opts.Seed {
			if dr == sr {
				return fmt.Errorf("has include seed in delimiter: %s", dr)
			}
		}
	}

	return nil
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
