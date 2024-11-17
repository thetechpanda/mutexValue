package internal_test

import (
	"reflect"
	"sync"
	"testing"

	"github.com/thetechpanda/mutex/internal"
)

func TestSlice_Comparable(t *testing.T) {
	s := internal.NewSliceComparable([]int{1, 2, 3})

	// Test Load
	v := s.Load()
	if !reflect.DeepEqual(v, []int{1, 2, 3}) {
		t.Errorf("Expected %v, got %v", []int{1, 2, 3}, v)
	}

	// Test Store
	s.Store([]int{4, 5, 6})
	v = s.Load()
	if !reflect.DeepEqual(v, []int{4, 5, 6}) {
		t.Errorf("Expected %v, got %v", []int{4, 5, 6}, v)
	}

	// Test LoadAndDelete
	v = s.LoadAndDelete()
	if !reflect.DeepEqual(v, []int{4, 5, 6}) {
		t.Errorf("Expected %v, got %v", []int{4, 5, 6}, v)
	}
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}

	// Test Remove
	s.Store([]int{1, 2, 3})
	ok := s.Remove(1)
	if !ok {
		t.Errorf("Expected true, got false")
	}
	v = s.Load()
	if !reflect.DeepEqual(v, []int{1, 3}) {
		t.Errorf("Expected %v, got %v", []int{1, 3}, v)
	}

	// Test CompareAndSwap
	ok = s.CompareAndSwap(3, 4)
	if !ok {
		t.Errorf("Expected true, got false")
	}
	v = s.Load()
	if !reflect.DeepEqual(v, []int{1, 4}) {
		t.Errorf("Expected %v, got %v", []int{1, 4}, v)
	}

	// Test CompareAndDelete
	ok = s.CompareAndDelete(4)
	if !ok {
		t.Errorf("Expected true, got false")
	}
	v = s.Load()
	if !reflect.DeepEqual(v, []int{1}) {
		t.Errorf("Expected %v, got %v", []int{1}, v)
	}

	// Test Range
	s.Store([]int{1, 2, 3})
	sum := 0
	s.Range(func(pos int, val int) bool {
		sum += val
		return true
	})
	if sum != 6 {
		t.Errorf("Expected sum 6, got %d", sum)
	}

	// Test UpdateRange
	s.UpdateRange(func(pos int, old int) (int, bool) {
		return old * 2, true
	})
	v = s.Load()
	if !reflect.DeepEqual(v, []int{2, 4, 6}) {
		t.Errorf("Expected %v, got %v", []int{2, 4, 6}, v)
	}

	// Test Exclusive
	s.Exclusive(func(slice *[]int) {
		*slice = append(*slice, 8)
	})
	v = s.Load()
	if !reflect.DeepEqual(v, []int{2, 4, 6, 8}) {
		t.Errorf("Expected %v, got %v", []int{2, 4, 6, 8}, v)
	}

	// Test Clear
	s.Clear()
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}

	// Test Contains
	s.Store([]int{1, 2, 3})
	pos := s.Contains(2)
	if pos != 1 {
		t.Errorf("Expected position 1, got %d", pos)
	}

	// Test Len
	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}

	// Test Append
	s.Append(4, 5)
	v = s.Load()
	if !reflect.DeepEqual(v, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Expected %v, got %v", []int{1, 2, 3, 4, 5}, v)
	}

	// Test Prepend
	s.Prepend(0)
	v = s.Load()
	if !reflect.DeepEqual(v, []int{0, 1, 2, 3, 4, 5}) {
		t.Errorf("Expected %v, got %v", []int{0, 1, 2, 3, 4, 5}, v)
	}

	// Test Pop
	val, ok := s.Pop()
	if !ok || val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}

	// Test Shift
	val, ok = s.Shift()
	if !ok || val != 0 {
		t.Errorf("Expected 0, got %d", val)
	}
}

