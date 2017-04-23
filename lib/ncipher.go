// Package ncipher provides "Nyanpasu Cipher (N-Cipher)" encode & decode library.
package ncipher

func NewConverter(opts *Options) (*Converter, error) {
	return &Converter{opts: opts}, nil
}
