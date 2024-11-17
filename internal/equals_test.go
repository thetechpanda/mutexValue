package internal_test

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/thetechpanda/mutex/internal"
)

type equal struct {
	value  int
	called bool
}

// Implement Equal(T) bool
func (t *equal) Equal(other *equal) bool {
	t.called = true
	return t.value == other.value
}

func TestEquals_EqualMethod(t *testing.T) {
	a := &equal{value: 1}
	b := &equal{value: 1}
	c := &equal{value: 2}

	internal.CompareStrategy = internal.Lax

	if !internal.Equals(a, b) {
		t.Errorf("Expected a and b to be equal using Equal method")
	}
	if internal.Equals(a, c) {
		t.Errorf("Expected a and c to not be equal using Equal method")
	}
	if !a.called {
		t.Errorf("Expected a compare function to be called, but wasn't")
	}
}

type equals struct {
	value  int
	called bool
}

// Implement Equals(T) bool
func (t *equals) Equals(other *equals) bool {
	t.called = true
	return t.value == other.value
}
func TestEquals_EqualsMethod(t *testing.T) {

	a := &equals{value: 2}
	b := &equals{value: 2}
	c := &equals{value: 3}

	internal.CompareStrategy = internal.Lax

	if !internal.Equals(a, b) {
		t.Errorf("Expected a and b to be equal using Equals method")
	}
	if internal.Equals(a, c) {
		t.Errorf("Expected a and c to not be equal using Equals method")
	}
	if !a.called {
		t.Errorf("Expected a compare function to be called, but wasn't")
	}
}

type compare struct {
	value  int
	called bool
}

// Implement Compare(T) int
func (t *compare) Compare(other *compare) int {
	t.called = true
	return t.value - other.value
}

func TestEquals_CompareMethod(t *testing.T) {

	a := &compare{value: 3}
	b := &compare{value: 3}
	c := &compare{value: 4}

	internal.CompareStrategy = internal.Lax

	if !internal.Equals(a, b) {
		t.Errorf("Expected a and b to be equal using Compare method")
	}
	if internal.Equals(a, c) {
		t.Errorf("Expected a and c to not be equal using Compare method")
	}
	if !a.called {
		t.Errorf("Expected a compare function to be called, but wasn't")
	}
}

type cmp struct {
	value  int
	called bool
}

// Implement Cmp(T) int
func (t *cmp) Cmp(other *cmp) int {
	t.called = true
	return t.value - other.value
}
func TestEquals_CmpMethod(t *testing.T) {

	a := &cmp{value: 4}
	b := &cmp{value: 4}
	c := &cmp{value: 5}

	internal.CompareStrategy = internal.Lax

	if !internal.Equals(a, b) {
		t.Errorf("Expected a and b to be equal using Cmp method")
	}
	if internal.Equals(a, c) {
		t.Errorf("Expected a and c to not be equal using Cmp method")
	}
	if !a.called {
		t.Errorf("Expected a compare function to be called, but wasn't")
	}
}

type is struct {
	value  int
	called bool
}

// Implement Is(T) bool
func (t *is) Is(other *is) bool {
	t.called = true
	return t.value == other.value
}

func TestEquals_IsMethod(t *testing.T) {

	a := &is{value: 5}
	b := &is{value: 5}
	c := &is{value: 6}

	internal.CompareStrategy = internal.Lax

	if !internal.Equals(a, b) {
		t.Errorf("Expected a and b to be equal using Is method")
	}
	if internal.Equals(a, c) {
		t.Errorf("Expected a and c to not be equal using Is method")
	}
	if !a.called {
		t.Errorf("Expected a compare function to be called, but wasn't")
	}
}

type isSame struct {
	value  int
	called bool
}

// Implement IsSame(T) bool
func (t *isSame) IsSame(other *isSame) bool {
	t.called = true
	return t.value == other.value
}
func TestEquals_IsSameMethod(t *testing.T) {

	a := &isSame{value: 6}
	b := &isSame{value: 6}
	c := &isSame{value: 7}

	internal.CompareStrategy = internal.Lax

	if !internal.Equals(a, b) {
		t.Errorf("Expected a and b to be equal using IsSame method")
	}
	if internal.Equals(a, c) {
		t.Errorf("Expected a and c to not be equal using IsSame method")
	}
	if !a.called {
		t.Errorf("Expected a compare function to be called, but wasn't")
	}
}

