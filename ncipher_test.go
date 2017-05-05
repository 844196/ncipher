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
	expected := "ぱすすゃぱに〜ぱすすゃぱす〜すにすんゃゃゃ"
	actual := enc.Encode(in)

	if actual != expected {
		t.Errorf("expected: %q, got: %q", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	enc := StdEncoding

	if _, err := enc.Decode("aaaa,bbbb,cccc"); err == nil {
		t.Fatal("expected error")
	}

	if _, err := enc.Decode("ぱすすゃぱに〜"); err == nil {
		t.Fatal("expected error")
	}

	in := "ぱすすゃぱに〜ぱすすゃぱす〜すにすんゃゃゃ"
	expected := "みゃ！"
	actual, err := enc.Decode(in)

	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	if actual != expected {
		t.Errorf("expected: %q, got: %q", expected, actual)
	}
}
