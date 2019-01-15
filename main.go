package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"regexp"
)

func main() {
	var dnaPattern string

	dna := flag.String("input", "ACGT", "a dna value")
	flag.Parse()

	createDNAPattern(*dna, &dnaPattern)

	findKeys := createFindKey(dnaPattern)

	distinct(findKeys)

	mapCounter := counter(dnaPattern, findKeys)

	println(&mapCounter)
}

func createDNAPattern(dna string, dnaPattern *string) {
	i := 0
	for j := i; j < len(dna); j++ {
		swaped := swap(dna, i, j)
		permutations(swaped, 0, len(swaped), dnaPattern)
	}
}

func createFindKey(dnaPattern string) *[]string {
	findKey := make([]string, 0)
	var startString int
	for {
		if startString == len(dnaPattern)-10 {
			break
		}
		var temp string
		subPattern := dnaPattern[startString:]
		for i, a := range subPattern {
			temp += string(a)
			if i < 2 {
				continue
			}
			findKey = append(findKey, temp)

			if i == 9 {
				startString++
				break
				// next curcer
			}
		}
	}
	return &findKey
}

func println(mapCounter *map[string]int) {
	i := 1
	for k, v := range *mapCounter {
		fmt.Printf("%d. pattern '%s' matches %d\n", i, k, v)
		i++
	}
}

func printJSON(mapCounter *map[string]int) {
	bytes, _ := json.Marshal(mapCounter)
	fmt.Println(string(bytes))
}

func counter(dnaPattern string, arr *[]string) map[string]int {
	resultMap := make(map[string]int, 0)
	for _, p := range *arr {
		reg := regexp.MustCompile(p)
		matches := reg.FindAllStringIndex(dnaPattern, -1)
		resultMap[p] = len(matches)
	}
	return resultMap
}

func distinct(arr *[]string) {
	tempMap := make(map[string]string, 0)
	arrResult := make([]string, 0)
	for _, v := range *arr {
		tempMap[v] = v
	}

	for key := range tempMap {
		arrResult = append(arrResult, key)
	}

	*arr = arrResult
}

func swap(s string, i, j int) string {
	var result []byte
	for k := 0; k < len(s); k++ {
		if k == i {
			result = append(result, s[j])
		} else if k == j {
			result = append(result, s[i])
		} else {
			result = append(result, s[k])
		}
	}
	return string(result)
}

// Function to find all Permutations of a given string str[i:n]
// containing all distinct characters
func permutations(str string, i, n int, out *string) {
	// base condition
	if i == n-1 {
		*out += str
		return
	}

	// process each character of the remaining string
	for j := i; j < n; j++ {
		// swap character at index i with current character
		str = swap(str, i, j)

		// recursion for string [i+1:n]
		permutations(str, i+1, n, out)

		// backtrack (restore the string to its original state)
		str = swap(str, i, j)
	}
}
