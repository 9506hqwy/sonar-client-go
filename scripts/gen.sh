#!/bin/bash
set -euo pipefail

BASE_DIR=$(dirname $(dirname "$0"))
TMP_DIR=/tmp

curl -fsSL \
    -o "${TMP_DIR}/openapi.yml" \
    https://github.com/9506hqwy/openapi-spec-sonarqube-v1/raw/refs/heads/main/openapi.yml

go tool oapi-codegen -config "${BASE_DIR}/cfg.yml" "${TMP_DIR}/openapi.yml"
