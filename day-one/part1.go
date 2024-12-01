package main

import (
	"fmt"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left...)
			left = nil
		} else if len(right) > 0 {
			result = append(result, right...)
			right = nil
		}
	}
	return result
}

func mergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	mid := len(list) / 2
	left := mergeSort(list[:mid])
	right := mergeSort(list[mid:])

	return merge(left, right)
}

func ExtractNumbers(line string) (int, int, error) {
	numbers := []int{}

	for _, numStr := range strings.Fields(line) {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println(err)
			return 0, 0, err
		}
		numbers = append(numbers, num)
	}

	return numbers[0], numbers[1], nil
}

func TotalDistance(left, right []int) int {
	// Sort list1 and list2
	left = mergeSort(left)
	right = mergeSort(right)

	totalDistance := 0

	for i := 0; i < len(left); i++ {
		totalDistance += abs(left[i] - right[i])
	}

	return totalDistance
}
