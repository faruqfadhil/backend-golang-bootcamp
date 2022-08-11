package main

import "fmt"

func main() {
	fmt.Println("menggunakan variable counter beserta kondisinya.")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("dengan hanya kondisi")
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("tanpa argumen")
	i = 0
	for {
		if i > 5 {
			break
		}
		fmt.Println(i)
		i++
	}

	fmt.Println("penggunaan continue")
	i = 0
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("pengunaan label")

outerlopp:
	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			if i == 3 {
				break outerlopp
			}
			fmt.Printf("i = %d | j= %d\n", i, j)
		}
	}

}
