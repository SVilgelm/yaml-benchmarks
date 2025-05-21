# benchmarks of yaml libraries

```shell
% go test -bench=. -benchmem -benchtime 10000x

BenchmarkUnmarshal/json-8                  10000             17238 ns/op            9296 B/op        190 allocs/op
BenchmarkUnmarshal/yaml.V3-8               10000             55256 ns/op           49349 B/op        772 allocs/op
BenchmarkUnmarshal/yaml.V2-8               10000             45868 ns/op           40703 B/op        693 allocs/op
BenchmarkUnmarshal/ghodss-8                10000             63666 ns/op           55549 B/op       1024 allocs/op
BenchmarkUnmarshal/goccy-8                 10000             79901 ns/op          106389 B/op       1927 allocs/op
BenchmarkUnmarshal/k8s-8                   10000             64616 ns/op           57970 B/op       1032 allocs/op
BenchmarkUnmarshal/k8s:_Number-8           10000             64631 ns/op           58001 B/op       1032 allocs/op
BenchmarkUnmarshal/invopop-8               10000             72919 ns/op           59991 B/op       1073 allocs/op
```

## libraries

* <https://github.com/ghodss/yaml> v1.0.0
* <https://github.com/goccy/go-yaml> v1.17.1
* <https://github.com/invopop/yaml> v0.3.1
* <https://gopkg.in/yaml.v2> v2.4.0
* <https://gopkg.in/yaml.v3> v3.0.1
* <https://github.com/kubernetes-sigs/yaml> v1.4.0
