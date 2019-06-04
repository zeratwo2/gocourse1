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

	triangle()
	procent()
}

func triangle() {

	var (
		katet1 float64
		katet2 float64
	)
	var gipoteanuza float64

	fmt.Println("Find S")
	fmt.Println("vvedite katet1")
	fmt.Scanln(&katet1)
	fmt.Println("vvedite katet2")
	fmt.Scanln(&katet2)

	gipoteanuza = math.Sqrt(math.Pow(katet1, 2) + math.Pow(katet2, 2))

	fmt.Printf("%.2f is S, %.2f is gipoteanuza , %.2f is perimetr\n", 0.5*katet1*katet2, gipoteanuza, katet1+katet2+gipoteanuza)
}

func procent() {
	var (
		clientmoney float64
		procent     float64
	)
	fmt.Println("vvedite summu clienta")
	fmt.Scanln(&clientmoney)
	fmt.Println("vvedite godovoi %")
	fmt.Scanln(&procent)
	//	105                  100      5           100
	clientmoney = (clientmoney * procent / 100) + clientmoney
	//	                        5
	//	                                    105
	fmt.Printf("%f vklad 1 goda\n", clientmoney)
	//	110.25               105      5            105
	clientmoney = (clientmoney * procent / 100) + clientmoney
	//                  5.25
	//                               110.25
	fmt.Printf("%f vklad 2 goda \n", clientmoney)
	//  115.7265                110.25   5             110.25
	clientmoney = (clientmoney * procent / 100) + clientmoney
	// 5.5125
	//                 115.7265
	fmt.Printf("%f vklad 3 goda \n", clientmoney)

	//   121.51               115.7265  5           115.7265
	clientmoney = (clientmoney * procent / 100) + clientmoney
	//  5.7863	        121.5
	fmt.Printf("%f vklad 4 goda \n", clientmoney)
	//    127.5855            121.51     5              121.51
	clientmoney = (clientmoney * procent / 100) + clientmoney
	// 6.07       127.5855
	fmt.Printf("%f vklad 5 goda \n", clientmoney)

}
