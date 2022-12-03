package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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
	ruckSacks := ParseRuckSacks(inputs)

	total := 0

	for _, r := range ruckSacks {
		total += r.Points()
	}

	fmt.Printf("MyPoints are %d", total)
	total2 := 0
	for i, _ := range ruckSacks {

		if i%3 == 0 && len(ruckSacks) >= i+3 {
			//log.Printf("fixing index %d total length %d", i, len(ruckSacks))
			c := FindComman(ruckSacks[i : i+3])
			total2 += GetPoints(c)
		}
	}

	fmt.Printf("MyPoints2 are %d", total2)
}

func ParseRuckSacks(input []string) []Rucksack {

	ruckSacks := []Rucksack{}

	for _, l := range input {

		a := l[0 : len(l)/2]
		b := l[len(l)/2 : len(l)]

		r := Rucksack{
			PocketA: a,
			PocketB: b,
		}
		ruckSacks = append(ruckSacks, r)
	}

	return ruckSacks
}

type Rucksack struct {
	PocketA string
	PocketB string
}

func (r *Rucksack) FullLoad() string {
	return r.PocketA + r.PocketB
}

func (r *Rucksack) findCommon() string {

	for _, p1 := range r.PocketA {
		for _, p2 := range r.PocketB {
			if p1 == p2 {
				return string(p1)
			}
		}
	}
	return ""
}

func (r *Rucksack) Points() int {

	c := r.findCommon()
	return GetPoints(c)
}

func FindComman(ruckSack []Rucksack) string {
	for _, r1 := range ruckSack[0].FullLoad() {
		for _, r2 := range ruckSack[1].FullLoad() {
			for _, r3 := range ruckSack[2].FullLoad() {
				if r1 == r2 && r1 == r3 {
					return string(r1)
				}
			}
		}

	}
	return ""
}

func GetPoints(input string) int {
	scores := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	points := 0

	for i, s := range scores {
		if input == string(s) {
			points = i + 1
		}
	}
	return points
}
