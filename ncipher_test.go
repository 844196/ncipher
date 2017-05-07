package ncipher

import (
	"testing"
)

func TestNewEncoding(t *testing.T) {
	if _, err := NewEncoding(&Config{Seed: "世", Delimiter: "-"}); err == nil {
		t.Error("expected error: seed to short")
	}

	s := "0123456789abcdefghijklmnopqrstuvwxyzA"
	if _, err := NewEncoding(&Config{Seed: s, Delimiter: "-"}); err == nil {
		t.Error("expected error: seed to long")
	}

	if _, err := NewEncoding(&Config{Seed: "世世", Delimiter: "-"}); err == nil {
		t.Error("expected error: duplicate seed char")
	}

	if _, err := NewEncoding(&Config{Seed: "世界", Delimiter: ""}); err == nil {
		t.Error("expected error: delimiter too short")
	}

	if _, err := NewEncoding(&Config{Seed: "世界", Delimiter: "世"}); err == nil {
		t.Error("expected error: has include seed in delimiter")
	}

	if _, err := NewEncoding(&Config{Seed: "'こんにちは", Delimiter: "世界"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestEncode(t *testing.T) {
	enc := StdEncoding

	in := "みゃ！"
	expected := "ぱすすゃぱに〜ぱすすゃぱす〜すにすんゃゃゃ〜"
	actual := enc.Encode(in)

	if actual != expected {
		t.Errorf("expected: %q, got: %q", expected, actual)
	}
}

func TestEncodeMultipleline(t *testing.T) {
	enc, err := NewEncoding(&Config{
		Seed:      "0123456789abcdef",
		Delimiter: ".",
	})
	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	in := "あのイーハトーヴォのすきとおった風\n夏でも底に冷たさをもつ青いそら\nうつくしい森で飾られたモリーオ市\n郊外のぎらぎらひかる草の波"
	expected := "3042.306e.30a4.30fc.30cf.30c8.30fc.30f4.30a9.306e.3059.304d.3068.304a.3063.305f.98a8.a.590f.3067.3082.5e95.306b.51b7.305f.3055.3092.3082.3064.9752.3044.305d.3089.a.3046.3064.304f.3057.3044.68ee.3067.98fe.3089.308c.305f.30e2.30ea.30fc.30aa.5e02.a.90ca.5916.306e.304e.3089.304e.3089.3072.304b.308b.8349.306e.6ce2."
	actual := enc.Encode(in)

	if actual != expected {
		t.Errorf("expected: %q, got: %q", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	enc := StdEncoding

	in := "ぱすすゃぱに〜ぱすすゃぱす〜すにすんゃゃゃ〜"
	expected := "みゃ！"
	actual, err := enc.Decode(in)

	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	if actual != expected {
		t.Errorf("expected: %q, got: %q", expected, actual)
	}
}

func TestDecodeMultipleline(t *testing.T) {
	enc, err := NewEncoding(&Config{
		Seed:      "0123456789abcdef",
		Delimiter: ".",
	})
	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	in := "3042.306e.30a4.30fc.30cf.30c8.30fc.30f4.30a9.306e.3059.304d.3068.304a.3063.305f.98a8.a.590f.3067.3082.5e95.306b.51b7.305f.3055.3092.3082.3064.9752.3044.305d.3089.a.3046.3064.304f.3057.3044.68ee.3067.98fe.3089.308c.305f.30e2.30ea.30fc.30aa.5e02.a.90ca.5916.306e.304e.3089.304e.3089.3072.304b.308b.8349.306e.6ce2."
	expected := "あのイーハトーヴォのすきとおった風\n夏でも底に冷たさをもつ青いそら\nうつくしい森で飾られたモリーオ市\n郊外のぎらぎらひかる草の波"
	actual, err := enc.Decode(in)

	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	if actual != expected {
		t.Errorf("expected: %q, got: %q", expected, actual)
	}
}
