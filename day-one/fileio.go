package main

import (
	"bufio"
	"os"
	"strings"
)

func ExtractLists(filename string) ([]int, []int, error) {
	list1 := []int{}
	list2 := []int{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close() // This ensures that the file is closed when the surrounding function (main) returns, freeing resources.

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			continue
		}

		num1, num2, err := ExtractNumbers(line)

		if err != nil {
			return nil, nil, err
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	return list1, list2, nil
}
