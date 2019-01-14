
##### RUN and Save log
```
./tandem-repeat.sh > out.log 
```



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
            // next curcer
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