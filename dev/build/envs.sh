#!/usr/bin/env /bin/sh

function app_name() {
	basename ${PWD}
}

function app_cmd_path() {
	echo './cmd/cli/main.go'
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

export APP_LDFLAGS=$(app_ldflags)
export APP_NAME=$(app_name)
export APP_CMD_PATH=$(app_cmd_path)

echo "APP_LDFLAGS is '${APP_LDFLAGS}'"
echo "APP_NAME is '${APP_NAME}'"
echo "APP_CMD_PATH is '${APP_CMD_PATH}'"
