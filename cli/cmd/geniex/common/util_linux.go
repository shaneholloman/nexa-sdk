// Copyright 2024-2026 Qualcomm Technologies, Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-3-Clause

package common

import (
	"golang.org/x/sys/unix"
)

func GetTerminalWidth() int {
	fd := int(unix.Stdout)
	ws, err := unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
	if err != nil {
		return 80
	}
	return int(ws.Col)
}
