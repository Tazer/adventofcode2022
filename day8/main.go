package main

import (
	"bufio"
	"flag"
	"log"
	"os"
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

	f := NewForest(inputs)

	log.Printf("visible trees %d", f.FindVisibleTrees())

}

func NewForest(inputs []string) *Forest {
	f := Forest{}

	f.Trees = map[int]map[int]int{}

	for i, in := range inputs {
		f.Trees[i] = map[int]int{}
		for i2, c := range in {
			h, _ := strconv.Atoi(string(c))
			f.Trees[i][i2] = h
		}
	}

	return &f
}

type Forest struct {
	Trees map[int]map[int]int
}

func (f *Forest) IsVisible(y, x, h int) bool {
	//Check TOP

	topVisible := true
	leftVisible := true
	rightVisible := true
	bottomVisible := true

	for y2 := y - 1; y2 >= 0; y2-- {
		if h < f.Trees[y2][x] {
			topVisible = false
			break
		}
	}
	//CHECK LEFT
	for x2 := x - 1; x2 >= 0; x2-- {
		if h < f.Trees[y][x2] {
			leftVisible = false
			break
		}
	}
	//TODO: CHECK RIGHT
	x2 := x + 1
	for {
		if _, ok := f.Trees[y][x2]; !ok {
			break
		}

		if h < f.Trees[y][x2] {
			rightVisible = false
			break
		}
		x2++
	}
	//TODO: CHECK BOTTOM
	y2 := y + 1
	for {
		if _, ok := f.Trees[y2][x]; !ok {
			break
		}

		if h < f.Trees[y2][x] {
			bottomVisible = false
			break
		}
		y2++
	}
	return topVisible || leftVisible || rightVisible || bottomVisible
}

func (f *Forest) FindVisibleTrees() int {
	total := 0
	for k, v := range f.Trees {
		for k2, v2 := range v {
			if k2 == 0 || k == 0 {
				total++
				continue
			}
			if f.IsVisible(k, k2, v2) {
				total++
			}
		}
	}

	return total
}
