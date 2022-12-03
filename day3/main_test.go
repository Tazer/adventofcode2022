package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestGame(t *testing.T) {
	input := []string{"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw"}

	ruckSacks := ParseRuckSacks(input)

	r := ruckSacks[0]

	require.Equal(t, "vJrwpWtwJgWr", r.PocketA)
	require.Equal(t, "hcsFMMfFFhFp", r.PocketB)
	require.Equal(t, "p", r.findCommon())
	require.Equal(t, 16, r.Points())
	total := 0
	total2 := 0
	for i, r := range ruckSacks {

		if i%3 == 0 && len(ruckSacks) < i+3 {
			c := FindComman(ruckSacks[i:3])
			total2 += GetPoints(c)
		}

		total += r.Points()
	}

	assert.Equal(t, 157, total)

	//points, points2 := game.MyPoints()
	//
	//assert.Equal(t, 15, points)
	//assert.Equal(t, 12, points2)

}

func TestGame2(t *testing.T) {
	input := []string{"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw"}

	ruckSacks := ParseRuckSacks(input)

	total := 0
	for i, _ := range ruckSacks {

		if i%3 == 0 && len(ruckSacks) >= i+3 {
			log.Printf("fixing index %d total length %d", i, len(ruckSacks))
			c := FindComman(ruckSacks[i : i+3])
			total += GetPoints(c)
		}
	}

	assert.Equal(t, 70, total)
}
