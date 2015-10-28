package randomizer

import (
	"math"
	"testing"
)

func TestRandom(t *testing.T) {
	pronounce := false

	//Length 0
	v, err := Random(0, NUMERIC|SMALL|CAPITAL|SYMBOL, pronounce)
	if err == nil {
		t.Error("Length cannot be zero")
	}

	length, _ := RandomMinMax(1, 18)
	//Pronounce false
	v, err = Random(int(length), NUMERIC|SMALL|CAPITAL|SYMBOL, pronounce)
	if err != nil {
		t.Error(err.Error())
	}

	//Pronounce true
	pronounce = true
	v, err = Random(int(length), NUMERIC|SMALL|CAPITAL|SYMBOL, pronounce)
	if err != nil {
		t.Error(err.Error())
	}

	if len(v) != int(length) {
		t.Errorf("Expected %d, got %d", length, len(v))
	}

	t.Log(v)
}

func TestRandomInt(t *testing.T) {
	//Length 0
	v, err := RandomInt(0)
	if err == nil {
		t.Error("Length cannot be zero")
	}

	length, _ := RandomMinMax(1, 18)
	v, err = RandomInt(int(length))
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
	v, err := RandomMinMax(math.MaxInt64, math.MinInt64)
	if err == nil {
		t.Error("Boundary is Integer 64")
	}

	min, _ := RandomMinMax(-20, 0)
	max, _ := RandomMinMax(1, 20)

	v, err = RandomMinMax(min, max)
	if err != nil {
		t.Error(err.Error())
	}

	if min > v || max < v {
		t.Errorf("Min %d Max %d got %d", min, max, v)
	}

	t.Log(v)
}
