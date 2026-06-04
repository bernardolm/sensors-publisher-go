#!/bin/sh

set -eu

if command -v rc-update >/dev/null 2>&1; then
	rc-update add ${APP_PACKAGE_NAME} default >/dev/null 2>&1 || true
fi

exit 0
