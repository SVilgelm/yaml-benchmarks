//go:build !goexperiment.jsonv2

package yaml_benchmarks

import (
	_ "embed"
	"testing"
)

func TestUnmarshalToInterface(t *testing.T) {
	testUnmarshalToInterface(t)
}

func BenchmarkUnmarshalToInterface(b *testing.B) {
	benchmarkUnmarshalToInterface(b)
}
