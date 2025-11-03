# benchmarks of yaml libraries

This repository contains benchmarks of various YAML libraries in Go, comparing their performance in unmarshaling YAML data into Go structs. 
The benchmarks also include the standard JSON library for reference.

## benchmark results

Go version: go1.25.3

```shell
go test -bench=. -benchmem -benchtime 10000x

goos: darwin
goarch: arm64
cpu: Apple M1
BenchmarkUnmarshalToInterface/json-8                               10000             16853 ns/op            9296 B/op        190 allocs/op
BenchmarkUnmarshalToInterface/json_input/yaml.v4-8                 10000             82736 ns/op           49403 B/op        761 allocs/op
BenchmarkUnmarshalToInterface/json_input/yaml.v3-8                 10000             75787 ns/op           49402 B/op        761 allocs/op
BenchmarkUnmarshalToInterface/json_input/yaml.v2-8                 10000             63324 ns/op           43232 B/op        736 allocs/op
BenchmarkUnmarshalToInterface/json_input/ghodss-8                  10000             75374 ns/op           58110 B/op       1069 allocs/op
BenchmarkUnmarshalToInterface/json_input/goccy-8                   10000            109940 ns/op          163580 B/op       2760 allocs/op                                  
BenchmarkUnmarshalToInterface/json_input/k8s-8                     10000             74313 ns/op           60523 B/op       1077 allocs/op                                  
BenchmarkUnmarshalToInterface/json_input/k8s:_Number-8             10000             75481 ns/op           60547 B/op       1077 allocs/op                                  
BenchmarkUnmarshalToInterface/json_input/invopop-8                 10000             85184 ns/op           59746 B/op       1062 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/yaml.v4-8                 10000             65295 ns/op           49349 B/op        772 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/yaml.v3-8                 10000             64130 ns/op           49349 B/op        772 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/yaml.v2-8                 10000             54479 ns/op           40704 B/op        693 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/ghodss-8                  10000             65123 ns/op           55548 B/op       1024 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/goccy-8                   10000             80036 ns/op          106334 B/op       1924 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/k8s-8                     10000             66237 ns/op           57971 B/op       1032 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/k8s:_Number-8             10000             65447 ns/op           57999 B/op       1032 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/invopop-8                 10000             73241 ns/op           59993 B/op       1073 allocs/op                                  
```

## benchmark results with jsonv2 enabled

```shell
GOEXPERIMENT=jsonv2 go test -bench=. -benchmem -benchtime 10000x

goos: darwin
goarch: arm64
cpu: Apple M1
BenchmarkUnmarshalToInterface/json.v2-8                            10000             17540 ns/op            8676 B/op        158 allocs/op                                          
BenchmarkUnmarshalToInterface/json-8                               10000             17003 ns/op            9963 B/op        225 allocs/op                                          
BenchmarkUnmarshalToInterface/yaml_input/yaml.v4-8                 10000             64493 ns/op           49367 B/op        774 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/yaml.v3-8                 10000             64970 ns/op           49367 B/op        774 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/yaml.v2-8                 10000             55439 ns/op           40722 B/op        695 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/ghodss-8                  10000             69453 ns/op           53082 B/op        971 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/goccy-8                   10000             79986 ns/op          106335 B/op       1924 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/k8s-8                     10000             72127 ns/op           55529 B/op        983 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/k8s:_Number-8             10000             70801 ns/op           55752 B/op        985 allocs/op                                  
BenchmarkUnmarshalToInterface/yaml_input/invopop-8                 10000             79767 ns/op           57522 B/op       1023 allocs/op                                  
BenchmarkUnmarshalToInterface/json_input/yaml.v4-8                 10000             76469 ns/op           49402 B/op        761 allocs/op                                  
BenchmarkUnmarshalToInterface/json_input/yaml.v3-8                 10000             76784 ns/op           49402 B/op        761 allocs/op                                  
BenchmarkUnmarshalToInterface/json_input/yaml.v2-8                 10000             63802 ns/op           43232 B/op        736 allocs/op                                  
BenchmarkUnmarshalToInterface/json_input/ghodss-8                  10000             77760 ns/op           55641 B/op       1017 allocs/op                                  
BenchmarkUnmarshalToInterface/json_input/goccy-8                   10000            109799 ns/op          163567 B/op       2760 allocs/op                                  
BenchmarkUnmarshalToInterface/json_input/k8s-8                     10000             80475 ns/op           58066 B/op       1028 allocs/op                                  
BenchmarkUnmarshalToInterface/json_input/k8s:_Number-8             10000             80224 ns/op           58299 B/op       1030 allocs/op                                  
BenchmarkUnmarshalToInterface/json_input/invopop-8                 10000             92082 ns/op           57277 B/op       1012 allocs/op                                  
```

## libraries

* <https://pkg.go.dev/encoding/json>
* <https://pkg.go.dev/encoding/json/v2>
* <https://github.com/ghodss/yaml> v1.0.0
* <https://github.com/goccy/go-yaml> v1.18.0
* <https://github.com/invopop/yaml> v0.3.1
* <https://github.com/yaml/go-yaml> v4.0.0-rc.2
* <https://github.com/go-yaml/yaml> v2.4.0
* <https://github.com/go-yaml/yaml> v3.0.1
* <https://github.com/kubernetes-sigs/yaml> v1.6.0
