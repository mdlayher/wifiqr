// Command wifiqr generates a WiFi QR code for an example network.
package main

import (
	"image/png"
	"log"
	"os"

	"github.com/mdlayher/wifiqr"
)

func main() {
	img, err := wifiqr.New(wifiqr.Config{
		Authentication: wifiqr.WPA,
		SSID:           "Example",
		Password:       "thisisanexample",
	})
	if err != nil {
		log.Fatalf("failed to create QR code: %v", err)
	}

	if err := png.Encode(os.Stdout, img.Image()); err != nil {
		log.Fatalf("failed to encode PNG: %v", err)
	}
}
