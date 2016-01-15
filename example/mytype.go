package example

// Uses 'stringer': go get golang.org/x/tools/cmd/stringer

//go:generate stringer -type=MyType

// MyType is the type we are generating the bitset for.
type MyType int

// Example values of the type
const (
	A MyType = 1 << iota
	B
	C
	D
)
