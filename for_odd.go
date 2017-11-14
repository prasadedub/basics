package main

import "fmt"

func main() {
	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("hello gopher\n")
		//		continue
		break
	}
	for j := 0; j <= 9; {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
		j++
	}
}
