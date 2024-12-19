package day7

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type AssignmentAttempt struct{}

func GetAssignment() AssignmentAttempt {
	return AssignmentAttempt{}
}

func (AssignmentAttempt) Handle(input []string, c chan<- string) {
	calibrationEquations := MapInput(input)

	startTime := time.Now()
	first := CalculateTotalCalibrationResult(calibrationEquations)
	firstTime := time.Since(startTime)

	startTime = time.Now()
	second := 0
	secondTime := time.Since(startTime)

	c <- "Day 7"
	c <- fmt.Sprintf("First result: %d in %s", first, firstTime)
	c <- fmt.Sprintf("Second result: %d in %s", second, secondTime)

	close(c)
}

func (AssignmentAttempt) FileName() string {
	return "day7.txt"
}

type Operation interface {
	Solve() uint64
}

type (
	CalibrationEquations []CalibrationEquation
	CalibrationEquation  struct {
		TestValue  uint64
		Operations []Operation
	}
)

func MapInput(input []string) CalibrationEquations {
	calibrationEquations := make(CalibrationEquations, len(input))
	for i, line := range input {
		firstSplit := strings.Split(line, ":")
		testValue, _ := strconv.ParseUint(firstSplit[0], 10, 64)

		operations := make([]Operation, 0)
		for _, stringNum := range strings.Fields(firstSplit[1]) {
			operationNumber, _ := strconv.ParseUint(stringNum, 10, 64)
			operations = append(operations, NumberOperation{operationNumber})
		}

		calibrationEquations[i] = CalibrationEquation{testValue, operations}
	}

	return calibrationEquations
}

func CalculateTotalCalibrationResult(calibrationEquations CalibrationEquations) uint64 {
	var sum uint64 = 0
	for _, equation := range calibrationEquations {
		operationSequences := CreateOperationMarkerSequences(len(equation.Operations) - 1)
		oldSum := sum

		for _, operationSequence := range operationSequences {
			operation := createLeftToRightOperation(equation, operationSequence)
			solvedOperation := operation.Solve()

			if solvedOperation == equation.TestValue {
				sum += solvedOperation
				break
			}
		}

		if sum == oldSum {
			fmt.Printf("Could not find a solution for %v\n", equation)
		}
	}

	return sum
}

type (
	OperationMarkerSequence []OperationMarker
	OperationMarker         int
)

const (
	Add OperationMarker = iota
	Multiply
)

func CreateOperationMarkerSequences(depth int) []OperationMarkerSequence {
	if depth < 1 {
		return make([]OperationMarkerSequence, 0)
	}

	sequences := []OperationMarkerSequence{
		{Add},
		{Multiply},
	}

	for i := 2; i <= depth; i++ {
		sequences = append(sequences, sequences...)

		for j := range sequences {
			if j < len(sequences)/2 {
				sequences[j] = append(sequences[j], Add)
			} else {
				sequences[j] = append(sequences[j], Multiply)
			}
		}
	}

	return sequences
}

func createLeftToRightOperation(equation CalibrationEquation, operationSequence OperationMarkerSequence) Operation {
	currentOperation := equation.Operations[0]

	for i, operation := range equation.Operations[1:] {
		sequenceIndex := i
		currentOperation = createOperation(currentOperation, operation, operationSequence[sequenceIndex])
	}

	return currentOperation
}

func createOperation(left Operation, right Operation, operationMarker OperationMarker) Operation {
	switch operationMarker {
	case Add:
		return AddOperation{left, right}
	case Multiply:
		return MultiplyOperation{left, right}
	default:
		panic(fmt.Sprintf("Unknown operation marker: %d", operationMarker))
	}
}
