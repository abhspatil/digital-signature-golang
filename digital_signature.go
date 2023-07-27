package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Generate a private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	// Create a hash of the data to be signed
	data := []byte("Pats")
	hash := sha256.Sum256(data)

	// Sign the hash with the private key
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Println(r, s)

	// Verify the signature with the public key
	if ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s) == false {
		panic("Signature verification failed")
	}

	fmt.Println("Signature verified successfully")
}
