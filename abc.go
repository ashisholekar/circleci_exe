package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")
	fmt.Fprintln(os.Stdout, 13.5)
}
