package randomizer_test

import (
	"math"
	"testing"

	r "github.com/kakilangit/randomizer"
)

func TestRandom(t *testing.T) {
	pronounce := false

	//Length 0
	v, err := r.Random(0, r.Numeric|r.Small|r.Capital|r.Symbol, pronounce)
	if err == nil {
		t.Error("Length cannot be zero")
	}

	length, _ := r.RandomMinMax(1, 18)
	//Pronounce false
	v, err = r.Random(int(length), r.Numeric|r.Small|r.Capital|r.Symbol, pronounce)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(length, v)

	//Pronounce true
	pronounce = true
	v, err = r.Random(int(length), r.Numeric|r.Small|r.Capital|r.Symbol, pronounce)
	if err != nil {
		t.Error(err.Error())
	}

	if len(v) != int(length) {
		t.Errorf("Expected %d, got %d", length, len(v))
	}

	t.Log(length, v)
}

func TestRandomInt(t *testing.T) {
	//Length 0
	v, err := r.RandomInt(0)
	if err == nil {
		t.Error("Length cannot be zero")
	}

	length, _ := r.RandomMinMax(1, 18)
	v, err = r.RandomInt(int(length))
	if err != nil {
		t.Error(err.Error())
	}

	outputLength := int(math.Floor(math.Log10(float64(v))) + 1)

	if outputLength != int(length) {
		t.Errorf("Expected %d, got %d", length, outputLength)
	}

	t.Log(v)
}

func TestRandomMinMax(t *testing.T) {
	//Min > Max
	//Boundary overflow
	v, err := r.RandomMinMax(math.MaxInt64, math.MinInt64)
	if err == nil {
		t.Error("Boundary is Integer 64")
	}

	min, _ := r.RandomMinMax(-20, 0)
	max, _ := r.RandomMinMax(1, 20)

	v, err = r.RandomMinMax(min, max)
	if err != nil {
		t.Error(err.Error())
	}

	if min > v || max < v {
		t.Errorf("Min %d Max %d got %d", min, max, v)
	}

	t.Log(v)
}

func BenchmarkRandomNumeric10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Numeric)
	}
}

func BenchmarkRandomSmall10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Small)
	}
}

func BenchmarkRandomCapital10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Capital)
	}
}

func BenchmarkRandomSymbol10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Symbol)
	}
}

func BenchmarkRandomNumericSmall10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Numeric|r.Small)
	}
}

func BenchmarkRandomNumericCapital10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Numeric|r.Capital)
	}
}

func BenchmarkRandomNumericSymbol10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Numeric|r.Symbol)
	}
}

func BenchmarkRandomSmallCapital10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Small|r.Capital)
	}
}

func BenchmarkRandomSmallSymbol10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Small|r.Symbol)
	}
}

func BenchmarkRandomCapitalSymbol10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Capital|r.Symbol)
	}
}

func BenchmarkRandomNumericSmallCapital10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Numeric|r.Small|r.Capital)
	}
}

func BenchmarkRandomNumericSmallSymbol10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Numeric|r.Small|r.Symbol)
	}
}

func BenchmarkRandomNumericCapitalSymbol10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Numeric|r.Capital|r.Symbol)
	}
}

func BenchmarkRandomSmallCapitalSymbol10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Small|r.Capital|r.Symbol)
	}
}

func BenchmarkRandom10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.Random(10, r.Numeric|r.Small|r.Capital|r.Symbol)
	}
}
