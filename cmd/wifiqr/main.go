// Command wifiqr generates a WiFi QR code image.
package main

import (
	"flag"
	"fmt"
	"image/png"
	"io"
	"log"
	"os"

	"github.com/mdlayher/wifiqr"
	"golang.org/x/term"
)

func main() {
	var (
		sFlag = flag.String("s", "Example", "SSID, or WiFi network name")
		pFlag = flag.String("p", "thisisanexample", "WiFi network password")
	)
	flag.Parse()

	cfg := wifiqr.Config{
		Authentication: wifiqr.WPA,
		SSID:           *sFlag,
		Password:       *pFlag,
	}

	img, err := wifiqr.New(cfg)
	if err != nil {
		log.Fatalf("failed to create QR code: %v", err)
	}

	if term.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Printf("SSID: %q, password: %q\n\n", cfg.SSID, cfg.Password)

		if _, err := io.WriteString(os.Stdout, img.String()); err != nil {
			log.Fatalf("failed to encode QR code for terminal: %v", err)
		}

		return
	}

	if err := png.Encode(os.Stdout, img.Image()); err != nil {
		log.Fatalf("failed to encode PNG: %v", err)
	}
}
