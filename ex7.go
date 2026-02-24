package main

import "fmt"

func main() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведення типів\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	// Завдання.
	// 1. Створіть 2 змінні різних типів. Виконайте арифметичні операції.

	var a int32 = 10
	var b float32 = 5.5

	// Для виконання операції приводимо 'a' до типу float32
	sum := float32(a) + b
	product := float32(a) * b

	fmt.Println("\nВиконання завдання:")
	fmt.Printf("Змінна a (int32): %d, Змінна b (float32): %f\n", a, b)
	fmt.Printf("Сума (float32(a) + b): %f\n", sum)
	fmt.Printf("Добуток (float32(a) * b): %f\n", product)
}
