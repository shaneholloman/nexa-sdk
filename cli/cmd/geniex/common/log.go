// Copyright 2024-2026 Qualcomm Technologies, Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-3-Clause

package common

import (
	"io"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"

	geniex_sdk "github.com/qualcomm/GenieX/bindings/go"
	"github.com/qualcomm/GenieX/cli/internal/config"
)

const (
	LogLevelNone  string = "none"
	LogLevelTrace string = "trace"
	LogLevelDebug string = "debug"
	LogLevelInfo  string = "info"
	LogLevelWarn  string = "warn"
	LogLevelError string = "error"
)

func ApplyLogLevel() {
	options := tint.Options{AddSource: true}

	if os.Getenv("NO_COLOR") == "1" {
		options.NoColor = true
	}

	level := config.Get().Log

	switch level {
	case LogLevelNone:
		geniex_sdk.SetLog(false)
		slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))
		return
	case LogLevelTrace:
		options.Level = slog.LevelDebug
		slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &options)))
		return
	case LogLevelDebug:
		options.Level = slog.LevelDebug
	case LogLevelInfo:
		options.Level = slog.LevelInfo
	case LogLevelWarn:
		options.Level = slog.LevelWarn
	case LogLevelError:
		options.Level = slog.LevelError
	}

	geniex_sdk.SetLog(true)
	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &options)))
}
