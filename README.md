##### Installation
Dowload [Golang 1.10.7](https://golang.org/dl/) and install.

##### Clone Repository
Go to your workspace and clone repository

`git clone git@github.com:ker13530018/tandem-repeat.git`


##### Build or RUN

1. go to project folder follow this command `cd tandem-repeat`

2. if your want to run please follow this command `go run main.go -input=ACGT`

3. if your want to build please follow this command `go build -o tandem-repeat.sh main.go`  and run `
./tandem-repeat.sh -input=ACGT > out.log`

4. if your want to run test follow this command `go test -v --cover  ./...`


##### Explain codeing

1. I create DNA pattern and save the output to `dnaPattern` parameter.

```golang
func createDNAPattern(dna string, dnaPattern *string) {
	i := 0
	for j := i; j < len(dna); j++ {
		swaped := swap(str, i, j)
		permutations(swaped, 0, len(swaped), dnaPattern)
	}
}

func permutations(str string, i, n int, out *string) {
	// base condition
	if i == n-1 {
		*out += str
		return
	}

	for j := i; j < n; j++ {
		str = swap(str, i, j)

		permutations(str, i+1, n, out)

		str = swap(str, i, j)
	}
}

```


2. I create forever loop for get string pattern and store pattern to `findKey` array.

```golang
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
```

3. I find the unique findKey.

```golang
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
```

4. Find the match pattern in `dnaPattern` and display output.
```golang
func println(mapCounter *map[string]int) {
	i := 1
	for k, v := range *mapCounter {
		fmt.Printf("%d. pattern '%s' matches %d\n", i, k, v)
		i++
	}
}
```