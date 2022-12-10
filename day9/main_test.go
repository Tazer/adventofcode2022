package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestForest(t *testing.T) {
	inputs := []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}

	g := NewGrid(inputs)

	g.ExecuteMoves()

	res := g.GetTailVisited()

	assert.Equal(t, 13, res)

}

func TestFull(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := []string{}

	for scanner.Scan() {

		inputs = append(inputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	g := NewGrid(inputs)

	g.ExecuteMoves()

	res := g.GetTailVisited()

	log.Printf("Tails %d", res)
}

func TestPart2(t *testing.T) {
	inputs := []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}

	g := NewGrid(inputs)

	g.ExecuteMoves()

	res := g.GetTailVisited2()

	assert.Equal(t, 1, res)
}

func TestPart2Bigger(t *testing.T) {
	inputs := []string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	}

	g := NewGrid(inputs)

	g.ExecuteMoves()

	res := g.GetTailVisited2()

	assert.Equal(t, 36, res)
}

func TestFull2(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := []string{}

	for scanner.Scan() {

		inputs = append(inputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	g := NewGrid(inputs)

	g.ExecuteMoves()

	res := g.GetTailVisited2()

	log.Printf("Tails %d", res)
}
