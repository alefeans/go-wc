# go-wc

`go-wc` is a Go implementation of the Unix CLI tool `wc` and the solution for the first challenge of [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc/).

### Step 1

```
wc -c test.txt
335040 test.txt

go run main.go -c test.txt
335040 test.txt
```

### Step 2

```sh
wc -l test.txt
7143 test.txt

go run main.go -l test.txt
7143 test.txt
```

### Step 3

```sh
wc -w test.txt
58164 test.txt

go run main.go -w test.txt
58164 test.txt
```

### Step 4

```sh
wc -m test.txt
332144 test.txt

go run main.go -m test.txt
332144 test.txt
```

### Step 5

```sh
wc test.txt
7143 58164 335040 test.txt

go run main.go test.txt
7143 58164 335040 test.txt
```

### Step 6

```sh
cat test.txt | wc
7143 58164 335040

cat test.txt | go run main.go
7143 58164 335040
```

### Tests

```sh
go test -v

=== RUN   TestCountAll
--- PASS: TestCountAll (0.00s)
PASS
ok      github.com/alefeans/go-wc       0.080s
```

### Benchmarks

```
go test -bench=. -benchmem

goos: darwin
goarch: arm64
pkg: github.com/alefeans/go-wc
BenchmarkCountAllSimpleInput-10          3123913               384.7 ns/op          4128 B/op             2 allocs/op
BenchmarkCountAllTestFile-10             1566600               757.4 ns/op          4128 B/op             2 allocs/op
PASS
ok      github.com/alefeans/go-wc       3.698s
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
