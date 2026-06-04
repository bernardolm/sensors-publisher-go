# sensors-publisher-go

## Raspberry Pi 3B Alpine deploy

This project ships as a native Alpine `.apk` package built by GoReleaser inside Docker. The package installs the binary in `/usr/bin`, an OpenRC service in `/etc/init.d/sensors-publisher-go`, and app config in `/etc/sensors-publisher-go/config.env`.

The Raspberry Pi builds intentionally use `CGO_ENABLED=0`. SQLite is provided by the pure-Go `modernc.org/sqlite` driver, so no Alpine ARM/musl cross compiler is required.

The development commands also run through Docker targets from `dev/docker/Dockerfile`:

```bash
make watch
make format
make lint
make package-pi
```

Create the deploy config:

```bash
cp deploy/raspberry-pi.env.sample deploy/raspberry-pi.env
```

Edit `deploy/raspberry-pi.env`, then deploy:

```bash
make deploy-pi
```

To build only the Raspberry Pi APKs:

```bash
make package-pi
```

For unsigned local packages, the deploy script uses `apk add --allow-untrusted`. For production, sign the APK and install the public key in `/etc/apk/keys`.

## ds18b20

to use with 1-wire, run before:

```bash
sudo modprobe w1-gpio
sudo modprobe w1-therm
```
