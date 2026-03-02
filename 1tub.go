//Sana: 2026.01.03
// Dasturr tuzuvchi: Xoliqov Abduvali;
//Maqsad: kiritilgan songacha mavjud bo'lgan tub sonlarni chop etuvchi dastur!

package main

import "fmt"

func main() {

	var a int

	fmt.Print("Sonni kiriting:= ")
	fmt.Scan(&a)
	fmt.Println("a gacha bo'lgan tub sonlar:")

	var tub bool

	for i := 2; i < a; i++ {
		tub = true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				tub = false
				break
			}
		}
		if tub {
			fmt.Println(i)
		}
	}

}
