package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Read the input file
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert hex-encoded transactions to byte slices
	var transactions [][]byte
	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		transaction, err := hex.DecodeString(string(line))
		if err != nil {
			log.Fatal(err)
		}
		transactions = append(transactions, transaction)
	}

	// Compute the Merkle Tree Root
	root := computeMerkleRoot(transactions)
	fmt.Println(hex.EncodeToString(root))
}

// Computes the Merkle Tree Root for the given transactions
func computeMerkleRoot(transactions [][]byte) []byte {
	if len(transactions) == 0 {
		return nil
	}

	// Compute the intermediate nodes of the Merkle Tree
	var nodes [][]byte
	for _, transaction := range transactions {
		hash := sha256.Sum256(transaction)
		nodes = append(nodes, hash[:])
	}
	for len(nodes) > 1 {
		if len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}
		var level []byte
		for i := 0; i < len(nodes); i += 2 {
			hash := sha256.Sum256(append(nodes[i], nodes[i+1]...))
			level = append(level, hash[:]...)
		}
		nodes = level
	}

	return nodes[0]
}
