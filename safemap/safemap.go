package safemap

import "sync"

type Key interface {
	~string | ~int | ~uint | ~uintptr | ~float32 | ~float64
}

type SafeMap[K Key, V any] struct {
	entry *sync.Map
	l     int
}

func New[K Key, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		entry: new(sync.Map),
		l:     0,
	}
}

func (m *SafeMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.entry.Load(key)
	if ok {
		value = v.(V)
	}
	return
}

func (m *SafeMap[K, V]) Store(key K, value V) {
	m.entry.Store(key, value)
	m.l++
}

func (m *SafeMap[K, V]) Delete(key K) {
	m.entry.Delete(key)
	m.l--
}

func (m *SafeMap[K, V]) Range(f func(key K, value V) bool) {
	m.entry.Range(func(key, value interface{}) bool {
		return f(key.(K), value.(V))
	})
}

func (m *SafeMap[K, V]) Len() int {
	return m.l
}

func (m *SafeMap[K, V]) Keys() []K {
	keys := make([]K, 0, m.Len())
	m.Range(func(key K, value V) bool {
		keys = append(keys, key)
		return true
	})
	return keys
}

func (m *SafeMap[K, V]) Values() []V {
	values := make([]V, 0, m.Len())
	m.Range(func(key K, value V) bool {
		values = append(values, value)
		return true
	})
	return values
}

func (m *SafeMap[K, V]) Clear() {
	m.entry = new(sync.Map)
	m.l = 0
}

func (m *SafeMap[K, V]) Clone() *SafeMap[K, V] {
	clone := New[K, V]()
	m.Range(func(key K, value V) bool {
		clone.Store(key, value)
		return true
	})
	return clone
}

func (m *SafeMap[K, V]) Merge(other *SafeMap[K, V]) {
	other.Range(func(key K, value V) bool {
		m.Store(key, value)
		return true
	})
}
