package main

import (
	"context"
	"fmt"
	"os"

	"github.com/nzoschke/codon/pkg/run"
)

func main() {
	ctx := context.Background()
	if err := run.Run(ctx, os.Args, os.Getenv, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
