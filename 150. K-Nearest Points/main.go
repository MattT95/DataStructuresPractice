package main

import (
	"fmt"
)

/*
*	Given a list of points, a central point, and an integer k
*	Find the nearest k points from the central point
*
*	For example, given
*		[(0,0), (5,4), (3,1)],
*		the central point (1,2)
*		k = 2
*
*	Return [(0,0), (3,1)]
 */

func main() {
	points := []Point{Point{0, 0}, Point{5, 4}, Point{3, 1}, Point{10, 10}, Point{10, 10}, Point{10, 10}, Point{10, 10}, Point{10, 10}}
	centralPoint := Point{1, 2}
	k := 4

	nearestPoints := BruteForceFindNearestPoints(points, centralPoint, k)

	for _, val := range nearestPoints {
		fmt.Printf("(%d, %d) ", (*val).X, (*val).Y)
	}
}

func BruteForceFindNearestPoints(points []Point, centralPoint Point, k int) []*Point {

	// We could iterate through the list and calculate the distance (euclidean)
	// To the central
	// Based off this, we can compare with our other min points and see if this is
	// nearer than points that we have considered as closest candidates so far
	// The problem with this solution is that it is largely bounded in time to k.

	// We have to iterate over every point to compute the distance, which is expected in all solutions
	// However, with this case we also have to compare this distance to the distance of all the points that we
	// are considering as k-nearest at that given point in the solution
	// Therefore, this solution is of time complexity O(n * k)
	// where n is the number of points we are given on input, and k is the number of nearest we wish to return
	nearestPoints := make([]*Point, k)

	for i := 0; i < len(points); i++ {
		point := points[i]
		distance := CalculateDistanceMetric(point, centralPoint)

		positionToAdd := -1
		for j := 0; j < len(nearestPoints); j++ {
			nearestPoint := nearestPoints[j]

			if nearestPoint == nil {
				positionToAdd = j
				break
			}

			nearestPointDistance := CalculateDistanceMetric(*nearestPoints[j], centralPoint)

			if distance < nearestPointDistance {
				positionToAdd = j
				distance = nearestPointDistance
			}
		}

		if positionToAdd != -1 {
			nearestPoints[positionToAdd] = &point
		}
	}

	return nearestPoints
}

func CalculateDistanceMetric(pointA Point, pointB Point) int {
	xDistance := pointA.X - pointB.X
	yDistance := pointA.Y - pointB.Y

	return (xDistance * xDistance) + (yDistance * yDistance)
}

type Point struct {
	X int
	Y int
}

func HeapFindNearestPoints(points []Point, centralPoint Point, k int) []*Point {

	Heapify(&points, centralPoint)

	nearestPoints := make([]*Point, k)

	for i := 0; i < len(nearestPoints); i++ {
		nearestPoints[i] = Extract(&points)
		Heapify(&points, centralPoint)
	}

	return nearestPoints
}

func Heapify(points *[]Point, centralPoint Point) {
	for i := len(*points) / 2; i >= 0; i-- {
		SiftDown(points, i, centralPoint)
	}
}

func SiftDown(points *[]Point, rootIdx int, centralPoint Point) {
	leftIdx := (rootIdx * 2) + 1
	rightIdx := (rootIdx * 2) + 2
	nearestIdx := rootIdx
	nearestDistance := CalculateDistanceMetric((*points)[nearestIdx], centralPoint)

	if leftIdx < len(*points) {
		leftDistance := CalculateDistanceMetric((*points)[leftIdx], centralPoint)
		if leftDistance < nearestDistance {
			nearestIdx = leftIdx
			nearestDistance = leftDistance
		}
	}

	if rightIdx < len(*points) {
		rightDistance := CalculateDistanceMetric((*points)[rightIdx], centralPoint)
		if rightDistance < nearestDistance {
			nearestIdx = rightIdx
			nearestDistance = rightDistance
		}
	}

	if nearestIdx != rootIdx {
		temp := (*points)[rootIdx]
		(*points)[rootIdx] = (*points)[nearestIdx]
		(*points)[nearestIdx] = temp

		SiftDown(points, nearestIdx, centralPoint)
	}
}

func Extract(points *[]Point) *Point {
	pointToTake := (*points)[0]

	(*points)[0] = (*points)[len(*points)-1]

	return &pointToTake
}
