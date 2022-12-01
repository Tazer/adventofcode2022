package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestElfs(t *testing.T) {
	input := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}

	elfs := ParseElfs(input)

	most := GetMostCalories(elfs)

	assert.Equal(t, 24000, most)

	top3 := Top3Calories(elfs)

	assert.Equal(t, 45000, top3)

}
