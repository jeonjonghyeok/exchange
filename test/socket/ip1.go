package main

import (
	. "fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	addr := net.ParseIP(name)
	if addr == nil {
		Println("Invalid address")
	} else {
		Println("The address is ", addr.String())
	}
	os.Exit(0)
}
