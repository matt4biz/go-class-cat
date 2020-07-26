package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// if there aren't any command-line args,
	// then the loop will not run

	for _, fname := range os.Args[1:] {
		f, err := os.Open(fname)

		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(os.Stdout, f)

		if err != nil {
			log.Fatal(err)
		}
	}
}
