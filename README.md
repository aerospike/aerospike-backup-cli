# Aerospike Backup CLI
[![Tests](https://github.com/aerospike/aerospike-backup-cli/actions/workflows/tests.yml/badge.svg)](https://github.com/aerospike/aerospike-backup-cli/actions/workflows/tests.yml/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/aerospike/aerospike-backup-cli)](https://pkg.go.dev/github.com/aerospike/aerospike-backup-cli)
[![codecov](https://codecov.io/gh/aerospike/aerospike-backup-cli/graph/badge.svg?token=29G65BU7QX)](https://codecov.io/gh/aerospike/aerospike-backup-cli)

The repository includes the [abs-backup-cli](cmd/backup) and [abs-restore-cli](cmd/restore) CLI tools,
built using [backup-go](https://github.com/aerospike/backup-go) library.
Refer to their respective README files for usage instructions.
Binaries for various platforms are released alongside the library and can be found under
[releases](https://github.com/aerospike/aerospike-backup-cli/releases).

## Core Features

### Standard Operations
- **Full backups**: Complete namespace or set backups
- **Incremental backups**: Time-based filtering for changed records
- **Parallel processing**: Configurable workers for optimal performance
- **Resume capability**: Continue interrupted backups from state files

### Advanced Filtering
- **Set-based**: Backup specific sets within namespaces
- **Bin filtering**: Include only specified bins
- **Time windows**: Records modified within date ranges
- **Partition filtering**: Backup specific partition ranges
- **Node/Rack targeting**: Geographic or hardware-specific backups

### Enterprise Features
- **Compression**: ZSTD compression for reduced storage
- **Encryption**: AES-128/256 encryption for data security
- **Cloud storage**: Direct backup to AWS S3, GCP Storage, Azure Blob
- **Secret management**: Integration with Aerospike Secret Agent
- **Rate limiting**: Bandwidth and RPS controls

## Quick Start

### Basic Backup
```bash
# Simple namespace backup
abs-backup-cli -h 127.0.0.1:3000 -n test -d /backup/test-namespace
```

### Basic Restore
```bash
# Restore from backup directory
abs-restore-cli -h 127.0.0.1:3000 -n test -d /backup/test-namespace
```

## Installation

### From Releases
Download pre-built binaries from [GitHub Releases](https://github.com/aerospike/aerospike-backup-cli/releases):

```bash
# Linux x64
wget https://github.com/aerospike/aerospike-backup-cli/releases/download/<version>/abs-restore-cli-<version>-<arch>.tar.gz
wget https://github.com/aerospike/aerospike-backup-cli/releases/download/<version>/abs-backup-cli-<version>-<arch>.tar.gz

# Extract
tar -xzvf abs-restore-cli-<version>-<arch>.tar.gz
tar -xzvf abs-backup-cli-<version>-<arch>.tar.gz

# Make executable
chmod +x abs-backup-cli abs-restore-cli
```

### Build from Source
```bash
# Build binaries
make build

# Install to /usr/bin
make install

# Uninstall
make uninstall
```

### Docker
Build and push a multi-platform Docker image:
```bash
DOCKER_USERNAME="<jfrog-username>" DOCKER_PASSWORD="<jfrog-password>" TAG="<tag>" make docker-buildx 
```

Build a Docker image for local use:
```bash
TAG="<tag>" make docker-build
```

A single docker image, including both tools, will be created.

Usage example:

```bash
docker run --rm aerospike-backup-tools:<tag> abs-restore-cli --help
docker run --rm aerospike-backup-tools:<tag> abs-backup-cli --help
```

### Linux Packages
To generate `.rpm` and `.deb` packages for supported Linux architectures (`linux/amd64`, `linux/arm64`):
```bash
make packages
```
The generated packages and their `sha256` checksum files will be located in the `/target` directory.

## Configuration Reference

Please look at [abs-backup-cli](cmd/backup/readme.md) and [abs-restore-cli](cmd/restore/readme.md) readme files for details.

## License

Apache License, Version 2.0. See [LICENSE](LICENSE) file for details.

## Support

- **Documentation**: [Aerospike Documentation](https://aerospike.com/docs/tools/backup/)
- **Issues**: [GitHub Issues](https://github.com/aerospike/aerospike-backup-cli/issues)
- **Community**: [Aerospike Community Forum](https://discuss.aerospike.com/)