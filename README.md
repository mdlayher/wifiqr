# wifiqr [![Test Status](https://github.com/mdlayher/wifiqr/workflows/Test/badge.svg)](https://github.com/mdlayher/wifiqr/actions) [![Go Reference](https://pkg.go.dev/badge/github.com/mdlayher/wifiqr.svg)](https://pkg.go.dev/github.com/mdlayher/wifiqr) [![Go Report Card](https://goreportcard.com/badge/github.com/mdlayher/wifiqr)](https://goreportcard.com/report/github.com/mdlayher/wifiqr)

Package `wifiqr` implements support for generating WiFi QR codes. MIT Licensed.

## Example

Generate a QR code image and redirect stdout to create a PNG file.

```
$ go run cmd/wifiqr/main.go > example.png
```

This produces:

![example](https://github.com/mdlayher/wifiqr/assets/1926905/6f46e6d1-a147-4d1a-8afb-0bd1e38034a7)

Alternatively, if stdout is a terminal (and not redirected) you can display the
QR code directly.

```
$ go run ./cmd/wifiqr/main.go
SSID: "Example", password: "thisisanexample"

█████████████████████████████████████
█████████████████████████████████████
████ ▄▄▄▄▄ ██ ▀ ▀▀▄ ▄▀▄█ █ ▄▄▄▄▄ ████
████ █   █ █ ▄▀ █ ▄▀▄▄ ▄▄█ █   █ ████
████ █▄▄▄█ █ █ ▀█▀▄▄▄▀▄ ▀█ █▄▄▄█ ████
████▄▄▄▄▄▄▄█ ▀▄█▄▀ █ ▀ ▀▄█▄▄▄▄▄▄▄████
████ ▀▄ ▄ ▄█▀▄█ ▄▄▀ ██▄███▄   ▄█▀████
████ ▄▄▄█▄▄▄▀▀ ▀▀██▄ ▄██ █▀▀ ▀▄█ ████
████ █▀ █▀▄▄▄ ▄ ▄▀█  ▄▄▀▄ ▄██    ████
████▀███▀▀▄▀▄▀▀▄▄ ▄█▀ ▀ ▄▀▀▄▀▀ ██████
████▀▀ ▀▄▀▄█▀ ████▄ ▄█▄█ ▄█▄▀▀▄▀▀████
████ █ ██▀▄▄▄▀ ▄▀█▀▄▀▄█▀▀▀▀  ▀█▄▀████
████▄██▄▄█▄█▀▄█ ▀▀▀ █▄▄█ ▄▄▄ ▀ ▄▄████
████ ▄▄▄▄▄ █▀█▄▀▄ ▄█ ▄▀▀ █▄█  ▄██████
████ █   █ █  ▄▀██▄  █▄▄ ▄  ▄██▀█████
████ █▄▄▄█ █▄▄█▄▀▀ ▄▀ ▀▄███▀█ ▀ ▀████
████▄▄▄▄▄▄▄█▄▄▄▄▄██▄█▄████▄█▄█▄▄█████
█████████████████████████████████████
▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
```
