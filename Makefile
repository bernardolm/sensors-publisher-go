ifneq (,$(wildcard ./.env))
	include ./.env
	export
endif

.EXPORT_ALL_VARIABLES:

SSH_PUB_KEY_PATH="${HOME}/.ssh/id_ed25519.pub"

DOCKER_CONTAINER_NAME=alpine-sensors
DOCKER_IMAGE_NAME="${GITHUB_USER}/${DOCKER_CONTAINER_NAME}:latest"
SSH_PUB_KEY=$(shell [ -f ${SSH_PUB_KEY_PATH} ] && /bin/cat ${SSH_PUB_KEY_PATH})

reset:
	@reset

air: reset
	@command -v air 1>/dev/null || go install github.com/air-verse/air@latest
	@command -v goimports 1>/dev/null || go install golang.org/x/tools/cmd/goimports@latest
	@[ "${OS}" = "darwin" ] && air -build.exclude_dir=tmp -c=dev/.air.toml -d || \
		(command -v expect_unbuffer || sudo apt install -y expect ; \
		expect_unbuffer air -build.exclude_dir=tmp -c=dev/.air.toml -d)

start:
	@go run ./cmd/console/main.go

install:
	@./dev/installer/start.sh

install-alpine:
	@ARCH="arm GOARM=7" OS=alpine make install

install-debian:
	@ARCH=amd64 OS=debian make install

install-docker:
	@(docker stop ${DOCKER_CONTAINER_NAME} 2>/dev/null && sleep 1) || true
	@docker build --no-cache \
		--progress=tty \
		--build-arg SSH_PUB_KEY="${SSH_PUB_KEY}" \
		-f "docker/Dockerfile" \
		-t ${DOCKER_IMAGE_NAME} \
		docker
	@docker run -d --rm --privileged \
		--name ${DOCKER_CONTAINER_NAME} \
		${DOCKER_IMAGE_NAME} /bin/bash
	@ARCH=amd64 OS=alpine-docker make install

debug:
	@source ./dev/dist/install.dev
	@ssh "${SYSTEM_USER}@${SYSTEM_HOST}" "sudo -S tail -n50 -f ${LOG_FILEPATH}"

format:
	@echo "> formatting"
	@source ./.githooks/pre-commit

format-all:
	@FILES=all make format
