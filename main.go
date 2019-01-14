package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var str, dnaPattern string
	i := 0
	dna := []string{"A", "C", "G", "T"}
	str = strings.Join(dna, "")

	for j := i; j < len(dna); j++ {
		swaped := swap(str, i, j)
		permutations(swaped, 0, len(swaped), &dnaPattern)
	}

	pattern := make([]string, 0)

	var startString int
	for {
		if startString == len(dnaPattern)-10 {
			break
		}
		var temp string
		newPattern := dnaPattern[startString:]
		for i, a := range newPattern {
			temp += string(a)
			if i < 2 {
				continue
			}
			pattern = append(pattern, temp)

			if i == 9 {
				startString++
				break
				// next curcer
			}
		}
	}

	unique := distinct(pattern)
	for i, p := range unique {
		reg := regexp.MustCompile(p)
		matches := reg.FindAllStringIndex(dnaPattern, -1)
		fmt.Println(i+1, "pattern", fmt.Sprintf(`'%s'`, p), "matches", len(matches))
	}

}

func distinct(arr []string) []string {
	tempMap := make(map[string]string, 0)
	arrResult := make([]string, 0)
	for _, v := range arr {
		tempMap[v] = v
	}

	for key := range tempMap {
		arrResult = append(arrResult, key)
	}

	return arrResult
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
