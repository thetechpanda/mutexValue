package internal

import (
	"reflect"
	"slices"
)

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type TStrategy int

var CompareStrategy TStrategy

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
	if CompareStrategy != Reflect {
		src := any(a)
		// lax or marshal
		if CompareStrategy == Lax || CompareStrategy == LaxMarshal {
			if equal, ok := src.(interface{ Equal(T) bool }); ok {
				return equal.Equal(b)
			}
			if equals, ok := src.(interface{ Equals(T) bool }); ok {
				return equals.Equals(b)
			}
			if compare, ok := src.(interface{ Compare(T) int }); ok {
				return compare.Compare(b) == 0
			}
			if cmp, ok := src.(interface{ Cmp(T) int }); ok {
				return cmp.Cmp(b) == 0
			}
			if is, ok := src.(interface{ Is(T) bool }); ok {
				return is.Is(b)
			}
			if isSame, ok := src.(interface{ IsSame(T) bool }); ok {
				return isSame.IsSame(b)
			}
			if hash0, ok := src.(interface{ Hash() uint64 }); ok {
				if hash1, ok := any(b).(interface{ Hash() uint64 }); ok {
					return hash1.Hash() == hash0.Hash()
				}
			}
		}
		if CompareStrategy == LaxMarshal || CompareStrategy == Marshal {
			if str0, ok := src.(interface{ String() string }); ok {
				if str1, ok := any(b).(interface{ String() string }); ok {
					return str1.String() == str0.String()
				}
			}
			if mrsh0, ok := src.(interface{ MarshalBinary() ([]byte, error) }); ok {
				if mrsh1, ok := any(b).(interface{ MarshalBinary() ([]byte, error) }); ok {
					if m0, err := mrsh0.MarshalBinary(); err == nil {
						if m1, err := mrsh1.MarshalBinary(); err == nil {
							return slices.Compare(m0, m1) == 0
						}
					}
				}
			}
			if mrsh0, ok := src.(interface{ MarshalJSON() ([]byte, error) }); ok {
				if mrsh1, ok := any(b).(interface{ MarshalJSON() ([]byte, error) }); ok {
					if m0, err := mrsh0.MarshalJSON(); err == nil {
						if m1, err := mrsh1.MarshalJSON(); err == nil {
							return slices.Compare(m0, m1) == 0
						}
					}
				}
			}
		}
	}
	return reflect.DeepEqual(a, b)
}
