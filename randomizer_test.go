package randomizer

import (
	"math"
	"testing"
)

func TestRandom(t *testing.T) {
	length := 10
	v, err := Random(length, NUMERIC|SMALL|CAPITAL)
	if err != nil {
		t.Error(err.Error())
	}

	if len(v) != length {
		t.Errorf("Expected %d, got %d", length, len(v))
	}

	t.Log(v)
}

func TestRandomInt(t *testing.T) {
	length := 15
	v, err := RandomInt(length)
	if err != nil {
		t.Error(err.Error())
	}

	outputLength := int(math.Floor(math.Log10(float64(v))) + 1)

	if outputLength != length {
		t.Errorf("Expected %d, got %d", length, outputLength)
	}

	t.Log(v)
}
