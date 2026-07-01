// Copyright 2024-2026 Qualcomm Technologies, Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-3-Clause

// Package testutil hosts shared helpers used only by *_test.go files.
package testutil

import (
	"io"
	"os"
	"testing"
)

// CaptureOutput swaps os.Stdout and os.Stderr with pipes for the duration of
// fn, then returns whatever fn wrote to each plus fn's own error. Use it to
// test code that reaches for fmt.Print* directly instead of an injectable
// io.Writer.
func CaptureOutput(t *testing.T, fn func() error) (stdout, stderr string, fnErr error) {
	t.Helper()

	origOut, origErr := os.Stdout, os.Stderr
	rOut, wOut, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe stdout: %v", err)
	}
	rErr, wErr, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe stderr: %v", err)
	}
	os.Stdout, os.Stderr = wOut, wErr

	outCh := make(chan string, 1)
	errCh := make(chan string, 1)
	go func() { b, _ := io.ReadAll(rOut); outCh <- string(b) }()
	go func() { b, _ := io.ReadAll(rErr); errCh <- string(b) }()

	fnErr = fn()

	wOut.Close()
	wErr.Close()
	os.Stdout, os.Stderr = origOut, origErr
	return <-outCh, <-errCh, fnErr
}
