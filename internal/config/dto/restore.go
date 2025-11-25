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

package dto

import (
	"strings"

	"github.com/aerospike/aerospike-backup-cli/internal/models"
)

// Restore is used to map yaml config.
type Restore struct {
	App         App           `yaml:"app"`
	Cluster     Cluster       `yaml:"cluster"`
	Restore     RestoreConfig `yaml:"restore"`
	Compression Compression   `yaml:"compression"`
	Encryption  Encryption    `yaml:"encryption"`
	SecretAgent SecretAgent   `yaml:"secret-agent"`
	Aws         struct {
		S3 AwsS3 `yaml:"s3"`
	} `yaml:"aws"`
	Gcp struct {
		Storage GcpStorage `yaml:"storage"`
	} `yaml:"gcp"`
	Azure struct {
		Blob AzureBlob `yaml:"blob"`
	} `yaml:"azure"`
}

func (r *Restore) ToModelRestore() *models.Restore {
	return &models.Restore{
		Common: models.Common{
			Directory:                     derefString(r.Restore.Directory),
			Namespace:                     derefString(r.Restore.Namespace),
			SetList:                       strings.Join(r.Restore.SetList, ","),
			BinList:                       strings.Join(r.Restore.BinList, ","),
			Parallel:                      derefInt(r.Restore.Parallel),
			NoRecords:                     derefBool(r.Restore.NoRecords),
			NoIndexes:                     derefBool(r.Restore.NoIndexes),
			NoUDFs:                        derefBool(r.Restore.NoUDFs),
			RecordsPerSecond:              derefInt(r.Restore.RecordsPerSecond),
			MaxRetries:                    derefInt(r.Restore.MaxRetries),
			TotalTimeout:                  derefInt64(r.Restore.TotalTimeout),
			SocketTimeout:                 derefInt64(r.Restore.SocketTimeout),
			Bandwidth:                     derefInt64(r.Restore.Bandwidth),
			InfoTimeout:                   derefInt64(r.Restore.InfoTimeout),
			InfoMaxRetries:                derefUint(r.Restore.InfoMaxRetries),
			InfoRetriesMultiplier:         derefFloat64(r.Restore.InfoRetriesMultiplier),
			InfoRetryIntervalMilliseconds: derefInt64(r.Restore.InfoRetryIntervalMilliseconds),
			StdBufferSize:                 derefInt(r.Restore.StdBufferSize),
		},
		InputFile:          derefString(r.Restore.InputFile),
		DirectoryList:      strings.Join(r.Restore.DirectoryList, ","),
		ParentDirectory:    derefString(r.Restore.ParentDirectory),
		DisableBatchWrites: r.Restore.DisableBatchWrites,
		BatchSize:          r.Restore.BatchSize,
		MaxAsyncBatches:    r.Restore.MaxAsyncBatches,
		WarmUp:             r.Restore.WarmUp,
		ExtraTTL:           r.Restore.ExtraTTL,
		IgnoreRecordError:  r.Restore.IgnoreRecordError,
		Uniq:               r.Restore.Uniq,
		Replace:            r.Restore.Replace,
		NoGeneration:       r.Restore.NoGeneration,
		RetryBaseInterval:  r.Restore.RetryBaseInterval,
		RetryMultiplier:    r.Restore.RetryMultiplier,
		RetryMaxAttempts:   r.Restore.RetryMaxAttempts,
		ValidateOnly:       r.Restore.ValidateOnly,
		ApplyMetadataLast:  r.Restore.ApplyMetadataLast,
	}
}

