package main

import (
	"bufio"
	"encoding/json"
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

	stacks := map[int]*Stack{}

	stack1 := Stack{}
	stack1.Push("N")
	stack1.Push("D")
	stack1.Push("M")
	stack1.Push("Q")
	stack1.Push("B")
	stack1.Push("P")
	stack1.Push("Z")
	stacks[1] = &stack1

	stack2 := Stack{}
	stack2.Push("C")
	stack2.Push("L")
	stack2.Push("Z")
	stack2.Push("Q")
	stack2.Push("M")
	stack2.Push("D")
	stack2.Push("H")
	stack2.Push("V")
	stacks[2] = &stack2

	stack3 := Stack{}
	stack3.Push("Q")
	stack3.Push("H")
	stack3.Push("R")
	stack3.Push("D")
	stack3.Push("V")
	stack3.Push("F")
	stack3.Push("Z")
	stack3.Push("G")
	stacks[3] = &stack3

	stack4 := Stack{}
	stack4.Push("H")
	stack4.Push("G")
	stack4.Push("D")
	stack4.Push("F")
	stack4.Push("N")
	stacks[4] = &stack4

	stack5 := Stack{}
	stack5.Push("N")
	stack5.Push("F")
	stack5.Push("Q")
	stacks[5] = &stack5

	stack6 := Stack{}
	stack6.Push("D")
	stack6.Push("Q")
	stack6.Push("V")
	stack6.Push("Z")
	stack6.Push("F")
	stack6.Push("B")
	stack6.Push("T")
	stacks[6] = &stack6

	stack7 := Stack{}
	stack7.Push("Q")
	stack7.Push("M")
	stack7.Push("T")
	stack7.Push("Z")
	stack7.Push("D")
	stack7.Push("V")
	stack7.Push("S")
	stack7.Push("H")
	stacks[7] = &stack7

	stack8 := Stack{}
	stack8.Push("M")
	stack8.Push("G")
	stack8.Push("F")
	stack8.Push("P")
	stack8.Push("N")
	stack8.Push("Q")
	stacks[8] = &stack8

	stack9 := Stack{}
	stack9.Push("B")
	stack9.Push("W")
	stack9.Push("R")
	stack9.Push("M")
	stacks[9] = &stack9

	s := NewShip(stacks, moves)

	s.MoveCrates()

	top := s.GetTopCrates()

	log.Printf("Top crates %s", top)

	top2 := Run2(moves)

	log.Printf("Top crates2 %s", top2)

}

func Run2(moves []Move) string {
	stacks := map[int]*Stack{}

	stack1 := Stack{}
	stack1.Push("N")
	stack1.Push("D")
	stack1.Push("M")
	stack1.Push("Q")
	stack1.Push("B")
	stack1.Push("P")
	stack1.Push("Z")
	stacks[1] = &stack1

	stack2 := Stack{}
	stack2.Push("C")
	stack2.Push("L")
	stack2.Push("Z")
	stack2.Push("Q")
	stack2.Push("M")
	stack2.Push("D")
	stack2.Push("H")
	stack2.Push("V")
	stacks[2] = &stack2

	stack3 := Stack{}
	stack3.Push("Q")
	stack3.Push("H")
	stack3.Push("R")
	stack3.Push("D")
	stack3.Push("V")
	stack3.Push("F")
	stack3.Push("Z")
	stack3.Push("G")
	stacks[3] = &stack3

	stack4 := Stack{}
	stack4.Push("H")
	stack4.Push("G")
	stack4.Push("D")
	stack4.Push("F")
	stack4.Push("N")
	stacks[4] = &stack4

	stack5 := Stack{}
	stack5.Push("N")
	stack5.Push("F")
	stack5.Push("Q")
	stacks[5] = &stack5

	stack6 := Stack{}
	stack6.Push("D")
	stack6.Push("Q")
	stack6.Push("V")
	stack6.Push("Z")
	stack6.Push("F")
	stack6.Push("B")
	stack6.Push("T")
	stacks[6] = &stack6

	stack7 := Stack{}
	stack7.Push("Q")
	stack7.Push("M")
	stack7.Push("T")
	stack7.Push("Z")
	stack7.Push("D")
	stack7.Push("V")
	stack7.Push("S")
	stack7.Push("H")
	stacks[7] = &stack7

	stack8 := Stack{}
	stack8.Push("M")
	stack8.Push("G")
	stack8.Push("F")
	stack8.Push("P")
	stack8.Push("N")
	stack8.Push("Q")
	stacks[8] = &stack8

	stack9 := Stack{}
	stack9.Push("B")
	stack9.Push("W")
	stack9.Push("R")
	stack9.Push("M")
	stacks[9] = &stack9

	s := NewShip(stacks, moves)

	s.MoveCrates2()

	top := s.GetTopCrates()

	return top
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

	m, _ := json.Marshal(moves)
	log.Println("Moves")
	log.Print(string(m))

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
		log.Printf("Move executed %+v \n", m)
		for k, v := range s.Stacks {
			log.Printf("Stack %d top: %s \n", k, v.Peek())
		}
	}

}

func (s *Ship) MoveCrates2() {
	for _, m := range s.Moves {
		f := s.Stacks[m.From].Pop(m.Crates)
		s.Stacks[m.To].PushM(f)

		log.Printf("Move executed %+v \n", m)
		for k, v := range s.Stacks {
			log.Printf("Stack %d top: %s \n", k, v.Peek())
		}
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
	log.Printf("PoP: %s", p)
	return p
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
