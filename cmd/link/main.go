package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/movaua/link/pkg/link"
)

func main() {
	filename := flag.String("file", "", "an HTML file name")
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	links, err := link.Find(f)
	if err != nil {
		log.Fatal(err)
	}

	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "  ")
	if err := e.Encode(links); err != nil {
		log.Fatal(err)
	}
}
