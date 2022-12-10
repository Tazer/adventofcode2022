package main

import (
	"bufio"
	"flag"
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

	g := NewGrid(inputs)

	g.ExecuteMoves()

	res := g.GetTailVisited()

	log.Printf("Tails %d", res)

}

type Grid struct {
	Head     Knot
	Tail     Knot
	HeadFull *Knot
	Moves    []string
}

func (g *Grid) GetTailVisited() int {
	return len(g.Tail.VisitedPositions)
}

func (g *Grid) GetTailVisited2() int {
	if g.HeadFull != nil {
		curr := g.HeadFull

		for {
			if curr.Child == nil {
				break
			}
			curr = curr.Child
		}
		return len(curr.VisitedPositions)
	}
	return 0
}

func (g *Grid) ExecuteMoves() {
	for _, m := range g.Moves {
		g.executeMove(m)
		g.executeMove2(m)
	}
}

func (g *Grid) executeMove2(m string) {
	mSplit := strings.Split(m, " ")
	direction := mSplit[0]
	length, _ := strconv.Atoi(mSplit[1])
	for i := 0; i < length; i++ {
		var parent *Knot
		curr := g.HeadFull
		for {

			switch direction {
			case "R":
				if parent == nil {
					curr.MoveRight()
				}
				if parent != nil && !curr.Touches(*parent) {
					curr.MoveTo(*parent)
					if curr.Child != nil && !curr.Child.Touches(*curr) {
						curr.Child.MoveTo(*curr)
					}
				}

			case "L":
				if parent == nil {
					curr.MoveLeft()
				}
				if parent != nil && !curr.Touches(*parent) {
					curr.MoveTo(*parent)
					if curr.Child != nil && !curr.Child.Touches(*curr) {
						curr.Child.MoveTo(*curr)
					}
				}
			case "U":
				if parent == nil {
					curr.MoveUp()
				}
				if parent != nil && !curr.Touches(*parent) {
					curr.MoveTo(*parent)
					if curr.Child != nil && !curr.Child.Touches(*curr) {
						curr.Child.MoveTo(*curr)
					}
				}
			case "D":
				if parent == nil {
					curr.MoveDown()
				}
				if parent != nil && !curr.Touches(*parent) {
					curr.MoveTo(*parent)
					if curr.Child != nil && !curr.Child.Touches(*curr) {
						curr.Child.MoveTo(*curr)
					}
				}
			}
			if curr.Child == nil {
				break
			}
			parent = curr
			curr = curr.Child
		}
	}

}

func (g *Grid) executeMove(m string) {
	mSplit := strings.Split(m, " ")
	direction := mSplit[0]
	length, _ := strconv.Atoi(mSplit[1])
	for i := 0; i < length; i++ {

		switch direction {
		case "R":
			g.Head.MoveRight()
			if !g.Tail.Touches(g.Head) {
				g.Tail.MoveTo(g.Head)
			}

		case "L":
			g.Head.MoveLeft()
			if !g.Tail.Touches(g.Head) {
				g.Tail.MoveTo(g.Head)
			}
		case "U":
			g.Head.MoveUp()
			if !g.Tail.Touches(g.Head) {
				g.Tail.MoveTo(g.Head)
			}
		case "D":
			g.Head.MoveDown()
			if !g.Tail.Touches(g.Head) {
				g.Tail.MoveTo(g.Head)
			}

		}
	}

}

func NewGrid(inputs []string) *Grid {
	g := &Grid{
		Moves: inputs,
		Head: Knot{
			Position: Position{
				Y: 1000,
				X: 1000,
			},
			VisitedPositions: map[string]bool{
				"1000x1000": true,
			},
		},
		Tail: Knot{
			Position: Position{
				Y: 1000,
				X: 1000,
			},
			VisitedPositions: map[string]bool{
				"1000x1000": true,
			},
		},
	}

	h := &Knot{
		Position: Position{
			Y: 1000,
			X: 1000,
		},
		VisitedPositions: map[string]bool{
			"1000x1000": true,
		},
	}

	currKnot := h

	for i := 0; i <= 8; i++ {
		n := &Knot{
			Position: Position{
				Y: 1000,
				X: 1000,
			},
			VisitedPositions: map[string]bool{
				"1000x1000": true,
			},
		}
		currKnot.Child = n
		currKnot = n

	}

	g.HeadFull = h

	return g
}

