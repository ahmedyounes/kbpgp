
package main

import (
	"golang.org/x/crypto/openpgp"
	"strings"
	"fmt"
)

func main() {
	portwood := `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: SKS 1.1.5
Comment: Hostname: pgp.mit.edu

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
6/qHgQG3LRJDFdxR81/RSYc/8qk=
=Zjq/
-----END PGP PUBLIC KEY BLOCK-----`
	reader := strings.NewReader(portwood)
	el, err := openpgp.ReadArmoredKeyRing(reader)
	if err != nil {
		fmt.Printf("error reading key: %v\n", err)
	} else {
		fmt.Printf("Got %d keys\n", len(el))
	}
}