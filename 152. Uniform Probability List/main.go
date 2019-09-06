package main

import (
	"fmt"
	"math/rand"
)

/*
*	You are given n numbers as well as n probabilities that sum up to 1
*
*	Write a function to genrate one of the numers with it corresponding probability
*
*	E.g Given the nunbers [1, 2, 3, 4] and [0.1, 0.5, 0.2, 0.2]
*		Your function should return 1 10% of the time, 2 50% of the time, and 3+4 20% of the time
 */

func main() {

	numbers := []int{1, 2, 3, 4}
	probabilities := []float64{0.1, 0.5, 0.2, 0.2}
	frequencies := make(map[int]int)

	cycles := 100000000
	for i := 0; i < cycles; i++ {
		number := generateNumber(numbers, probabilities)

		_, ok := frequencies[number]

		if ok {
			frequencies[number]++
		} else {
			frequencies[number] = 1
		}
	}

	printResults(frequencies)
}

func generateNumber(numbers []int, probabilities []float64) int {
	randomNumber := rand.Float64()

	probabilitySum := 0.0
	for i := 0; i < len(numbers); i++ {

		probabilitySum += probabilities[i]

		if randomNumber <= probabilitySum {
			return numbers[i]
		}
	}

	fmt.Printf("?? - %f \n", randomNumber)
	return -1
}

func normalise(original *map[int]int) map[int]float64 {
	sum := sum(original)

	normalised := make(map[int]float64)

	for key, value := range *original {
		normalised[key] = float64(value) / float64(sum)
	}

	return normalised
}

func sum(values *map[int]int) int {
	sum := 0

	for _, value := range *values {
		sum += value
	}

	return sum
}

func printResults(frequencies map[int]int) {

	fmt.Println("Original")
	for key, val := range frequencies {
		fmt.Printf("[%d]: %d \n", key, val)
	}

	normalised := normalise(&frequencies)
	fmt.Println("Normalised")
	for key, val := range normalised {
		fmt.Printf("[%d]: %f \n", key, val)
	}
}
