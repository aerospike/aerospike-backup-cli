// Copyright 2024 Aerospike, Inc.
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

package models

import (
	"fmt"
	"strings"

	sa "github.com/aerospike/backup-go/pkg/secret-agent"
)

// SecretAgent contains flags that will be mapped to SecretAgentConfig for backup and restore operations.
type SecretAgent struct {
	ConnectionType     string
	Address            string
	Port               int
	TimeoutMillisecond int

	CaFile   string
	TLSName  string
	CertFile string
	KeyFile  string

	IsBase64 bool
}

// Validate checks if SecretAgent params are valid.
func (s *SecretAgent) Validate() error {
	if s == nil {
		return nil
	}

	if s.ConnectionType == "" {
		return fmt.Errorf("missing connection type")
	}

	if !strings.EqualFold(s.ConnectionType, sa.ConnectionTypeTCP) &&
		!strings.EqualFold(s.ConnectionType, sa.ConnectionTypeUDS) {
		return fmt.Errorf("unsupported connection type: %s", s.ConnectionType)
	}

	return nil
}
