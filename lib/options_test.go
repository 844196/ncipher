package ncipher

import (
	"testing"
)

func TestSeedSize(t *testing.T) {
	opts := Options{
		Seed:      "あbうdお",
		Delimiter: "-",
	}

	expected := 5
	actual := opts.seedSize()

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestEncodeMap(t *testing.T) {
	opts := Options{
		Seed:      "あbうdお",
		Delimiter: "-",
	}

	expected := map[string]string{
		"0": "あ",
		"1": "b",
		"2": "う",
		"3": "d",
		"4": "お",
	}
	actual := opts.encodeMap()

	for k, v := range expected {
		if e, a := v, actual[k]; a != e {
			t.Errorf("expected: %v, actual: %v", e, a)
		}
	}
}

func TestDecodeMap(t *testing.T) {
	opts := Options{
		Seed:      "あbうdお",
		Delimiter: "-",
	}

	expected := map[string]string{
		"あ": "0",
		"b": "1",
		"う": "2",
		"d": "3",
		"お": "4",
	}
	actual := opts.decodeMap()

	for k, v := range expected {
		if e, a := v, actual[k]; a != e {
			t.Errorf("expected: %v, actual: %v", e, a)
		}
	}
}
