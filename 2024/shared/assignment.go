package shared

type Assignment interface {
	Handle(input []string, c chan<- int)
	FileName() string
}
