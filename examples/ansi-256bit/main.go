package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"github.com/go-compile/mim"
)

func main() {
	fingerprint := sha256.Sum256([]byte("certificate contents would typically go here"))

	fmt.Printf("Fingerprint: %X\n\n", fingerprint)
	fmt.Println(mim.New(fingerprint[:], sha256.New).ANSI())

	fmt.Println()

	fmt.Println(mim.New(fingerprint[:], sha512.New).ANSI())
}
