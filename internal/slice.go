package internal

import (
	"slices"
	"sync"
)

type Slice[E any] struct {
	slice []E
	mu    sync.RWMutex
}

func NewSlice[T any](s []T) *Slice[T] {
	return &Slice[T]{
		slice: s,
		mu:    sync.RWMutex{},
	}
}

// Load returns the slice.
// ok is true if the slice wasn't nil
func (s *Slice[E]) Load() (v []E) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v = make([]E, len(s.slice))
	copy(v, s.slice)
	return v
}

// Store stores a new slice.
// Elements of values are copied into a new slice.
func (s *Slice[E]) Store(values []E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = make([]E, len(values))
	copy(s.slice, values)
}

// LoadAndDelete empties the slice and returns its old values
func (s *Slice[E]) LoadAndDelete() (values []E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.slice != nil {
		values = make([]E, len(s.slice))
		copy(values, s.slice)
		s.slice = make([]E, 0)
	}
	return values
}

// Remove removes the element at pos int.
// ok is true if an element was at pos.
func (s *Slice[E]) Remove(pos int) (ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.slice) <= pos {
		return false
	}
	s.slice = append(s.slice[:pos], s.slice[pos+1:]...)
	return true
}

// CompareAndSwap swaps search for old in the slice and if found swaps its value with new.
// Returns true if the swap was performed.
//
// ! this function uses reflect.DeepEqual if E isn't comparable and does not implement any comparable interface handled by Equals(a,b T) bool
func (s *Slice[E]) CompareAndSwap(old, new E) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for pos, elm := range s.slice {
		if Equals(elm, old) {
			s.slice[pos] = new
			return true
		}
	}
	return false
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
//
// If old is not found in the slice CompareAndDelete returns false.
//
// ! this function uses reflect.DeepEqual if E isn't comparable and does not implement any comparable interface handled by Equals(a,b T) bool
func (s *Slice[E]) CompareAndDelete(old E) (deleted bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for pos, elm := range s.slice {
		if Equals(elm, old) {
			s.slice = append(s.slice[:pos], s.slice[pos+1:]...)
			return true
		}
	}
	return false
}

// Range calls f sequentially for each value present in the slice.
// If f returns false, range stops the iteration.
//
// Range does not block any other operations on the receiver, for this reason Range does not necessarily correspond to any consistent snapshot of the slice content, the underlying slice is copied before ranging the elements.
func (s *Slice[E]) Range(f func(pos int, val E) bool) {
	for pos, val := range s.Load() {
		if !f(pos, val) {
			break
		}
	}
}

// UpdateRange is a thread-safe version of Range that locks the slice for the duration of the iteration and allows for the modification of the values.
// If f returns false, UpdateRange stops the iteration, without updating the corresponding value in the slice.
//
// ! Do not invoke any Slice functions within 'f' to prevent a deadlock.
func (s *Slice[E]) UpdateRange(f func(pos int, old E) (new E, update bool)) {
	s.Exclusive(func(s *[]E) {
		for pos, old := range *s {
			if new, update := f(pos, old); update {
				(*s)[pos] = new
				continue
			}
			break
		}
	})

}

// Exclusive provides a way to perform  operations on the slice ensuring that no other operation is performed on the slice during the execution of the function.
//
// ! Do not invoke any Slice functions within 'f' to prevent a deadlock.
func (s *Slice[E]) Exclusive(f func(s *[]E)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	f(&s.slice)
}

// Clear removes all items from the slice.
func (s *Slice[E]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = make([]E, 0)
}

// Contains returns true if the slice contains value.
//
// ! this function uses reflect.DeepEqual if E isn't comparable and does not implement any comparable interface handled by Equals(a,b T) bool
func (s *Slice[E]) Contains(value E) int {
	return slices.IndexFunc(s.Load(), func(t E) bool {
		return Equals(value, t)
	})
}

// Len returns the number of elements in the slice.
func (s *Slice[E]) Len() (n int) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.slice)
}

// Append appends the elements to the slice
func (s *Slice[E]) Append(values ...E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = append(s.slice, values...)
}

// Prepend prepends the elements to the slice
func (s *Slice[E]) Prepend(values ...E) {
	s.mu.Lock()
	defer s.mu.Unlock()
	dst := make([]E, len(values)+len(s.slice))
	copy(dst, values)
	copy(dst[len(values):], s.slice)
	s.slice = dst
}

// Pop returns and delete the last element from the slice
func (s *Slice[E]) Pop() (value E, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.slice) == 0 {
		return value, false
	}
	value = s.slice[len(s.slice)-1]
	s.slice = s.slice[0 : len(s.slice)-1]
	return value, true
}

// Shift returns and delete the first element from the slice
func (s *Slice[E]) Shift() (value E, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.slice) == 0 {
		return value, false
	}
	value = s.slice[0]
	s.slice = s.slice[1:]
	return value, true

}

type SliceComparable[E comparable] struct {
	*Slice[E]
}

func NewSliceComparable[E comparable](s []E) *SliceComparable[E] {
	return &SliceComparable[E]{
		Slice: &Slice[E]{
			slice: s,
			mu:    sync.RWMutex{},
		},
	}
}

// CompareAndSwap swaps search for old in the slice and if found swaps its value with new.
// Returns true if the swap was performed.
func (s *SliceComparable[E]) CompareAndSwap(old, new E) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for pos, elm := range s.slice {
		if elm == old {
			s.slice[pos] = new
			return true
		}
	}
	return false
}

// Contains returns true if the slice contains value.
func (s *SliceComparable[E]) Contains(value E) int {
	return slices.Index(s.Load(), value)
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
//
// If old is not found in the slice CompareAndDelete returns false.
func (s *SliceComparable[E]) CompareAndDelete(old E) (deleted bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for pos, elm := range s.slice {
		if elm == old {
			s.slice = append(s.slice[:pos], s.slice[pos+1:]...)
			return true
		}
	}
	return false
}
