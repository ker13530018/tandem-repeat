##### Installation
Dowload [Golang 1.10.7](https://golang.org/dl/) and install.

##### Clone Repository
Go to your workspace and clone repository

`git clone git@github.com:ker13530018/tandem-repeat.git`


##### Build or RUN

1. go to project folder follow this command `cd tandem-repeat`

2. if your want to run please follow this command `go run main.go`

3. if your want to build please follow this command `go build -o tandem-repeat.sh main.go`  and run `
./tandem-repeat.sh > out.log`


##### Explain codeing

1. I create DNA pattern and save the output to `dnaPattern` parameter.

```golang
for j := i; j < len(dna); j++ {
    swaped := swap(str, i, j)
    permutations(swaped, 0, len(swaped), &dnaPattern)
}
```

2. I create a new array string for store the key 3 - 10 length.

```golang
    pattern := make([]string, 0)
```

3. I create forever loop for get string pattern and store pattern to `pattern` array.

```golang
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
        }
    }
}
```

4. I find the unique pattern.
```golang
unique := distinct(pattern)
```

5. Find the match pattern in `dnaPattern` and display output.
```golang
    for i, p := range unique {
        reg := regexp.MustCompile(p)
        matches := reg.FindAllStringIndex(dnaPattern, -1)
        fmt.Println(i+1, "pattern", fmt.Sprintf(`'%s'`, p), "matches", len(matches))
    }

```