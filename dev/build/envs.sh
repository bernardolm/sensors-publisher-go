#!/usr/bin/env /bin/sh

function app_name() {
	basename ${PWD}
}

function app_ldflags() {
	local commit_hash=$(git rev-parse --short HEAD)
	local is_dirty="no"

	if [ -n "$(git diff --name-only -- '*.go')" ] ||
		[ -n "$(git diff --cached --name-only -- '*.go')" ] ||
		[ -n "$(git ls-files --others --exclude-standard -- '*.go')" ]; then
		is_dirty="yes"
	fi

	local version=$(date +%y%m%d%H%M%S)

	[ -n "${commit_hash}" ] && version="${version}_${commit_hash}"
	[ "${is_dirty}" = "yes" ] && version="${version}_dirty"

	local ns='github.com/bernardolm/sensors-publisher-go'

	echo "-X ${ns}/pkg/infrastructure/config.Version=${version}"
}

export APP_CMD_PATH='cmd/cli/main.go'
export APP_LDFLAGS="-w -s $(app_ldflags)"
export APP_NAME=$(app_name)
export APP_TAGS='sonic avx'
export APP_BIN_PATH="bin/${APP_NAME}-${GOOS}-${GOARCH}${GOARM:+v${GOARM}}"

test -z "$DEBUG" || echo "APP_CMD_PATH:	${APP_CMD_PATH}"
test -z "$DEBUG" || echo "APP_LDFLAGS:	${APP_LDFLAGS}"
test -z "$DEBUG" || echo "APP_NAME:	${APP_NAME}"
test -z "$DEBUG" || echo "APP_TAGS:	${APP_TAGS}"
test -z "$DEBUG" || echo "APP_BIN_PATH:	${APP_BIN_PATH}"
