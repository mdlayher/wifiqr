package wifiqr

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
