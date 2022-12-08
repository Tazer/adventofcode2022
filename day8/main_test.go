package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForest(t *testing.T) {
	inputs := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	f := NewForest(inputs)

	assert.Equal(t, 21, f.FindVisibleTrees())

	assert.Equal(t, 8, f.FindMostScenicTree())
}
