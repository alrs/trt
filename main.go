package main

import (
	"github.com/aeden/traceroute"
	"github.com/davecgh/go-spew/spew"
	"log"
	"os"
)

func main() {
	host, ok := os.LookupEnv("TRTARGET")
	if !ok {
		log.Fatal("need TRTARGET env variable")
	}
	log.Printf("traceroute to %s", host)
	options := &traceroute.TracerouteOptions{}

	r, err := traceroute.Traceroute(host, options)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(r)
}
