package ncipher

import (
	"testing"
)

func TestEncode(t *testing.T) {
	opts := Options{
		Seed:      "0123456789ABCDEF",
		Delimiter: ".",
	}
	conv, _ := NewConverter(&opts)

	target := "0123" // [U+0030 U+0031 U+0032 U+0033]
	expected := "30.31.32.33"
	actual := conv.Encode(target)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	opts := Options{
		Seed:      "0123456789ABCDEF",
		Delimiter: ".",
	}
	conv, _ := NewConverter(&opts)

	target := "30.31.32.33"
	expected := "0123" // [U+0030 U+0031 U+0032 U+0033]
	actual := conv.Decode(target)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
