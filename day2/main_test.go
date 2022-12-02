package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame(t *testing.T) {
	input := []string{"A Y", "B X", "C Z"}

	game := ParseGame(input)

	points, points2 := game.MyPoints()

	assert.Equal(t, 15, points)
	assert.Equal(t, 12, points2)

}
