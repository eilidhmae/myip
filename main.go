package main

import (
	"log"
	"fmt"
)

func main() {
	ip, err := New()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-20s\t%s\n", ip.Address, ip.Label())
}
