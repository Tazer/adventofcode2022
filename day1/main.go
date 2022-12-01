package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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
	elfs := ParseElfs(inputs)

	most := GetMostCalories(elfs)

	top3 := Top3Calories(elfs)

	fmt.Printf("Elf carrying the most %d and top3 %d", most, top3)
}

func ParseElfs(input []string) []Elf {
	elfs := []Elf{}
	e := Elf{}
	for _, in := range input {

		iIn, err := strconv.Atoi(in)

		if err != nil {
			elfs = append(elfs, e)
			e = Elf{}
		}

		e.Items = append(e.Items, iIn)
	}

	elfs = append(elfs, e)

	return elfs
}

func GetMostCalories(elfs []Elf) int {
	most := 0
	for _, e := range elfs {
		if e.Total() > most {
			most = e.Total()
		}
	}

	return most
}

func Top3Calories(elfs []Elf) int {
	allElfTotal := []int{}

	for _, e := range elfs {
		allElfTotal = append(allElfTotal, e.Total())
	}

	sort.Slice(allElfTotal, func(i, j int) bool {
		return allElfTotal[i] > allElfTotal[j]
	})
	return allElfTotal[0] + allElfTotal[1] + allElfTotal[2]
}

type Elf struct {
	Items []int
}

func (e *Elf) Total() int {
	total := 0

	for _, item := range e.Items {
		total += item
	}
	return total
}
