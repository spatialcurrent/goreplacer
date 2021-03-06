// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package replacer

import (
	"encoding/json"
	"fmt"
)

type Replacement struct {
	Old []byte
	New []byte
}

func (r *Replacement) UnmarshalJSON(data []byte) error {
	values := make([]string, 0)
	err := json.Unmarshal(data, &values)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON into slice: %w", err)
	}
	if len(values) != 2 {
		return fmt.Errorf("error unmarshaling replacement with length %d, expecting 2 values", len(values))
	}
	r.Old = []byte(values[0])
	r.New = []byte(values[1])
	return nil
}

func (r *Replacement) UnmarshalYAML(unmarshal func(object interface{}) error) error {
	values := make([]string, 0)
	err := unmarshal(&values)
	if err != nil {
		return fmt.Errorf("error unmarshaling YAML into slice: %w", err)
	}
	if len(values) != 2 {
		return fmt.Errorf("error unmarshaling replacement with length %d, expecting 2 values", len(values))
	}
	r.Old = []byte(values[0])
	r.New = []byte(values[1])
	return nil
}

func (r Replacement) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string{string(r.Old), string(r.New)})
}

func (r Replacement) MarshalYAML() (interface{}, error) {
	return []string{string(r.Old), string(r.New)}, nil
}
