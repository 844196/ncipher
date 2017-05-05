// Package ncipher provides "Nyanpasu Cipher (N-Cipher)" encoder & decoder
package ncipher

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	SeedMin      = 2
	SeedMax      = 36
	DelimiterMin = 1
)

var (
	basemap        = [36]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	StdConfig      = Config{Seed: "にゃんぱす", Delimiter: "〜"}
	StdEncoding, _ = NewEncoding(&StdConfig)
)

type Config struct {
	Seed, Delimiter string
}

type Encoding struct {
	c *Config
}

func NewEncoding(cnf *Config) (*Encoding, error) {
	sl := utf8.RuneCountInString(cnf.Seed)

	if sl < SeedMin {
		return nil, errors.New("seed too short")
	}

	if sl > SeedMax {
		return nil, errors.New("seed too long")
	}

	if utf8.RuneCountInString(cnf.Delimiter) < DelimiterMin {
		return nil, errors.New("delimiter too short")
	}

	for _, r := range cnf.Seed {
		if s := string(r); strings.Count(cnf.Seed, s) != 1 {
			return nil, fmt.Errorf(fmt.Sprintf("duplicate seed char \"%s\"", s))
		}
	}

	if isDuplicate := strings.ContainsAny(cnf.Seed, cnf.Delimiter); isDuplicate {
		return nil, errors.New("has include seed char in delimiter")
	}

	return &Encoding{c: cnf}, nil
}

func (e *Encoding) Encode(src string) (dst string) {
	seedL := utf8.RuneCountInString(e.c.Seed)

	sub := make([]string, utf8.RuneCountInString(src))
	i := 0
	for _, srcR := range src {
		sub[i] = strconv.FormatInt(int64(srcR), seedL)
		i++
	}
	total := strings.Join(sub, e.c.Delimiter)

	j, k, l := 0, 0, 1
	pair := make([]string, seedL*2)
	for _, seedR := range e.c.Seed {
		pair[k] = basemap[j]
		pair[l] = string(seedR)
		j++
		k += 2
		l += 2
	}

	replacer := strings.NewReplacer(pair...)
	dst = replacer.Replace(total)

	return
}

func (e *Encoding) Decode(src string) (dst string, err error) {
	seedL := utf8.RuneCountInString(e.c.Seed)

	cipherCharSet := e.c.Seed + e.c.Delimiter
	for _, srcR := range src {
		if false == strings.ContainsAny(string(srcR), cipherCharSet) {
			return dst, errors.New("invalid cipher string")
		}
	}

	i, j, k := 0, 0, 1
	pair := make([]string, seedL*2)
	for _, seedR := range e.c.Seed {
		pair[j] = string(seedR)
		pair[k] = basemap[i]
		i++
		j += 2
		k += 2
	}

	replacer := strings.NewReplacer(pair...)
	sub := strings.Split(replacer.Replace(src), e.c.Delimiter)

	for l := 0; l < len(sub); l++ {
		cp, err := strconv.ParseInt(sub[l], seedL, 0)
		if err != nil {
			return dst, err
		}
		sub[l] = string(rune(cp))
	}
	dst = strings.Join(sub, "")

	return
}
