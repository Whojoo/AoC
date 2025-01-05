package day10

import (
	"fmt"
	"github.com/Whojoo/AoC/2024/shared"
	"strconv"
	"strings"
	"time"
)

const (
	startPositionHeight int = 0
	targetHeight        int = 9
)

type Assignment struct{}

func NewAssignment() Assignment { return Assignment{} }

func (Assignment) FileName() string { return "day10.txt" }

func (a Assignment) Handle(input []string, c chan<- string) {
	defer close(c)

	startTime := time.Now()

	part1, part2 := a.Part1(input), a.Part2(input)

	elapsed := time.Since(startTime)

	c <- "Day 10"
	c <- fmt.Sprintf("%d", part1)
	c <- fmt.Sprintf("%d", part2)
	c <- fmt.Sprintf("Took %s", elapsed)
}

func (a Assignment) Part1(input []string) int {
	world := GenerateWorld(input)
	trailHeads := shared.Filter(world, func(n *Node) bool { return n.Height == targetHeight })

	score := 0
	for _, trailHead := range trailHeads {
		uniqueTrailsPerHike := CountTrailheadScore2(trailHead)
		score += len(uniqueTrailsPerHike)
	}

	return score
}

func (a Assignment) Part2(input []string) int {
	world := GenerateWorld(input)
	trailHeads := shared.Filter(world, func(n *Node) bool { return n.Height == targetHeight })

	score := 0
	for _, trailHead := range trailHeads {
		uniqueTrailsPerHike := CountTrailheadScore2(trailHead)

		for _, v := range uniqueTrailsPerHike {
			score += v
		}
	}

	return score
}

func CountTrailheadScore2(trailHead *Node) map[int]int {
	nodeStack := NewNodeStack()
	nodeStack.Push(trailHead)

	uniqueTrailsPerHike := make(map[int]int)

	for nodeStack.Len() > 0 {
		current := nodeStack.Pop()

		if current.Height == startPositionHeight {
			uniqueTrailsPerHike[current.ID]++
			continue
		}

		newHeight := current.Height - 1
		for _, neighbour := range current.Neighbours {
			if neighbour.Height == newHeight {
				nodeStack.Push(neighbour)
			}
		}
	}

	return uniqueTrailsPerHike
}

func GenerateWorld(input []string) []*Node {
	worldNodes := make([]*Node, 0)
	var previousNodes []*Node

	for y, line := range input {
		nodes := shared.Project(strings.Split(line, ""), func(s string, x int) *Node {
			height, _ := strconv.Atoi(s)
			return NewNode(height, x, y)
		})

		for i, n := range nodes {
			if i < len(nodes)-1 {
				n.AddNeighbour(nodes[i+1])
				nodes[i+1].AddNeighbour(nodes[i])
			}

			if len(previousNodes) > 0 {
				n.AddNeighbour(previousNodes[i])
				previousNodes[i].AddNeighbour(n)
			}
		}

		previousNodes = nodes
		worldNodes = append(worldNodes, nodes...)
	}

	return worldNodes
}

type Node struct {
	Neighbours []*Node
	Height     int
	ID         int
}

func NewNode(height, x, y int) *Node {
	return &Node{
		Neighbours: make([]*Node, 0),
		Height:     height,
		ID:         x*1000 + y,
	}
}

func (n *Node) AddNeighbour(neighbour *Node) {
	n.Neighbours = append(n.Neighbours, neighbour)
}

type NodeSet map[int]*Node

func (ns *NodeSet) Add(n *Node) { (*ns)[n.ID] = n }
func (ns NodeSet) Len() int     { return len(ns) }

type NodeStack []*Node

func NewNodeStack() *NodeStack { return &NodeStack{} }

func (n NodeStack) Len() int    { return len(n) }
func (n NodeStack) Peek() *Node { return n[0] }

func (n *NodeStack) Pop() *Node {
	old := *n
	oldLen := n.Len()
	node := old[oldLen-1]
	*n = old[:oldLen-1]
	return node
}

func (n *NodeStack) Push(node *Node) {
	*n = append(*n, node)
}
