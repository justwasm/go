// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !wasm

package ssa

func init() {
	registerArch("amd64", func(c *Config) {
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = rewriteBlockAMD64
		c.lowerValue = rewriteValueAMD64
		c.lateLowerBlock = rewriteBlockAMD64latelower
		c.lateLowerValue = rewriteValueAMD64latelower
		c.splitLoad = rewriteValueAMD64splitload
		c.registers = registersAMD64[:]
		c.gpRegMask = gpRegMaskAMD64
		c.fpRegMask = fpRegMaskAMD64
		c.simdRegMask = simdRegMaskAMD64
		c.specialRegMask = specialRegMaskAMD64
		c.intParamRegs = paramIntRegAMD64
		c.floatParamRegs = paramFloatRegAMD64
		c.FPReg = framepointerRegAMD64
		c.LinkReg = linkRegAMD64
		c.hasGReg = true
		c.unalignedOK = true
		c.haveBswap64 = true
		c.haveBswap32 = true
		c.haveBswap16 = true
		c.haveCondSelect = true
	})
}
