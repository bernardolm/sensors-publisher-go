# syntax=docker/dockerfile:1

FROM golang:1.26.5-alpine AS go-base

WORKDIR /workspace

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# ---

FROM go-base AS tool

COPY dev/packages.golang /tmp/packages.golang
RUN --mount=type=cache,target=/go/pkg/mod xargs -n 1 go install < /tmp/packages.golang

# ---

FROM go-base AS build

ARG APP_GOOS
ARG APP_GOARCH
ARG APP_GOARM
ARG APP_VERSION=dev

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
	CGO_ENABLED=0 \
	GOOS="${APP_GOOS}" \
	GOARCH="${APP_GOARCH}" \
	GOARM="${APP_GOARM}" \
	go build \
		-trimpath \
		-tags="sonic avx" \
		-ldflags="-w -s -X github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/config.Version=${APP_VERSION}" \
		-o /out/sensors-publisher-go \
		./cmd/cli

# ---

FROM scratch AS binary-output

ARG APP_BINARY_NAME

COPY --from=build /out/sensors-publisher-go /${APP_BINARY_NAME}

# ---

FROM alpine:3.23 AS runtime

COPY bin/sensors-publisher-go-linux-armv7 /usr/local/bin/sensors-publisher-go

ENTRYPOINT ["/usr/local/bin/sensors-publisher-go"]

# ---

FROM alpine:3.23 AS apk

ARG APK_RELEASE=0
ARG APK_VERSION

RUN apk add --no-cache abuild gettext tar

COPY dev/docker/apk /templates
COPY bin/sensors-publisher-go-linux-armv7 /input/sensors-publisher-go
RUN APK_RELEASE="${APK_RELEASE}" \
	APK_VERSION="${APK_VERSION}" \
	PACKAGE_TEMPLATE_DIR=/templates \
	PACKAGE_BINARY=/input/sensors-publisher-go \
	sh /templates/package.sh

# ---

FROM scratch AS apk-output

COPY --from=apk /out/ /
