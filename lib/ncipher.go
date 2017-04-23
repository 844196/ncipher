// Package ncipher provides "Nyanpasu Cipher (N-Cipher)" encoder & decoder
package ncipher

func NewConverter(opts *Options) (*Converter, error) {
	return &Converter{opts: opts}, nil
}
