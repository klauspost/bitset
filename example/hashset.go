package example

import (
	"fmt"
	"strings"
)

// This will fail if the underlying type cannot be converted to a uint64.
// Your underlying type MUST be convertible to uint64.
var _ = uint64(int(0))


// This will fail if the base type cannot be converted to a the new type.
var _ = HashSet(Hash(0))

type HashSet int

// NewHashSet will create a new Hash set with the Hash types supplied
func NewHashSet(t ...Hash) HashSet {
	x := HashSet(0)
	return x.Set(t...)
}

// Set one or more Hash types to the set.
// Returns the modified Hash set.
func (x *HashSet) Set(t ...Hash) HashSet {
	for _, v := range t {
		*x |= HashSet(v)
	}
	return *x
}

// Clear one or more types from the set.
// Returns the modified Hash set.
func (x *HashSet) Clear(t ...Hash) HashSet {
	for _, v := range t {
		*x &= ^HashSet(v)
	}
	return *x
}

// Toggle one or more types from the set.
// Returns the modified Hash set.
func (x *HashSet) Toggle(t ...Hash) HashSet {
	for _, v := range t {
		*x ^= HashSet(v)
	}
	return *x
}

// Contains returns true if the
func (x HashSet) Contains(t Hash) bool {
	return int(x)&int(t) != 0
}

// Overlap returns the overlapping set bits
func (x HashSet) Overlap(t HashSet) HashSet {
	return HashSet(int(x) & int(t))
}

// Combine returns the combined bit sets
func (x HashSet) Combine(t HashSet) HashSet {
	return HashSet(int(x) | int(t))
}

// Same returns the bit values that are the same
func (x HashSet) Same(t HashSet) HashSet {
	return HashSet(int(x) ^ int(t))
}

// SubsetOf will return true if all types of h
// is present in the set c
func (x HashSet) SubsetOf(y HashSet) bool {
	return int(x)|int(y) == int(y)
}

// GetOne will return a Hash type.
// Currently the first is returned, but it could be
// improved to return the strongest.
func (x HashSet) GetOne() Hash {
	v := int(x)
	i := uint(0)
	for v != 0 {
		if v&1 != 0 {
			return Hash(1 << i)
		}
		i++
		v >>= 1
	}
	return Hash(0)
}

// Array returns an array of all Hash types in the set
func (x HashSet) Array() (ht []Hash) {
	v := uint64(x)
	i := uint(0)
	for v != 0 {
		if v&1 != 0 {
			ht = append(ht, Hash(1<<i))
		}
		i++
		v >>= 1
	}
	return ht
}

// Count returns the number of Hash types in the set
func (x HashSet) Count() int {
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

// String returns a string representation of the Hash set.
func (x HashSet) String() string {
	a := x.Array()
	var r []string
	for _, v := range a {
		s := fmt.Sprintf("%v", v)
		r = append(r, s)
	}
	return "[" + strings.Join(r, ", ") + "]"
}
