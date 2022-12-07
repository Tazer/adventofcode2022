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

	f := Filesystem{}

	cmds := ParseCommands(inputs)
	var n *Node
	for _, cmd := range cmds {
		n = cmd.Execute(&f, n)
	}

	log.Printf("total %d", f.SumPart1())
	log.Printf("total2 %d", f.SumPart2())

}

func ParseCommands(input []string) []Command {
	cmds := []Command{}
	cmd := Command{}

	for i, in := range input {
		if strings.HasPrefix(in, "$") {
			if i != 0 {
				cmds = append(cmds, cmd)
			}
			cmd = Command{
				Input:  in,
				Output: []string{},
			}
		} else {
			cmd.Output = append(cmd.Output, in)
		}
	}
	cmds = append(cmds, cmd)
	return cmds
}

type Filesystem struct {
	Node *Node
}

func (fs *Filesystem) AddRoot(n *Node) {
	fs.Node = n
}

func (fs *Filesystem) SumPart1() int {
	totalSize := 0

	d := map[string]int{}

	d = fs.Node.WalkDirs(d, "")

	for _, v := range d {
		if v <= 100000 {
			totalSize += v
		}

	}

	return totalSize
}

func (fs *Filesystem) SumPart2() int {
	d := map[string]int{}

	d = fs.Node.WalkDirs(d, "")
	totalDisk := 70000000
	usedDisk := d["/"]

	diskLeft := totalDisk - usedDisk
	diskNeeded := 30000000

	log.Print(diskLeft)
	log.Print(diskNeeded)

	diskToClean := 70000000

	for _, dr := range d {
		if dr+diskLeft >= diskNeeded && dr < diskToClean {
			diskToClean = dr
		}
	}

	return diskToClean
}

type Node struct {
	IsDir  bool
	Name   string
	Size   int
	Parent *Node
	Nodes  []*Node
}

func (n *Node) AddChild(cn *Node) {
	n.Nodes = append(n.Nodes, cn)
}

func (n *Node) WalkDirs(d map[string]int, path string) map[string]int {
	if path != "" {
		path += n.Name + "/"
	} else {
		path += n.Name
	}

	d[path] = n.RDirSize()

	for _, nc := range n.Nodes {
		if nc.IsDir {
			d = nc.WalkDirs(d, path)
		}
	}
	return d
}

func (n *Node) RDirSize() int {
	total := 0
	for _, nc := range n.Nodes {
		if nc.IsDir {
			total += nc.RDirSize()
		}
		total += nc.Size
	}
	return total
}

func NewNode(parent *Node, name string, isDir bool, size int) *Node {
	return &Node{
		Parent: parent,
		Name:   name,
		IsDir:  isDir,
		Size:   size,
		Nodes:  []*Node{},
	}
}

type Command struct {
	Input  string
	Output []string
}

func (c *Command) Execute(f *Filesystem, curNode *Node) *Node {

	params := strings.Split(c.Input, " ")

	cmd := params[1]

	switch cmd {
	case "cd":
		switch params[2] {
		case "/":
			curNode = NewNode(nil, "/", true, 0)
			f.AddRoot(curNode)
			break
		case "..":
			curNode = curNode.Parent

		default:
			cn := NewNode(curNode, params[2], true, 0)
			curNode.AddChild(cn)
			curNode = cn
		}

	case "ls":
		for _, l := range c.Output {
			lSplit := strings.Split(l, " ")

			if lSplit[0] == "dir" {
				continue
			}

			name := lSplit[1]
			size, _ := strconv.Atoi(lSplit[0])

			curNode.AddChild(NewNode(curNode, name, false, size))

		}
		break

	}
	return curNode
}
