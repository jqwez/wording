package finder

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/jqwez/wording/utils"
)

type WordConditionFunc func(word string) bool

type Dictionary struct {
	AllWords []string
}

func NewDictionary() (*Dictionary, error) {
	dictionary := &Dictionary{}
	words, err := LoadWords()
	if err != nil {
		return dictionary, err
	}
	dictionary.AllWords = wordSlicer(words)
	return dictionary, nil
}

func (d *Dictionary) GetWordByPos(position int) string {
	if position >= len(d.AllWords) {
		return ""
	}
	return strings.ReplaceAll(d.AllWords[position], "\r", "")
}

func (d *Dictionary) ReturnWordIf(condFunc WordConditionFunc) []string {
	words := make([]string, 0)
	for _, word := range d.AllWords {
		word = strings.ReplaceAll(word, "\r", "")
		if condFunc(word) == true {
			words = append(words, word)
		}
	}
	return words
}

func LoadWords() (string, error) {
	public, err := utils.GetPublicPath()
	if err != nil {
		return "", err
	}
	_path := filepath.Join(public, "words")
	return wordLoader(_path)
}

func wordLoader(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func wordSlicer(words string) []string {
	return strings.Split(words, "\n")
}
