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
BenchmarkUnmarshal/json-8                  10000             17139 ns/op            9297 B/op        190 allocs/op                                                          
BenchmarkUnmarshal/yaml.v4-8               10000             63663 ns/op           49349 B/op        772 allocs/op                                                          
BenchmarkUnmarshal/yaml.v3-8               10000             63183 ns/op           49348 B/op        772 allocs/op                                                          
BenchmarkUnmarshal/yaml.v2-8               10000             54073 ns/op           40703 B/op        693 allocs/op                                                          
BenchmarkUnmarshal/ghodss-8                10000             64281 ns/op           55551 B/op       1024 allocs/op                                                          
BenchmarkUnmarshal/goccy-8                 10000             79373 ns/op          106342 B/op       1924 allocs/op                                                          
BenchmarkUnmarshal/k8s-8                   10000             65548 ns/op           57967 B/op       1032 allocs/op                                                          
BenchmarkUnmarshal/k8s:_Number-8           10000             65088 ns/op           57999 B/op       1032 allocs/op                                                          
BenchmarkUnmarshal/invopop-8               10000             72514 ns/op           59994 B/op       1073 allocs/op                                                          
```

## benchmark results with jsonv2

```shell
GOEXPERIMENT=jsonv2 go test -bench=. -benchmem -benchtime 10000x

goos: darwin                                                                                                                                                                
goarch: arm64                                                                                                                                                               
cpu: Apple M1                                                                                                                                                               
BenchmarkUnmarshalJsonV2/json.v2-8                 10000             10453 ns/op            8675 B/op        158 allocs/op                                                  
BenchmarkUnmarshalJsonV2/json-8                    10000             16685 ns/op            9963 B/op        225 allocs/op                                                  
BenchmarkUnmarshalJsonV2/yaml.v4-8                 10000             64145 ns/op           49367 B/op        774 allocs/op                                                  
BenchmarkUnmarshalJsonV2/yaml.v3-8                 10000             64219 ns/op           49368 B/op        774 allocs/op                                                  
BenchmarkUnmarshalJsonV2/yaml.v2-8                 10000             56330 ns/op           40722 B/op        695 allocs/op                                                  
BenchmarkUnmarshalJsonV2/ghodss-8                  10000             69396 ns/op           53096 B/op        972 allocs/op                                                  
BenchmarkUnmarshalJsonV2/goccy-8                   10000             81734 ns/op          106331 B/op       1924 allocs/op                                                  
BenchmarkUnmarshalJsonV2/k8s-8                     10000             71388 ns/op           55529 B/op        983 allocs/op                                                  
BenchmarkUnmarshalJsonV2/k8s:_Number-8             10000             71687 ns/op           55744 B/op        985 allocs/op                                                  
BenchmarkUnmarshalJsonV2/invopop-8                 10000             82664 ns/op           57540 B/op       1024 allocs/op                                                  
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
