package main

import "fmt"

func main() {
	var nilai = 10

	switch {
	case nilai > 0:
		fmt.Println("lebih dari 0")
		fallthrough
	case nilai > 5:
		fmt.Println("lebih dari 5")
	}
}
