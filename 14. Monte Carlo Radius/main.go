package main

import "fmt"

// Good morning! Here's your coding interview problem for today.
// This problem was asked by Google.
// The area of a circle is defined as πr^2.
//
// Estimate π to 3 decimal places using a Monte Carlo method.
// Hint: The basic equation of a circle is x2 + y2 = r2.

// pi = r^2 / area
// pi = (x^2) + (y^2) / area

func main() {

	// 100 x 100 square. Area => 10000
	// The area of the circle that fits the edges of this square. 50^2 * pi => 2500 * pi

	// Ratio of space between circle and square :
	// pi * 2500 / 10000 => pi / 4
	//
	// So, pi = (Ratio of space between points) x 4
	// We'll sample a bunch of points within a circle x=1, y=1, and track the proportion of points

	result := CalculatePiMonteCarlo()

	fmt.Println(result)
}

func CalculatePiMonteCarlo() float64 {

	squareSize := 100000
	squareArea := squareSize * squareSize
	innerCircleRadius := squareSize / 2
	radiusSquared := float64(innerCircleRadius) * float64(innerCircleRadius)

	inCircleCount := 0
	count := 0

	for i := 0; i < squareSize; i++ {
		for j := 0; j < squareSize; j++ {
			pointRadiusSquared := (i * i) + (j * j)

			if float64(pointRadiusSquared) >= radiusSquared {
				inCircleCount++
			}
			count++
		}
	}

	fmt.Println(inCircleCount)
	fmt.Println(squareArea)
	ratio := float64(inCircleCount) / float64(squareArea)
	return ratio * 4
}
