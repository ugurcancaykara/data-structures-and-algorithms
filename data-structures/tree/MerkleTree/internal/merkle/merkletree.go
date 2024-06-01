package merkle

import (
	"crypto/sha256"
	"encoding/hex"
)

// we need a merkle node struct
type MerkleNode struct {
	Children []*MerkleNode // holds the children nodes of the node
	Data     string        // holds the hashed data of the node
}

// we need a merkle tree which holds a list of merklenodes
type MerkleTree struct {
	Root *MerkleNode // holds the root node of the Merkle tree
}

// calculateHashSHA256 calculates the SHA256 hash of the input data
func calculateHash256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// NewMerkleNode creates a new MerkleNode with the given data
func NewMerkleNode(children []*MerkleNode, data string) (*MerkleNode, error) {
	node := &MerkleNode{}
	// check the children count, if its empty then it's a leaf node
	if len(children) == 0 {
		node.Data = calculateHash256(data)
	} else {
		var combinedHash string
		for _, child := range children {
			combinedHash += child.Data // combine the hash of the children
		}

		node.Data = calculateHash256(combinedHash) // calculate the hash of the combined hash and set it as the data of the node
		node.Children = children
	}
	return node, nil
}

// NewMerkleTree creates a new MerkleTree with the given data
func NewMerkleTree(data []string) (*MerkleTree, error) {
	nodes := make([]*MerkleNode, len(data)) // create a slice of MerkleNode with the length of the data

	// create leaf nodes
	for i, datum := range data {
		node, err := NewMerkleNode(nil, datum) // create a new leaf node for each data
		if err != nil {
			return nil, err
		}
		nodes[i] = node
	}

	// now creates the internal nodes until reacing to root node

	for len(nodes) > 1 {
		var newLevel []*MerkleNode

		// create 1 parent from 2 node
		for i := 0; i < len(nodes); i += 2 {
			end := i + 2
			if end > len(nodes) {
				end = len(nodes)
			}
			node, err := NewMerkleNode(nodes[i:end], "") // combining the nodes in pairs
			if err != nil {
				return nil, err
			}
			newLevel = append(newLevel, node) // add the new node to the new level of nodes slice
		}
		nodes = newLevel
	}
	tree := &MerkleTree{nodes[0]} // the last node in the nodes slice is the root node of the Merkle tree
	return tree, nil
}
