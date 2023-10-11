package safemap

import (
	"strconv"
	"sync"
	"testing"
)

func TestGenericSafeMap(t *testing.T) {
	var m = New[string, int]()
	var i = 25
	m.Store("a", i)
	i2, ok := m.Load("a")
	if !ok {
		t.Errorf("Load() failed")
	}
	if i2 != i {
		t.Errorf("Load() failed")
	}
	m.Store("b", 26)
	if m.Len() != 2 {
		t.Errorf("Len() failed")
	}
	m.Delete("a")
	if m.Len() != 1 {
		t.Errorf("Delete() failed")
	}
	m.Range(func(key string, value int) bool {
		if key != "b" || value != 26 {
			t.Errorf("Range() failed")
		}
		return true
	})
	var newM = New[string, int]()
	newM.Store("a", 1)
	m.Merge(newM)
	i3, ok := m.Load("a")
	if !ok {
		t.Errorf("Merge() failed")
	}
	if i3 != 1 {
		t.Errorf("Merge() failed")
	}
}

func TestConcurrentSafeMap(t *testing.T) {
	var m = New[string, int]()
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(strconv.Itoa(i), i)
		}(i)
	}
	m.Range(func(key string, value int) bool {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Delete(key)
		}()
		return true
	})
	wg.Wait()
}
