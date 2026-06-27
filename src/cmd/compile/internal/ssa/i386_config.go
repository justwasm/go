// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !wasm

package ssa

func init() {
	registerArch("386", func(c *Config) {
		c.PtrSize = 4
		c.RegSize = 4
		c.lowerBlock = rewriteBlock386
		c.lowerValue = rewriteValue386
		c.splitLoad = rewriteValue386splitload
		c.registers = registers386[:]
		c.gpRegMask = gpRegMask386
		c.fpRegMask = fpRegMask386
		c.FPReg = framepointerReg386
		c.LinkReg = linkReg386
		c.hasGReg = false
		c.unalignedOK = true
		c.haveBswap32 = true
		c.haveBswap16 = true
	})
}
