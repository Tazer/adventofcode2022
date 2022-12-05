package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
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

	stackInputs := []string{}
	inputs := []string{}

	emptyStringFound := false

	for scanner.Scan() {
		if scanner.Text() == "" {
			emptyStringFound = true
			continue
		}
		if emptyStringFound {
			inputs = append(inputs, scanner.Text())
		} else {
			stackInputs = append(stackInputs, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	moves := ParseMoves(inputs)
	stacks := ParseStacks(stackInputs)

	s := NewShip(stacks, moves)

	s.MoveCrates()

	top := s.GetTopCrates()

	log.Printf("Top crates %s", top)
	stacks2 := ParseStacks(stackInputs)
	s2 := NewShip(stacks2, moves)

	s2.MoveCrates2()
	top2 := s2.GetTopCrates()

	log.Printf("Top crates2 %s", top2)

}

func ParseStacks(input []string) map[int]*Stack {
	stacks := map[int]*Stack{}

	lastLine := input[len(input)-1]

	numberOfStacks, _ := strconv.Atoi(string(lastLine[len(lastLine)-1]))

	for i := 1; i <= numberOfStacks; i++ {
		stacks[i] = &Stack{}
	}

	for _, l := range input[:len(input)-1] {
		stackPosition := 1
		for s := 1; s <= numberOfStacks; s++ {
			if len(l) < stackPosition {
				break
			}

			if string(l[stackPosition]) != " " {
				stacks[s].Append(string(l[stackPosition]))
			}
			stackPosition += 4

		}
	}

	return stacks
}

var moveExp = regexp.MustCompile(`move (?P<crates>\d+) from (?P<from>\d+) to (?P<to>\d+)`)

func findNamedMatches(regex *regexp.Regexp, str string) map[string]int {
	match := regex.FindStringSubmatch(str)

	results := map[string]int{}
	for i, name := range match {
		r, _ := strconv.Atoi(name)
		results[regex.SubexpNames()[i]] = r
	}
	return results
}

func ParseMoves(input []string) []Move {
	// move 1 from 2 to 1
	moves := []Move{}
	for _, in := range input {
		res := findNamedMatches(moveExp, in)
		moves = append(moves, Move{Crates: res["crates"], To: res["to"], From: res["from"]})
	}

	return moves
}

func NewShip(stacks map[int]*Stack, moves []Move) *Ship {
	return &Ship{
		Stacks: stacks,
		Moves:  moves,
	}
}

type Ship struct {
	Stacks map[int]*Stack
	Moves  []Move
}

func (s *Ship) MoveCrates() {
	for _, m := range s.Moves {
		for i := 1; i <= m.Crates; i++ {
			f := s.Stacks[m.From].Pop(1)
			s.Stacks[m.To].PushM(f)
		}
	}

}

func (s *Ship) MoveCrates2() {
	for _, m := range s.Moves {
		f := s.Stacks[m.From].Pop(m.Crates)
		s.Stacks[m.To].PushM(f)
	}

}

func (s *Ship) GetTopCrates() string {
	res := ""
	for i := 1; i <= len(s.Stacks); i++ {
		res += s.Stacks[i].Peek()
	}
	return res
}

type Stack struct {
	list []string
}

func (s *Stack) Pop(n int) []string {
	if len(s.list) == 1 {
		p := s.list[0:n]
		s.list = []string{}
		return p
	}
	p := s.list[0:n]
	s.list = s.list[n:]
	return p
}

func (s *Stack) Append(input string) {
	s.list = append(s.list, input)
}

func (s *Stack) Push(input string) {
	if len(s.list) == 0 {
		s.list = append(s.list, input)
	}

	var newStack []string
	newStack = append(newStack, input)
	newStack = append(newStack, s.list...)

	s.list = newStack
}

func (s *Stack) PushM(input []string) {
	if len(s.list) == 0 {
		s.list = append(s.list, input...)
	}

	var newStack []string
	newStack = append(newStack, input...)
	newStack = append(newStack, s.list...)

	s.list = newStack
}

func (s *Stack) Peek() string {
	if len(s.list) == 0 {
		return ""
	}

	return s.list[0]
}

type Move struct {
	Crates int
	From   int
	To     int
}
