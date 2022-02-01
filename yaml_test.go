package yaml_benchmarks

import (
	_ "embed"
	"encoding/json"
	"testing"

	ghodss "github.com/ghodss/yaml"
	goccy "github.com/goccy/go-yaml"
	"github.com/stretchr/testify/require"
	yamlV2 "gopkg.in/yaml.v2"
	yamlV3 "gopkg.in/yaml.v3"
)

var (
	//go:embed testdata/test.yaml
	yamlFile []byte
	//go:embed testdata/test.json
	jsonFile []byte
)

func yamlJSON(input interface{}) (interface{}, error) {
	switch x := input.(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{}, len(x))
		nonstrings := make(map[interface{}]interface{})

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

	case map[string]interface{}:
		m := make(map[string]interface{}, len(x))

		// Check sub items
		for k, v := range x {
			jv, err := yamlJSON(v)
			if err != nil {
				return nil, err
			}

			m[k] = jv
		}

		return m, nil

	case []interface{}:
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

func compareWithJSON(t testing.TB, v interface{}) {
	data, err := json.Marshal(v)
	t.Logf("json: %s", string(data))
	require.NoError(t, err)
	require.JSONEq(t, string(jsonFile), string(data))
}

func TestUnmarshalToInterface(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		var v interface{}
		require.NoError(t, json.Unmarshal(jsonFile, &v))
		compareWithJSON(t, v)
	})
	t.Run("yaml.v3", func(t *testing.T) {
		var v interface{}
		require.NoError(t, yamlV3.Unmarshal(yamlFile, &v))
		j, err := yamlJSON(v)
		require.NoError(t, err)
		compareWithJSON(t, j)
	})
	t.Run("yaml.v2", func(t *testing.T) {
		var v interface{}
		require.NoError(t, yamlV2.Unmarshal(yamlFile, &v))
		j, err := yamlJSON(v)
		require.NoError(t, err)
		compareWithJSON(t, j)
	})
	t.Run("ghodss", func(t *testing.T) {
		var v interface{}
		require.NoError(t, ghodss.Unmarshal(yamlFile, &v))
		compareWithJSON(t, v)
	})
	t.Run("goccy", func(t *testing.T) {
		var v interface{}
		require.NoError(t, goccy.Unmarshal(yamlFile, &v))
		compareWithJSON(t, v)
	})
}

var benchmarkRes interface{}

func BenchmarkUnmarshal(b *testing.B) {
	b.Run("json", func(b *testing.B) {
		var res interface{}
		for i := 0; i < b.N; i++ {
			var v interface{}
			_ = json.Unmarshal(jsonFile, &v)
			res = v
		}
		benchmarkRes = res
	})
	b.Run("yaml.V3", func(b *testing.B) {
		var res interface{}
		for i := 0; i < b.N; i++ {
			var v interface{}
			_ = yamlV3.Unmarshal(yamlFile, &v)
			j, _ := yamlJSON(v)
			res = j
		}
		benchmarkRes = res
	})
	b.Run("yaml.V2", func(b *testing.B) {
		var res interface{}
		for i := 0; i < b.N; i++ {
			var v interface{}
			_ = yamlV2.Unmarshal(yamlFile, &v)
			j, _ := yamlJSON(v)
			benchmarkRes = j
		}
		benchmarkRes = res
	})
	b.Run("ghodss", func(b *testing.B) {
		var res interface{}
		for i := 0; i < b.N; i++ {
			var v interface{}
			_ = ghodss.Unmarshal(yamlFile, &v)
			benchmarkRes = v
		}
		benchmarkRes = res
	})
	b.Run("goccy", func(b *testing.B) {
		var res interface{}
		for i := 0; i < b.N; i++ {
			var v interface{}
			_ = goccy.Unmarshal(yamlFile, &v)
			benchmarkRes = v
		}
		benchmarkRes = res
	})
	b.Logf("Res type: %T", benchmarkRes)
}
