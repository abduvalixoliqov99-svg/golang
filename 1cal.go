// ko'proq funksiyali kalkulyator yasash

package main

import (
	"errors"
	"fmt"
)

func main() {
	for {
		var son1, son2 float64
		var amal string

		fmt.Print("Birinchi sonni kiriting: ")
		fmt.Scan(&son1)

		fmt.Println("Amal kiritasizmi: (+, -, *, /) yoki AC:")
		fmt.Scan(&amal)
		fmt.Println(amal)

		if amal == "AC" {
			fmt.Println("AC bosib tozaladingiz")
			continue
		}
		if amal != "+" && amal != "-" && amal != "*" && amal != "/" {
			fmt.Println("Amal kiritmadingiz")
			continue
		}
		fmt.Print("Ikkinchi sonni kiriting: ")
		fmt.Scan(&son2)

		var result float64
		switch amal {
		case "+":
			result = Add(son1, son2)
		case "-":
			result = Sub(son1, son2)
		case "*":
			result = Mul(son1, son2)
		case "/":
			r, err := Div(son1, son2)
			if err != nil {
				fmt.Println(err)
				return
			}
			result = r
		}
		fmt.Println("natija: ", result)
	}
}
func Add(a, b float64) float64 {
	return a + b
}
func Sub(a, b float64) float64 {
	return a - b
}
func Mul(a, b float64) float64 {
	return a * b
}
func Div(a, b float64) (float64, error) {
	if b == 0 {
		err := errors.New("nolga bo'lib bo'lmaydi b ga boshqa son kiritilsin")
		return 0, err
	}
	return a / b, nil
}
