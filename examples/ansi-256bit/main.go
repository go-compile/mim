package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"github.com/go-compile/mim"
)

func main() {
	fingerprint := sha256.Sum256([]byte("certificate contents would typically go here"))

	fmt.Println(mim.New(fingerprint[:]).ANSI256())

	fmt.Println()

	sha512Mozaic()
}

func sha512Mozaic() {
	fingerprint := sha512.Sum512([]byte("certificate contents would typically go here"))

	fmt.Println(mim.New(fingerprint[:]).ANSI256())
}
