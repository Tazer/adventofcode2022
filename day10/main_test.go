package main

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestForest(t *testing.T) {
	inputs := []string{
		"noop",
		"addx 3",
		"addx -5",
	}

	cpu := NewCPU()
	cmds := ParseCommands(inputs)

	cpu.RunAllInsturctions(cmds)

	assert.Equal(t, -1, cpu.X)
	assert.Equal(t, 5, cpu.Cycle)

}

func TestLargeTest(t *testing.T) {
	file, err := os.Open("input-test.txt")
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
	cpu := NewCPU()
	cmds := ParseCommands(inputs)

	res := cpu.RunAllInsturctions(cmds)
	fmt.Println("Result")
	for _, l := range cpu.CrtLines {
		fmt.Println(l)
	}
	fmt.Println("End")

	assert.Equal(t, 13140, res)

}

func TestFull(t *testing.T) {
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
	cpu := NewCPU()
	cmds := ParseCommands(inputs)

	res := cpu.RunAllInsturctions(cmds)

	log.Printf("res: %d \n", res)
	fmt.Println("Result")
	for _, l := range cpu.CrtLines {
		fmt.Println(l)
	}
	fmt.Println("End")

}
