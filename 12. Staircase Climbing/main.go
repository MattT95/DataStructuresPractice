package main

import "fmt"

// This problem was asked by Amazon.
// There exists a staircase with N steps, and you can climb up either 1 or 2 steps at a time. Given N, write a function that returns the number of unique ways you can climb the staircase. The order of the steps matters.
// For example, if N is 4, then there are 5 unique ways:
// •	1, 1, 1, 1
// •	2, 1, 1
// •	1, 2, 1
// •	1, 1, 2
// •	2, 2
// What if, instead of being able to climb 1 or 2 steps at a time, you could climb any number from a set of positive integers X? For example,
// if X = {1, 3, 5}, you could climb 1, 3, or 5 steps at a time

func main() {

	steps := 50

	// possibilities := CountStaircasePossibilities(steps)

	// fmt.Println(possibilities)

	memo := make([]int, steps)
	for idx, _ := range memo {
		memo[idx] = -1
	}

	// possibilities := CountStaircasePossibilities2(steps, []int{1, 2, 3, 4}, memo)
	possibilities := CountStaircasePossibilitiesNoMemo(steps, []int{1, 2, 3, 4})

	fmt.Println(possibilities)
}

func CountStaircasePossibilitiesNoMemo(steps int, possibleSteps []int) int {

	count := 0

	if steps == 0 {
		return 1
	}

	for _, val := range possibleSteps {

		if steps >= val {
			count += CountStaircasePossibilitiesNoMemo(steps-val, possibleSteps)
		}
	}

	return count
}

func CountStaircasePossibilities2(steps int, possibleSteps []int, memo []int) int {

	count := 0

	if steps == 0 {
		return 1
	}

	if memo[steps-1] != -1 {
		return memo[steps]
	}

	for _, val := range possibleSteps {

		if steps >= val {
			count += CountStaircasePossibilities2(steps-val, possibleSteps, memo)
		}
	}

	memo[steps-1] = count

	return count
}

func CountStaircasePossibilities(steps int) int {

	if steps == 1 || steps == 0 {
		return 1
	}

	if steps == 2 {
		return 2
	}

	return CountStaircasePossibilities(steps-1) + CountStaircasePossibilities(steps-2)
}
