package main

import (
	"fmt"
	"crypto/sha1"
)

func main() {
	// Create the hashes
	h := sha1.New() 
	// Write into the hasher
	h.Write([]byte("data"))
	v := h.Sum([]byte{})
	//Print the result
	fmt.Println(v)
}