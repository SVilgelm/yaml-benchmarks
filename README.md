# benchmarks of yaml libraries

```shell
% go test -bench=. -benchmem -benchtime 10000x

BenchmarkUnmarshal/json-16                 10000              6348 ns/op            2168 B/op         60 allocs/op
BenchmarkUnmarshal/yaml.V3-16              10000             41960 ns/op           19535 B/op        291 allocs/op
BenchmarkUnmarshal/yaml.V2-16              10000             30319 ns/op           15666 B/op        260 allocs/op
BenchmarkUnmarshal/ghodss-16               10000             42560 ns/op           20956 B/op        379 allocs/op
BenchmarkUnmarshal/goccy-16                10000             74175 ns/op           79236 B/op       1090 allocs/op
```

## libraries

* github.com/ghodss/yaml v1.0.0
* github.com/goccy/go-yaml v1.9.5
* gopkg.in/yaml.v2 v2.4.0
* gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
