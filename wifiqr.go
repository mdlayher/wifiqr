// Package wifiqr implements support for generating WiFi QR codes.
package wifiqr

import (
	"errors"
	"fmt"
	"image"
	"strings"

	qrcode "github.com/skip2/go-qrcode"
)

// An Image is a WiFi QR code image which may be encoded in a variety of formats
// for display.
type Image struct{ qr *qrcode.QRCode }

// New generates an image containing a WiFi QR code using the parameters defined
// in Config. See the documentation of Config for details.
func New(cfg Config) (*Image, error) {
	s, err := cfg.encode()
	if err != nil {
		return nil, err
	}

	qr, err := qrcode.New(s, cfg.RecoveryLevel.convert())
	if err != nil {
		return nil, err
	}

	return &Image{qr: qr}, nil
}

// Image returns an image which may be encoded for external use.
func (i *Image) Image() image.Image {
	// TODO(mdlayher): make dimensions configurable.
	return i.qr.Image(-10)
}

// String returns a Unicode string for an Image, suitable for display in a
// terminal.
func (i *Image) String() string { return i.qr.ToSmallString(false) }

// Authentication defines the type of WiFi authentication used by a network.
type Authentication int

// Possible Authentication values.
const (
	None Authentication = iota
	WEP
	WPA
)

// RecoveryLevel defines the QR code recovery and error detection level.
type RecoveryLevel int

// Possible RecoveryLevel values. Medium is a good default for most
// applications.
const (
	Medium RecoveryLevel = iota
	Low
	High
	Highest
)

// convert converts a RecoveryLevel to a qrcode.RecoveryLevel.
func (rl RecoveryLevel) convert() qrcode.RecoveryLevel {
	// This conversion exists because the zero value for qrcode.RecoveryLevel is
	// "low" while we'd like medium to be the default instead, per the docs
	// stating it is a reasonable default.
	switch rl {
	case Low:
		return qrcode.Low
	case High:
		return qrcode.High
	case Highest:
		return qrcode.Highest
	case Medium:
		fallthrough
	default:
		// Medium or unspecified value.
		return qrcode.Medium
	}
}

// A Config defines the parameters for generating a WiFi QR code.
type Config struct {
	// Authentication specifies the type of WiFi authentication used by a
	// network. The zero value is "None", meaning an open network.
	Authentication Authentication

	// SSID and Password define the WiFi network name and password,
	// respectively.
	//
	// SSID is required and an error will be returned if it is unset.
	//
	// Password must be set for WEP or WPA authentication. It must be empty if
	// Authentication is set to None.
	SSID, Password string

	// Hidden defines whether the WiFi network is hidden.
	Hidden bool

	// RecoveryLevel specifies the QR code recovery and error detection level.
	// If unset, Medium is used as the default.
	RecoveryLevel RecoveryLevel
}

// A kv holds a key/value string pair used to generate WiFi QR code values.
type kv struct{ Key, Value string }

// authKV generates the kv for the WiFi authentication type.
func (c Config) authKV() (kv, error) {
	var v string
	switch c.Authentication {
	case None:
		// None has no value, but must also have no Password set.
		if c.Password != "" {
			return kv{}, errors.New("cannot set a password with no authentication type")
		}
	// WEP and WPA require passwords.
	case WEP:
		if c.Password == "" {
			return kv{}, errors.New("a password must be set for WEP authentication")
		}

		v = "WEP"
	case WPA:
		if c.Password == "" {
			return kv{}, errors.New("a password must be set for WPA authentication")
		}

		v = "WPA"
	default:
		return kv{}, errors.New("invalid authentication type")
	}

	return kv{Key: "T", Value: v}, nil
}

// encode encodes a Config as text suitable for generating a WiFi QR code. For
// documentation on the text format, see:
// https://www.qr-code-generator.com/solutions/wifi-qr-code/.
func (c Config) encode() (string, error) {
	// All configs set authentication type and SSID.
	auth, err := c.authKV()
	if err != nil {
		return "", err
	}

	if c.SSID == "" {
		return "", errors.New("no SSID is set")
	}

	kvs := []kv{auth, {Key: "S", Value: c.SSID}}

	// Password and hidden are optional depending on the network.
	if c.Password != "" {
		kvs = append(kvs, kv{Key: "P", Value: c.Password})
	}
	if c.Hidden {
		kvs = append(kvs, kv{Key: "H", Value: "true"})
	}

	// Combine each key/value pair with a colon and semicolon terminator.
	var sb strings.Builder
	for _, kv := range kvs {
		fmt.Fprintf(&sb, "%s:%s;", kv.Key, kv.Value)
	}

	return fmt.Sprintf("WIFI:%s;", sb.String()), nil
}
