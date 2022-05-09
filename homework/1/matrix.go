// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func transpose(a [][]int) [][]int {
	newArr := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}

	return newArr
}

func main() {
	var a = [][]int{
		{6, 11, 4},
		{3, 5, 8},
		{33, 9, 2},
	}

	fmt.Println(transpose(a))
}
