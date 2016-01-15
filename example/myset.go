package example

import (
	"fmt"
	"strings"
)

// This will fail if the underlying type cannot be converted to a uint64.
// Your underlying type MUST be convertible to uint64.
var _ = uint64(int(0))


// This will fail if the base type cannot be converted to a the new type.
var _ = MySet(MyType(0))

type MySet int

// NewMySet will create a new MyType set with the MyType types supplied
func NewMySet(t ...MyType) MySet {
	x := MySet(0)
	return x.Set(t...)
}

// Set one or more MyType types to the set.
// Returns the modified MyType set.
func (x *MySet) Set(t ...MyType) MySet {
	for _, v := range t {
		*x |= MySet(v)
	}
	return *x
}

// Clear one or more types from the set.
// Returns the modified MyType set.
func (x *MySet) Clear(t ...MyType) MySet {
	for _, v := range t {
		*x &= ^MySet(v)
	}
	return *x
}

// Toggle one or more types from the set.
// Returns the modified MyType set.
func (x *MySet) Toggle(t ...MyType) MySet {
	for _, v := range t {
		*x ^= MySet(v)
	}
	return *x
}

// Contains returns true if the
func (x MySet) Contains(t MyType) bool {
	return int(x)&int(t) != 0
}

// Overlap returns the overlapping set bits
func (x MySet) Overlap(t MySet) MySet {
	return MySet(int(x) & int(t))
}

// Combine returns the combined bit sets
func (x MySet) Combine(t MySet) MySet {
	return MySet(int(x) | int(t))
}

// Same returns the bit values that are the same
func (x MySet) Same(t MySet) MySet {
	return MySet(int(x) ^ int(t))
}

// SubsetOf will return true if all types of h
// is present in the set c
func (x MySet) SubsetOf(y MySet) bool {
	return int(x)|int(y) == int(y)
}

// GetOne will return a MyType type.
// Currently the first is returned, but it could be
// improved to return the strongest.
func (x MySet) GetOne() MyType {
	v := int(x)
	i := uint(0)
	for v != 0 {
		if v&1 != 0 {
			return MyType(1 << i)
		}
		i++
		v >>= 1
	}
	return MyType(0)
}

// Array returns an array of all MyType types in the set
func (x MySet) Array() (ht []MyType) {
	v := uint64(x)
	i := uint(0)
	for v != 0 {
		if v&1 != 0 {
			ht = append(ht, MyType(1<<i))
		}
		i++
		v >>= 1
	}
	return ht
}

// Count returns the number of MyType types in the set
func (x MySet) Count() int {
	if int(x) == 0 {
		return 0
	}
	// credit: https://code.google.com/u/arnehormann/
	y := uint64(x)
	y -= (y >> 1) & 0x5555555555555555
	y = (y>>2)&0x3333333333333333 + y&0x3333333333333333
	y += y >> 4
	y &= 0x0f0f0f0f0f0f0f0f
	y *= 0x0101010101010101
	return int(y >> 56)
}

// String returns a string representation of the MyType set.
func (x MySet) String() string {
	a := x.Array()
	var r []string
	for _, v := range a {
		s := fmt.Sprintf("%v", v)
		r = append(r, s)
	}
	return "[" + strings.Join(r, ", ") + "]"
}
