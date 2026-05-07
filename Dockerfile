FROM golang:alpine AS base
WORKDIR /usr/app
COPY go.mod go.sum ./
RUN go mod download -x

FROM base AS dev
WORKDIR /usr/app
COPY packages.golang .
RUN go get -tool $(/bin/cat packages.golang)

FROM base AS basebuild
WORKDIR /usr/app
ARG APP_LDFLAGS
ARG APP_NAME
ARG APP_CMD_PATH
RUN true \
	&& go build -tags="sonic avx" \
	-ldflags "-w -s ${APP_LDFLAGS}" \
	-o /usr/app/bin/${APP_NAME} ${APP_CMD_PATH}

FROM scratch AS build
ARG APP_NAME
COPY --from=basebuild /usr/app/bin/${APP_NAME} /${APP_NAME}

FROM alpine AS prod
ARG APP_NAME
COPY --from=basebuild /usr/app/bin/${APP_NAME} /${APP_NAME}
ENTRYPOINT [ "/${APP_NAME}" ]
