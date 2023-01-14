package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	set := NewSet([]string{"a", "b", "c"}...)
	assert.Len(t, set, 3)
}

func TestHas(t *testing.T) {
	set := NewSet([]string{"a", "b", "c"}...)

	// examples
	assert.True(t, set.Has("a"))
	assert.True(t, set.Has("b"))
	assert.True(t, set.Has("c"))
	assert.True(t, set.Has("a", "c"))

	// counter-examples
	assert.False(t, set.Has("1"))
	assert.False(t, set.Has("w"))
	assert.False(t, set.Has("a", "w"))
}

func TestSetEquals(t *testing.T) {
	A := NewSet([]string{"a", "b", "c"}...)
	B := NewSet([]string{"c", "b", "a"}...)

	assert.True(t, A.Equals(B))
	assert.True(t, B.Equals(A))
}

func TestSetAdd(t *testing.T) {
	set := NewSet([]string{"a", "b", "c"}...)
	set.Add("a")
	set.Add("w")

	assert.Len(t, set, 4)
}

func TestSetDelete(t *testing.T) {
	set := NewSet([]string{"a", "b", "c"}...)
	set.Delete("a")
	assert.Len(t, set, 2)
	assert.False(t, set.Has("a"))
}

func TestSetToList(t *testing.T) {
	listy := []string{"a", "b", "c"}
	setty := NewSet(listy...)
	settyListy := setty.ToList()

	assert.ElementsMatch(t, listy, settyListy)
}

func TestSetString(t *testing.T) {
	set := NewSet([]string{"a", "b", "c"}...)

	assert.Len(t, set.String(), 7)
}

func TestSetStringEmpty(t *testing.T) {
	set := NewSet[rune]()
	assert.Len(t, set.String(), 2)
}
