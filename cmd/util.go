package cmd

import (
	"fmt"
	"os"
)

func exitOnError(err error) {
	fmt.Printf("error: %v\n", err)
	os.Exit(1)
}
