// Sana: 11.03.2026
// Dastur tuzuvchi: Xoliqov Abduvali
// Maqsad: Array yoki slice elementlari ichidan toqlarnini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	var num [4]int = [4]int{5, 6, 9, 4}

	for i := 0; i < len(num); i++ {

		if num[i]%2 != 0 {
			fmt.Println(num[i])
		}
	}
}
