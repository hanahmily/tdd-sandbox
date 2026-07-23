// Package e2e drives the difference feature end to end: it builds the real
// cmd/difference binary and runs it as a subprocess, asserting on its stdout.
// This simulates a genuine user invoking the CLI. RED until Sub is implemented.
package e2e

import (
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestE2EDifferenceCLI builds and runs the difference command against the
// three required operand classes (positive, negative, zero). Named with the
// TestE2E prefix so the e2e gate (`go test -run TestE2E ./...`) matches it.
func TestE2EDifferenceCLI(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "difference")
	build := exec.Command("go", "build", "-o", bin, "github.com/hanahmily/tdd-sandbox/cmd/difference")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("building cmd/difference failed: %v\n%s", err, out)
	}

	cases := []struct {
		name string
		a, b string
		want string
	}{
		{"positive operands", "7", "3", "4"},
		{"negative operands", "-5", "-2", "-3"},
		{"zero operands", "0", "0", "0"},
		{"positive minus larger is negative", "3", "8", "-5"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := exec.Command(bin, tc.a, tc.b).CombinedOutput()
			if err != nil {
				t.Fatalf("running difference %s %s failed: %v\n%s", tc.a, tc.b, err, out)
			}
			if got := strings.TrimSpace(string(out)); got != tc.want {
				t.Errorf("difference %s %s = %q, want %q", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
