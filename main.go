package main

import "fmt"

func main() {
	keysToLookup := []string{
		"abc",
		"Node CC",
		"Test",
		"Bla",
		"im",
		"running",
		"out",
		"of",
		"ideas",
	}

	// withModulus(keysToLookup)

	withConsistentHash(keysToLookup)
}

func withModulus(keysToLookup []string) {
	fmt.Println("Example with modulus")

	nodes := []string{
		"Node A",
		"Node B",
		"Node C",
		"Node D",
		"Node E",
		"Node F",
	}

	lookupKeysWithModulus(nodes, keysToLookup)

	nodes = append(nodes, "Node G")

	lookupKeysWithModulus(nodes, keysToLookup)
}

func withConsistentHash(keysToLookup []string) {
	fmt.Println("Example with consistent hashing")

	ring := NewRing()

	ring.AddNode("Node A")
	ring.AddNode("Node B")
	ring.AddNode("Node C")
	ring.AddNode("Node D")
	ring.AddNode("Node E")
	ring.AddNode("Node F")

	printNodes(ring)
	lookupKeys(ring, keysToLookup)

	ring.AddNode("Node G")

	printNodes(ring)
	lookupKeys(ring, keysToLookup)
}

func printNodes(ring *Ring) {
	fmt.Println(ring.Nodes.Len())
	for i, node := range ring.Nodes {
		fmt.Printf("Node %d: %+v\n", i, node)
	}
}

func lookupKeys(ring *Ring, keys []string) {
	fmt.Println("Fetching keys")

	for _, key := range keys {
		node := ring.Get(key)
		fmt.Printf("%s\t%d\t%+v\n", key, Hash([]byte(key)), node)
	}
}

func lookupKeysWithModulus(nodes []string, keys []string) {
	fmt.Println("Fetching keys")

	for _, key := range keys {
		node := Hash([]byte(key)) % uint64(len(nodes))
		fmt.Printf("%s\t%d\t%+v\n", key, node, nodes[node])
	}
}
