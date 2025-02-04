// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE-GO file.

//go:build freebsd && 386

package memory

import (
	"syscall"
)

// https://cs.opensource.google/go/go/+/refs/tags/go1.17.8:src/syscall/zsyscall_freebsd_386.go
func mmapSyscall(addr uintptr, length uintptr, prot int, flag int, fd int, pos int64) (ret uintptr, err error) {
	r0, _, e1 := syscall.Syscall9(syscall.SYS_MMAP, uintptr(addr), uintptr(length), uintptr(prot), uintptr(flag), uintptr(fd), uintptr(pos), uintptr(pos>>32), 0, 0)
	ret = uintptr(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
