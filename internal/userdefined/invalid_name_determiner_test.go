package userdefined

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueIsKeyword(t *testing.T) {
	keyword := "keyword"
	keywords := map[string]bool {keyword: true}
	invalidNameDeterminer := NewInvalidNameDeterminer(keywords)

	actual := invalidNameDeterminer.IsInvalidName(keyword)
	assert.True(t, actual)
}

func TestValueIsNumeric(t *testing.T) {
	keyword := "keyword"
	keywords := map[string]bool {keyword: true}
	invalidNameDeterminer := NewInvalidNameDeterminer(keywords)

	actual := invalidNameDeterminer.IsInvalidName("1234")
	assert.True(t, actual)
}

func TestValueIsValid(t *testing.T) {
	keyword := "keyword"
	keywords := map[string]bool {keyword: true}
	invalidNameDeterminer := NewInvalidNameDeterminer(keywords)

	actual := invalidNameDeterminer.IsInvalidName("ERIC")
	assert.False(t, actual)
}
