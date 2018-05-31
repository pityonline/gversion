package main

import (
	"fmt"
)

func main() {
	version := Version("./")
	fmt.Printf("Current tag is %q.\n", version)
}
