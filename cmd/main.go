package main

import (
	"fmt"
	"os"

	"mtoohey.com/which"
)

func main() {
	for _, v := range os.Args[1:] {
		path, err := which.Which(v)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			defer os.Exit(1)
		}
		fmt.Println(path)
	}
}
