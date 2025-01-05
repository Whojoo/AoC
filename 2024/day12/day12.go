package day12

type Assignment struct{}

func NewAssignment() *Assignment { return new(Assignment) }

func (a Assignment) FileName() string { return "day12.txt" }

func (a Assignment) Part1(input []string) int {
	return 0
}

func (a Assignment) Part2(input []string) int {
	return 0
}
