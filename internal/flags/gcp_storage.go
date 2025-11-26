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

package flags

import (
	"github.com/aerospike/aerospike-backup-cli/internal/models"
	"github.com/spf13/pflag"
)

type GcpStorage struct {
	operation int
	models.GcpStorage
}

func NewGcpStorage(operation int) *GcpStorage {
	return &GcpStorage{
		operation: operation,
	}
}

func (f *GcpStorage) NewFlagSet() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}

	flagSet.StringVar(&f.KeyFile, "gcp-key-path",
		models.DefaultGcpKeyFile,
		"Path to file containing service account JSON key.")

	flagSet.StringVar(&f.BucketName, "gcp-bucket-name",
		models.DefaultGcpBucketName,
		"Name of the Google cloud storage bucket.")

	flagSet.StringVar(&f.Endpoint, "gcp-endpoint-override",
		models.DefaultGcpEndpoint,
		"An alternate url endpoint to send GCP API calls to.")

	switch f.operation {
	case OperationBackup:
		flagSet.IntVar(&f.ChunkSize, "gcp-chunk-size",
			models.DefaultGcpChunkSize,
			"Chunk size controls the maximum number of megabytes of the object that the app will attempt to send to\n"+
				"the storage in a single request. Objects smaller than the size will be sent in a single request,\n"+
				"while larger objects will be split over multiple requests.")

		flagSet.BoolVar(&f.CalculateChecksum, "gcp-calculate-checksum",
			models.DefaultCloudCalculateChecksum,
			"Calculate checksum for each uploaded object.")
	case OperationRestore:
		flagSet.IntVar(&f.RetryReadBackoffSeconds, "gcp-retry-read-backoff",
			models.DefaultCloudRetryReadBackoffSeconds,
			"The initial delay in seconds between retry attempts. In case of connection errors\n"+
				"tool will retry reading the object from the last known position.")

		flagSet.Float64Var(&f.RetryReadMultiplier, "gcp-retry-read-multiplier",
			models.DefaultCloudRetryReadMultiplier,
			"Multiplier is used to increase the delay between subsequent retry attempts.\n"+
				"Used in combination with initial delay.")

		flagSet.UintVar(&f.RetryReadMaxAttempts, "gcp-retry-read-max-attempts",
			models.DefaultCloudRetryReadMaxAttempts,
			"The maximum number of retry attempts that will be made. If set to 0, no retries will be performed.")
	}

	flagSet.IntVar(&f.RetryMaxAttempts, "gcp-retry-max-attempts",
		models.DefaultGcpRetryMaxAttempts,
		"Max retries specifies the maximum number of attempts a failed operation will be retried\n"+
			"before producing an error.")

	flagSet.IntVar(&f.RetryBackoffMaxSeconds, "gcp-retry-max-backoff",
		models.DefaultGcpRetryBackoffMaxSeconds,
		"Max backoff is the maximum value in seconds of the retry period.")

	flagSet.IntVar(&f.RetryBackoffInitSeconds, "gcp-retry-init-backoff",
		models.DefaultGcpRetryBackoffInitSeconds,
		"Initial backoff is the initial value in seconds of the retry period.")

	flagSet.Float64Var(&f.RetryBackoffMultiplier, "gcp-retry-backoff-multiplier",
		models.DefaultGcpRetryBackoffMultiplier,
		"Multiplier is the factor by which the retry period increases.\n"+
			"It should be greater than 1.")

	flagSet.IntVar(&f.MaxConnsPerHost, "gcp-max-conns-per-host",
		models.DefaultCloudMaxConnsPerHost,
		"MaxConnsPerHost optionally limits the total number of connections per host,\n"+
			"including connections in the dialing, active, and idle states. On limit violation, dials will block.\n"+
			"0 means no limit.")

	flagSet.IntVar(&f.RequestTimeoutSeconds, "gcp-request-timeout",
		models.DefaultCloudRequestTimeoutSeconds,
		"Timeout in seconds specifies a time limit for requests made by this Client.\n"+
			"The timeout includes connection time, any redirects, and reading the response body.\n"+
			"0 means no limit.")

	return flagSet
}

func (f *GcpStorage) GetGcpStorage() *models.GcpStorage {
	return &f.GcpStorage
}
