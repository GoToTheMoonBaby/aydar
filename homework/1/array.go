// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func BubbleSort(arr []int) (res []int) {
	l := len(arr) - 1
	for i := 0; i < l; i++ {
		if arr[i] > arr[i+1] {
			tmp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = tmp
			i = -1
		}
	}

	return arr
}

func main() {
	arr := []int{48, 96, 178, 68, 4}

	fmt.Println(BubbleSort(arr))

}
