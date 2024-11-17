package mutex

import (
	"github.com/thetechpanda/mutex/internal"
)

type TStrategy internal.TStrategy

const (
	// Lax uses all Equals interfaces, default
	Lax TStrategy = iota
	// LaxMarshal uses all Equals interfaces, then fallback on MarshalJSON and MarshalBinary before falling back on reflect.DeepEqual
	LaxMarshal
	// Marshal uses String, MarshalBinary, MarshalJSON before falling back on reflect.DeepEqual
	Marshal
	// Reflect forces the comparison to be performed using reflect.DeepEqual
	Reflect
)

// Equals does its best to not resolve to reflect.DeepEqual when comparing two values.
// This function tries the following interfaces:
//
// With CompareStrategy set to Lax or LaxMarshal
//   - Equal(T) bool
//   - Equals(T) bool
//   - Compare(T) int
//   - Cmp(T) int
//   - Is(T) bool
//   - IsSame(T) bool
//   - Hash() uint64
//
// With Compare Strategy set to LaxMarshal or Marshal
//   - String() string
//   - MarshalBinary ([]byte, error)
//   - MarshalJSON ([]byte, error)
//
// If no interface is implemented this function falls back to reflect.DeepEqual
func Equals[T any](a, b T) bool {
	return internal.Equals(a, b)
}

// Sets Equals() compare strategy
func SetStrategy(s TStrategy) {
	internal.CompareStrategy = internal.TStrategy(s)
}
