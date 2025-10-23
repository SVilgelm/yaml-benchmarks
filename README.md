# benchmarks of yaml libraries

```shell
go test -bench=. -benchmem -benchtime 10000x

goos: darwin                                                                                                                                                                
goarch: arm64                                                                                                                                                               
cpu: Apple M1                                                                                                                                                               
BenchmarkUnmarshal/json-8                  10000             17308 ns/op            9297 B/op        190 allocs/op                                                          
BenchmarkUnmarshal/yaml.v4-8               10000             55333 ns/op           49349 B/op        772 allocs/op                                                          
BenchmarkUnmarshal/yaml.v3-8               10000             55483 ns/op           49349 B/op        772 allocs/op                                                          
BenchmarkUnmarshal/yaml.v2-8               10000             45383 ns/op           40704 B/op        693 allocs/op                                                          
BenchmarkUnmarshal/ghodss-8                10000             63939 ns/op           55550 B/op       1024 allocs/op                                                          
BenchmarkUnmarshal/goccy-8                 10000             79807 ns/op          106338 B/op       1924 allocs/op                                                          
BenchmarkUnmarshal/k8s-8                   10000             64948 ns/op           57962 B/op       1032 allocs/op                                                          
BenchmarkUnmarshal/k8s:_Number-8           10000             64572 ns/op           57999 B/op       1032 allocs/op                                                          
BenchmarkUnmarshal/invopop-8               10000             72498 ns/op           59992 B/op       1073 allocs/op                                                          
```

## libraries

* <https://github.com/ghodss/yaml> v1.0.0
* <https://github.com/goccy/go-yaml> v1.17.1
* <https://github.com/invopop/yaml> v0.3.1
* <https://gopkg.in/yaml.v2> v2.4.0
* <https://gopkg.in/yaml.v3> v3.0.1
* <https://github.com/kubernetes-sigs/yaml> v1.4.0
