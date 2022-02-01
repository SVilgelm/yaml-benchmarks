# benchmarks of yaml libraries

```shell
% go test -bench=. -benchmem -benchtime 10000x

BenchmarkUnmarshal/json-16                 10000             19422 ns/op            9263 B/op        188 allocs/op
BenchmarkUnmarshal/yaml.V2-16              10000             80364 ns/op           40299 B/op        683 allocs/op
BenchmarkUnmarshal/yaml.V3-16              10000            101712 ns/op           48229 B/op        752 allocs/op
BenchmarkUnmarshal/ghodss-16               10000            121582 ns/op           60256 B/op       1091 allocs/op
BenchmarkUnmarshal/k8s-16                  10000            122444 ns/op           62658 B/op       1099 allocs/op
BenchmarkUnmarshal/k8s:_Number-16          10000            121836 ns/op           62694 B/op       1099 allocs/op
BenchmarkUnmarshal/goccy-16                10000            210116 ns/op          311603 B/op       2601 allocs/op
```

## libraries

* <https://github.com/ghodss/yaml> v1.0.0
* <https://github.com/goccy/go-yaml> v1.9.5
* <https://gopkg.in/yaml.v2> v2.4.0
* <https://gopkg.in/yaml.v3> v3.0.0-20210107192922-496545a6307b
