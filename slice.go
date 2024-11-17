package mutex

import "github.com/thetechpanda/mutex/internal"

type Slice[E any] interface {
	// Load returns the slice.
	Load() (v []E)

	// Store stores a new slice.
	// Elements of values are copied into a new slice.
	Store(values []E)

	// LoadAndDelete empties the slice and returns its old values
	LoadAndDelete() (values []E)

	// Remove removes the element at pos int.
	// ok is true if an element was at pos.
	Remove(pos int) (ok bool)

	// CompareAndSwap swaps search for old in the slice and if found swaps its value with new.
	// Returns true if the swap was performed.
	//
	// ! this function uses reflect.DeepEqual if E isn't comparable and does not implement any comparable interface handled by Equals(a,b T) bool
	CompareAndSwap(old, new E) bool

	// CompareAndDelete deletes the entry for key if its value is equal to old.
	//
	// If old is not found in the slice CompareAndDelete returns false.
	//
	// ! this function uses reflect.DeepEqual if E isn't comparable and does not implement any comparable interface handled by Equals(a,b T) bool
	CompareAndDelete(old E) (deleted bool)

	// Range calls f sequentially for each value present in the slice.
	// If f returns false, range stops the iteration.
	//
	// Range does not block any other operations on the receiver, for this reason Range does not necessarily correspond to any consistent snapshot of the slice content, the underlying slice is copied before ranging the elements.
	Range(f func(pos int, val E) bool)

	// UpdateRange is a thread-safe version of Range that locks the slice for the duration of the iteration and allows for the modification of the values.
	// If f returns false, UpdateRange stops the iteration, without updating the corresponding value in the slice.
	//
	// ! Do not invoke any Slice functions within 'f' to prevent a deadlock.
	UpdateRange(f func(pos int, old E) (new E, update bool))

	// Exclusive provides a way to perform  operations on the slice ensuring that no other operation is performed on the slice during the execution of the function.
	//
	// ! Do not invoke any Slice functions within 'f' to prevent a deadlock.
	Exclusive(f func(s *[]E))

	// Clear removes all items from the slice.
	Clear()

	// Contains returns true if the slice contains value.
	//
	// ! this function uses reflect.DeepEqual if E isn't comparable and does not implement any comparable interface handled by Equals(a,b T) bool
	Contains(value E) (pos int)

	// Len returns the number of elements in the slice.
	Len() (n int)

	// Append appends the elements to the slice
	Append(values ...E)

	// Prepend prepends the elements to the slice
	Prepend(values ...E)

	// Pop returns and delete the last element from the slice
	Pop() (value E, ok bool)

	// Shift returns and delete the first element from the slice
	Shift() (value E, ok bool)
}

// creates a new slice, elements must be comparable
func NewSlice[E comparable]() Slice[E] {
	return internal.NewSliceComparable([]E{})
}

// creates a new slice, elements must be comparable, slice is initialised with v
func NewSliceWithValue[E comparable](v []E) Slice[E] {
	return internal.NewSliceComparable[E](v)
}

// creates a new slice, elements are compared using Equals()
func NewSliceAny[E any]() Slice[E] {
	return internal.NewSlice([]E{})
}

// creates a new slice, elements are compared using Equals(), slice is initialised with v
func NewSliceAnyWithValue[E any](v []E) Slice[E] {
	return internal.NewSlice(v)
}
