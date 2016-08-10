package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
  Example

  # f1
  hello
  there

  # f2
  hello
  goodbye
  cool
  neat

  # output
  there

*/

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: ./main <f1> <f2>")
	}

	f1Path := os.Args[1]
	f2Path := os.Args[2]

	// Create hash from f1
	f1Lines := make(map[string]bool)
	f1, err := os.Open(f1Path)
	if err != nil {
		log.Fatal("Failed to open", f1Path)
	}
	defer f1.Close()

	f1Scanner := bufio.NewScanner(f1)
	for f1Scanner.Scan() {
		f1Lines[f1Scanner.Text()] = true
	}
	if err := f1Scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Iterate over f2 and print all the lines f1 has but f2 doesnt
	f2, err := os.Open(f2Path)
	if err != nil {
		log.Fatal("Failed to open", f2Path)
	}
	defer f2.Close()

	f2Scanner := bufio.NewScanner(f2)
	for f2Scanner.Scan() {
		if _, ok := f1Lines[f2Scanner.Text()]; !ok {
			fmt.Println(f2Scanner.Text())
		}
	}
	if err := f2Scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
