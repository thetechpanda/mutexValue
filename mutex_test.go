package mutex_test

import (
	"testing"

	"github.com/thetechpanda/mutex"
)

func TestValue(t *testing.T) {

	t.Run("new without value", func(t *testing.T) {
		mv := mutex.NewValue[string]()
		_, ok := mv.Load()
		if ok {
			t.Errorf("Expected ok to be false, got true")
		}
		mv.Store("42")
		v, ok := mv.Load()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != "42" {
			t.Errorf("Expected value to be 42, got %v", v)
		}
	})
	t.Run("new with value", func(t *testing.T) {
		mv := mutex.NewWithValue[string]("42")
		v, ok := mv.Load()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != "42" {
			t.Errorf("Expected value to be 42, got %v", v)
		}
	})
}

func TestNumeric(t *testing.T) {
	t.Run("new without value", func(t *testing.T) {
		mv := mutex.NewNumeric[int]()
		_, ok := mv.Load()
		if ok {
			t.Errorf("Expected ok to be false, got true")
		}
		mv.Store(42)
		v, ok := mv.Load()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != 42 {
			t.Errorf("Expected value to be 42, got %v", v)
		}
	})
	t.Run("new with value", func(t *testing.T) {
		mv := mutex.NewNumericWithValue(42)
		v, ok := mv.Load()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != 42 {
			t.Errorf("Expected value to be 42, got %v", v)
		}
		mv.Add(1)
		v, ok = mv.Load()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != 43 {
			t.Errorf("Expected value to be 43, got %v", v)
		}
	})

}

func TestMap(t *testing.T) {
	t.Run("new without value", func(t *testing.T) {
		mv := mutex.NewMap[string, string]()
		_, ok := mv.Load("key")
		if ok {
			t.Errorf("Expected ok to be false, got true")
		}
		mv.Store("key", "42")
		v, ok := mv.Load("key")
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != "42" {
			t.Errorf("Expected value to be 42, got %v", v)
		}
	})

	t.Run("new with value", func(t *testing.T) {
		m := map[string]string{"key": "42"}
		mv := mutex.NewMapWithValue(m)
		v, ok := mv.Load("key")
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != "42" {
			t.Errorf("Expected value to be 42, got %v", v)
		}
	})
}

func TestSlice(t *testing.T) {
	t.Run("new without value (append, shift)", func(t *testing.T) {
		mv := mutex.NewSlice[string]()
		if len(mv.Load()) > 0 {
			t.Errorf("Expected ok to be false, got true")
		}
		mv.Append("41", "42")
		v, ok := mv.Shift()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != "41" {
			t.Errorf("Expected value to be 41, got %v", v)
		}
		v, ok = mv.Shift()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != "42" {
			t.Errorf("Expected value to be 42, got %v", v)
		}

	})

	t.Run("new without value (prepend, pop)", func(t *testing.T) {
		mv := mutex.NewSlice[string]()
		if len(mv.Load()) > 0 {
			t.Errorf("Expected ok to be false, got true")
		}
		mv.Prepend("42", "41")
		v, ok := mv.Pop()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != "41" {
			t.Errorf("Expected value to be 41, got %v", v)
		}
		v, ok = mv.Pop()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != "42" {
			t.Errorf("Expected value to be 42, got %v", v)
		}

	})

	t.Run("new with value", func(t *testing.T) {
		m := []string{"41", "42"}
		mv := mutex.NewSliceWithValue(m)
		v, ok := mv.Shift()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != "41" {
			t.Errorf("Expected value to be 41, got %v", v)
		}

		v, ok = mv.Pop()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != "42" {
			t.Errorf("Expected value to be 42, got %v", v)
		}

	})
	t.Run("new any", func(t *testing.T) {
		m := []string{"41", "42"}
		mv := mutex.NewSliceAny[string]()
		mv.Store(m)
		v, ok := mv.Shift()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != "41" {
			t.Errorf("Expected value to be 41, got %v", v)
		}

		v, ok = mv.Pop()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != "42" {
			t.Errorf("Expected value to be 42, got %v", v)
		}

	})
	t.Run("new any with value", func(t *testing.T) {
		mv := mutex.NewSliceAnyWithValue([]string{"41", "42"})
		v, ok := mv.Shift()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != "41" {
			t.Errorf("Expected value to be 41, got %v", v)
		}

		v, ok = mv.Pop()
		if !ok {
			t.Errorf("Expected ok to be true, got false")
		}
		if v != "42" {
			t.Errorf("Expected value to be 42, got %v", v)
		}

	})
}

func TestEquals(t *testing.T) {
	if mutex.Equals("A", "B") {
		t.Error("A should not match B")
	}
}
