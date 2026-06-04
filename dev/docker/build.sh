#!/usr/bin/env /bin/sh

export GOOS=linux
export GOARCH=arm
export GOARM=7

. ./dev/build/build.sh

cp ${APP_BIN_PATH} ./dev/deploy/app

docker buildx build \
	--platform linux/arm/v7 \
	--tag bernardolm/${APP_NAME}:latest \
	--push \
	./dev/deploy

test -z "$DEBUG" || docker run -t --rm \
	--env-file ./.env \
	bernardolm/${APP_NAME}:latest
