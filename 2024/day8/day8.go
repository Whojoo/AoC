package day8

import (
	"fmt"
	"strings"
	"time"
)

type Assignment struct{}

func GetAssignment() Assignment { return Assignment{} }

func (a Assignment) Handle(input []string, c chan<- string) {
	defer close(c)

	startTime := time.Now()

	part1, part2 := a.Part1(input), a.Part2(input)

	elapsed := time.Since(startTime)

	c <- "Day 8"
	c <- fmt.Sprintf("%d", part1)
	c <- fmt.Sprintf("%d", part2)
	c <- fmt.Sprintf("Took %s", elapsed)
}

func (Assignment) FileName() string { return "day8.txt" }

func (a Assignment) Part1(input []string) int {
	return performLogic(input, false)
}

func (a Assignment) Part2(input []string) int {
	return performLogic(input, true)
}

//nolint:gocognit
func performLogic(input []string, includePart2 bool) int {
	worldHeight, worldWidth := len(input), len(input[0])
	antennas := getAntennas(input)

	groupedAntennas := make(map[string][]Antenna)

	for _, a := range antennas {
		if groupedAntennas[a.frequency] == nil {
			groupedAntennas[a.frequency] = make([]Antenna, 0)
		}

		groupedAntennas[a.frequency] = append(groupedAntennas[a.frequency], a)
	}

	antiNodeSet := NewAntiNodeSet()

	for _, antennaGroup := range groupedAntennas {
		for i, currentAntenna := range antennaGroup {
			if includePart2 {
				antiNodeSet.Add(currentAntenna.position)
			}

			for _, antenna := range antennaGroup[i+1:] {
				// Vector math b - a means going from a -> b, so we're pointing to current
				diffVec := currentAntenna.position.Sub(antenna.position)

				antiNode1Pos := currentAntenna.position.Add(diffVec)
				antiNode2Pos := antenna.position.Sub(diffVec)

				if isInBounds(antiNode1Pos, worldWidth, worldHeight) {
					antiNodeSet.Add(antiNode1Pos)
				}

				if isInBounds(antiNode2Pos, worldWidth, worldHeight) {
					antiNodeSet.Add(antiNode2Pos)
				}

				if !includePart2 {
					continue
				}

				resonantHarmonicsAntiNodePos := antiNode1Pos.Add(diffVec)
				for isInBounds(resonantHarmonicsAntiNodePos, worldWidth, worldHeight) {
					antiNodeSet.Add(resonantHarmonicsAntiNodePos)
					resonantHarmonicsAntiNodePos = resonantHarmonicsAntiNodePos.Add(diffVec)
				}

				resonantHarmonicsAntiNodePos = antiNode2Pos.Sub(diffVec)
				for isInBounds(resonantHarmonicsAntiNodePos, worldWidth, worldHeight) {
					antiNodeSet.Add(resonantHarmonicsAntiNodePos)
					resonantHarmonicsAntiNodePos = resonantHarmonicsAntiNodePos.Sub(diffVec)
				}
			}
		}
	}

	return antiNodeSet.Length()
}

func isInBounds(pos Vector, width, height int) bool {
	return pos.X >= 0 && pos.X < width && pos.Y >= 0 && pos.Y < height
}

func getAntennas(input []string) []Antenna {
	antennas := make([]Antenna, 0)
	for y, line := range input {
		fields := strings.Split(line, "")
		for x, frequency := range fields {
			if frequency != "." {
				position := Vector{x, y}
				antenna := Antenna{frequency, position}
				antennas = append(antennas, antenna)
			}
		}
	}

	return antennas
}

type AntiNodeSet struct {
	elements map[int]struct{}
}

func NewAntiNodeSet() *AntiNodeSet {
	return &AntiNodeSet{make(map[int]struct{})}
}

func (s *AntiNodeSet) GetSetValue(v Vector) int {
	return v.X + v.Y*100
}

func (s *AntiNodeSet) Add(value Vector) {
	s.elements[s.GetSetValue(value)] = struct{}{}
}

func (s *AntiNodeSet) Length() int {
	return len(s.elements)
}

type Antenna struct {
	frequency string
	position  Vector
}

type Vector struct {
	X, Y int
}

func (v Vector) Add(o Vector) Vector {
	return Vector{v.X + o.X, v.Y + o.Y}
}

func (v Vector) Sub(o Vector) Vector {
	return v.Add(o.Scale(-1))
}

func (v Vector) Scale(s int) Vector {
	return Vector{v.X * s, v.Y * s}
}
