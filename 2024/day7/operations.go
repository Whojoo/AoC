package day7

type NumberOperation struct {
	number uint64
}

func (n NumberOperation) Solve() uint64 {
	return n.number
}

type AddOperation struct {
	left, right Operation
}

func (a AddOperation) Solve() uint64 {
	return a.left.Solve() + a.right.Solve()
}

type MultiplyOperation struct {
	left, right Operation
}

func (m MultiplyOperation) Solve() uint64 {
	return m.left.Solve() * m.right.Solve()
}
