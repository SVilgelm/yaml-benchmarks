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
BenchmarkUnmarshalToInterface/json-8               10000             16897 ns/op            9297 B/op        190 allocs/op                                                  
BenchmarkUnmarshalToInterface/yaml.v4-8            10000             64740 ns/op           49349 B/op        772 allocs/op                                                  
BenchmarkUnmarshalToInterface/yaml.v3-8            10000             63798 ns/op           49348 B/op        772 allocs/op                                                  
BenchmarkUnmarshalToInterface/yaml.v2-8            10000             53949 ns/op           40703 B/op        693 allocs/op                                                  
BenchmarkUnmarshalToInterface/ghodss-8             10000             64315 ns/op           55550 B/op       1024 allocs/op                                                  
BenchmarkUnmarshalToInterface/goccy-8              10000             79215 ns/op          106334 B/op       1924 allocs/op                                                  
BenchmarkUnmarshalToInterface/k8s-8                10000             65694 ns/op           57966 B/op       1032 allocs/op                                                  
BenchmarkUnmarshalToInterface/k8s:_Number-8        10000             64948 ns/op           58004 B/op       1032 allocs/op                                                  
BenchmarkUnmarshalToInterface/invopop-8            10000             72977 ns/op           59993 B/op       1073 allocs/op                                                  
```

## benchmark results with jsonv2 enabled

```shell
GOEXPERIMENT=jsonv2 go test -bench=. -benchmem -benchtime 10000x

goos: darwin
goarch: arm64
cpu: Apple M1
BenchmarkUnmarshalToInterface/json.v2-8                    10000              8331 ns/op            8676 B/op        158 allocs/op                                          
BenchmarkUnmarshalToInterface/json-8                       10000             16243 ns/op            9963 B/op        225 allocs/op                                          
BenchmarkUnmarshalToInterface/yaml.v4-8                    10000             64649 ns/op           49367 B/op        774 allocs/op                                          
BenchmarkUnmarshalToInterface/yaml.v3-8                    10000             64925 ns/op           49369 B/op        774 allocs/op                                          
BenchmarkUnmarshalToInterface/yaml.v2-8                    10000             55159 ns/op           40722 B/op        695 allocs/op                                          
BenchmarkUnmarshalToInterface/ghodss-8                     10000             68947 ns/op           53109 B/op        972 allocs/op                                          
BenchmarkUnmarshalToInterface/goccy-8                      10000             81587 ns/op          106337 B/op       1924 allocs/op                                          
BenchmarkUnmarshalToInterface/k8s-8                        10000             72701 ns/op           55527 B/op        983 allocs/op                                          
BenchmarkUnmarshalToInterface/k8s:_Number-8                10000             73656 ns/op           55744 B/op        985 allocs/op                                          
BenchmarkUnmarshalToInterface/invopop-8                    10000             80729 ns/op           57534 B/op       1024 allocs/op                                          
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
