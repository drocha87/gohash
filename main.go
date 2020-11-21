// Simple hash map implementation to help me learn golang

package main

import (
	"fmt"
	"log"
)

type HashPair struct {
	key   string
	value interface{}
}

type HashMap struct {
	capacity int32
	length   int32
	pairs    []*HashPair
	hasher   func(string) int32
}

func default_hasher(key string) int32 {
	log.Fatalln("default_hasher not implemented")
	return 0
}

func New_HashMap() HashMap {
	return New_HashMap(1024)
}

func New_HashMap(capacity int32) HashMap {
	pairs := make([]HashPair, 0, capacity)
	return HashMap{
		capacity: capacity,
		length:   0,
		pairs:    pairs,
		hasher:   default_hasher,
	}
}

func (hm *HashMap) insert(pair HashPair) error {
	index := hm.hasher(pair.key)
	hm.pairs[index] = &pair
	hm.length += 1
	return nil
}

func (hm *HashMap) remove(pair HashPair) error {
	index := hm.hasher(pair.key)
	hm.pairs[index] = nil
	hm.length -= 1
	return nil
}

func (*HashMap) get(pair HashPair) error {
	return nil
}

func main() {
	fmt.Println("Diego Rocha")
}
