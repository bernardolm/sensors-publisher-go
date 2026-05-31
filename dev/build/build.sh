#!/usr/bin/env /bin/sh

. ./dev/build/envs.sh

eval "
	go build \
		-o ${APP_BIN_PATH} \
		-ldflags '${APP_LDFLAGS}' \
		-tags '${APP_TAGS}' \
		${APP_CMD_PATH}
"

test -z "$DEBUG" || /bin/ls -lah ${APP_BIN_PATH}
