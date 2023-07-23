package wifiqr

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/skip2/go-qrcode"
)

func TestImageString(t *testing.T) {
	const want = `█████████████████████████████████████
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
`

	img, err := New(Config{
		Authentication: WPA,
		SSID:           "Example",
		Password:       "thisisanexample",
	})
	if err != nil {
		t.Fatalf("failed to create image: %v", err)
	}

	// Display to terminal for manual verification with a phone.
	got := img.String()
	t.Logf("\n%s", got)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("unexpected QR code string (-want +got):\n%s", diff)
	}
}

func TestConfig_encode(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		str  string
		ok   bool
	}{
		{
			name: "no SSID",
			cfg:  Config{SSID: ""},
		},
		{
			name: "bad Authentication",
			cfg:  Config{Authentication: -1},
		},
		{
			name: "None with password",
			cfg:  Config{Password: "xxx"},
		},
		{
			name: "WEP no password",
			cfg:  Config{Authentication: WEP},
		},
		{
			name: "WPA no password",
			cfg:  Config{Authentication: WPA},
		},
		{
			name: "ok None",
			cfg:  Config{SSID: "Foo"},
			str:  "WIFI:T:;S:Foo;;",
			ok:   true,
		},
		{
			name: "ok WEP",
			cfg: Config{
				Authentication: WEP,
				SSID:           "Bar",
				Password:       "abc",
			},
			str: "WIFI:T:WEP;S:Bar;P:abc;;",
			ok:  true,
		},
		{
			name: "ok WPA",
			cfg: Config{
				Authentication: WPA,
				SSID:           "Baz",
				Password:       "def",
			},
			str: "WIFI:T:WPA;S:Baz;P:def;;",
			ok:  true,
		},
		{
			name: "ok hidden",
			cfg: Config{
				Authentication: WPA,
				SSID:           "Qux",
				Password:       "ghi",
				Hidden:         true,
			},
			str: "WIFI:T:WPA;S:Qux;P:ghi;H:true;;",
			ok:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str, err := tt.cfg.encode()
			if tt.ok && err != nil {
				t.Fatalf("failed to encode: %v", err)
			}
			if !tt.ok && err == nil {
				t.Fatal("expected an error, but none occurred")
			}
			if err != nil {
				t.Logf("err: %v", err)
				return
			}

			if diff := cmp.Diff(tt.str, str); diff != "" {
				t.Fatalf("unexpected encoded string (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRecoveryLevel_convert(t *testing.T) {
	tests := []struct {
		name string
		in   RecoveryLevel
		out  qrcode.RecoveryLevel
	}{
		{
			name: "zero",
			out:  qrcode.Medium,
		},
		{
			name: "unhandled",
			in:   100,
			out:  qrcode.Medium,
		},
		{
			name: "low",
			in:   Low,
			out:  qrcode.Low,
		},
		{
			name: "medium",
			in:   Medium,
			out:  qrcode.Medium,
		},
		{
			name: "high",
			in:   High,
			out:  qrcode.High,
		},
		{
			name: "highest",
			in:   Highest,
			out:  qrcode.Highest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.out, tt.in.convert()); diff != "" {
				t.Fatalf("unexpected recovery level (-want +got):\n%s", diff)
			}
		})
	}
}
