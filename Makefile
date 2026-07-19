.PHONY: \
	arch-check \
	build \
	config \
	docker-destroy \
	format \
	image-pi \
	lint \
	package-pi \
	run \
	test \
	watch

COMPOSE := docker compose --project-directory $(CURDIR) --file $(CURDIR)/dev/docker/compose.yaml
BAKE := docker buildx bake --file $(CURDIR)/dev/docker/docker-bake.hcl
APP_VERSION ?= $(shell date +%y%m%d%H%M%S)_$(shell git rev-parse --short HEAD)
APK_VERSION ?= 0.1.0.$(shell date +%Y%m%d%H%M%S)

build:
	rm -rf bin
	$(BAKE) --set "*.args.APP_VERSION=$(APP_VERSION)" build-artifacts

package-pi: build
	$(BAKE) --set "apk.args.APK_VERSION=$(APK_VERSION)" apk

image-pi: build
	$(BAKE) image-pi

run:
	$(COMPOSE) up --build app

watch:
	$(COMPOSE) up --build --watch app

format:
	$(COMPOSE) build tool
	$(COMPOSE) run --rm tool golangci-lint --verbose fmt ./...

lint:
	$(COMPOSE) build tool
	$(COMPOSE) run --rm tool golangci-lint --verbose run --fix ./...
	$(COMPOSE) run --rm tool go-arch-lint check --project-path /workspace

arch-check:
	$(COMPOSE) run --rm --build tool go-arch-lint check --project-path /workspace

test:
	$(COMPOSE) run --rm --build tool go test -count=1 ./...

docker-destroy:
	$(COMPOSE) down --remove-orphans

config:
	ln -sf "$(CURDIR)/.githooks/pre-commit" "$(CURDIR)/.git/hooks/pre-commit"
