#!/bin/sh

set -eu

APP_PACKAGE_NAME="sensors-publisher-go"
APP_PACKAGE_ARCH="armv7"
APP_CONFIG_PATH="/etc/${APP_PACKAGE_NAME}/config.env"
APP_SQLITE_PATH="/var/lib/${APP_PACKAGE_NAME}/data.db"
APP_PACKAGE_DIR="/tmp/package"
APP_PACKAGE_ROOT="${APP_PACKAGE_DIR}/root"
APP_PACKAGE_CONTROL="${APP_PACKAGE_DIR}/control"
APP_PACKAGE_DATA_FILE="${APP_PACKAGE_DIR}/data.tar.gz"
APP_PACKAGE_CONTROL_FILE="${APP_PACKAGE_DIR}/control.tar.gz"
APP_PACKAGE_FILE="/out/${APP_PACKAGE_NAME}-${APK_VERSION}-r${APK_RELEASE}.apk"
APP_BINARY_TARGET="${APP_PACKAGE_ROOT}/usr/bin/${APP_PACKAGE_NAME}"
APP_CONFIG_DIR="${APP_PACKAGE_ROOT}/etc/${APP_PACKAGE_NAME}"
APP_INIT_FILE="${APP_PACKAGE_ROOT}/etc/init.d/${APP_PACKAGE_NAME}"
APP_DATA_DIR="${APP_PACKAGE_ROOT}/var/lib/${APP_PACKAGE_NAME}"
APP_MODULES_FILE="${APP_PACKAGE_ROOT}/etc/modules-load.d/w1.conf"
APP_PACKAGE_VARIABLES='${APP_PACKAGE_NAME} ${APK_VERSION} ${APK_RELEASE} ${APP_PACKAGE_ARCH} ${APP_PACKAGE_SIZE} ${APP_PACKAGE_BUILDDATE} ${APP_PACKAGE_DATAHASH} ${APP_CONFIG_PATH} ${APP_SQLITE_PATH}'

export APK_RELEASE
export APK_VERSION
export APP_CONFIG_PATH
export APP_PACKAGE_ARCH
export APP_PACKAGE_NAME
export APP_SQLITE_PATH

render_template() {
	app_template="${1}"
	app_target="${2}"
	app_mode="${3}"

	envsubst "${APP_PACKAGE_VARIABLES}" < "${PACKAGE_TEMPLATE_DIR}/${app_template}" > "${app_target}"
	chmod "${app_mode}" "${app_target}"
}

mkdir -p \
	"${APP_PACKAGE_CONTROL}" \
	"$(dirname "${APP_BINARY_TARGET}")" \
	"${APP_CONFIG_DIR}" \
	"$(dirname "${APP_INIT_FILE}")" \
	"$(dirname "${APP_MODULES_FILE}")" \
	"${APP_DATA_DIR}" \
	"/out"

install -m 0755 "${PACKAGE_BINARY}" "${APP_BINARY_TARGET}"
render_template "config.tpl.env" "${APP_CONFIG_DIR}/config.env" 0644
render_template "init.tpl.openrc" "${APP_INIT_FILE}" 0755
render_template "w1.tpl.conf" "${APP_MODULES_FILE}" 0644
render_template "post-install.tpl.sh" "${APP_PACKAGE_CONTROL}/.post-install" 0755
render_template "post-upgrade.tpl.sh" "${APP_PACKAGE_CONTROL}/.post-upgrade" 0755
render_template "pre-deinstall.tpl.sh" "${APP_PACKAGE_CONTROL}/.pre-deinstall" 0755

(
	cd "${APP_PACKAGE_ROOT}"
	tar --xattrs -cf - * 2>/dev/null | abuild-tar --hash | gzip -9 > "${APP_PACKAGE_DATA_FILE}"
)

APP_PACKAGE_SIZE=$(du -sk "${APP_PACKAGE_ROOT}" | awk '{print $1 * 1024}')
APP_PACKAGE_BUILDDATE=$(date +%s)
APP_PACKAGE_DATAHASH=$(sha256sum "${APP_PACKAGE_DATA_FILE}" | awk '{print $1}')
export APP_PACKAGE_SIZE
export APP_PACKAGE_BUILDDATE
export APP_PACKAGE_DATAHASH
render_template "pkginfo.tpl.PKGINFO" "${APP_PACKAGE_CONTROL}/.PKGINFO" 0644

(
	cd "${APP_PACKAGE_CONTROL}"
	tar -cf - .PKGINFO .post-install .post-upgrade .pre-deinstall | abuild-tar --cut | gzip -9 > "${APP_PACKAGE_CONTROL_FILE}"
)

cat "${APP_PACKAGE_CONTROL_FILE}" "${APP_PACKAGE_DATA_FILE}" > "${APP_PACKAGE_FILE}"
