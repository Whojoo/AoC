package shared

type Assignment interface {
	Handle(input []string, c chan<- string)
	FileName() string
}
