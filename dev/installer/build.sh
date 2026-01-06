#!/bin/bash

echo "sh: building..."

go mod tidy

[ ! -d "./dev/bin" ] && mkdir ./dev/bin

eval $(echo GOOS=linux GOARCH=${ARCH} go build -ldflags \"-s -w\" -o ./dev/bin/sensors-publisher-go ./cmd/console/main.go)

[ ! -d "./dev/dist" ] && mkdir ./dev/dist

# command -v upx 1>/dev/null || sudo apt install upx-ucl
# upx --lzma -o ./dev/dist/sensors-publisher-go ./dev/bin/sensors-publisher-go
cp -f ./dev/bin/sensors-publisher-go ./dev/dist/

cp -f ./dev/installer/prod.env ./dev/dist/.env
cat ./dev/installer/install.env >> ./dev/dist/.env
cp -f "./dev/installer/${OS}/local.sh" ./dev/dist/

echo "in dist" && ls ./dev/dist

[ -f "./dev/installer/${OS}/pre-install.sh" ] && "./dev/installer/${OS}/pre-install.sh"

export $(grep -v '^#' ./dev/dist/.env | xargs)

envsubst < ./dev/installer/${OS}/autostart > ./dev/dist/autostart

cat ./dev/dist/autostart

[ ! -d "tmp" ] && mkdir -p tmp

if [ ! -f tmp/makeself/makeself.sh ]; then
	curl -sL -o tmp/makeself.run https://github.com/megastep/makeself/releases/download/release-2.7.1/makeself-2.7.1.run
	chmod u+x tmp/makeself.run
	tmp/makeself.run --quiet --target tmp/makeself
	chmod u+x tmp/makeself/makeself.sh
fi

tmp/makeself/makeself.sh --notemp --nox11 ./dev/dist/ "tmp/${RUN_FILE}" \
	"ms: installing..." "./local.sh" 1>/dev/null

echo "in tmp" && ls tmp

echo "sh: building done"
