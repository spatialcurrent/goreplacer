// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package replacer

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceBytesOne(t *testing.T) {
	in := "Hello World"
	s, err := New(Replacement{Old: []byte("World"), New: []byte("Planet")})
	assert.NoError(t, err)
	out, err := s.ReplaceBytes([]byte(in))
	assert.NoError(t, err)
	assert.Equal(t, "Hello Planet", string(out))
}

func TestReplaceStringOne(t *testing.T) {
	in := "Hello World"
	s, err := New(Replacement{Old: []byte("World"), New: []byte("Planet")})
	assert.NoError(t, err)
	out, err := s.ReplaceString(in)
	assert.NoError(t, err)
	assert.Equal(t, "Hello Planet", out)
}

func TestReplaceBytesCutoff(t *testing.T) {
	in := `""\""Hello\""""`
	s, err := New(Replacement{Old: []byte("\"\""), New: []byte("\"")})
	assert.NoError(t, err)
	out, err := s.ReplaceBytes([]byte(in))
	assert.NoError(t, err)
	assert.Equal(t, `"\"Hello\""`, string(out))
}

func TestReplaceStringMultiple(t *testing.T) {
	in := "Hello\\nWorld\\tCiao\\\\Ciao"
	s, err := New(
		Replacement{Old: []byte("World"), New: []byte("Planet")},
		Replacement{Old: []byte("\\n"), New: []byte("\n")},
		Replacement{Old: []byte("\\t"), New: []byte("\t")},
		Replacement{Old: []byte("\\\\"), New: []byte("\\")},
	)
	assert.NoError(t, err)
	out, err := s.ReplaceString(in)
	assert.NoError(t, err)
	assert.Equal(t, "Hello\nPlanet\tCiao\\Ciao", out)
}

func TestReplaceStringN(t *testing.T) {
	in := "Hello World"
	s, err := New(
		Replacement{Old: []byte("Hello"), New: []byte("Ciao")},
		Replacement{Old: []byte("World"), New: []byte("Planet")},
	)
	assert.NoError(t, err)
	out, err := s.ReplaceStringN(in, 1)
	assert.NoError(t, err)
	assert.Equal(t, "Ciao World", out)
}

func BenchmarkReplacer(b *testing.B) {
	r, err := New(
		Replacement{Old: []byte("\\n"), New: []byte("\n")},
		Replacement{Old: []byte("\\t"), New: []byte("\t")},
		Replacement{Old: []byte("\\\\"), New: []byte("\\")},
	)
	if err != nil {
		panic(err)
	}
	for n := 0; n < b.N; n++ {
		b := make([]byte, 4096)
		_, err := rand.Read(b)
		if err != nil {
			panic(err)
		}
		r.MustReplaceBytes(b)
	}
}
