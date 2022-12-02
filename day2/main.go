package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//var version = flag.Int("version", 1, "first or second part of the assignment")

	flag.Parse()

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
	g := ParseGame(inputs)

	myPoints, myPointsPart2 := g.MyPoints()

	fmt.Printf("MyPoints are %d , part 2: %d", myPoints, myPointsPart2)
}

func ParseGame(input []string) Game {
	g := Game{}

	for _, l := range input {

		sl := strings.Split(l, " ")
		r := Round{
			EnemeyAction: sl[0],
			MyAction:     sl[1],
		}
		g.Rounds = append(g.Rounds, r)
	}

	return g
}

type Round struct {
	EnemeyAction string
	MyAction     string
}

func (r *Round) Points() int {
	if r.EnemeyAction == "A" {
		switch r.MyAction {
		case "X":
			return 1 + 3
		case "Y":
			return 2 + 6
		case "Z":
			return 3
		}
	}
	if r.EnemeyAction == "B" {
		switch r.MyAction {
		case "X":
			return 1
		case "Y":
			return 2 + 3
		case "Z":
			return 3 + 6
		}
	}

	if r.EnemeyAction == "C" {
		switch r.MyAction {
		case "X":
			return 1 + 6
		case "Y":
			return 2
		case "Z":
			return 3 + 3
		}
	}
	return 0
}

func (r *Round) PointsPart2() int {
	if r.EnemeyAction == "A" {
		switch r.MyAction {
		case "X": // lose
			r.MyAction = "Z"
		case "Y": // draw
			r.MyAction = "X"
		case "Z": // win
			r.MyAction = "Y"
		}
	}
	if r.EnemeyAction == "B" {
		switch r.MyAction {
		case "X": // lose
			r.MyAction = "X"
		case "Y": // draw
			r.MyAction = "Y"
		case "Z": // win
			r.MyAction = "Z"
		}
	}

	if r.EnemeyAction == "C" {
		switch r.MyAction {
		case "X": // lose
			r.MyAction = "Y"
		case "Y": // draw
			r.MyAction = "Z"
		case "Z": // win
			r.MyAction = "X"
		}
	}
	return r.Points()
}

type Game struct {
	Rounds []Round
}

func (g *Game) MyPoints() (int, int) {
	total := 0
	for _, r := range g.Rounds {
		total += r.Points()
	}
	totalPart2 := 0
	for _, r := range g.Rounds {
		totalPart2 += r.PointsPart2()
	}
	return total, totalPart2
}
