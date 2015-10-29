package randomizer

import (
	"errors"
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	_SET_VOWEL             = "aiueo"
	_SET_CONSONANT         = "bcdfghjklmnpqrstvwqyz"
	_SET_NUMERIC_VOWEL     = "0134"
	_SET_NUMERIC_CONSONANT = "256789"
	_SET_SYMBOL            = "~!@#$%^&*()_+`=-{}|[]\\;':\",./<>?"
	_SEED_VOWEL            = 1
	_SEED_CONSONANT        = 2
	_SEED_ALL              = 3
	NUMERIC                = 1
	SMALL                  = 1 << NUMERIC
	CAPITAL                = 1 << SMALL
	SYMBOL                 = 1 << CAPITAL
)

type Character struct {
	Vocal, Consonant int
	mux              *sync.Mutex
}

//Random
//Parameters desired length, and mask, const CAPITAL, SMALL, NUMERIC, SYMBOL
func Random(length int, mask uint64, args ...interface{}) (string, error) {
	if length < 1 {
		return "", errors.New("Invalid length")
	}

	pronounce := false
	for _, arg := range args {
		switch v := arg.(type) {
		case bool:
			pronounce = v
		}
	}

	output := ""
	strchan := make(chan string, length)
	defer close(strchan)

	if pronounce == true {

		var wg sync.WaitGroup
		wg.Add(length)

		char := &Character{0, 0, &sync.Mutex{}}

		for i := 0; i < length; i++ {
			go char.RandomPronounce(mask, strchan, &wg)
		}

		wg.Wait()

	} else {
		seedbox := _populate(mask, _SEED_ALL)

		for i := 0; i < length; i++ {
			go func() {
				strchan <- _randomize(seedbox)
			}()
		}

	}

	for i := 0; i < length; i++ {
		select {
		case str := <-strchan:
			output += str
		}
	}

	return output, nil
}

//RandomInt output is int64
//Parameters desired length, max length = 19, max random = (10 pow 20) - 1
func RandomInt(length int) (int64, error) {
	if length < 1 || length > 18 {
		return 0, errors.New("Invalid length")
	}

	min := int64(math.Pow10(length - 1))
	max := int64(math.Pow10(length) - 1)

	return RandomMinMax(min, max)
}

//RandomMinMax output is int64
//Parameters min, max
func RandomMinMax(min, max int64) (int64, error) {
	if min > max {
		min, max = max, min
	}

	if min <= math.MinInt64 || max >= math.MaxInt64 {
		return 0, errors.New("Invalid parameter(s)")
	}

	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min), nil
}

func (char *Character) RandomPronounce(mask uint64, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	seedtype, _ := RandomMinMax(_SEED_VOWEL, _SEED_CONSONANT)

	//rule max consonant and vowel is 2
	if seedtype == _SEED_VOWEL {
		char.mux.Lock()
		char.Vocal += 1
		char.Consonant = 0
		char.mux.Unlock()
	}

	if seedtype == _SEED_CONSONANT {
		char.mux.Lock()
		char.Vocal = 0
		char.Consonant += 1
		char.mux.Unlock()

	}

	if char.Vocal > 2 {
		char.mux.Lock()
		char.Vocal = 0
		char.Consonant = 1
		char.mux.Unlock()

		seedtype = _SEED_CONSONANT
	}

	if char.Consonant > 2 {
		char.mux.Lock()
		char.Vocal = 1
		char.Consonant = 0
		char.mux.Unlock()

		seedtype = _SEED_VOWEL
	}

	seedbox := _populate(mask, seedtype)
	ch <- _randomize(seedbox)
}

func _populate(mask uint64, seedtype int64) string {
	seedbox := ""

	if mask&NUMERIC > 0 {
		switch seedtype {
		case _SEED_VOWEL:
			seedbox += _SET_NUMERIC_VOWEL
		case _SEED_CONSONANT:
			seedbox += _SET_NUMERIC_CONSONANT
		case _SEED_ALL:
			seedbox += _SET_NUMERIC_CONSONANT + _SET_NUMERIC_VOWEL
		}
	}

	if mask&SMALL > 0 {
		switch seedtype {
		case _SEED_VOWEL:

			seedbox += _SET_VOWEL
		case _SEED_CONSONANT:
			seedbox += _SET_CONSONANT
		case _SEED_ALL:
			seedbox += _SET_VOWEL + _SET_CONSONANT
		}

	}

	if mask&CAPITAL > 0 {
		switch seedtype {
		case _SEED_VOWEL:
			seedbox += strings.ToUpper(_SET_VOWEL)
		case _SEED_CONSONANT:
			seedbox += strings.ToUpper(_SET_CONSONANT)
		case _SEED_ALL:
			seedbox += strings.ToUpper(_SET_VOWEL + _SET_CONSONANT)
		}
	}

	if (mask&SYMBOL > 0) && (seedtype == _SEED_ALL) {
		seedbox += _SET_SYMBOL
	}

	return seedbox
}

func _randomize(seedbox string) string {
	if len(seedbox) < 1 {
		return ""
	}

	r, e := RandomMinMax(0, int64(len(seedbox)-1))
	if e != nil {
		return ""
	}

	return string([]byte(seedbox)[r])
}
