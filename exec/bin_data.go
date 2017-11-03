// bin_data.go
// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT

// Copyright 2017 Elasticsearch Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exec

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
)

var assets map[string][]byte

func asset(key string) ([]byte, error) {
	if assets == nil {
		assets = map[string][]byte{}

		var value []byte
		value, _ = base64.StdEncoding.DecodeString("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")
		value, _ = gzipDecode(value)
		assets["exec.o"] = value
	}

	if value, found := assets[key]; found {
		return value, nil
	}
	return nil, fmt.Errorf("asset not found for key=%v", key)
}

func gzipDecode(data []byte) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	out := new(bytes.Buffer)
	if _, err = io.Copy(out, gz); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
