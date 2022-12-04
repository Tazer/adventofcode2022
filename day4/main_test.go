package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame(t *testing.T) {
	inputs := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}

	pairs := ParsePairs(inputs)
	assert.Equal(t, 2, pairs.Score())

	assert.Equal(t, 4, pairs.Score2())

}
