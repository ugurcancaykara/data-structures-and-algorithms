package main

import (
	"fmt"
	"github.com/ugurcancaykara/data-structures-and-algorithms/data-structures/tree/MerkleTree/internal/merkle"
	"log"
	"strconv"
)

func main() {
	fmt.Print("An example of Merkle Tree\n")

	data := make([]string, 100)
	for i := range data {
		data[i] = "data" + strconv.Itoa(i)
	}
	tree, err := merkle.NewMerkleTree(data)
	if err != nil {
		log.Fatalf("Failed to create Merkle Tree: %v", err)
	}
	fmt.Println("Root hash of the merkle tree: ", tree.Root.Data)
}
