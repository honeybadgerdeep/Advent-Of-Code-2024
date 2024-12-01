package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument")
		return
	}
	filename := os.Args[1]

	list1, list2, err := ExtractLists(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(TotalDistance(list1, list2))
	fmt.Println(SimilarityScore(list1, list2))
}
