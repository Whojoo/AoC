package day10

import (
	"fmt"
	"github.com/Whojoo/AoC/2024/shared"
	"math"
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

func (Assignment) Handle(input []string, c chan<- string) {
	defer close(c)

	startTime := time.Now()

	world := GenerateWorld(input)
	trailHeadScores, uniqueTrails := CountTrailheadScores(world)

	elapsed := time.Since(startTime)

	c <- "Day 10"
	c <- fmt.Sprintf("%d", trailHeadScores)
	c <- fmt.Sprintf("%d", uniqueTrails)
	c <- fmt.Sprintf("Took %s", elapsed)
}

func CountTrailheadScores(world []*Node) (score, uniqueTrails int) {
	trailHeads := shared.Filter(world, func(n *Node) bool { return n.Height == targetHeight })

	for _, startNode := range trailHeads {
		trailScore, trailCount := CountTrailheadScore(startNode)
		score += trailScore
		uniqueTrails += trailCount
	}

	return score, uniqueTrails
}

func CountTrailheadScore(trailHead *Node) (score, uniqueTrails int) {
	nodeStack := NewNodeStack()
	nodeStack.Push(trailHead)

	startingNodeSet := make(NodeSet)

	for nodeStack.Len() > 0 {
		current := nodeStack.Pop()

		if current.Height == startPositionHeight {
			startingNodeSet.Add(current)
			uniqueTrails++
			continue
		}

		newHeight := current.Height - 1
		for _, neighbour := range current.Neighbours {
			if neighbour.Height == newHeight {
				nodeStack.Push(neighbour)
			}
		}
	}

	score = startingNodeSet.Len()

	return score, uniqueTrails
}

func GenerateWorld(input []string) []*Node {
	worldNodes := make([]*Node, 0)
	var previousNodes []*Node

	for y, line := range input {
		nodes := shared.Project(strings.Split(line, ""), func(s string, x int) *Node {
			height, _ := strconv.Atoi(s)
			return NewNode(height, float64(x), float64(y))
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
	Position   Vector
	Height     int
	ID         int
}

func NewNode(height int, x, y float64) *Node {
	return &Node{
		Neighbours: make([]*Node, 0),
		Position:   NewVector(x, y),
		Height:     height,
		ID:         int(x)*1000 + int(y),
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

type Vector struct{ X, Y float64 }

func NewVector(x, y float64) Vector { return Vector{x, y} }
func (v Vector) Magnitude() float64 { return math.Sqrt(v.X*v.X + v.Y*v.Y) }
