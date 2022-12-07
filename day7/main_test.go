package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame(t *testing.T) {
	inputs := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}

	cmds := ParseCommands(inputs)
	f := Filesystem{}
	var n *Node
	for _, cmd := range cmds {
		n = cmd.Execute(&f, n)
	}

	assert.Equal(t, 95437, f.SumPart1())

	assert.Equal(t, 24933642, f.SumPart2())

}
