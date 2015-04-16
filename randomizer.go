package randomizer

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

const (
	_SET_CAPITAL = "ABCDEFGHIJKLMNOPQRSTUVWQYZ"
	_SET_SMALL   = "abcdefghijklmnopqrstuvwqyz"
	_SET_NUMERIC = "01234567890"
	NUMERIC      = 1
	SMALL        = 1 << NUMERIC
	CAPITAL      = 1 << SMALL
)

//Random
//Parameters desired length, and mask, const CAPITAL, SMALL, NUMERIC
func Random(length int, mask uint64) (string, error) {
	if length < 1 {
		return "", errors.New("Invalid length")
	}

	rand.Seed(time.Now().Unix())

	seedbox := _populate(length, mask)
	output := ""
	for i := 0; i < length; i++ {
		output += _randomize(seedbox)
	}

	return output, nil
}

//RandomInt output is int64
//Parameters desired length, max length = 19, max random = (10 pow 20) - 1
func RandomInt(length int) (int64, error) {
	if length < 1 || length > 19 {
		return 0, errors.New("Invalid length")
	}

	rand.Seed(time.Now().Unix())

	min := math.Pow10(length - 1)
	max := math.Pow10(length) - 1

	return int64(min) + rand.Int63n(int64(max-min)), nil
}

func _populate(length int, mask uint64) string {
	seedbox := ""

	if mask&NUMERIC > 0 {
		seedbox += _SET_NUMERIC
	}

	if mask&SMALL > 0 {
		seedbox += _SET_SMALL
	}

	if mask&CAPITAL > 0 {
		seedbox += _SET_CAPITAL
	}

	return seedbox
}

func _randomize(seedbox string) string {
	if len(seedbox) < 1 {
		return ""
	}

	return string([]byte(seedbox)[rand.Intn(len(seedbox)-1)])
}
