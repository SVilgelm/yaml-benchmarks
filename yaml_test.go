package yaml_benchmarks

import (
	_ "embed"
	"encoding/json"
	"testing"

	ghodss "github.com/ghodss/yaml"
	goccy "github.com/goccy/go-yaml"
	invopop "github.com/invopop/yaml"
	"github.com/stretchr/testify/require"
	yamlV4 "go.yaml.in/yaml/v4"
	yamlV2 "gopkg.in/yaml.v2"
	yamlV3 "gopkg.in/yaml.v3"
	k8s "sigs.k8s.io/yaml"
)

var (
	//go:embed testdata/test.yaml
	yamlFile []byte
	//go:embed testdata/test.json
	jsonFile []byte
)

func yamlJSON(input any) (any, error) {
	switch x := input.(type) {
	case map[any]any:
		m := make(map[string]any, len(x))
		nonstrings := make(map[any]any)

		// Store string keys and values as such.
		for k, v := range x {
			jv, err := yamlJSON(v)
			if err != nil {
				return nil, err
			}

			if s, ok := k.(string); ok {
				m[s] = jv
			} else {
				nonstrings[k] = jv
			}
		}

		// Convert non-string keys to JSON and then to strings, if possible, while giving preference to the string keys.
		for k, jv := range nonstrings {
			jk, err := yamlJSON(k)
			if err != nil {
				return nil, err
			}

			if raw, err := json.Marshal(jk); err != nil {
				return nil, err
			} else {
				sk := string(raw)
				if _, ok := m[sk]; !ok {
					m[sk] = jv
				}
			}
		}

		return m, nil

	case map[string]any:
		m := make(map[string]any, len(x))

		// Check sub items
		for k, v := range x {
			jv, err := yamlJSON(v)
			if err != nil {
				return nil, err
			}

			m[k] = jv
		}

		return m, nil

	case []any:
		for i, v := range x {
			if o, err := yamlJSON(v); err != nil {
				return nil, err
			} else {
				x[i] = o
			}
		}
	}

	return input, nil
}

func compareWithJSON(t testing.TB, v any) {
	data, err := json.Marshal(v)
	t.Logf("json: %s", string(data))
	require.NoError(t, err)
	require.JSONEq(t, string(jsonFile), string(data))
}

func k8sOpt(d *json.Decoder) *json.Decoder {
	d.UseNumber()
	return d
}

func TestUnmarshalToInterface(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		var v any
		require.NoError(t, json.Unmarshal(jsonFile, &v))
		compareWithJSON(t, v)
	})
	t.Run("yaml.v4", func(t *testing.T) {
		var v any
		require.NoError(t, yamlV4.Unmarshal(yamlFile, &v))
		j, err := yamlJSON(v)
		require.NoError(t, err)
		compareWithJSON(t, j)
	})
	t.Run("yaml.v3", func(t *testing.T) {
		var v any
		require.NoError(t, yamlV3.Unmarshal(yamlFile, &v))
		j, err := yamlJSON(v)
		require.NoError(t, err)
		compareWithJSON(t, j)
	})
	t.Run("yaml.v2", func(t *testing.T) {
		var v any
		require.NoError(t, yamlV2.Unmarshal(yamlFile, &v))
		j, err := yamlJSON(v)
		require.NoError(t, err)
		compareWithJSON(t, j)
	})
	t.Run("ghodss", func(t *testing.T) {
		var v any
		require.NoError(t, ghodss.Unmarshal(yamlFile, &v))
		compareWithJSON(t, v)
	})
	t.Run("goccy", func(t *testing.T) {
		var v any
		require.NoError(t, goccy.Unmarshal(yamlFile, &v))
		compareWithJSON(t, v)
	})
	t.Run("k8s", func(t *testing.T) {
		var v any
		require.NoError(t, k8s.Unmarshal(yamlFile, &v))
		compareWithJSON(t, v)
	})
	t.Run("k8s: Number", func(t *testing.T) {
		var v any
		require.NoError(t, k8s.Unmarshal(yamlFile, &v, k8sOpt))
		compareWithJSON(t, v)
	})
	t.Run("invopop", func(t *testing.T) {
		var v any
		require.NoError(t, invopop.Unmarshal(yamlFile, &v))
		compareWithJSON(t, v)
	})
}

var benchmarkRes any

func BenchmarkUnmarshal(b *testing.B) {
	b.Run("json", func(b *testing.B) {
		var res any
		for i := 0; i < b.N; i++ {
			var v any
			_ = json.Unmarshal(jsonFile, &v)
			res = v
		}
		benchmarkRes = res
	})
	b.Run("yaml.v4", func(b *testing.B) {
		var res any
		for i := 0; i < b.N; i++ {
			var v any
			_ = yamlV4.Unmarshal(yamlFile, &v)
			j, _ := yamlJSON(v)
			res = j
		}
		benchmarkRes = res
	})
	b.Run("yaml.v3", func(b *testing.B) {
		var res any
		for i := 0; i < b.N; i++ {
			var v any
			_ = yamlV3.Unmarshal(yamlFile, &v)
			j, _ := yamlJSON(v)
			res = j
		}
		benchmarkRes = res
	})
	b.Run("yaml.v2", func(b *testing.B) {
		var res any
		for i := 0; i < b.N; i++ {
			var v any
			_ = yamlV2.Unmarshal(yamlFile, &v)
			j, _ := yamlJSON(v)
			benchmarkRes = j
		}
		res = res
	})
	b.Run("ghodss", func(b *testing.B) {
		var res any
		for i := 0; i < b.N; i++ {
			var v any
			_ = ghodss.Unmarshal(yamlFile, &v)
			res = v
		}
		benchmarkRes = res
	})
	b.Run("goccy", func(b *testing.B) {
		var res any
		for i := 0; i < b.N; i++ {
			var v any
			_ = goccy.Unmarshal(yamlFile, &v)
			res = v
		}
		benchmarkRes = res
	})
	b.Run("k8s", func(b *testing.B) {
		var res any
		for i := 0; i < b.N; i++ {
			var v any
			_ = k8s.Unmarshal(yamlFile, &v)
			res = v
		}
		benchmarkRes = res
	})
	b.Run("k8s: Number", func(b *testing.B) {
		var res any
		for i := 0; i < b.N; i++ {
			var v any
			_ = k8s.Unmarshal(yamlFile, &v, k8sOpt)
			res = v
		}
		benchmarkRes = res
	})
	b.Run("invopop", func(b *testing.B) {
		var res any
		for i := 0; i < b.N; i++ {
			var v any
			_ = invopop.Unmarshal(yamlFile, &v)
			res = v
		}
		benchmarkRes = res
	})
	b.Logf("Res type: %T", benchmarkRes)
}
