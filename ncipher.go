package ncipher

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	SEED_MIN      = 2
	SEED_MAX      = 36
	DELIMITER_MIN = 1
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

func NewEncoding(config *Config) (*Encoding, error) {
	sl := utf8.RuneCountInString(config.Seed)

	if sl < SEED_MIN {
		return nil, errors.New("seed too short")
	}

	if sl > SEED_MAX {
		return nil, errors.New("seed too long")
	}

	if utf8.RuneCountInString(config.Delimiter) < DELIMITER_MIN {
		return nil, errors.New("delimiter too short")
	}

	for _, r := range config.Seed {
		if s := string(r); strings.Count(config.Seed, s) != 1 {
			return nil, fmt.Errorf(fmt.Sprintf("duplicate seed char \"%s\"", s))
		}
	}

	if isDuplicate := strings.ContainsAny(config.Seed, config.Delimiter); isDuplicate {
		return nil, errors.New("has include seed char in delimiter")
	}

	return &Encoding{c: config}, nil
}

func (c *Encoding) Encode(src string) (dst string) {
	seedL := utf8.RuneCountInString(c.c.Seed)

	sub := make([]string, utf8.RuneCountInString(src))
	i := 0
	for _, srcR := range src {
		sub[i] = strconv.FormatInt(int64(srcR), seedL)
		i++
	}
	total := strings.Join(sub, c.c.Delimiter)

	j, k, l := 0, 0, 1
	pair := make([]string, seedL*2)
	for _, seedR := range c.c.Seed {
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

func (c *Encoding) Decode(src string) (dst string, err error) {
	seedL := utf8.RuneCountInString(c.c.Seed)

	cipherCharSet := c.c.Seed + c.c.Delimiter
	for _, srcR := range src {
		if false == strings.ContainsAny(string(srcR), cipherCharSet) {
			return dst, errors.New("invalid cipher string")
		}
	}

	i, j, k := 0, 0, 1
	pair := make([]string, seedL*2)
	for _, seedR := range c.c.Seed {
		pair[j] = string(seedR)
		pair[k] = basemap[i]
		i++
		j += 2
		k += 2
	}

	replacer := strings.NewReplacer(pair...)
	sub := strings.Split(replacer.Replace(src), c.c.Delimiter)

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
