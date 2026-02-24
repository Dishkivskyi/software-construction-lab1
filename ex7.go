package main

import "fmt"

const (
	MaxInt8Value  = 127
	Custom16Value = 16383
)

func main() {
	variable8 := int8(MaxInt8Value)
	variable16 := int16(Custom16Value)

	fmt.Println("Приведення типів\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	performCalculations()
}

func performCalculations() {
	var intVal int32 = 10
	var floatVal float32 = 5.5

	sum := float32(intVal) + floatVal
	product := float32(intVal) * floatVal

	fmt.Println("\nВиконання завдання:")
	fmt.Printf("Змінна intVal: %d, Змінна floatVal: %f\n", intVal, floatVal)
	fmt.Printf("Сума: %f\n", sum)
	fmt.Printf("Добуток: %f\n", product)
}
