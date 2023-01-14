package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	length := 10
	rndStr := RandomString(length)
	assert.Equal(t, length, len(rndStr))
}
