// Copyright 2022 Robert S. Muhlestein
  // SPDX-License-Identifier: Apache-2.0
  
//go:build !aix && !js && !nacl && !plan9 && !windows && !android && !solaris

package term

import (
	"syscall"
	"unsafe"
)

func WinSizeUpdate() {
	syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(0), uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&WinSize)))
}
