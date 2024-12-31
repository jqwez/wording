package finder_test

import (
	"testing"

	"github.com/jqwez/wording/finder"
	"github.com/stretchr/testify/assert"
)

func TestLoadAllWords(t *testing.T) {
	words, err := finder.LoadWords()
	assert.Nil(t, err)
	assert.NotNil(t, words)
}

func TestNewDictionaryLoadsWordsSlice(t *testing.T) {
	dictionary, err := finder.NewDictionary()
	assert.Nil(t, err)
	assert.NotNil(t, dictionary)
	assert.NotEmpty(t, dictionary.AllWords)
	assert.EqualValues(t, "a", dictionary.GetWordByPos(0))
}

func TestDictionaryReturnWordIf(t *testing.T) {
	dictionary, _ := finder.NewDictionary()
	w := dictionary.ReturnWordIf(func(word string) bool {
		return true
	})
	assert.NotEmpty(t, w)
	assert.Contains(t, w, "a")
	assert.Equal(t, len(w), len(dictionary.AllWords))
	w = dictionary.ReturnWordIf(func(word string) bool {
		if word == "a" {
			return true
		}
		return false
	})
	assert.NotEmpty(t, w)
	assert.Contains(t, w, "a")
	assert.NotEqual(t, len(w), len(dictionary.AllWords))
}
