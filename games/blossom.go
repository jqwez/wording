package games

import (
	"encoding/json"
	"errors"
	"strings"
	"unicode"

	"github.com/jqwez/wording/finder"
)

type Blossom struct {
	Center  string
	Petals  string
	letters map[rune]int
}

type BlossomWordInfo struct {
	IsPangram bool         `json:"is_pangram"`
	Length    int          `json:"word_length"`
	Scoring   map[rune]int `json:"scoring"`
}

type BlossomWordInfoCollection map[string]BlossomWordInfo

func (b BlossomWordInfo) MarshalJSON() ([]byte, error) {
	stringScoring := make(map[string]int)
	for key, value := range b.Scoring {
		stringScoring[string(key)] = value
	}
	type Alias BlossomWordInfo
	return json.Marshal(&struct {
		Scoring map[string]int `json:"scoring"`
		*Alias
	}{
		Scoring: stringScoring,
		Alias:   (*Alias)(&b),
	})
}

func NewBlossom(center string, petals string) (*Blossom, error) {
	blossom := &Blossom{}
	blossom.letters = alphaMap()
	center = strings.ToLower(center)
	blossom.Center = center
	petals = strings.ToLower(petals)
	blossom.Petals = petals

	if len(center) != 1 || !isValidCharacter(rune(center[0])) {
		return nil, errors.New("center letter must be a single valid character")
	}
	blossom.letters[rune(center[0])] = 1

	if len(petals) != 6 {
		return nil, errors.New("petals must be six characters")
	}

	for _, ch := range petals {
		if isValidCharacter(ch) == false {
			return nil, errors.New("must be valid character")
		}
		blossom.letters[ch] = blossom.letters[ch] + 1
		if blossom.letters[ch] > 1 {
			return nil, errors.New("each character must be unique")
		}
	}
	return blossom, nil
}

func (b *Blossom) WordsWithInfo(validWords []string) BlossomWordInfoCollection {
	wordsWithInfo := make(map[string]BlossomWordInfo)
	for _, word := range validWords {
		scoring := make(map[rune]int)
		for _, letter := range b.Petals {
			score := b.ScoreWord(word, letter)
			scoring[letter] = score
		}
		info := BlossomWordInfo{
			IsPangram: b.IsPangram(word),
			Length:    len(word),
			Scoring:   scoring,
		}
		wordsWithInfo[word] = info
	}
	return wordsWithInfo
}

func (b *Blossom) WordsWithInfoJSON(validWords []string) ([]byte, error) {
	wordsWithInfo := b.WordsWithInfo(validWords)
	data, err := json.Marshal(wordsWithInfo)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func (b *Blossom) ValidWordsFunc() finder.WordConditionFunc {
	return func(word string) bool {
		if len(word) < 4 {
			return false
		}
		usesCenter := false
		_center := rune(b.Center[0])
		for _, ch := range word {
			if b.letters[ch] == 0 {
				return false
			}
			if ch == _center {
				usesCenter = true
			}
		}
		if usesCenter == false {
			return false
		}
		return true
	}
}

func (b *Blossom) FindWords(dictionary *finder.Dictionary) []string {
	return dictionary.ReturnWordIf(b.ValidWordsFunc())
}

func (b *Blossom) ScoreWord(word string, bonus rune) int {
	score := 0
	baseScoring := map[int]int{
		4: 2,
		5: 4,
		6: 6,
		7: 12,
	}
	if len(word) > 7 {
		score = (len(word)-7)*3 + 12
	} else {
		score = baseScoring[len(word)]
	}
	for _, ch := range word {
		if ch == bonus {
			score = score + 5
		}
	}
	if b.IsPangram(word) == true {
		score = score + 7
	}

	return score
}

func (b *Blossom) IsPangram(word string) bool {
	usedLetters := make(map[rune]bool)
	for _, ch := range word {
		usedLetters[ch] = true
	}
	_center := rune(b.Center[0])
	if val, ok := usedLetters[_center]; ok != true || val == false {
		return false
	}
	for _, ch := range b.Petals {
		if val, ok := usedLetters[ch]; ok != true || val == false {
			return false
		}
	}
	return true
}

func isValidCharacter(ch rune) bool {
	ch = unicode.ToLower(ch)
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	return false
}

func alphaMap() map[rune]int {
	am := make(map[rune]int)
	for ch := 'a'; ch <= 'z'; ch++ {
		am[ch] = 0
	}
	return am
}
