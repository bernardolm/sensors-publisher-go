#!/bin/bash

echo "sh: building..."

go mod tidy

[ ! -d "bin" ] && mkdir bin

eval $(echo GOOS=linux GOARCH=${ARCH} go build -ldflags \"-s -w\" -o bin/sensors-publisher-go cmd/console/main.go)

[ ! -d "dist" ] && mkdir dist

# command -v upx 1>/dev/null || sudo apt install upx-ucl
# upx --lzma -o dist/sensors-publisher-go bin/sensors-publisher-go
cp -f bin/sensors-publisher-go dist/

cp -f installer/prod.env dist/.env
cat installer/install.env >> dist/.env
cp -f "installer/${OS}/local.sh" dist/

echo "in dist" && ls dist

[ -f "./installer/${OS}/pre-install.sh" ] && "./installer/${OS}/pre-install.sh"

export $(grep -v '^#' ./dist/.env | xargs)

envsubst < installer/${OS}/autostart > dist/autostart

cat dist/autostart

[ ! -d "tmp" ] && mkdir -p tmp

if [ ! -f tmp/makeself/makeself.sh ]; then
	curl -sL -o tmp/makeself.run https://github.com/megastep/makeself/releases/download/release-2.7.1/makeself-2.7.1.run
	chmod u+x tmp/makeself.run
	tmp/makeself.run --quiet --target tmp/makeself
	chmod u+x tmp/makeself/makeself.sh
fi

tmp/makeself/makeself.sh --notemp --nox11 dist/ "tmp/${RUN_FILE}" \
	"ms: installing..." "./local.sh" 1>/dev/null

echo "in tmp" && ls tmp

echo "sh: building done"
