#!/bin/sh

set -eu

if command -v rc-update >/dev/null 2>&1; then
	rc-update add modules boot >/dev/null 2>&1 || true
	rc-update add ${APP_PACKAGE_NAME} default >/dev/null 2>&1 || true
fi

if command -v modprobe >/dev/null 2>&1; then
	modprobe -q w1-gpio >/dev/null 2>&1 || true
	modprobe -q w1-therm >/dev/null 2>&1 || true
fi

mkdir -p /boot
touch /boot/usercfg.txt

if ! grep -qx 'dtoverlay=w1-gpio' /boot/usercfg.txt; then
	if [ -s /boot/usercfg.txt ] &&
		[ "$(tail -c 1 /boot/usercfg.txt | od -An -t x1 | tr -d ' ')" != "0a" ]; then
		printf '\n' >> /boot/usercfg.txt
	fi

	printf 'dtoverlay=w1-gpio\n' >> /boot/usercfg.txt
	echo "Added dtoverlay=w1-gpio to /boot/usercfg.txt. Reboot to enable Raspberry Pi 1-Wire GPIO."
fi

exit 0
