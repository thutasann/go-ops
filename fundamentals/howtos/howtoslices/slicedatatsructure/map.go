package slicedatatsructure

import "fmt"

type KVPair[K comparable, V any] struct {
	Key   K
	Value V
}

type SimpleMap[K comparable, V any] struct {
	data []KVPair[K, V]
}

func (m *SimpleMap[K, V]) Set(key K, val V) {
	for i, kv := range m.data {
		if kv.Key == key {
			m.data[i].Value = val
			return
		}
	}
	m.data = append(m.data, KVPair[K, V]{key, val})
}

func (m *SimpleMap[K, V]) Get(key K) (V, bool) {
	for _, kv := range m.data {
		if kv.Key == key {
			return kv.Value, true
		}
	}
	var zero V
	return zero, false
}

func Map_Slice_DSA() {
	fmt.Println("\n==> Map Slice DSA")
	m := &SimpleMap[string, int]{}
	m.Set("a", 1)
	m.Set("b", 2)
	fmt.Println(m.Get("a")) // 1, true
}
