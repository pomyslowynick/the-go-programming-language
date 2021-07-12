package main

import "fmt"

// Exercise 5.16
func main() {
	result := VariadicJoin(" ", []string{"Hello", "I", "like", "it"}, []string{"What", "I", "don't", "like", "it"})
	fmt.Println(result)
}

func VariadicJoin(sep string, stringSlices ...[]string) string {
	finalString := ""

	for _, stringSlice := range stringSlices {
		for index, element := range stringSlice {
			finalString += element
			if index != len(stringSlice)-1 {
				finalString += sep
			}
		}
	}

	return finalString
}
