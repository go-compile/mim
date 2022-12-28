# MIM (Minimal Identity Mozaic)

MIM is a Hash Visualization format utilising 4x4 colour matrixes. This provides a quick and easy method to compare fingerprints, e.g. SSH keys, x509 certs etc.

## Properties
- Pre Image Resistant
- Fixed Length Output
- Collision Resistant
- Fast & Efficient
- Customisable CHF
- Identical Colours Cross Platform

## Output

MIM outputs coloured **ANSI escape codes** or **SVGs**.

![MIM Mozaic output](./.github/images/mim2.png)

## Example

```go
package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/go-compile/mim"
)

func main() {
	fingerprint := sha256.Sum256([]byte("certificate contents would typically go here"))

	fmt.Printf("Fingerprint: %X\n\n", fingerprint)
	
	fmt.Println(mim.New(fingerprint[:], sha256.New).ANSI())
}
```