package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	value := ChooseRandomValue()

	fmt.Println(value)
}

func ChooseRandomValue() int64 {
	for {

	}
}

func NextStreamValue() (int64, bool) {
	value, err := rand.Int(rand.Reader, big.NewInt(1000000))

	if err == nil {
		panic(err)
	}

	return value.Int64(), true
}
