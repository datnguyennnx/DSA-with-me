package main

import (
	"bytes"
	"fmt"
)

type slice struct {
	arr1 []int
}

func Swap(numbers []int, i, j int) {
	tempnum := numbers[i]
	numbers[i] = numbers[j]
	numbers[j] = tempnum
}

func BubbleSort(numbers []int) []int {
	for j := 0; j < len(numbers); j++ {
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				Swap(numbers, i, i+1)
			}
		}
	}
	return numbers
}

func (s *slice) Scan(state fmt.ScanState, verb rune) error {
	arr1, err := state.Token(false, func(r rune) bool { return r != '\n' })
	if err != nil {
		return err
	}
	if _, _, err := state.ReadRune(); err != nil {
		if len(arr1) == 0 {
			panic(err)
		}
	}
	b := bytes.NewReader(arr1)
	for {
		var d int
		_, err := fmt.Fscan(b, &d)
		if err != nil {
			break
		}
		s.arr1 = append(s.arr1, d)
	}
	return nil
}

func main() {
	fmt.Print("Give me a array need sorting: ")
	var arr slice
	fmt.Scan(&arr)
	fmt.Println(BubbleSort(arr.arr1))
}
