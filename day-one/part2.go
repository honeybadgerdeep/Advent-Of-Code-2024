package main

func listToOccMap(list []int) map[int]int {
	result := map[int]int{}

	for _, num := range list {
		if _, ok := result[num]; ok {
			result[num]++
		} else {
			result[num] = 1
		}
	}

	return result
}

func SimilarityScore(left, right []int) int {
	score := 0

	rightOccMap := listToOccMap(right)

	for _, num := range left {
		if occ, exists := rightOccMap[num]; exists {
			score += occ * num
		}
	}

	return score
}
