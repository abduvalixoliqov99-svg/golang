// Sana: 11.03.2026
// Dastur tuzuvchi: Xoliqov Abduvali
// Maqsad: Array yoki slice elementlarini teskari tartibda dastur tuzing!

package main

import "fmt"

func main() {

	num := []int{1, 5, 8, 6, 4, 2}

	for i := len(num) - 1; i >= 0; i-- {
		fmt.Println(num[i])
	}
}
