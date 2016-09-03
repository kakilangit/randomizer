package randomizer

import (
	"errors"
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"
)

//Constant
const (
	setVowel            = "aiueo"
	setConsonant        = "bcdfghjklmnpqrstvwqyz"
	setNumericVowel     = "0134"
	setNumericConsonant = "256789"
	setSymbol           = "~!@#$%^&*()_+`=-{}|[]\\;':\",./<>?"
	seedVowel           = 1
	seedConsonant       = 2
	seedAll             = 3
	Numeric             = 1
	Small               = 1 << Numeric
	Capital             = 1 << Small
	Symbol              = 1 << Capital
)

//Character type
type Character struct {
	Vocal, Consonant int
	mux              *sync.Mutex
}

//Random function
//Parameters desired length, and mask, const Capital, Small, Numeric, Symbol
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
		seedbox := _populate(mask, seedAll)

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

//RandomPronounce to produce pronounce random char
func (char *Character) RandomPronounce(mask uint64, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	seedtype, _ := RandomMinMax(seedVowel, seedConsonant)

	//rule max consonant and vowel is 2
	if seedtype == seedVowel {
		char.mux.Lock()
		char.Vocal++
		char.Consonant = 0
		char.mux.Unlock()
	}

	if seedtype == seedConsonant {
		char.mux.Lock()
		char.Vocal = 0
		char.Consonant++
		char.mux.Unlock()

	}

	if char.Vocal > 2 {
		char.mux.Lock()
		char.Vocal = 0
		char.Consonant = 1
		char.mux.Unlock()

		seedtype = seedConsonant
	}

	if char.Consonant > 2 {
		char.mux.Lock()
		char.Vocal = 1
		char.Consonant = 0
		char.mux.Unlock()

		seedtype = seedVowel
	}

	seedbox := _populate(mask, seedtype)
	ch <- _randomize(seedbox)
}

func _populate(mask uint64, seedtype int64) string {
	seedbox := ""

	if mask&Numeric > 0 {
		switch seedtype {
		case seedVowel:
			seedbox += setNumericVowel
		case seedConsonant:
			seedbox += setNumericConsonant
		case seedAll:
			seedbox += setNumericConsonant + setNumericVowel
		}
	}

	if mask&Small > 0 {
		switch seedtype {
		case seedVowel:

			seedbox += setVowel
		case seedConsonant:
			seedbox += setConsonant
		case seedAll:
			seedbox += setVowel + setConsonant
		}

	}

	if mask&Capital > 0 {
		switch seedtype {
		case seedVowel:
			seedbox += strings.ToUpper(setVowel)
		case seedConsonant:
			seedbox += strings.ToUpper(setConsonant)
		case seedAll:
			seedbox += strings.ToUpper(setVowel + setConsonant)
		}
	}

	if (mask&Symbol > 0) && (seedtype == seedAll) {
		seedbox += setSymbol
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
