package main

import (
	"bufio"
	"flag"
	"log"
	"os"
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

	res := ProccessSignal(inputs[0])
	res2 := ProccessMessage(inputs[0])

	log.Printf("signal index %d res 2: %d", res, res2)

}

func ProccessSignal(signal string) int {
	b := Buffer{}
	for i, _ := range signal {
		r := b.Check(signal[i:i+4], i+4)
		if r != -1 {
			return r
		}
	}

	return 0
}

func ProccessMessage(signal string) int {
	b := Buffer{}
	for i, _ := range signal {
		r := b.Check(signal[i:i+14], i+14)
		if r != -1 {
			return r
		}
	}

	return 0
}

type Buffer struct {
	startIndex int
	characters []string
}

func (b *Buffer) Check(s string, i int) int {
	m := map[int32]bool{}
	for _, c := range s {
		if ok, _ := m[c]; ok {
			return -1
		}
		m[c] = true
	}
	return i
}

func (b *Buffer) Add(s string, i int) int {
	if len(b.characters) == 0 {
		b.startIndex = i
	}
	b.characters = append(b.characters, s)

	if len(b.characters) == 4 {
		m := map[string]bool{}
		for _, c := range b.characters {
			if ok, _ := m[c]; ok {
				b.characters = []string{}
				b.startIndex = 0
				return -1
			}
			m[c] = true
		}
		return b.startIndex
	}
	return -1
}
