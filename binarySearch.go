package main

import "fmt"

func Search(numbers []int, desiredNumber int) int {
	firtsValue := numbers[0]
	lastValue := numbers[len(numbers)-1]
	middleValue := (firtsValue + lastValue) / 2

	for i := 0; i < len(numbers); i++ {
		if desiredNumber < middleValue {
			middleValue = lastValue
		} else if desiredNumber > middleValue {
			middleValue = firtsValue
		} else if desiredNumber == middleValue {
			return desiredNumber
		}

		if desiredNumber == numbers[i] {
			fmt.Println("Число есть в массиве")
			return desiredNumber
		}

	}
	return 0
}
