// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package replacer

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplacementJSON(t *testing.T) {
	in := `[["World","Planet"],["Hello","Ciao"]]`
	expected := []Replacement{
		Replacement{Old: []byte("World"), New: []byte("Planet")},
		Replacement{Old: []byte("Hello"), New: []byte("Ciao")},
	}
	out := make([]Replacement, 0)
	err := json.Unmarshal([]byte(in), &out)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
	j, err := json.Marshal(out)
	assert.Equal(t, in, string(j))
}
