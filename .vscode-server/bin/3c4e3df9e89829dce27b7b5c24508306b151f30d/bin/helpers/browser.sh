#!/usr/bin/env sh
#
# Copyright (c) Microsoft Corporation. All rights reserved.
#
ROOT=$(dirname "$(dirname "$(dirname "$0")")")

APP_NAME="code"
VERSION="1.55.2"
COMMIT="3c4e3df9e89829dce27b7b5c24508306b151f30d"
EXEC_NAME="code"
CLI_SCRIPT="$ROOT/out/vs/server/cli.js"
"$ROOT/node" "$CLI_SCRIPT" "$APP_NAME" "$VERSION" "$COMMIT" "$EXEC_NAME" "--openExternal" "$@"
