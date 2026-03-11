// Sana: 11.03.2026
// Dastur tuzuvchi: Xoliqov Abduvali
// Maqsad: Array yoki slice elementlari ichidan eng kattasini topuvchi dastur tuzing!
package main

import "fmt"

func main() {
	num := []int{5, 60, 20}
	if num[0] < num[1] && num[1] > num[2] {
		fmt.Println(num[1])

	} else if num[1] < num[0] && num[0] > num[2] {
		fmt.Println(num[0])

	} else {
		fmt.Println(num[2])
	}
}
