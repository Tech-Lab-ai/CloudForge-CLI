package main

import (
	"fmt"
	"os"

	"cloudforge-backend/internal/engine"
)

func main() {
	if err := engine.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start CloudForge Engine: %s\n", err)
		os.Exit(1)
	}
}
