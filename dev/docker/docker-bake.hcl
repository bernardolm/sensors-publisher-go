target "base" {
  context    = "cwd://."
  dockerfile = "dev/docker/Dockerfile"
  args = {
    APP_VERSION = "dev"
  }
}

target "binary-darwin-arm64" {
  inherits = ["base"]
  target   = "binary-output"
  args = {
    APP_BINARY_NAME = "sensors-publisher-go-darwin-arm64"
    APP_GOARCH      = "arm64"
    APP_GOOS        = "darwin"
  }
  output = ["type=local,dest=bin"]
}

target "binary-linux-amd64" {
  inherits = ["base"]
  target   = "binary-output"
  args = {
    APP_BINARY_NAME = "sensors-publisher-go-linux-amd64"
    APP_GOARCH      = "amd64"
    APP_GOOS        = "linux"
  }
  output = ["type=local,dest=bin"]
}

target "binary-linux-armv7" {
  inherits = ["base"]
  target   = "binary-output"
  args = {
    APP_BINARY_NAME = "sensors-publisher-go-linux-armv7"
    APP_GOARCH      = "arm"
    APP_GOARM       = "7"
    APP_GOOS        = "linux"
  }
  output = ["type=local,dest=bin"]
}

group "build-artifacts" {
  targets = [
    "binary-darwin-arm64",
    "binary-linux-amd64",
    "binary-linux-armv7",
  ]
}

target "apk" {
  inherits = ["base"]
  target   = "apk-output"
  args = {
    APK_VERSION = "0.1.0-dev"
  }
  output = ["type=local,dest=bin"]
}

target "image-pi" {
  inherits  = ["base"]
  target    = "runtime"
  platforms = ["linux/arm/v7"]
  tags      = ["sensors-publisher-go:armv7"]
  output    = ["type=docker"]
}
