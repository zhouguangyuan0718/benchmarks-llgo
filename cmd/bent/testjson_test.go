// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func writeTestJSONFixture(t *testing.T, directory string) string {
	t.Helper()
	contents, err := scripts.ReadFile("scripts/testjson")
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(directory, "testjson")
	if err := os.WriteFile(path, contents, 0o755); err != nil {
		t.Fatal(err)
	}
	return path
}

func writeExecutableFixture(t *testing.T, path, contents string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(contents), 0o755); err != nil {
		t.Fatal(err)
	}
}

func fixtureCommand(t *testing.T, directory, wrapper string, args ...string) *exec.Cmd {
	t.Helper()
	cmd := exec.Command(wrapper, args...)
	cmd.Env = replaceEnv(os.Environ(), "PATH", directory+string(os.PathListSeparator)+os.Getenv("PATH"))
	cmd.Env = replaceEnv(cmd.Env, "BENT_BENCH", "fixture")
	return cmd
}

func TestTestJSONWrapperWithoutWatchdog(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("the embedded wrapper requires bash")
	}
	directory := t.TempDir()
	wrapper := writeTestJSONFixture(t, directory)
	writeExecutableFixture(t, filepath.Join(directory, "go"), "#!/usr/bin/env bash\nprintf '%s\\n' \"$@\"\n")

	output, err := fixtureCommand(t, directory, wrapper, "/tmp/testbin", "-test.short").CombinedOutput()
	if err != nil {
		t.Fatalf("testjson failed: %v\n%s", err, output)
	}
	got := strings.Fields(string(output))
	want := []string{"tool", "test2json", "-p", "fixture", "/tmp/testbin", "-test.short"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("testjson arguments = %q, want %q", got, want)
	}
}

func TestTestJSONWrapperWatchdog(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("the embedded wrapper requires bash")
	}
	directory := t.TempDir()
	wrapper := writeTestJSONFixture(t, directory)
	writeExecutableFixture(t, filepath.Join(directory, "timeout"), "#!/usr/bin/env bash\nprintf '%s\\n' \"$@\"\nexit 23\n")

	output, err := fixtureCommand(t, directory, wrapper, "--timeout", "6m", "/tmp/testbin", "-test.short").CombinedOutput()
	exitError, ok := err.(*exec.ExitError)
	if !ok || exitError.ExitCode() != 23 {
		t.Fatalf("testjson error = %v, want exit status 23\n%s", err, output)
	}
	got := strings.Fields(string(output))
	want := []string{
		"--signal=TERM", "--kill-after=30s", "6m",
		"go", "tool", "test2json", "-p", "fixture", "/tmp/testbin", "-test.short",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("timeout arguments = %q, want %q", got, want)
	}
}
