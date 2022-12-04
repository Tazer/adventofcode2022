package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
	pairs := ParsePairs(inputs)
	fmt.Printf("Score %d and score 2 %d", pairs.Score(), pairs.Score2())
}

func ParsePairs(input []string) Result {
	pairs := []Pairs{}
	for _, i := range input {
		iSplit := strings.Split(i, ",")

		p1Split := strings.Split(iSplit[0], "-")
		p2Split := strings.Split(iSplit[1], "-")

		p1From, _ := strconv.Atoi(p1Split[0])
		p1To, _ := strconv.Atoi(p1Split[1])

		p2From, _ := strconv.Atoi(p2Split[0])
		p2To, _ := strconv.Atoi(p2Split[1])

		p := Pairs{
			P1: Pair{From: p1From, To: p1To},
			P2: Pair{From: p2From, To: p2To},
		}

		pairs = append(pairs, p)
	}
	return Result{Pairs: pairs}
}

type Result struct {
	Pairs []Pairs
}

func (r *Result) Score() int {
	total := 0
	for _, p := range r.Pairs {
		total += p.Score()
	}
	return total
}

func (r *Result) Score2() int {
	total := 0
	for _, p := range r.Pairs {
		total += p.Score2()
	}
	return total
}

type Pairs struct {
	P1 Pair
	P2 Pair
}

func (p *Pairs) Score() int {

	if p.P1.From <= p.P2.From && p.P1.To >= p.P2.To {
		return 1
	}

	if p.P2.From <= p.P1.From && p.P2.To >= p.P1.To {
		return 1
	}

	return 0
}

func (p *Pairs) Score2() int {

	if p.P1.From <= p.P2.From && p.P1.To >= p.P2.From {
		return 1
	}

	if p.P2.From <= p.P1.From && p.P2.To >= p.P1.From {
		return 1
	}

	return 0
}

type Pair struct {
	From int
	To   int
}
