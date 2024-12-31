package finder_test

import (
	"testing"

	"github.com/jqwez/wording/finder"
	"github.com/jqwez/wording/games"
	"github.com/stretchr/testify/assert"
)

func TestNewBlossom(t *testing.T) {
	blossom, err := games.NewBlossom("a", "bcdefg")
	assert.Nil(t, err)
	assert.NotNil(t, blossom)
	blossom, err = games.NewBlossom("a", "abcedf")
	assert.NotNil(t, err, "non-unique set of letters must error")
	blossom, err = games.NewBlossom("", "abcedf")
	assert.NotNilf(t, err, "no center letter must error")
	blossom, err = games.NewBlossom("g", "abczdfi")
	assert.NotNilf(t, err, "too many petals must error")
	blossom, err = games.NewBlossom("g", "a9cefi")
	assert.NotNilf(t, err, "numbers must error")
	blossom, err = games.NewBlossom("8", "awcefi")
	assert.NotNilf(t, err, "numbers must error")
}

func TestValidWordsFunc(t *testing.T) {
	blossom, err := games.NewBlossom("w", "eriabs")
	assert.Nil(t, err)
	assert.NotNil(t, blossom)
	_func := blossom.ValidWordsFunc()
	dictionary, err := finder.NewDictionary()
	assert.Nil(t, err)
	assert.NotNil(t, dictionary)
	words := dictionary.ReturnWordIf(_func)
	assert.NotEmpty(t, words)
}

func TestFindWords(t *testing.T) {
	blossom, _ := games.NewBlossom("w", "eriabs")
	dictionary, _ := finder.NewDictionary()
	words := blossom.FindWords(dictionary)
	assert.NotEmpty(t, words)
}

func TestIsPangram(t *testing.T) {
	blossom, _ := games.NewBlossom("e", "dvlpro")
	assert.True(t, blossom.IsPangram("developer"))
	assert.False(t, blossom.IsPangram("rope"))
}

func TestScoreWord(t *testing.T) {
	blossom, _ := games.NewBlossom("e", "dvlpro")
	score := blossom.ScoreWord("developer", 'd')
	assert.Equal(t, 30, score)
	score = blossom.ScoreWord("develop", 'e')
	assert.Equal(t, 22, score)
	score = blossom.ScoreWord("looper", 'l')
	assert.Equal(t, 11, score)
}

func TestWordsWithInfo(t *testing.T) {
	blossom, _ := games.NewBlossom("e", "dvlpro")
	dictionary, _ := finder.NewDictionary()
	words := blossom.FindWords(dictionary)
	assert.NotEmpty(t, words)
	withInfo := blossom.WordsWithInfo(words)
	assert.NotEmpty(t, withInfo)
	score := withInfo["developer"].Scoring['d']
	assert.Equal(t, 30, score)
}

func TestWordsWithInfoJSON(t *testing.T) {
	blossom, _ := games.NewBlossom("e", "dvlpro")
	dictionary, _ := finder.NewDictionary()
	words := blossom.FindWords(dictionary)
	assert.NotEmpty(t, words)
	data, err := blossom.WordsWithInfoJSON(words)
	assert.Nil(t, err)
	assert.NotEmpty(t, data)
}
