package luhn

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	prefix := "000000"
	num := New().Generate(16, prefix)
	if !Verify(num) {
		t.Error("invalid number")
	}
	fmt.Println(num)
}

func BenchmarkGenerator_Generate(b *testing.B) {
	gen := New()
	for i := 0; i < b.N; i++ {
		gen.Generate(16)
	}
}

func BenchmarkGenerator_GenerateWithPrefix(b *testing.B) {
	gen := New()
	for i := 0; i < b.N; i++ {
		gen.Generate(16, "000000")
	}
}

func BenchmarkGenerator_GenerateWithPrefixAndVerify(b *testing.B) {
	gen := New()
	for i := 0; i < b.N; i++ {
		num := gen.Generate(16, "000000")
		Verify(num)
	}
}
