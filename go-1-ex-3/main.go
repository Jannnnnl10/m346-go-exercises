package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello, Output!")
	fmt.Fprintln(os.Stderr, "Hello, Error!")
}
