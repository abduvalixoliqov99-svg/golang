// Sana: 26.02.2026
// Dastur tuzuvchi: Xoliqov Abduvali
/* Maqsad:	Doiraning radiusi berilgan r, uning aylanasi uzunligi va yuzini
hisblovchi dastur tuzing
( s = pi * r * r, L = 2 * pi * r) pi = 3.14*/

package main

import "fmt"

func uzl(r float32) {
	var Pi float32 = 3.14
	fmt.Println(2 * Pi * r)

}
func yus(r float32) {
	var Pi float32 = 3.14
	fmt.Println(Pi * r * r)

}
func main() {
	var r float32 = 3
	uzl(r)
	yus(r)
}
