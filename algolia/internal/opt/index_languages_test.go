// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestIndexLanguages(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.IndexLanguagesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.IndexLanguages([]string{}...),
		},
		{
			opts:     []interface{}{opt.IndexLanguages("value1")},
			expected: opt.IndexLanguages("value1"),
		},
		{
			opts:     []interface{}{opt.IndexLanguages("value1", "value2", "value3")},
			expected: opt.IndexLanguages("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractIndexLanguages(c.opts...)
			out opt.IndexLanguagesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}