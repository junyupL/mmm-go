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

func rset_map[V any](hasht *Map[V], key []byte, value V) {
	hash := 15 & FNV32a(key)
	for p_node := hasht.rdyn_array[hash]; p_node != nil; p_node = p_node.next {
		if bytes.Equal(p_node.data.key, key) {

			p_node.data.value = value
			return
		}
	}
	insert(&hasht.rdyn_array[hash], 0, Pair[V]{key, value})
	hasht.len++
}

func rget_map[V any](hasht *Map[V], key []byte) V {
	hash := 15 & FNV32a(key)
	for p_node := hasht.rdyn_array[hash]; p_node != nil; p_node = p_node.next {
		if bytes.Equal(p_node.data.key, key) {

			return p_node.data.value
		}
	}
	panic("no value for key provided")
}
