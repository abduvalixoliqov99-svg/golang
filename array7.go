// Sana: 11.03.2026
// Dastur tuzuvchi: Xoliqov Abduvali
// Maqsad: Array yoki slice elementlari ichidagi juft va toq sonlarni alohida slicelarga yig’uvchi dastur tuzing

package main

import "fmt"

func main() {
	var num []int = []int{4, 5, 9, 2, 3, 1, 7, 8}
	for i := 0; i < len(num); i++ {
		if num[i]%2 == 0 {
			fmt.Println(num[i])

		} else {
			fmt.Println(num[i])
		}
	}

}
