package main

import (
	"log"
	"sort"
	"strconv"
)

func main() {

}

type World struct {
	Monkeys map[int]*Monkey
}

func (w *World) Run() {
	for i := 1; i <= 20; i++ {
		for im := 0; im < len(w.Monkeys); im++ {
			w.Monkeys[im].Execute(w.Monkeys)
		}
	}
}

func (w *World) Run2() {
	for i := 1; i <= 10000; i++ {
		for im := 0; im < len(w.Monkeys); im++ {
			w.Monkeys[im].Execute2(w.Monkeys)
		}
	}
}

func (w *World) MostTwoActiveMonkeys() int {
	inspectionPerMonkey := []int{}

	for _, m := range w.Monkeys {
		inspectionPerMonkey = append(inspectionPerMonkey, m.Inspections)
	}

	sort.Slice(inspectionPerMonkey, func(i, j int) bool {
		return inspectionPerMonkey[i] > inspectionPerMonkey[j]
	})

	return inspectionPerMonkey[0] * inspectionPerMonkey[1]
}

type Monkey struct {
	Items       Stack[int64]
	Operation   Operation
	Test        Test
	Inspections int
}

func (m *Monkey) Execute2(monkeys map[int]*Monkey) {
	for {
		if len(m.Items.list) == 0 {
			break
		}

		item := m.Items.Pop(1)[0]
		oldItem := item
		log.Print(oldItem)
		item = m.Operation.Execute(item)

		if item%int64(m.Test.Number) == 0 {
			m2 := monkeys[m.Test.True]
			m2.Items.Append(item)
		} else {
			m2 := monkeys[m.Test.False]
			m2.Items.Append(item)
		}
		m.Inspections++
	}
}

func (m *Monkey) Execute(monkeys map[int]*Monkey) {
	for {
		if len(m.Items.list) == 0 {
			break
		}

		item := m.Items.Pop(1)[0]
		item = m.Operation.Execute(item)
		item = item / 3

		if item%int64(m.Test.Number) == 0 {
			m2 := monkeys[m.Test.True]
			m2.Items.Append(item)
		} else {
			m2 := monkeys[m.Test.False]
			m2.Items.Append(item)
		}
		m.Inspections++
	}
}

type Test struct {
	Number int
	True   int
	False  int
}

type Operation struct {
	Modifier string
	Val1     string
	Val2     string
}

func (o *Operation) Execute(old int64) int64 {
	val1 := int64(0)
	val2 := int64(0)

	if o.Val1 == "old" {
		val1 = old
	} else {
		val1, _ = strconv.ParseInt(o.Val1, 10, 64)
	}

	if o.Val2 == "old" {
		val2 = old
	} else {
		val2, _ = strconv.ParseInt(o.Val2, 10, 64)
	}

	switch o.Modifier {
	case "+":
		return val1 + val2
	case "*":
		return val1 * val2

	}
	log.Fatalf("Not supported modifier %v", o.Modifier)
	return 0
}

type Stack[T any] struct {
	list []T
}

func (s *Stack[T]) Pop(n int) []T {
	if len(s.list) == 1 {
		p := s.list[0:n]
		s.list = []T{}
		return p
	}
	p := s.list[0:n]
	s.list = s.list[n:]
	return p
}

func (s *Stack[T]) Append(input T) {
	s.list = append(s.list, input)
}

func (s *Stack[T]) Push(input T) {
	if len(s.list) == 0 {
		s.list = append(s.list, input)
	}

	var newStack []T
	newStack = append(newStack, input)
	newStack = append(newStack, s.list...)

	s.list = newStack
}

func (s *Stack[T]) Peek() T {
	return s.list[0]
}
