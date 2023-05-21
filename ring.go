package main

import (
	"errors"
	"sort"
	"sync"

	"github.com/spaolacci/murmur3"
)

const RingSize = 360

type Node struct {
	Key    string
	HashId uint64
}

func NewNode(key string) *Node {
	return &Node{
		Key:    key,
		HashId: Hash([]byte(key)),
	}
}

// Creating a type for a list often makes sense
// it allows you to create methods on top of the list
type Nodes []*Node

func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Nodes) Less(i, j int) bool { return n[i].HashId < n[j].HashId }

func Hash(data []byte) uint64 {
	return murmur3.Sum64(data) % RingSize
}

type Ring struct {
	Nodes Nodes
	sync.Mutex
}

func NewRing() *Ring {
	return &Ring{Nodes: Nodes{}}
}

func (r *Ring) AddNode(id string) {
	r.Lock()
	defer r.Unlock()

	node := NewNode(id)
	r.Nodes = append(r.Nodes, node)

	sort.Sort(r.Nodes)
}

func (r *Ring) RemoveNode(key string) error {
	r.Lock()
	defer r.Unlock()

	i := r.search(key)
	if i >= r.Nodes.Len() || r.Nodes[i].Key != key {
		return errors.New("node not found")
	}

	r.Nodes = append(r.Nodes[:i], r.Nodes[i+1:]...)

	return nil
}

func (r *Ring) Get(key string) string {
	i := r.search(key)
	if i >= r.Nodes.Len() {
		i = 0
	}

	return r.Nodes[i].Key
}

func (r *Ring) search(key string) int {
	searchfn := func(i int) bool {
		return r.Nodes[i].HashId >= Hash([]byte(key))
	}

	return sort.Search(r.Nodes.Len(), searchfn)
}
