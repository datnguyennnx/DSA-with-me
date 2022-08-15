package main

import "fmt"

func GenDisplaceFn(a, v, s float64) func(float64) float64 {

	return func(t float64) float64 {
		return (0.5*a*(t*t) + v*t + s)
	}
}

func main() {
	input := [4]float64{}
	fmt.Print("Input acceleration:  ")
	fmt.Scanln(&input[0])
	fmt.Print("Input velocity :  ")
	fmt.Scanln(&input[1])
	fmt.Print("Input displacement :  ")
	fmt.Scanln(&input[2])
	fmt.Print("Input time :  ")
	fmt.Scanln(&input[3])

	fn := GenDisplaceFn(input[0], input[1], input[2])
	fmt.Println("Result = ", fn(input[3]))

}