type Knot struct {
	Position         Position
	VisitedPositions map[string]bool
	Child            *Knot
}

func (kt *Knot) MoveTo(h Knot) {

	if kt.Position.Y == h.Position.Y && (kt.Position.X-1 == h.Position.X+1) {
		kt.MoveLeft()
	}

	if kt.Position.Y == h.Position.Y && kt.Position.X+1 == h.Position.X-1 {
		kt.MoveRight()
	}

	if kt.Position.X == h.Position.X && kt.Position.Y-1 == h.Position.Y+1 {
		kt.MoveUp()
	}

	if kt.Position.X == h.Position.X && kt.Position.Y+1 == h.Position.Y-1 {
		kt.MoveDown()
	}

	if kt.Position.X > h.Position.X && kt.Position.Y < h.Position.Y {
		kt.MoveDiagonallyDownLeft()
	}

	if kt.Position.X < h.Position.X && kt.Position.Y < h.Position.Y {
		kt.MoveDiagonallyDownRight()
	}

	if kt.Position.X > h.Position.X && kt.Position.Y > h.Position.Y {
		kt.MoveDiagonallyUpLeft()
	}

	if kt.Position.X < h.Position.X && kt.Position.Y > h.Position.Y {
		kt.MoveDiagonallyUpRight()
	}

}

func (k *Knot) MoveRight() {
	k.Position.X++
	k.VisitedPositions[k.Position.String()] = true

}
func (k *Knot) MoveLeft() {
	k.Position.X--
	k.VisitedPositions[k.Position.String()] = true
}

func (k *Knot) MoveUp() {
	k.Position.Y--
	k.VisitedPositions[k.Position.String()] = true
}

func (k *Knot) MoveDown() {
	k.Position.Y++
	k.VisitedPositions[k.Position.String()] = true
}

func (k *Knot) MoveDiagonallyUpRight() {
	k.Position.Y--
	k.Position.X++
	k.VisitedPositions[k.Position.String()] = true
}
func (k *Knot) MoveDiagonallyUpLeft() {
	k.Position.Y--
	k.Position.X--
	k.VisitedPositions[k.Position.String()] = true

}
func (k *Knot) MoveDiagonallyDownRight() {
	k.Position.Y++
	k.Position.X++
	k.VisitedPositions[k.Position.String()] = true
}
func (k *Knot) MoveDiagonallyDownLeft() {
	k.Position.Y++
	k.Position.X--
	k.VisitedPositions[k.Position.String()] = true
}

func (k *Knot) Touches(t Knot) bool {
	if k.Position.Y == t.Position.Y && k.Position.X == t.Position.X {
		return true
	}

	if k.Position.Y == t.Position.Y && (k.Position.X-1 == t.Position.X || k.Position.X+1 == t.Position.X) {
		return true
	}

	if k.Position.X == t.Position.X && (k.Position.Y-1 == t.Position.Y || k.Position.Y+1 == t.Position.Y) {
		return true
	}

	if k.Position.X == t.Position.X+1 && k.Position.Y == t.Position.Y+1 {
		return true
	}
	if k.Position.X == t.Position.X-1 && k.Position.Y == t.Position.Y+1 {
		return true
	}

	if k.Position.X == t.Position.X+1 && k.Position.Y == t.Position.Y-1 {
		return true
	}
	if k.Position.X == t.Position.X-1 && k.Position.Y == t.Position.Y-1 {
		return true
	}

	return false
}

type Position struct {
	X, Y int
}

func (p *Position) String() string {
	return strconv.Itoa(p.X) + "x" + strconv.Itoa(p.Y)
}
