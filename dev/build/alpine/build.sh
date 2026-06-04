#!/usr/bin/env /bin/sh

set -eu

APP_TEMPLATE_DIR=$(cd "$(dirname "$0")" && pwd)
APP_ROOT=$(cd "${APP_TEMPLATE_DIR}/../../.." && pwd)
APP_PACKAGE_DIR="${APP_ROOT}/tmp/alpine"
APP_PACKAGE_ROOT="${APP_PACKAGE_DIR}/root"
APP_PACKAGE_CONTROL="${APP_PACKAGE_DIR}/control"
APP_PACKAGE_DATA_FILE="${APP_PACKAGE_DIR}/data.tar.gz"
APP_PACKAGE_CONTROL_FILE="${APP_PACKAGE_DIR}/control.tar.gz"
APP_ALPINE_VERSION="3.23"
APP_PACKAGE_NAME="sensors-publisher-go"
APP_PACKAGE_VERSION="${APK_VERSION:-0.1.0.$(date +%Y%m%d%H%M%S)}"
APP_PACKAGE_RELEASE="${APK_RELEASE:-0}"
APP_PACKAGE_ARCH="armv7"
APP_PACKAGE_FILE="${APP_ROOT}/bin/${APP_PACKAGE_NAME}-${APP_PACKAGE_VERSION}-r${APP_PACKAGE_RELEASE}.apk"
APP_BINARY_SOURCE="${APP_ROOT}/bin/${APP_PACKAGE_NAME}-linux-armv7"
APP_BINARY_TARGET="${APP_PACKAGE_ROOT}/usr/bin/${APP_PACKAGE_NAME}"
APP_CONFIG_PATH="/etc/${APP_PACKAGE_NAME}/config.env"
APP_CONFIG_DIR="${APP_PACKAGE_ROOT}/etc/${APP_PACKAGE_NAME}"
APP_CONFIG_FILE="${APP_CONFIG_DIR}/config.env"
APP_INIT_FILE="${APP_PACKAGE_ROOT}/etc/init.d/${APP_PACKAGE_NAME}"
APP_DATA_DIR="${APP_PACKAGE_ROOT}/var/lib/${APP_PACKAGE_NAME}"
APP_MODULES_FILE="${APP_PACKAGE_ROOT}/etc/modules-load.d/w1.conf"
APP_SQLITE_PATH="/var/lib/${APP_PACKAGE_NAME}/queue.db"
APP_PACKAGE_VARIABLES='${APP_PACKAGE_NAME} ${APP_PACKAGE_VERSION} ${APP_PACKAGE_RELEASE} ${APP_PACKAGE_ARCH} ${APP_PACKAGE_SIZE} ${APP_PACKAGE_BUILDDATE} ${APP_PACKAGE_DATAHASH} ${APP_CONFIG_PATH} ${APP_SQLITE_PATH}'

render_template() {
	local template
	local target
	local mode

	template="${1}"
	target="${2}"
	mode="${3}"

	envsubst "${APP_PACKAGE_VARIABLES}" < "${APP_TEMPLATE_DIR}/${template}" > "${target}"
	chmod "${mode}" "${target}"
}

run_abuild_tar() {
	if command -v abuild-tar >/dev/null 2>&1; then
		abuild-tar "$@"

		return
	fi

	docker run --rm -i "alpine:${APP_ALPINE_VERSION}" sh -c \
		'apk add --no-cache abuild >/dev/null && abuild-tar "$@"' \
		sh "$@"
}

create_alpine_data_package() {
	docker run --rm -v "${APP_PACKAGE_DIR}:/work" "alpine:${APP_ALPINE_VERSION}" sh -eu -c '
		apk add --no-cache abuild tar >/dev/null

		cd /work/root
		tar --xattrs -cf - * 2>/dev/null | abuild-tar --hash | gzip -9 > /work/data.tar.gz
	'
}

create_alpine_control_package() {
	docker run --rm -v "${APP_PACKAGE_DIR}:/work" "alpine:${APP_ALPINE_VERSION}" sh -eu -c '
		apk add --no-cache abuild tar >/dev/null

		cd /work/control
		tar -cf - .PKGINFO .post-install .post-upgrade .pre-deinstall | abuild-tar --cut | gzip -9 > /work/control.tar.gz
	'
}

write_package_info() {
	export APP_PACKAGE_SIZE
	export APP_PACKAGE_BUILDDATE
	export APP_PACKAGE_DATAHASH

	APP_PACKAGE_SIZE=$(du -sk "${APP_PACKAGE_ROOT}" | awk '{print $1 * 1024}')
	APP_PACKAGE_BUILDDATE=$(date +%s)
	APP_PACKAGE_DATAHASH=$(sha256sum "${APP_PACKAGE_DATA_FILE}" | awk '{print $1}')

	render_template "pkginfo.tpl.PKGINFO" "${APP_PACKAGE_CONTROL}/.PKGINFO" 0644
}

create_data_package() {
	if ! command -v abuild-tar >/dev/null 2>&1; then
		create_alpine_data_package

		return
	fi

	(
		cd "${APP_PACKAGE_ROOT}"
		tar --xattrs -cf - * 2>/dev/null | run_abuild_tar --hash | gzip -9 > "${APP_PACKAGE_DATA_FILE}"
	)
}

create_control_package() {
	if ! command -v abuild-tar >/dev/null 2>&1; then
		create_alpine_control_package

		return
	fi

	(
		cd "${APP_PACKAGE_CONTROL}"
		tar -cf - .PKGINFO .post-install .post-upgrade .pre-deinstall | run_abuild_tar --cut | gzip -9 > "${APP_PACKAGE_CONTROL_FILE}"
	)
}

create_package() {
	cat "${APP_PACKAGE_CONTROL_FILE}" "${APP_PACKAGE_DATA_FILE}" > "${APP_PACKAGE_FILE}"
}

cd "${APP_ROOT}"

export APP_PACKAGE_NAME
export APP_PACKAGE_VERSION
export APP_PACKAGE_RELEASE
export APP_PACKAGE_ARCH
export APP_CONFIG_PATH
export APP_SQLITE_PATH

GOOS=linux GOARCH=arm GOARM=7 make build-linux-armv7

rm -rf "${APP_PACKAGE_DIR}"
mkdir -p \
	"${APP_PACKAGE_CONTROL}" \
	"$(dirname "${APP_BINARY_TARGET}")" \
	"${APP_CONFIG_DIR}" \
	"$(dirname "${APP_INIT_FILE}")" \
	"$(dirname "${APP_MODULES_FILE}")" \
	"${APP_DATA_DIR}"

install -m 0755 "${APP_BINARY_SOURCE}" "${APP_BINARY_TARGET}"
render_template "config.tpl.env" "${APP_CONFIG_FILE}" 0644
render_template "init.tpl.openrc" "${APP_INIT_FILE}" 0755
render_template "w1.tpl.conf" "${APP_MODULES_FILE}" 0644
render_template "post-install.tpl.sh" "${APP_PACKAGE_CONTROL}/.post-install" 0755
render_template "post-upgrade.tpl.sh" "${APP_PACKAGE_CONTROL}/.post-upgrade" 0755
render_template "pre-deinstall.tpl.sh" "${APP_PACKAGE_CONTROL}/.pre-deinstall" 0755
create_data_package
write_package_info
create_control_package
create_package

echo "${APP_PACKAGE_FILE}"
