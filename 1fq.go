// Sana: 26.02.2026
// Dastur tuzuvchi: Xoliqov Abduvali
/* Maqsad:  To’gri to’rtburchakni a va b tomonlari berilgan uning yuzini va perimetrini
topuvchi dastur tuzing! (s= a * b, p = 2 * (a + b))
*/
package main

import "fmt"

func main() {
	var a, b int = 4, 5

	s, p := H(a, b)

	fmt.Println("yuza: ", s)
	fmt.Println("Perimetr: ", p)
}
func H(a, b int) (int, int) {
	s := a * b
	p := 2 * (a + b)

	return s, p
}
