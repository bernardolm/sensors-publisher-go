.ONESHELL:
.PHONY: format lint bin docker reset config

ifneq (,$(wildcard ./.env))
	include ./.env
	export
endif

PWD = $(shell pwd)

define docker_build
	. ./dev/build/envs.sh \
	&& IMAGE="$${APP_NAME}:$${STAGE}" \
	&& if [ "$${STAGE}" = "dev" ] || ! docker image inspect "$${IMAGE}" >/dev/null 2>&1; then \
		echo "> docker building $${IMAGE}... "; \
		mkdir -p bin ; \
		docker build --target "$${STAGE}" \
			$${DOCKER_BUILD_EXTRA_ARGS} \
			--build-arg APP_CMD_PATH \
			--build-arg APP_LDFLAGS \
			--build-arg APP_NAME \
			--tag "$${IMAGE}" .; \
	fi
endef

define docker_run
	export STAGE=base \
	&& $(docker_build) \
	&& . ./dev/build/envs.sh \
	&& docker run -t --rm \
		--env-file .env \
		-v "$(PWD):/usr/app" \
		-v "$${APP_NAME}-go-build-cache:/root/.cache/go-build" \
		-w /usr/app \
		"$${APP_NAME}:$${STAGE}" \
		go run "$${APP_CMD_PATH}" $${CLI_COMMAND} $${CLI_ARGS}
endef

define docker_golangci
	export STAGE=dev \
	&& $(docker_build) \
	&& . ./dev/build/envs.sh \
	&& docker run -t --rm \
		-v "$(PWD):/usr/app" \
		-v "$${APP_NAME}-go-build-cache:/root/.cache/go-build" \
		-w /usr/app \
		"$${APP_NAME}:$${STAGE}" \
		golangci-lint --verbose "$${GOLANGCI_COMMAND}" "$${GOLANGCI_FILE_LIST}"
endef

bin: export STAGE=build
bin: export DOCKER_BUILD_EXTRA_ARGS=--output type=local,dest=./bin
bin:
	rm -rf bin || true
	$(docker_build)

format:
	@export GOLANGCI_COMMAND=fmt GOLANGCI_FILE_LIST=./... \
	&& $(docker_golangci)

lint:
	@export GOLANGCI_COMMAND=run GOLANGCI_FILE_LIST=./... \
	&& $(docker_golangci)

run:
	@export CLI_COMMAND=foobar \
	&& $(docker_run)

docker-destroy:
	docker compose stop
	docker compose rm -f

docker:
	eval docker compose up --remove-orphans $(DOCKER_ARGS)

reset:
	reset

config:
	ln -sf "$(PWD)/.githooks/pre-commit" "$(PWD)/.git/hooks/pre-commit"

gotools:
	go get -tool $(shell /bin/cat packages.golang)
