package main

import (
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	
	tests := []struct {
		input string
		expected int
	}{
		{ "", 0 },
		{ "1", 1 },
		{ "111", 3 },
		{ "131", 3 },
	}

	for _, test := range tests {
		actual := CountDecodesDP(test.input)
		
		if actual != test.expected {
			t.Errorf("Got %d, expected %d", actual, test.expected)
		}
	}
}

// Complexity of O(2^n)
// Think of as a binary tree, adding double the number of nodes
func CountDecodes(input string) int {
	
	if input == "" {
		return 0;
	}

	if len(input) == 1 {
		return 1
	}	

	occurrences := 0

	occurrences = 1 + CountDecodes(input[1:])

	if parsedTwoChars, _ := strconv.ParseInt(input[0:1], 10, 32); parsedTwoChars >= 10 && parsedTwoChars <= 26 {
		occurrences += 1 + CountDecodes(input[2:])
	}   

	return occurrences
}

// By using memoization, we save replayign computation
func CountDecodesDP(input string) int {

	memo := make([]int, len(input))

	for idx, _ := range memo {
		memo[idx] = -1
	}

	return Helper(input, memo);	
}


func Helper(input string, memo []int) int {
	
	if input == "" {
		return 0;
	}

	if len(input) == 1 {
		return 1
	}	

	if memo[len(input)-1] != -1 {
		return memo[len(input)-1]
	}

	occurrences := 0

	occurrences = 1 + CountDecodesDP(input[1:])

	if parsedTwoChars, _ := strconv.ParseInt(input[0:1], 10, 32); parsedTwoChars >= 10 && parsedTwoChars <= 26 {
		occurrences += 1 + CountDecodesDP(input[2:])
	}   

	memo[len(input)-1] = occurrences 
	return occurrences
}