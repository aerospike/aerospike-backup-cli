group default {
  targets = [
    "aerospike-backup-cli"
  ]
}

variable CONTEXT {
  default = null
}

variable LATEST {
  default = false
}

variable TAG {
  default = ""
}

variable GIT_BRANCH {
  default = null
}

variable GIT_COMMIT_SHA {
  default = null
}

variable VERSION {
  default = null
}

variable ISO8601 {
  default = null
}

variable REPO {
  default = "aerospike/aerospike-backup-cli"
}

variable PLATFORMS {
  default = "linux/amd64,linux/arm64"
}

variable REGISTRY {
  default = "docker.io"
}

variable GO_VERSION {
  default = "1.23.10"
}

variable CACHE_FROM {
  default = ""
}

variable CACHE_TO {
  default = ""
}

variable OUTPUT {
  default = "type=image,push=true"
}

function norm {
  params = [value]

  result = value == null || value == "" ? [] : length(regexall(" ", value)) > 0 ? split(" ", value) : [value]
}

function tags {
  params = [service]
  result = LATEST == true ? [
    "${REPO}/${service}:${TAG}",
    "${REPO}/${service}:latest"
  ] : ["${REPO}/${service}:${TAG}"]
}

target aerospike-backup-cli {
  labels = {
    "org.opencontainers.image.title"         = "Aerospike Backup CLI"
    "org.opencontainers.image.description"   = "Command-line tools for backing up and restoring Aerospike data"
    "org.opencontainers.image.documentation" = "https://github.com/aerospike/aerospike-backup-cli?tab=readme-ov-file#aerospike-backup-cli"
    "org.opencontainers.image.base.name"     = "registry.access.redhat.com/ubi9/ubi-minimal"
    "org.opencontainers.image.source"        = "https://github.com/aerospike/aerospike-backup-cli/tree/${GIT_BRANCH}"
    "org.opencontainers.image.vendor"        = "Aerospike"
    "org.opencontainers.image.version"       = "${VERSION}"
    "org.opencontainers.image.url"           = "https://github.com/aerospike/aerospike-backup-cli"
    "org.opencontainers.image.licenses"      = "Apache-2.0"
    "org.opencontainers.image.revision"      = "${GIT_COMMIT_SHA}"
    "org.opencontainers.image.created"       = "${ISO8601}"
  }

  args = {
    GO_VERSION = "${GO_VERSION}"
    REGISTRY   = "${REGISTRY}"
  }

  secret     = ["id=GOPROXY,env=GOPROXY"]
  context    = "${CONTEXT}"
  dockerfile = "Dockerfile"
  platforms  = split(",", replace("${PLATFORMS}", " ", ","))
  cache-to   = norm("${CACHE_TO}")
  cache-from = norm("${CACHE_FROM}")

  tags   = tags("aerospike-backup-cli")
  output = norm("${OUTPUT}")
}
