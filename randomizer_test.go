package randomizer

import (
	"math"
	"testing"
)

func TestRandom(t *testing.T) {
	length, _ := RandomMinMax(0, 18)
	v, err := Random(int(length), NUMERIC|SMALL|CAPITAL)
	if err != nil {
		t.Error(err.Error())
	}

	if len(v) != int(length) {
		t.Errorf("Expected %d, got %d", length, len(v))
	}

	t.Log(v)
}

func TestRandomInt(t *testing.T) {
	length, _ := RandomMinMax(0, 18)
	v, err := RandomInt(int(length))
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
	min, _ := RandomMinMax(-20, 0)
	max, _ := RandomMinMax(0, 20)

	v, err := RandomMinMax(min, max)
	if err != nil {
		t.Error(err.Error())
	}

	if min > v || max < v {
		t.Errorf("Min %d Max %d got %d", min, max, v)
	}

	t.Log(v)
}
