// Copyright 2024-2026 Qualcomm Technologies, Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-3-Clause

package readline

import "golang.org/x/sys/windows"

type Termios uint32

func getTermios() (*Termios, error) {
	termios := new(Termios)
	return termios, windows.GetConsoleMode(windows.Stdin, (*uint32)(termios))
}

func setTermios(termios *Termios) error {
	return windows.SetConsoleMode(windows.Stdin, uint32(*termios))
}

func applyRawMode(termios *Termios) {
	*termios &^= windows.ENABLE_ECHO_INPUT | windows.ENABLE_LINE_INPUT | windows.ENABLE_PROCESSED_INPUT
	*termios |= windows.ENABLE_VIRTUAL_TERMINAL_INPUT
}
