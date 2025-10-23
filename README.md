# benchmarks of yaml libraries

```shell
GOEXPERIMENT=jsonv2 go test -bench=. -benchmem -benchtime 10000x

goos: darwin                                                                                                                                                                
goarch: arm64                                                                                                                                                               
cpu: Apple M1                                                                                                                                                               
BenchmarkUnmarshal/json.v2-8               10000             10380 ns/op            8676 B/op        158 allocs/op                                                          
BenchmarkUnmarshal/json-8                  10000             15890 ns/op            9963 B/op        225 allocs/op                                                          
BenchmarkUnmarshal/yaml.v4-8               10000             55526 ns/op           49368 B/op        774 allocs/op                                                          
BenchmarkUnmarshal/yaml.v3-8               10000             55542 ns/op           49368 B/op        774 allocs/op                                                          
BenchmarkUnmarshal/yaml.v2-8               10000             45767 ns/op           40721 B/op        695 allocs/op                                                          
BenchmarkUnmarshal/ghodss-8                10000             68130 ns/op           53081 B/op        971 allocs/op                                                          
BenchmarkUnmarshal/goccy-8                 10000             79220 ns/op          106340 B/op       1924 allocs/op                                                          
BenchmarkUnmarshal/k8s-8                   10000             71295 ns/op           55534 B/op        983 allocs/op                                                          
BenchmarkUnmarshal/k8s:_Number-8           10000             70107 ns/op           55743 B/op        985 allocs/op                                                          
BenchmarkUnmarshal/invopop-8               10000             78911 ns/op           57520 B/op       1023 allocs/op                                                          
```

## libraries

* <https://github.com/ghodss/yaml> v1.0.0
* <https://github.com/goccy/go-yaml> v1.17.1
* <https://github.com/invopop/yaml> v0.3.1
* <https://gopkg.in/yaml.v2> v2.4.0
* <https://gopkg.in/yaml.v3> v3.0.1
* <https://github.com/kubernetes-sigs/yaml> v1.4.0
