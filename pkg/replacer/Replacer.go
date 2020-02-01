// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package replacer

import (
	"bytes"
)

type Replacer struct {
	replacements []Replacement
}

func New(replacements ...Replacement) (*Replacer, error) {
	r := &Replacer{replacements: replacements}
	return r, nil
}

func (r *Replacer) ReplaceBytes(in []byte) ([]byte, error) {
	out := make([]byte, 0, len(in))
	for _, x := range in {
		// append byte to output slice
		out = append(out, x)
		// iterate through possible replacements
		for _, repl := range r.replacements {
			if bytes.HasSuffix(out, repl.Old) {
				if len(repl.Old) == len(repl.New) {
					copy(out[len(out)-len(repl.Old):], repl.New)
				} else {
					out = append(out[0:len(out)-len(repl.Old)], repl.New...)
				}
				break
			}
		}
	}
	return out, nil
}

func (r *Replacer) ReplaceBytesN(in []byte, n int) ([]byte, error) {
	out := make([]byte, 0, len(in))
	count := 0
	for i, x := range in {
		// append byte to output slice
		out = append(out, x)
		// iterate through possible replacements
		for _, repl := range r.replacements {
			if bytes.HasSuffix(out, repl.Old) {
				if len(repl.Old) == len(repl.New) {
					copy(out[len(out)-len(repl.Old):], repl.New)
				} else {
					out = append(out[0:len(out)-len(repl.Old)], repl.New...)
				}
				count++
				break
			}
		}
		if count == n {
			if i+1 < len(in) {
				out = append(out, in[i+1:]...)
			}
			break
		}
	}
	return out, nil
}

func (r *Replacer) MustReplaceBytes(in []byte) []byte {
	out, err := r.ReplaceBytes(in)
	if err != nil {
		panic(err)
	}
	return out
}

func (r *Replacer) MustReplaceBytesN(in []byte, n int) []byte {
	out, err := r.ReplaceBytesN(in, n)
	if err != nil {
		panic(err)
	}
	return out
}

func (r *Replacer) ReplaceString(in string) (string, error) {
	out, err := r.ReplaceBytes([]byte(in))
	if err != nil {
		return "", nil
	}
	return string(out), nil
}

func (r *Replacer) ReplaceStringN(in string, n int) (string, error) {
	out, err := r.ReplaceBytesN([]byte(in), n)
	if err != nil {
		return "", nil
	}
	return string(out), nil
}

func (r *Replacer) MustReplaceString(in string) string {
	out, err := r.ReplaceString(in)
	if err != nil {
		panic(err)
	}
	return out
}

func (r *Replacer) MustReplaceStringN(in string, n int) string {
	out, err := r.ReplaceStringN(in, n)
	if err != nil {
		panic(err)
	}
	return out
}
