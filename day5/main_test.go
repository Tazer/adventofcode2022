package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame(t *testing.T) {
	inputs := []string{
		//"    [D]    ",
		//"[N] [C]    ",
		//"[Z] [M] [P]",
		//" 1   2   3 ",
		//"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	m := ParseMoves(inputs)

	stacks := map[int]*Stack{}

	stack1 := Stack{}
	stack1.Push("Z")
	stack1.Push("N")
	stacks[1] = &stack1

	stack2 := Stack{}
	stack2.Push("M")
	stack2.Push("C")
	stack2.Push("D")
	stacks[2] = &stack2

	stack3 := Stack{}
	stack3.Push("P")
	stacks[3] = &stack3

	s := NewShip(stacks, m)

	s.MoveCrates()

	res := s.GetTopCrates()

	assert.Equal(t, "CMZ", res)

}

func TestGame2(t *testing.T) {
	inputs := []string{
		//"    [D]    ",
		//"[N] [C]    ",
		//"[Z] [M] [P]",
		//" 1   2   3 ",
		//"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	m := ParseMoves(inputs)

	stacks := map[int]*Stack{}

	stack1 := Stack{}
	stack1.Push("Z")
	stack1.Push("N")
	stacks[1] = &stack1

	stack2 := Stack{}
	stack2.Push("M")
	stack2.Push("C")
	stack2.Push("D")
	stacks[2] = &stack2

	stack3 := Stack{}
	stack3.Push("P")
	stacks[3] = &stack3

	s := NewShip(stacks, m)

	s.MoveCrates2()

	res := s.GetTopCrates()

	assert.Equal(t, "MCD", res)

}
