# benchmarks of yaml libraries

```shell
% go test -bench=. -benchmem -benchtime 10000x

BenchmarkUnmarshal/json-16                 10000             20154 ns/op            9262 B/op        188 allocs/op
BenchmarkUnmarshal/yaml.V2-16              10000             80254 ns/op           40298 B/op        683 allocs/op
BenchmarkUnmarshal/yaml.V3-16              10000            105281 ns/op           48229 B/op        752 allocs/op
BenchmarkUnmarshal/ghodss-16               10000            122681 ns/op           60261 B/op       1091 allocs/op
BenchmarkUnmarshal/goccy-16                10000            213185 ns/op          311600 B/op       2601 allocs/op
```

## libraries

* github.com/ghodss/yaml v1.0.0
* github.com/goccy/go-yaml v1.9.5
* gopkg.in/yaml.v2 v2.4.0
* gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
