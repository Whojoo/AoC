package shared

type Assignment interface {
	FileName() string
	Part1(input []string) int
	Part2(input []string) int
}
