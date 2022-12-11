package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestMonkeysSmall(t *testing.T) {
	w := World{
		Monkeys: map[int]*Monkey{
			//Monkey 0
			0: {
				Items: Stack[int64]{list: []int64{79, 98}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "*",
					Val2:     "19",
				},
				Test: Test{
					Number: 23,
					True:   2,
					False:  3,
				},
			},
			//Monkey 1
			1: {
				Items: Stack[int64]{list: []int64{54, 65, 75, 74}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "+",
					Val2:     "6",
				},
				Test: Test{
					Number: 19,
					True:   2,
					False:  0,
				},
			},
			//Monkey 2
			2: {
				Items: Stack[int64]{list: []int64{79, 60, 97}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "*",
					Val2:     "old",
				},
				Test: Test{
					Number: 13,
					True:   1,
					False:  3,
				},
			},
			//Monkey 3
			3: {
				Items: Stack[int64]{list: []int64{74}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "+",
					Val2:     "3",
				},
				Test: Test{
					Number: 17,
					True:   0,
					False:  1,
				},
			},
		},
	}

	w.Run()

	res := w.MostTwoActiveMonkeys()
	assert.Equal(t, 10605, res)

}

func TestMonkeysSmall2(t *testing.T) {
	w := World{
		Monkeys: map[int]*Monkey{
			//Monkey 0
			0: {
				Items: Stack[int64]{list: []int64{79, 98}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "*",
					Val2:     "19",
				},
				Test: Test{
					Number: 23,
					True:   2,
					False:  3,
				},
			},
			//Monkey 1
			1: {
				Items: Stack[int64]{list: []int64{54, 65, 75, 74}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "+",
					Val2:     "6",
				},
				Test: Test{
					Number: 19,
					True:   2,
					False:  0,
				},
			},
			//Monkey 2
			2: {
				Items: Stack[int64]{list: []int64{79, 60, 97}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "*",
					Val2:     "old",
				},
				Test: Test{
					Number: 13,
					True:   1,
					False:  3,
				},
			},
			//Monkey 3
			3: {
				Items: Stack[int64]{list: []int64{74}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "+",
					Val2:     "3",
				},
				Test: Test{
					Number: 17,
					True:   0,
					False:  1,
				},
			},
		},
	}

	w.Run2()

	res := w.MostTwoActiveMonkeys()
	assert.Equal(t, 2713310158, res)

}

func TestMonkeysBig(t *testing.T) {
	w := World{
		Monkeys: map[int]*Monkey{
			//Monkey 0
			0: {
				Items: Stack[int64]{list: []int64{89, 73, 66, 57, 64, 80}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "*",
					Val2:     "3",
				},
				Test: Test{
					Number: 13,
					True:   6,
					False:  2,
				},
			},
			//Monkey 1
			1: {
				Items: Stack[int64]{list: []int64{83, 78, 81, 55, 81, 59, 69}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "+",
					Val2:     "1",
				},
				Test: Test{
					Number: 3,
					True:   7,
					False:  4,
				},
			},
			//Monkey 2
			2: {
				Items: Stack[int64]{list: []int64{76, 91, 58, 85}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "*",
					Val2:     "13",
				},
				Test: Test{
					Number: 7,
					True:   1,
					False:  4,
				},
			},
			//Monkey 3
			3: {
				Items: Stack[int64]{list: []int64{71, 72, 74, 76, 68}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "*",
					Val2:     "old",
				},
				Test: Test{
					Number: 2,
					True:   6,
					False:  0,
				},
			},
			//Monkey 4
			4: {
				Items: Stack[int64]{list: []int64{98, 85, 84}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "+",
					Val2:     "7",
				},
				Test: Test{
					Number: 19,
					True:   5,
					False:  7,
				},
			},
			//Monkey 5
			5: {
				Items: Stack[int64]{list: []int64{78}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "+",
					Val2:     "8",
				},
				Test: Test{
					Number: 5,
					True:   3,
					False:  0,
				},
			},
			//Monkey 6
			6: {
				Items: Stack[int64]{list: []int64{86, 70, 60, 88, 88, 78, 74, 83}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "+",
					Val2:     "4",
				},
				Test: Test{
					Number: 11,
					True:   1,
					False:  2,
				},
			},
			//Monkey 7
			7: {
				Items: Stack[int64]{list: []int64{81, 58}},
				Operation: Operation{
					Val1:     "old",
					Modifier: "+",
					Val2:     "5",
				},
				Test: Test{
					Number: 17,
					True:   3,
					False:  5,
				},
			},
		},
	}

	w.Run()

	res := w.MostTwoActiveMonkeys()
	log.Printf("Result %d", res)

}
