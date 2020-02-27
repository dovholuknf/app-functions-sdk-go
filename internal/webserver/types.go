//
// Copyright (c) 2020 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package webserver

import (
	"errors"
	"fmt"
	"net/url"
)

// SecretData the structure to store post secret requests
type SecretData struct {
	Path    string     `json:"path"`
	Secrets []KeyValue `json:"secrets"`
}

// KeyValue stores the key value pair of the secret
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (secretData SecretData) validateSecretData() error {

	if len(secretData.Secrets) == 0 {
		return errors.New("Missing required field 'Secrets'")
	}

	for _, kv := range secretData.Secrets {
		if kv.Key == "" {
			return errors.New("'Secrets' key should not be empty")
		}
	}
	if _, err := url.Parse(secretData.Path); err != nil {
		return fmt.Errorf("'Path' is invalid %v", err)
	}

	return nil
}
