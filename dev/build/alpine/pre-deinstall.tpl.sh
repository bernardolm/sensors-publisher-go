#!/bin/sh

set -eu

if command -v rc-service >/dev/null 2>&1; then
	rc-service ${APP_PACKAGE_NAME} stop >/dev/null 2>&1 || true
fi

if command -v rc-update >/dev/null 2>&1; then
	rc-update del ${APP_PACKAGE_NAME} default >/dev/null 2>&1 || true
fi

exit 0
