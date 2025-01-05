package day9

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Whojoo/AoC/2024/shared"
)

type Assignment struct{}

func GetAssignment() Assignment { return Assignment{} }

func (Assignment) FileName() string { return "day9.txt" }

func (a Assignment) Handle(input []string, c chan<- string) {
	defer close(c)

	if len(input) != 1 {
		panic("bad input, should be a single line")
	}

	startTime := time.Now()

	part1, part2 := a.Part1(input), a.Part2(input)

	elapsed := time.Since(startTime)

	c <- "Day 9"
	c <- fmt.Sprintf("%d", part1)
	c <- fmt.Sprintf("%d", part2)
	c <- fmt.Sprintf("Took %s", elapsed)
}

func (Assignment) Part1(input []string) int {
	diskMap := GenerateDiskMap(input)
	layout := GenerateDiskLayout(diskMap)
	fragmentedLayout := ShiftToFragmentedLayout(layout)
	return CalculateChecksum(fragmentedLayout)
}

func (Assignment) Part2(input []string) int {
	diskMap := GenerateDiskMap(input)
	layout := GenerateDiskLayout(diskMap)
	structuredLayout := ShiftToStructuredLayout(layout)
	return CalculateChecksum(structuredLayout)
}

func CalculateChecksum(layout []int) int {
	checkSum := 0

	for i := 0; i < len(layout); i++ {
		file := layout[i]
		if file <= freeSpace {
			continue
		}

		checkSum += file * i
	}

	return checkSum
}

func ShiftToStructuredLayout(layout []int) []int {
	// Find empty pockets and record them
	emptySpaces := GatherEmptySpaces(layout)

	// Find last file from some index starting at the end
	fileStartIndex, fileEndIndex, foundFile := FindNextFileToMove(layout, len(layout)-1)
	for foundFile {
		capacity := fileEndIndex - fileStartIndex + 1
		emptySpaceLength, ok := FindFittingEmptySpaceLength(capacity, fileStartIndex, emptySpaces)

		if ok {
			emptySpaceStartingIndex := heap.Pop(emptySpaces[emptySpaceLength]).(int)

			for i := range capacity {
				fileIndex := fileStartIndex + i
				emptySpaceIndex := emptySpaceStartingIndex + i
				layout[fileIndex], layout[emptySpaceIndex] = layout[emptySpaceIndex], layout[fileIndex]
			}

			if capacity < emptySpaceLength {
				emptySpaceStartingIndex += capacity
				emptySpaceLength -= capacity
				heap.Push(emptySpaces[emptySpaceLength], emptySpaceStartingIndex)
			}
		}

		// Repeat
		fileStartIndex, fileEndIndex, foundFile = FindNextFileToMove(layout, fileStartIndex-1)
	}

	return layout
}

const notFoundIndex int = 1000000000

func FindFittingEmptySpaceLength(minLength, maxIndex int, emptySpaces []*MinHeap) (index int, ok bool) {
	if minLength > len(emptySpaces) {
		return 0, false
	}

	bestLength := minLength
	currentIndex := notFoundIndex

	for i := minLength; i < len(emptySpaces); i++ {
		if emptySpaces[i].Len() == 0 {
			continue
		}

		index := emptySpaces[i].Peek()

		if index < currentIndex && index <= maxIndex {
			bestLength = i
			currentIndex = index
		}
	}

	if currentIndex == notFoundIndex {
		return 0, false
	}

	return bestLength, true
}

func FindNextFileToMove(layout []int, currentIndex int) (startIndex, endIndex int, ok bool) {
	currentSpaceToMoveIndex := currentIndex

	for layout[currentSpaceToMoveIndex] == freeSpace {
		currentSpaceToMoveIndex--

		if currentSpaceToMoveIndex <= 0 {
			return -1, -1, false
		}
	}

	lastFileEndIndex := currentSpaceToMoveIndex
	for layout[currentSpaceToMoveIndex] == layout[lastFileEndIndex] {
		if currentSpaceToMoveIndex <= 0 {
			return -1, -1, false
		}
		currentSpaceToMoveIndex--
	}
	lastFileStartIndex := currentSpaceToMoveIndex + 1

	return lastFileStartIndex, lastFileEndIndex, true
}

func GatherEmptySpaces(layout []int) []*MinHeap {
	var emptySpaces []struct{ startIndex, length int }
	maxLength := 0

	for i := 0; i < len(layout); i++ {
		if layout[i] != freeSpace {
			continue
		}

		length := 1
		j := i + 1
		for ; j < len(layout); j++ {
			if layout[j] == freeSpace {
				length++
			} else {
				j--
				break
			}
		}

		emptySpaces = append(emptySpaces, struct{ startIndex, length int }{i, length})
		i = j
		maxLength = max(maxLength, length)
	}

	heaps := make([]*MinHeap, maxLength+1)
	for i := range heaps {
		h := &MinHeap{}
		heaps[i] = h
		heap.Init(h)
	}

	for _, emptySpace := range emptySpaces {
		heap.Push(heaps[emptySpace.length], emptySpace.startIndex)
	}

	return heaps
}

func ShiftToFragmentedLayout(layout []int) []int {
	spaceToMoveIndex := len(layout) - 1
	for layout[spaceToMoveIndex] == freeSpace {
		spaceToMoveIndex--
	}

	for i := 0; i < len(layout); i++ {
		currentSpace := layout[i]
		if currentSpace != freeSpace {
			continue
		}

		if i >= spaceToMoveIndex {
			break
		}

		layout[i], layout[spaceToMoveIndex] = layout[spaceToMoveIndex], layout[i]

		for layout[spaceToMoveIndex] == freeSpace {
			spaceToMoveIndex--
		}
	}

	return layout
}

func GenerateDiskMap(input []string) []int {
	characters := strings.Split(input[0], "")
	return shared.Project(characters, func(c string, _ int) int {
		i, _ := strconv.Atoi(c)
		return i
	})
}

const freeSpace int = -1

func GenerateDiskLayout(diskMap []int) []int {
	layout := make([]int, 0)
	isFile := true
	currentID := 0

	for _, size := range diskMap {
		if isFile {
			for range size {
				layout = append(layout, currentID)
			}
			currentID++
		} else {
			for range size {
				layout = append(layout, freeSpace)
			}
		}

		isFile = !isFile
	}

	return layout
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Peek() int {
	return h[0]
}