type hash struct {
	value  uint64
	called bool
}

// Implement Hash() uint64
func (t *hash) Hash() uint64 {
	t.called = true
	return t.value
}
func TestEquals_HashMethod(t *testing.T) {

	a := &hash{value: 7}
	b := &hash{value: 7}
	c := &hash{value: 8}

	internal.CompareStrategy = internal.Lax

	if !internal.Equals(a, b) {
		t.Errorf("Expected a and b to be equal using Hash method")
	}
	if internal.Equals(a, c) {
		t.Errorf("Expected a and c to not be equal using Hash method")
	}
	if !a.called {
		t.Errorf("Expected a compare function to be called, but wasn't")
	}
}

type stringer struct {
	value  int
	called bool
}

// Implement String() string
func (t *stringer) String() string {
	t.called = true
	return fmt.Sprintf("%d", t.value)
}
func TestEquals_StringMethod(t *testing.T) {

	a := &stringer{value: 8}
	b := &stringer{value: 8}
	c := &stringer{value: 9}

	internal.CompareStrategy = internal.Marshal

	if !internal.Equals(a, b) {
		t.Errorf("Expected a and b to be equal using String method")
	}
	if internal.Equals(a, c) {
		t.Errorf("Expected a and c to not be equal using String method")
	}
	if !a.called {
		t.Errorf("Expected a compare function to be called, but wasn't")
	}
}

type marshalBinary struct {
	value  int
	called bool
}

// Implement MarshalBinary() ([]byte, error)
func (t *marshalBinary) MarshalBinary() ([]byte, error) {
	t.called = true
	return []byte{byte(t.value)}, nil
}

func TestEquals_MarshalBinaryMethod(t *testing.T) {

	a := &marshalBinary{value: 9}
	b := &marshalBinary{value: 9}
	c := &marshalBinary{value: 10}

	internal.CompareStrategy = internal.Marshal

	if !internal.Equals(a, b) {
		t.Errorf("Expected a and b to be equal using MarshalBinary method")
	}
	if internal.Equals(a, c) {
		t.Errorf("Expected a and c to not be equal using MarshalBinary method")
	}
	if !a.called {
		t.Errorf("Expected a compare function to be called, but wasn't")
	}
}

type marshalJSON struct {
	Value  int `json:"value"`
	called bool
}

// Implement MarshalJSON() ([]byte, error)
func (t *marshalJSON) MarshalJSON() ([]byte, error) {
	t.called = true
	return binary.LittleEndian.AppendUint64(nil, uint64(t.Value)), nil
}
func TestEquals_MarshalJSONMethod(t *testing.T) {

	a := &marshalJSON{Value: 10}
	b := &marshalJSON{Value: 10}
	c := &marshalJSON{Value: 11}

	internal.CompareStrategy = internal.Marshal

	if !internal.Equals(a, b) {
		t.Errorf("Expected a and b to be equal using MarshalJSON method")
	}
	if internal.Equals(a, c) {
		t.Errorf("Expected a and c to not be equal using MarshalJSON method")
	}

	if !a.called {
		t.Errorf("Expected a compare function to be called, but wasn't")
	}
}

func TestEquals_ReflectStrategy(t *testing.T) {
	type testStruct struct {
		value int
	}
	a := &testStruct{value: 11}
	b := &testStruct{value: 11}
	c := &testStruct{value: 12}

	internal.CompareStrategy = internal.Reflect

	if !internal.Equals(a, b) {
		t.Errorf("Expected a and b to be equal using Reflect strategy")
	}
	if internal.Equals(a, c) {
		t.Errorf("Expected a and c to not be equal using Reflect strategy")
	}
}

func TestEquals_DeepEqualFallback(t *testing.T) {
	type testStruct struct {
		value int
	}
	a := &testStruct{value: 12}
	b := &testStruct{value: 12}
	c := &testStruct{value: 13}

	// No methods implemented, should fallback to reflect.DeepEqual
	internal.CompareStrategy = internal.Lax

	if !internal.Equals(a, b) {
		t.Errorf("Expected a and b to be equal using DeepEqual fallback")
	}
	if internal.Equals(a, c) {
		t.Errorf("Expected a and c to not be equal using DeepEqual fallback")
	}
}
