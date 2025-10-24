//go:build goexperiment.jsonv2

package yaml_benchmarks

import (
	_ "embed"
	jsonV2 "encoding/json/v2"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalToInterface(t *testing.T) {
	t.Run("json.v2", func(t *testing.T) {
		var v any
		require.NoError(t, jsonV2.Unmarshal(jsonFile, &v))
		compareWithJSON(t, v)
	})
	testUnmarshalToInterface(t)
}

func BenchmarkUnmarshalToInterface(b *testing.B) {
	b.Run("json.v2", func(b *testing.B) {
		var res any
		for i := 0; i < b.N; i++ {
			var v any
			_ = jsonV2.Unmarshal(jsonFile, &v)
			res = v
		}
		benchmarkRes = res
	})
	benchmarkUnmarshalToInterface(b)
}
