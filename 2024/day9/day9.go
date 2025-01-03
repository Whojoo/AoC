package day9

import (
	"fmt"
	"github.com/Whojoo/AoC/2024/shared"
	"strconv"
	"strings"
	"time"
)

type Assignment struct{}

func GetAssignment() Assignment { return Assignment{} }

func (Assignment) Handle(input []string, c chan<- string) {
	defer close(c)

	if len(input) != 1 {
		panic("bad input, should be a single line")
	}

	startTime := time.Now()

	diskMap := GenerateDiskMap(input)
	fragmentedLayout := GenerateDiskLayout(diskMap)
	structuredLayout := make([]int, len(fragmentedLayout))
	copy(structuredLayout, fragmentedLayout)

	fragmentedLayout = ShiftToFragmentedLayout(fragmentedLayout)
	structuredLayout = ShiftToStructuredLayout(structuredLayout)

	fragmentedChecksum := CalculateChecksum(fragmentedLayout)
	structuredChecksum := CalculateChecksum(structuredLayout)

	elapsed := time.Since(startTime)

	c <- "Day 9"
	c <- fmt.Sprintf("%d", fragmentedChecksum)
	c <- fmt.Sprintf("%d", structuredChecksum)
	c <- fmt.Sprintf("Took %s", elapsed)
}

func CalculateChecksum(layout []int) uint64 {
	checkSum := uint64(0)

	for i := uint64(0); i < uint64(len(layout)); i++ {
		file := layout[i]
		if file <= freeSpace {
			continue
		}

		checkSum += uint64(file) * i //nolint:gosec
	}

	return checkSum
}

func ShiftToStructuredLayout(layout []int) []int {
	// Find empty pockets and record them
	emptyPockets := GatherEmptyPockets(layout)

	// Find last file from some index starting at the end
	fileStartIndex, fileEndIndex, foundFile := FindNextFileToMove(layout, len(layout)-1)
	for foundFile {
		capacity := fileEndIndex - fileStartIndex + 1

		// Find most left fitting empty pocket
		emptyPocket, ok := shared.First(&emptyPockets, func(e EmptyPocket) bool {
			return e.adjustedStartIndex < fileStartIndex && e.capacity >= capacity
		})

		foo := &emptyPockets[0]

		if ok && foo != nil {
			// Swap and update empty pocket data
			for i := range capacity {
				fileIndex := fileStartIndex + i
				emptySpaceIndex := emptyPocket.adjustedStartIndex + i
				layout[fileIndex], layout[emptySpaceIndex] = layout[emptySpaceIndex], layout[fileIndex]
			}

			emptyPocket.capacity -= capacity
			emptyPocket.adjustedStartIndex += capacity
		}

		// Repeat
		fileStartIndex, fileEndIndex, foundFile = FindNextFileToMove(layout, fileStartIndex-1)
	}

	return layout
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

type EmptyPocket struct {
	startIndex, endIndex, adjustedStartIndex, capacity int
}

func NewEmptyPocket(startIndex int, endIndex int, capacity int) EmptyPocket {
	return EmptyPocket{startIndex: startIndex, adjustedStartIndex: startIndex, endIndex: endIndex, capacity: capacity}
}

func GatherEmptyPockets(layout []int) []EmptyPocket {
	emptyPockets := make([]EmptyPocket, 0)

	for i := 0; i < len(layout); i++ {
		if layout[i] == freeSpace {
			j := i + 1
			for ; j < len(layout); j++ {
				if layout[j] != freeSpace {
					// Subtract 1 to make sure J is always equal to the last index including
					j--
					break
				}
			}

			emptyPocket := NewEmptyPocket(i, j, j-i+1)
			emptyPockets = append(emptyPockets, emptyPocket)
			i = j
		}
	}

	return emptyPockets
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

func (Assignment) FileName() string { return "day9.txt" }

func GenerateDiskMap(input []string) []int {
	characters := strings.Split(input[0], "")
	return shared.Project(characters, func(c string) int {
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
