// Copyright 2024-2026 Qualcomm Technologies, Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-3-Clause

package common

import (
	"log/slog"

	"golang.org/x/sys/windows"
)

func EnableUTF8Console() {
	if err := windows.SetConsoleOutputCP(65001); err != nil {
		slog.Debug("SetConsoleOutputCP failed", "err", err)
	}
	if err := windows.SetConsoleCP(65001); err != nil {
		slog.Debug("SetConsoleCP failed", "err", err)
	}
}
