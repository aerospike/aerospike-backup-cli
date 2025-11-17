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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAzureBlob_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		azure    *AzureBlob
		isBackup bool
		wantErr  string
	}{
		{
			name: "valid backup configuration",
			azure: &AzureBlob{
				ContainerName:        testBucketName,
				RetryMaxAttempts:     3,
				RetryTimeoutSeconds:  30,
				RetryDelaySeconds:    5,
				RetryMaxDelaySeconds: 60,
				BlockSize:            1024,
				UploadConcurrency:    1,
			},
			isBackup: true,
			wantErr:  "",
		},
		{
			name: "valid restore configuration",
			azure: &AzureBlob{
				ContainerName:        testBucketName,
				RetryMaxAttempts:     3,
				RetryTimeoutSeconds:  30,
				RetryDelaySeconds:    5,
				RetryMaxDelaySeconds: 60,
				BlockSize:            1024,
				RestorePollDuration:  10,
			},
			isBackup: false,
			wantErr:  "",
		},
		{
			name: "empty container name",
			azure: &AzureBlob{
				ContainerName: "",
			},
			isBackup: true,
			wantErr:  "container name is required",
		},
		{
			name: "negative retry max attempts",
			azure: &AzureBlob{
				ContainerName:    testBucketName,
				RetryMaxAttempts: -1,
			},
			isBackup: true,
			wantErr:  "retry maximum attempts must be non-negative",
		},
		{
			name: "negative retry timeout seconds",
			azure: &AzureBlob{
				ContainerName:       testBucketName,
				RetryMaxAttempts:    3,
				RetryTimeoutSeconds: -1,
			},
			isBackup: true,
			wantErr:  "retry try timeout must be non-negative",
		},
		{
			name: "negative retry delay seconds",
			azure: &AzureBlob{
				ContainerName:       testBucketName,
				RetryMaxAttempts:    3,
				RetryTimeoutSeconds: 30,
				RetryDelaySeconds:   -1,
			},
			isBackup: true,
			wantErr:  "retry delay must be non-negative",
		},
		{
			name: "negative retry max delay seconds",
			azure: &AzureBlob{
				ContainerName:        testBucketName,
				RetryMaxAttempts:     3,
				RetryTimeoutSeconds:  30,
				RetryDelaySeconds:    5,
				RetryMaxDelaySeconds: -1,
			},
			isBackup: true,
			wantErr:  "retry max delay must be non-negative",
		},
		{
			name: "negative block size",
			azure: &AzureBlob{
				ContainerName:        testBucketName,
				RetryMaxAttempts:     3,
				RetryTimeoutSeconds:  30,
				RetryDelaySeconds:    5,
				RetryMaxDelaySeconds: 60,
				BlockSize:            -1,
			},
			isBackup: true,
			wantErr:  "block size must be non-negative",
		},
		{
			name: "restore poll duration less than 1 for restore",
			azure: &AzureBlob{
				ContainerName:        testBucketName,
				RetryMaxAttempts:     3,
				RetryTimeoutSeconds:  30,
				RetryDelaySeconds:    5,
				RetryMaxDelaySeconds: 60,
				BlockSize:            1024,
				RestorePollDuration:  0,
			},
			isBackup: false,
			wantErr:  "restore poll duration can't be less than 1",
		},
		{
			name: "restore poll duration less than 1 for backup is ok",
			azure: &AzureBlob{
				ContainerName:        testBucketName,
				RetryMaxAttempts:     3,
				RetryTimeoutSeconds:  30,
				RetryDelaySeconds:    5,
				RetryMaxDelaySeconds: 60,
				BlockSize:            1024,
				RestorePollDuration:  0,
				UploadConcurrency:    1,
			},
			isBackup: true,
			wantErr:  "",
		},
		{
			name: "zero values are valid",
			azure: &AzureBlob{
				ContainerName:        testBucketName,
				RetryMaxAttempts:     0,
				RetryTimeoutSeconds:  0,
				RetryDelaySeconds:    0,
				RetryMaxDelaySeconds: 0,
				BlockSize:            0,
				RestorePollDuration:  1,
			},
			isBackup: false,
			wantErr:  "",
		},
		{
			name: "restore poll duration exactly 1 is valid",
			azure: &AzureBlob{
				ContainerName:        testBucketName,
				RetryMaxAttempts:     0,
				RetryTimeoutSeconds:  0,
				RetryDelaySeconds:    0,
				RetryMaxDelaySeconds: 0,
				BlockSize:            0,
				RestorePollDuration:  1,
			},
			isBackup: false,
			wantErr:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.azure.LoadSecrets(nil)
			require.NoError(t, err)

			err = tt.azure.Validate(tt.isBackup)

			if tt.wantErr != "" {
				switch {
				case err == nil:
					t.Errorf("Validate() expected error %q, got nil", tt.wantErr)
				case err.Error() != tt.wantErr:
					t.Errorf("Validate() error = %q, want %q", err.Error(), tt.wantErr)
				}

				return
			}

			if err != nil {
				t.Errorf("Validate() unexpected error: %v", err)
			}
		})
	}
}
