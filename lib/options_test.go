package ncipher

import (
	"fmt"
	"testing"
)

func TestNewOptions(t *testing.T) {
	actual, err := NewOptions("こんにちは", "世界")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if fmt.Sprintf("%T", actual) != "*ncipher.Options" {
		t.Errorf("expected: *ncipher.Options, actual:%T", actual)
	}
}

func TestEnsureValid(t *testing.T) {
	var o *Options

	o = &Options{Seed: "世", Delimiter: "-"}
	if err := o.ensureValid(); err == nil {
		t.Errorf("expected error: seed to short")
	}

	o = &Options{Seed: "世世", Delimiter: "-"}
	if err := o.ensureValid(); err == nil {
		t.Errorf("expected error: duplicate seed char")
	}

	o = &Options{Seed: "世界", Delimiter: ""}
	if err := o.ensureValid(); err == nil {
		t.Errorf("expected error: delimiter too short")
	}

	o = &Options{Seed: "世界", Delimiter: "世"}
	if err := o.ensureValid(); err == nil {
		t.Errorf("expected error: has include seed in delimiter")
	}

	o = &Options{Seed: "こんにちは", Delimiter: "世界"}
	if err := o.ensureValid(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

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
