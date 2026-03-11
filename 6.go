// Sana: 11.03.2026
// Dastur tuzuvchi: Xoliqov Abduvali
// Maqsad: Array yoki slice elementlari ichidan juftlarini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	num := []int{6, 3, 9, 45, 6, 2}

	for i := 0; i < len(num); i++ {
		if num[i]%2 == 0 {
			fmt.Println(num[i])
		}
	}

}
