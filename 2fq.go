// Sana: 27.02.2026
// Dastur tuzuvchi: Xoliqov Abduvali
/* Maqsad:	Doiraning radiusi berilgan r, uning aylanasi uzunligi va yuzini
hisblovchi dastur tuzing
( s = pi * r * r, L = 2 * pi * r) pi = 3.14*/

package main

import "fmt"

func main() {
	var r, p float32 = 4, 3.14

	s, l := A(r, p)

	fmt.Println("yuza: ", s)
	fmt.Println("Perimetr: ", l)
}
func A(r, p float32) (float32, float32) {
	s := r * r * p
	l := 2 * r * p

	return s, l
}