func TestSlice_Any(t *testing.T) {
	s := internal.NewSlice([]int{1, 2, 3})

	// Test Load
	v := s.Load()
	if !reflect.DeepEqual(v, []int{1, 2, 3}) {
		t.Errorf("Expected %v, got %v", []int{1, 2, 3}, v)
	}

	// Test Store
	s.Store([]int{4, 5, 6})
	v = s.Load()
	if !reflect.DeepEqual(v, []int{4, 5, 6}) {
		t.Errorf("Expected %v, got %v", []int{4, 5, 6}, v)
	}

	// Test LoadAndDelete
	v = s.LoadAndDelete()
	if !reflect.DeepEqual(v, []int{4, 5, 6}) {
		t.Errorf("Expected %v, got %v", []int{4, 5, 6}, v)
	}
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}

	// Test Remove
	s.Store([]int{1, 2, 3})
	ok := s.Remove(1)
	if !ok {
		t.Errorf("Expected true, got false")
	}
	v = s.Load()
	if !reflect.DeepEqual(v, []int{1, 3}) {
		t.Errorf("Expected %v, got %v", []int{1, 3}, v)
	}

	// Test CompareAndSwap
	ok = s.CompareAndSwap(3, 4)
	if !ok {
		t.Errorf("Expected true, got false")
	}
	v = s.Load()
	if !reflect.DeepEqual(v, []int{1, 4}) {
		t.Errorf("Expected %v, got %v", []int{1, 4}, v)
	}

	// Test CompareAndDelete
	ok = s.CompareAndDelete(4)
	if !ok {
		t.Errorf("Expected true, got false")
	}
	v = s.Load()
	if !reflect.DeepEqual(v, []int{1}) {
		t.Errorf("Expected %v, got %v", []int{1}, v)
	}

	// Test Range
	s.Store([]int{1, 2, 3})
	sum := 0
	s.Range(func(pos int, val int) bool {
		sum += val
		return true
	})
	if sum != 6 {
		t.Errorf("Expected sum 6, got %d", sum)
	}

	// Test UpdateRange
	s.UpdateRange(func(pos int, old int) (int, bool) {
		return old * 2, true
	})
	v = s.Load()
	if !reflect.DeepEqual(v, []int{2, 4, 6}) {
		t.Errorf("Expected %v, got %v", []int{2, 4, 6}, v)
	}

	// Test Exclusive
	s.Exclusive(func(slice *[]int) {
		*slice = append(*slice, 8)
	})
	v = s.Load()
	if !reflect.DeepEqual(v, []int{2, 4, 6, 8}) {
		t.Errorf("Expected %v, got %v", []int{2, 4, 6, 8}, v)
	}

	// Test Clear
	s.Clear()
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}

	// Test Contains
	s.Store([]int{1, 2, 3})
	pos := s.Contains(2)
	if pos != 1 {
		t.Errorf("Expected position 1, got %d", pos)
	}

	// Test Len
	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}

	// Test Append
	s.Append(4, 5)
	v = s.Load()
	if !reflect.DeepEqual(v, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Expected %v, got %v", []int{1, 2, 3, 4, 5}, v)
	}

	// Test Prepend
	s.Prepend(0)
	v = s.Load()
	if !reflect.DeepEqual(v, []int{0, 1, 2, 3, 4, 5}) {
		t.Errorf("Expected %v, got %v", []int{0, 1, 2, 3, 4, 5}, v)
	}

	// Test Pop
	val, ok := s.Pop()
	if !ok || val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}

	// Test Shift
	val, ok = s.Shift()
	if !ok || val != 0 {
		t.Errorf("Expected 0, got %d", val)
	}
}

func TestSlice_Concurrent(t *testing.T) {
	s := internal.NewSliceComparable[int]([]int{})
	var wg sync.WaitGroup
	numGoroutines := 1000
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Append(i)
			s.Prepend(i)
			s.Remove(i)
			s.Contains(i)
			s.CompareAndSwap(i, i+1)
			s.CompareAndDelete(i + 1)
			s.Load()
			s.Store([]int{i})
			s.LoadAndDelete()
			s.Clear()
			s.Len()
			s.Exclusive(func(slice *[]int) {
				*slice = append(*slice, i)
			})
			s.Range(func(pos int, val int) bool {
				return true
			})
			s.UpdateRange(func(pos int, old int) (int, bool) {
				return old + 1, true
			})
		}(i)
	}
	wg.Wait()
}

