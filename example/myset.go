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

// NewMySet will create a new hash set with the hash types supplied
func NewMySet(t ...MyType) MySet {
	h := MySet(0)
	return h.Set(t...)
}

// Set one or more hash types to the set.
// Returns the modified hash set.
func (h *MySet) Set(t ...MyType) MySet {
	for _, v := range t {
		*h |= MySet(v)
	}
	return *h
}

// Clear one or more types from the set.
// Returns the modified hash set.
func (h *MySet) Clear(t ...MyType) MySet {
	for _, v := range t {
		*h &= ^MySet(v)
	}
	return *h
}

// Toggle one or more types from the set.
// Returns the modified hash set.
func (h *MySet) Toggle(t ...MyType) MySet {
	for _, v := range t {
		*h ^= MySet(v)
	}
	return *h
}

// Contains returns true if the
func (h MySet) Contains(t MyType) bool {
	return int(h)&int(t) != 0
}

// Overlap returns the overlapping set bits
func (h MySet) Overlap(t MySet) MySet {
	return MySet(int(h) & int(t))
}

// Combine returns the combined bit sets
func (h MySet) Combine(t MySet) MySet {
	return MySet(int(h) | int(t))
}

// Same returns the bit values that are the same
func (h MySet) Same(t MySet) MySet {
	return MySet(int(h) ^ int(t))
}

// SubsetOf will return true if all types of h
// is present in the set c
func (h MySet) SubsetOf(c MySet) bool {
	return int(h)|int(c) == int(c)
}

// GetOne will return a hash type.
// Currently the first is returned, but it could be
// improved to return the strongest.
func (h MySet) GetOne() MyType {
	v := int(h)
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

// Array returns an array of all hash types in the set
func (h MySet) Array() (ht []MyType) {
	v := uint64(h)
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

// Count returns the number of hash types in the set
func (h MySet) Count() int {
	if int(h) == 0 {
		return 0
	}
	// credit: https://code.google.com/u/arnehormann/
	x := uint64(h)
	x -= (x >> 1) & 0x5555555555555555
	x = (x>>2)&0x3333333333333333 + x&0x3333333333333333
	x += x >> 4
	x &= 0x0f0f0f0f0f0f0f0f
	x *= 0x0101010101010101
	return int(x >> 56)
}

// String returns a string representation of the hash set.
func (h MySet) String() string {
	a := h.Array()
	var r []string
	for _, v := range a {
		s := fmt.Sprintf("%v", v)
		r = append(r, s)
	}
	return "[" + strings.Join(r, ", ") + "]"
}
