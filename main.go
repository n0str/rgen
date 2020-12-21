package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var letters58 = []rune("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
var letters58_lower = []rune("123456789abcdefghijkmnopqrstuvwxyz")

func randSeq(n int, letters []rune) string {
	if n < 0 {
		n = 10
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var numb int

	mySet := flag.NewFlagSet("",flag.ExitOnError)
	mySet.IntVar(&numb,"n",10,"how many symbols to generate")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s: <command> <arguments>\nCommands: id58, id58l\nArguments:\n", os.Args[0])

		mySet.PrintDefaults()
		os.Exit(0)
	}

	mySet.Parse(os.Args[2:])

	switch os.Args[1] {
	case "id58": fmt.Println(randSeq(numb, letters58))
	case "id58l": fmt.Println(randSeq(numb, letters58_lower))
	}

}
