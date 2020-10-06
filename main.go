package main

import (
	"fmt"

	m "github.com/therealfakemoot/gomarkov"
)

func main() {
	c := m.NewChain(2)
	fmt.Printf("%+v", c)
}
