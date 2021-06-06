package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	wordsCount := 0
	countMap := make(map[string]int)
	frequencyMap := make(map[string]float32)

	input.Split(bufio.ScanWords)
	for input.Scan() {
		line := input.Text()
		countMap[line] += 1
	}

	wordsCount = len(countMap)
	
	for key, value := range countMap {
		frequencyMap[key] = float32(value) / float32(wordsCount)
	}

	fmt.Println(frequencyMap)
}