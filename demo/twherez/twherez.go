package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/agustim/wherez"
)

// port for the wherez protocol (UDP+TCP).
const port = 40000

func main() {
	flag.Parse()
	if len(flag.Args()) != 2 {
		log.Fatalln("Usage: demo [options] <app port> <passphrase>")
	}
	appPort, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalf("Invalid port parameter: %v", err)
	}
	passphrase := flag.Arg(1)
	c := wherez.FindAuthenticatedPeers(port, appPort, 1, []byte(passphrase), "10.139.40.82:40000")
	for p := range c {
		// Peer found!
		fmt.Println("peer found:", p.String())
	}
}
