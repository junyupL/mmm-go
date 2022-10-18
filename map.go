package main

import (
	"bytes"
	"hash/fnv"
)

type Pair[V any] struct {
	key   []byte
	value V
}

func FNV32a(bytes []byte) uint32 {
	algorithm := fnv.New32a()
	algorithm.Write(bytes)
	return algorithm.Sum32()
}
func set_map[V any](hasht *dyn_array[*Node[Pair[V]]], key []byte, value V) {
	hash := 15 & FNV32a(key)
	for p_node := (*hasht)[hash]; p_node != nil; p_node = p_node.next {
		if bytes.Equal(p_node.data.key, key) {

			p_node.data.value = value
			return
		}
	}
	insert(&(*hasht)[hash], 0, Pair[V]{key, value})
}

func rset_map[V any](hasht *rdyn_array[*Node[Pair[V]]], key []byte, value V) {
	hash := 15 & FNV32a(key)
	for p_node := (*hasht)[hash]; p_node != nil; p_node = p_node.next {
		if bytes.Equal(p_node.data.key, key) {

			p_node.data.value = value
			return
		}
	}
	insert(&(*hasht)[hash], 0, Pair[V]{key, value})
}

func get_map[V any](hasht *dyn_array[*Node[Pair[V]]], key []byte) V {
	hash := 15 & FNV32a(key)
	for p_node := (*hasht)[hash]; p_node != nil; p_node = p_node.next {
		if bytes.Equal(p_node.data.key, key) {

			return p_node.data.value
		}
	}
	panic("no value for key provided")
}

func rget_map[V any](hasht *rdyn_array[*Node[Pair[V]]], key []byte) V {
	hash := 15 & FNV32a(key)
	for p_node := (*hasht)[hash]; p_node != nil; p_node = p_node.next {
		if bytes.Equal(p_node.data.key, key) {

			return p_node.data.value
		}
	}
	panic("no value for key provided")
}
