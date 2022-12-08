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

	log.Printf("scenic trees %d", f.FindMostScenicTree())

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

	for yTop := y - 1; yTop >= 0; yTop-- {
		if h <= f.Trees[yTop][x] {
			topVisible = false
			break
		}
	}
	//CHECK LEFT

	for xLeft := x - 1; xLeft >= 0; xLeft-- {
		if h <= f.Trees[y][xLeft] {
			leftVisible = false
			break
		}
	}
	//TODO: CHECK RIGHT
	xRight := x + 1
	for {
		if _, ok := f.Trees[y][xRight]; !ok {
			break
		}

		if h <= f.Trees[y][xRight] {
			rightVisible = false
			break
		}
		xRight++
	}
	//TODO: CHECK BOTTOM
	yBottom := y + 1
	for {
		if _, ok := f.Trees[yBottom][x]; !ok {
			break
		}

		if h <= f.Trees[yBottom][x] {
			bottomVisible = false
			break
		}
		yBottom++
	}

	if topVisible {
		return true
	}

	if leftVisible {
		return true
	}

	if rightVisible {
		return true
	}

	if bottomVisible {
		return true
	}

	return false
}

func (f *Forest) GetScenicScore(y, x, h int) int {
	topScore := 0
	leftScore := 0
	rightScore := 0
	bottomScore := 0

	for yTop := y - 1; yTop >= 0; yTop-- {
		if h <= f.Trees[yTop][x] {
			topScore++
			break
		}
		topScore++
	}
	//CHECK LEFT

	for xLeft := x - 1; xLeft >= 0; xLeft-- {
		if h <= f.Trees[y][xLeft] {
			leftScore++
			break
		}
		leftScore++
	}
	//TODO: CHECK RIGHT
	xRight := x + 1
	for {
		if _, ok := f.Trees[y][xRight]; !ok {
			break
		}

		if h <= f.Trees[y][xRight] {
			rightScore++
			break
		}
		xRight++
		rightScore++
	}
	//TODO: CHECK BOTTOM
	yBottom := y + 1
	for {
		if _, ok := f.Trees[yBottom][x]; !ok {
			break
		}

		if h <= f.Trees[yBottom][x] {
			bottomScore++
			break
		}
		yBottom++
		bottomScore++
	}

	return topScore * leftScore * rightScore * bottomScore
}

func (f *Forest) FindVisibleTrees() int {
	total := 0
	for k, v := range f.Trees {
		for k2, v2 := range v {
			if f.IsVisible(k, k2, v2) {
				total++
			}
		}
	}

	return total
}
func (f *Forest) FindMostScenicTree() int {
	highScore := 0
	for k, v := range f.Trees {
		for k2, v2 := range v {
			treeScore := f.GetScenicScore(k, k2, v2)
			if treeScore > highScore {
				highScore = treeScore
			}
		}
	}
	return highScore
}
