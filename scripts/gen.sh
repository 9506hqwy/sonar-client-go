#!/bin/bash
set -euo pipefail

VERSION=v9.9

BASE_DIR=$(dirname "$(dirname "$0")")
TMP_DIR=/tmp

curl -fsSL \
    -o "${TMP_DIR}/openapi-orig.yml" \
    "https://github.com/9506hqwy/openapi-spec-sonarqube-v1/raw/refs/heads/${VERSION}/openapi.yml"

# shellcheck disable=SC2016
cat "${TMP_DIR}/openapi-orig.yml" | \
    yq '.. |= select(has("parameters")) |= .parameters |= map(.description as $x | .x-oapi-codegen-extra-tags.jsonschema = "description=" + $x)' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= trim' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= sub("\"", "\\\"")' \
    > "${TMP_DIR}/openapi.yml"

go tool oapi-codegen -config "${BASE_DIR}/cfg.yml" "${TMP_DIR}/openapi.yml"
