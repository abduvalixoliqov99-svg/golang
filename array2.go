// Sana: 11.03.2026
// Dastur tuzuvchi: Xoliqov Abduvali
// Maqsad: Array yoki slice elementlari yig’indisni topuvchi dastur tuzing

package main

import "fmt"

func main() {
	num := [5]int{1, 5, 6, 5, 9}
	var s int
	s = 0

	for i := 0; i < len(num); i++ {
		s += num[i]
	}
	fmt.Println(s)
}
