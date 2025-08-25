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

package backup

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path"
	"testing"
	"time"

	"github.com/aerospike/aerospike-backup-cli/internal/config"
	"github.com/aerospike/aerospike-backup-cli/internal/models"
	"github.com/aerospike/aerospike-backup-cli/internal/storage"
	"github.com/aerospike/aerospike-client-go/v8"
	"github.com/aerospike/backup-go"
	"github.com/aerospike/tools-common-go/client"
	"github.com/stretchr/testify/require"
)

const (
	testNamespace       = "test"
	testSet             = "test"
	testSetXDR          = "test-xdr"
	testStateFile       = "state"
	testASLoginPassword = "admin"
	testDC              = "dc1"
	testXDRHost         = "172.17.0.1"
	testXDRPort         = 8066
	testAckQueueSize    = 256
	testResultQueueSize = 256
	testRewind          = "all"
)

func Test_BackupWithState(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	dir := path.Join(t.TempDir(), "plain")
	hostPort := client.NewDefaultHostTLSPort()

	asbParams := &config.BackupServiceConfig{
		App: &models.App{},
		ClientConfig: &client.AerospikeConfig{
			Seeds: client.HostTLSPortSlice{
				hostPort,
			},
			User:     testASLoginPassword,
			Password: testASLoginPassword,
		},
		ClientPolicy: &models.ClientPolicy{
			Timeout:      1000,
			IdleTimeout:  1000,
			LoginTimeout: 1000,
		},
		Backup: &models.Backup{
			StateFileDst: testStateFile,
			ScanPageSize: 10,
			FileLimit:    100000,
			Common: models.Common{
				Directory:                     dir,
				Namespace:                     testNamespace,
				Parallel:                      1,
				InfoMaxRetries:                3,
				InfoRetriesMultiplier:         1,
				InfoRetryIntervalMilliseconds: 1000,
			},
		},
		Compression: &models.Compression{
			Mode: backup.CompressNone,
		},
		Encryption:  &models.Encryption{},
		SecretAgent: &models.SecretAgent{},
		AwsS3:       &models.AwsS3{},
		GcpStorage:  &models.GcpStorage{},
		AzureBlob:   &models.AzureBlob{},
	}

	err := createRecords(asbParams.ClientConfig, asbParams.ClientPolicy, testNamespace, testSet)
	require.NoError(t, err)

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	asb, err := NewService(ctx, asbParams, logger)
	require.NoError(t, err)

	err = asb.Run(ctx)
	require.NoError(t, err)
}

func Test_BackupXDR(t *testing.T) {
	// Do not parallel this test. We have multiply xdr tests, so they should be executed sequentially.
	ctx := context.Background()
	dir := path.Join(t.TempDir(), "xdr")
	hostPort := client.NewDefaultHostTLSPort()

	asbParams := &config.BackupServiceConfig{
		App: &models.App{},
		ClientConfig: &client.AerospikeConfig{
			Seeds: client.HostTLSPortSlice{
				hostPort,
			},
			User:     testASLoginPassword,
			Password: testASLoginPassword,
		},
		ClientPolicy: &models.ClientPolicy{
			Timeout:      1000,
			IdleTimeout:  1000,
			LoginTimeout: 1000,
		},
		BackupXDR: &models.BackupXDR{
			FileLimit:                     100000,
			InfoMaxRetries:                3,
			InfoRetriesMultiplier:         1,
			InfoRetryIntervalMilliseconds: 1000,
			Directory:                     dir,
			Namespace:                     testNamespace,
			DC:                            testDC,
			LocalAddress:                  testXDRHost,
			LocalPort:                     testXDRPort,
			MaxConnections:                10,
			Rewind:                        testRewind,
			InfoPolingPeriodMilliseconds:  100,
			ReadTimeoutMilliseconds:       10000,
			WriteTimeoutMilliseconds:      10000,
			ResultQueueSize:               testAckQueueSize,
			AckQueueSize:                  testResultQueueSize,
			StartTimeoutMilliseconds:      10000,
		},
		Compression: &models.Compression{
			Mode: backup.CompressNone,
		},
		Encryption:  &models.Encryption{},
		SecretAgent: &models.SecretAgent{},
		AwsS3:       &models.AwsS3{},
		GcpStorage:  &models.GcpStorage{},
		AzureBlob:   &models.AzureBlob{},
	}

	err := createRecords(asbParams.ClientConfig, asbParams.ClientPolicy, testNamespace, testSetXDR)
	require.NoError(t, err)

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	asb, err := NewService(ctx, asbParams, logger)
	require.NoError(t, err)

	err = asb.Run(ctx)
	require.NoError(t, err)
}

func Test_BackupEstimates(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	hostPort := client.NewDefaultHostTLSPort()

	asbParams := &config.BackupServiceConfig{
		App: &models.App{},
		ClientConfig: &client.AerospikeConfig{
			Seeds: client.HostTLSPortSlice{
				hostPort,
			},
			User:     testASLoginPassword,
			Password: testASLoginPassword,
		},
		ClientPolicy: &models.ClientPolicy{
			Timeout:      1000,
			IdleTimeout:  1000,
			LoginTimeout: 1000,
		},
		Backup: &models.Backup{
			FileLimit: 100000,
			Common: models.Common{
				Namespace:                     testNamespace,
				Parallel:                      1,
				InfoMaxRetries:                3,
				InfoRetriesMultiplier:         1,
				InfoRetryIntervalMilliseconds: 1000,
			},
			Estimate:        true,
			EstimateSamples: 100,
		},
		Compression: &models.Compression{
			Mode: backup.CompressNone,
		},
		Encryption:  &models.Encryption{},
		SecretAgent: &models.SecretAgent{},
		AwsS3:       &models.AwsS3{},
		GcpStorage:  &models.GcpStorage{},
		AzureBlob:   &models.AzureBlob{},
	}

	err := createRecords(asbParams.ClientConfig, asbParams.ClientPolicy, testNamespace, testSet)
	require.NoError(t, err)

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	asb, err := NewService(ctx, asbParams, logger)
	require.NoError(t, err)

	err = asb.Run(ctx)
	require.NoError(t, err)
}

func createRecords(cfg *client.AerospikeConfig, cp *models.ClientPolicy, namespace, set string) error {
	client, err := storage.NewAerospikeClient(cfg, cp, "", 0, slog.Default())
	if err != nil {
		return fmt.Errorf("failed to create aerospike client: %w", err)
	}

	wp := aerospike.NewWritePolicy(0, 0)

	for i := 0; i < 10; i++ {
		key, err := aerospike.NewKey(namespace, set, fmt.Sprintf("map-key-%d", i))
		if err != nil {
			return fmt.Errorf("failed to create aerospike key: %w", err)
		}

		bin := aerospike.NewBin("time", time.Now().Unix())

		if err = client.PutBins(wp, key, bin); err != nil {
			return fmt.Errorf("failed to create aerospike key: %w", err)
		}
	}

	return nil
}