type RestoreConfig struct {
	Directory                     *string  `yaml:"directory"`
	Namespace                     *string  `yaml:"namespace"`
	SetList                       []string `yaml:"set-list"`
	BinList                       []string `yaml:"bin-list"`
	Parallel                      *int     `yaml:"parallel"`
	NoRecords                     *bool    `yaml:"no-records"`
	NoIndexes                     *bool    `yaml:"no-indexes"`
	NoUDFs                        *bool    `yaml:"no-udfs"`
	RecordsPerSecond              *int     `yaml:"records-per-second"`
	MaxRetries                    *int     `yaml:"max-retries"`
	TotalTimeout                  *int64   `yaml:"total-timeout"`
	SocketTimeout                 *int64   `yaml:"socket-timeout"`
	Bandwidth                     *int64   `yaml:"bandwidth"`
	InputFile                     *string  `yaml:"input-file"`
	DirectoryList                 []string `yaml:"directory-list"`
	ParentDirectory               *string  `yaml:"parent-directory"`
	DisableBatchWrites            *bool    `yaml:"disable-batch-writes"`
	BatchSize                     *int     `yaml:"batch-size"`
	MaxAsyncBatches               *int     `yaml:"max-async-batches"`
	WarmUp                        *int     `yaml:"warm-up"`
	ExtraTTL                      *int64   `yaml:"extra-ttl"`
	IgnoreRecordError             *bool    `yaml:"ignore-record-error"`
	Uniq                          *bool    `yaml:"unique"`
	Replace                       *bool    `yaml:"replace"`
	NoGeneration                  *bool    `yaml:"no-generation"`
	RetryBaseInterval             *int64   `yaml:"retry-base-interval"`
	RetryMultiplier               *float64 `yaml:"retry-multiplier"`
	RetryMaxAttempts              *uint    `yaml:"retry-max-attempts"`
	ValidateOnly                  *bool    `yaml:"validate-only"`
	InfoTimeout                   *int64   `yaml:"info-timeout"`
	InfoMaxRetries                *uint    `yaml:"info-max-retries"`
	InfoRetriesMultiplier         *float64 `yaml:"info-retry-multiplier"`
	InfoRetryIntervalMilliseconds *int64   `yaml:"info-retry-interval"`
	ApplyMetadataLast             *bool    `yaml:"apply-metadata-last"`
	StdBufferSize                 *int     `yaml:"std-buffer-size"`
}

func defaultRestoreConfig() *RestoreConfig {
	return &RestoreConfig{
		Directory:                     stringPtr(models.DefaultCommonDirectory),
		Namespace:                     stringPtr(models.DefaultCommonNamespace),
		SetList:                       []string{},
		BinList:                       []string{},
		NoRecords:                     boolPtr(models.DefaultCommonNoRecords),
		NoIndexes:                     boolPtr(models.DefaultCommonNoIndexes),
		NoUDFs:                        boolPtr(models.DefaultCommonNoUDFs),
		RecordsPerSecond:              intPtr(models.DefaultCommonRecordsPerSecond),
		MaxRetries:                    intPtr(models.DefaultCommonMaxRetries),
		SocketTimeout:                 int64Ptr(models.DefaultCommonSocketTimeout),
		InfoTimeout:                   int64Ptr(models.DefaultCommonInfoTimeout),
		InfoMaxRetries:                uintPtr(models.DefaultCommonInfoMaxRetries),
		InfoRetriesMultiplier:         float64Ptr(models.DefaultCommonInfoRetriesMultiplier),
		InfoRetryIntervalMilliseconds: int64Ptr(models.DefaultCommonInfoRetryIntervalMilliseconds),
		Bandwidth:                     int64Ptr(models.DefaultCommonBandwidth),
		StdBufferSize:                 intPtr(models.DefaultCommonStdBufferSize),
		TotalTimeout:                  int64Ptr(models.DefaultRestoreTotalTimeout),
		Parallel:                      intPtr(models.DefaultRestoreParallel),
		InputFile:                     stringPtr(models.DefaultRestoreInputFile),
		DirectoryList:                 []string,
		ParentDirectory:               stringPtr(models.DefaultRestoreParentDirectory),
		DisableBatchWrites:            boolPtr(models.DefaultRestoreDisableBatchWrites),
		BatchSize:                     intPtr(models.DefaultRestoreBatchSize),
		MaxAsyncBatches:               intPtr(models.DefaultRestoreMaxAsyncBatches),
		WarmUp:                        intPtr(models.DefaultRestoreWarmUp),
		ExtraTTL:                      int64Ptr(models.DefaultRestoreExtraTTL),
		IgnoreRecordError:             boolPtr(models.DefaultRestoreIgnoreRecordError),
		Uniq:                          boolPtr(models.DefaultRestoreUniq),
		Replace:                       boolPtr(models.DefaultRestoreReplace),
		NoGeneration:                  boolPtr(models.DefaultRestoreNoGeneration),
		RetryBaseInterval:             int64Ptr(models.DefaultRestoreRetryBaseInterval),
		RetryMultiplier:               float64Ptr(models.DefaultRestoreRetryMultiplier),
		RetryMaxAttempts:              uintPtr(models.DefaultRestoreRetryMaxAttempts),
		ValidateOnly:                  boolPtr(models.DefaultRestoreValidateOnly),
		ApplyMetadataLast:             boolPtr(models.DefaultRestoreApplyMetadataLast),
	}
}
