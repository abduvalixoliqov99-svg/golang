// Sana: 26.02.2026
// Dastur tuzuvchi: Xoliqov Abduvali
/* Maqsad:  To’gri to’rtburchakni a va b tomonlari berilgan uning yuzini va perimetrini
topuvchi dastur tuzing! (s= a * b, p = 2 * (a + b))
*/
package main

import "fmt"

func yuza(x, y int) {
	fmt.Println(x * y)

}
func perm(x, y int) {
	fmt.Println(2 * (x + y))
}

func main() {
	var a, b int = 3, 4
	yuza(a, b)
	perm(a, b)
}