type eint int

func (a eint) Equal(b eint) bool {
	return b == a
}

func TestSlice_AnyEquals(t *testing.T) {

	s := internal.NewSlice([]eint{1, 2, 3})

	// Test Load
	v := s.Load()
	if !reflect.DeepEqual(v, []eint{1, 2, 3}) {
		t.Errorf("Expected %v, got %v", []eint{1, 2, 3}, v)
	}

	// Test Store
	s.Store([]eint{4, 5, 6})
	v = s.Load()
	if !reflect.DeepEqual(v, []eint{4, 5, 6}) {
		t.Errorf("Expected %v, got %v", []eint{4, 5, 6}, v)
	}

	// Test LoadAndDelete
	v = s.LoadAndDelete()
	if !reflect.DeepEqual(v, []eint{4, 5, 6}) {
		t.Errorf("Expected %v, got %v", []eint{4, 5, 6}, v)
	}
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}

	// Test Remove
	s.Store([]eint{1, 2, 3})
	ok := s.Remove(1)
	if !ok {
		t.Errorf("Expected true, got false")
	}
	v = s.Load()
	if !reflect.DeepEqual(v, []eint{1, 3}) {
		t.Errorf("Expected %v, got %v", []eint{1, 3}, v)
	}

	// Test CompareAndSwap
	ok = s.CompareAndSwap(3, 4)
	if !ok {
		t.Errorf("Expected true, got false")
	}
	v = s.Load()
	if !reflect.DeepEqual(v, []eint{1, 4}) {
		t.Errorf("Expected %v, got %v", []eint{1, 4}, v)
	}

	// Test CompareAndDelete
	ok = s.CompareAndDelete(4)
	if !ok {
		t.Errorf("Expected true, got false")
	}
	v = s.Load()
	if !reflect.DeepEqual(v, []eint{1}) {
		t.Errorf("Expected %v, got %v", []eint{1}, v)
	}

	// Test Range
	s.Store([]eint{1, 2, 3})
	sum := eint(0)
	s.Range(func(pos int, val eint) bool {
		sum += val
		return true
	})
	if sum != 6 {
		t.Errorf("Expected sum 6, got %d", sum)
	}

	// Test UpdateRange
	s.UpdateRange(func(pos int, old eint) (eint, bool) {
		return old * 2, true
	})
	v = s.Load()
	if !reflect.DeepEqual(v, []eint{2, 4, 6}) {
		t.Errorf("Expected %v, got %v", []eint{2, 4, 6}, v)
	}

	// Test Exclusive
	s.Exclusive(func(slice *[]eint) {
		*slice = append(*slice, 8)
	})
	v = s.Load()
	if !reflect.DeepEqual(v, []eint{2, 4, 6, 8}) {
		t.Errorf("Expected %v, got %v", []eint{2, 4, 6, 8}, v)
	}

	// Test Clear
	s.Clear()
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}

	// Test Contains
	s.Store([]eint{1, 2, 3})
	pos := s.Contains(2)
	if pos != 1 {
		t.Errorf("Expected position 1, got %d", pos)
	}

	// Test Len
	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}

	// Test Append
	s.Append(4, 5)
	v = s.Load()
	if !reflect.DeepEqual(v, []eint{1, 2, 3, 4, 5}) {
		t.Errorf("Expected %v, got %v", []eint{1, 2, 3, 4, 5}, v)
	}

	// Test Prepend
	s.Prepend(0)
	v = s.Load()
	if !reflect.DeepEqual(v, []eint{0, 1, 2, 3, 4, 5}) {
		t.Errorf("Expected %v, got %v", []eint{0, 1, 2, 3, 4, 5}, v)
	}

	// Test Pop
	val, ok := s.Pop()
	if !ok || val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}

	// Test Shift
	val, ok = s.Shift()
	if !ok || val != 0 {
		t.Errorf("Expected 0, got %d", val)
	}
}
