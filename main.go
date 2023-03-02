package main

import (
	"log"
	"fmt"
	"github.com/troby/ipcheck"
)

func main() {
	ip, err := New()
	if err != nil {
		log.Fatal(err)
	}

	// get ISP
	wtf, err := ipcheck.WTFMyIP()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-20s\t%-16s\t%s\n", ip.Address, ip.Label(), wtf.YourFuckingISP)
}
