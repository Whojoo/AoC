package shared

type Assignment interface {
	FileName() string
	Handle(input []string, c chan<- string)
}
