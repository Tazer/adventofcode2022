package main

import (
	"strconv"
	"strings"
)

func main() {

}

func ParseCommands(inputs []string) []Command {
	cmds := []Command{}
	for _, l := range inputs {
		sl := strings.Split(l, " ")

		switch sl[0] {
		case "noop":
			cmds = append(cmds, Command{
				Command:   "noop",
				X:         0,
				CycleTime: 1,
			})
		case "addx":
			x := 0
			if len(sl) == 2 {
				x, _ = strconv.Atoi(sl[1])
			}
			cmds = append(cmds, Command{
				Command:   "addx",
				X:         x,
				CycleTime: 2,
			})
		}
	}
	return cmds
}

type CPU struct {
	Cycle       int
	X           int
	CrtLine     string
	CrtPosition int
	CrtLines    []string
	Val         map[int]int
}

func NewCPU() *CPU {
	return &CPU{
		X:        1,
		Val:      map[int]int{},
		CrtLines: []string{},
	}
}

func (cpu *CPU) RunAllInsturctions(cmds []Command) int {
	for _, c := range cmds {
		cpu.RunInstruction(c)
	}

	cpu.CrtLines = append(cpu.CrtLines, cpu.CrtLine)

	total := 0

	for _, v := range cpu.Val {
		total += v
	}

	return total

}

func (cpu *CPU) RunInstruction(cmd Command) {
	cycleTime := cmd.CycleTime
	for {

		cycleTime--

		if cycleTime != -1 {
			if cpu.CrtPosition == cpu.X || cpu.CrtPosition == cpu.X-1 || cpu.CrtPosition == cpu.X-2 {
				cpu.CrtLine += "&"
			} else {
				cpu.CrtLine += "."
			}
			if cpu.Cycle != 1 {
				cpu.CrtPosition++
			}
		}

		if cycleTime == -1 {
			cpu.X += cmd.X
			if cpu.Cycle == 20 || cpu.Cycle == 60 || cpu.Cycle == 100 || cpu.Cycle == 140 || cpu.Cycle == 180 || cpu.Cycle == 220 {
				if _, ok := cpu.Val[cpu.Cycle]; !ok {
					cpu.Val[cpu.Cycle] = cpu.X * cpu.Cycle
				}
			}

			break
		}

		cpu.Cycle++

		if cpu.Cycle == 20 || cpu.Cycle == 60 || cpu.Cycle == 100 || cpu.Cycle == 140 || cpu.Cycle == 180 || cpu.Cycle == 220 {
			cpu.Val[cpu.Cycle] = cpu.X * cpu.Cycle
		}

		if cpu.Cycle == 40 || cpu.Cycle == 80 || cpu.Cycle == 120 || cpu.Cycle == 160 || cpu.Cycle == 200 {
			cpu.CrtLines = append(cpu.CrtLines, cpu.CrtLine)
			cpu.CrtLine = ""
			cpu.CrtPosition = -1
		}

	}
}

type Command struct {
	Command   string
	CycleTime int
	X         int
}
