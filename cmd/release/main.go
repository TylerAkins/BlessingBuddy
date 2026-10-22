package main

import (
	"fmt"
	"os"

	"github.com/TylerAkins/BlessingBuddy/internal/release"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] != "validate" {
		fmt.Fprintln(os.Stderr, "usage: go run ./cmd/release validate")
		os.Exit(2)
	}
	version, err := release.ValidateRepository(".")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(version)
}
