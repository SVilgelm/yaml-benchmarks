# benchmarks of yaml libraries

```shell
% go test -bench=. -benchmem -benchtime 10000x

BenchmarkUnmarshal/json-16                 10000             21107 ns/op            9263 B/op        188 allocs/op
BenchmarkUnmarshal/yaml.V2-16              10000             83208 ns/op           40298 B/op        683 allocs/op
BenchmarkUnmarshal/yaml.V3-16              10000             95817 ns/op           48549 B/op        761 allocs/op
BenchmarkUnmarshal/ghodss-16               10000            126235 ns/op           60254 B/op       1072 allocs/op
BenchmarkUnmarshal/k8s-16                  10000            128281 ns/op           62656 B/op       1080 allocs/op
BenchmarkUnmarshal/k8s:_Number-16          10000            133754 ns/op           62691 B/op       1080 allocs/op
BenchmarkUnmarshal/invopop-16              10000            141098 ns/op           64318 B/op       1122 allocs/op
BenchmarkUnmarshal/goccy-16                10000            205752 ns/op          308235 B/op       2461 allocs/op
```

## libraries

* <https://github.com/ghodss/yaml> v1.0.0
* <https://github.com/goccy/go-yaml> v1.9.5
* <https://gopkg.in/yaml.v2> v2.4.0
* <https://gopkg.in/yaml.v3> v3.0.1
* <https://github.com/kubernetes-sigs/yaml> v1.3.0
* <https://github.com/invopop/yaml> v0.1.0
