package main

import (
	"fmt"
	"math"
)

func main() {
	const dollar float64 = 65.17
	var rub float64

	fmt.Println("vvedite summu rublei:")
	fmt.Scanln(&rub)
	fmt.Println("otvet:", rub/dollar)

	add()
}

func add() {

	var (
		katet1 float64
		katet2 float64
	)
	var gipoteanuza float64

	fmt.Println("Find S ?:")
	fmt.Scanln(&katet1)
	fmt.Scanln(&katet2)

	gipoteanuza = math.Sqrt(math.Pow(katet1, 2) + math.Pow(katet2, 2))

	fmt.Printf("%.2f is S, %.2f is gipoteanuza , %.2f is perimetr\n", 0.5*katet1*katet2, gipoteanuza, katet1+katet2+gipoteanuza)
}
