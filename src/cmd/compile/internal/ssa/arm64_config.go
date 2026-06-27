// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !wasm

package ssa

func init() {
	registerArch("arm64", func(c *Config) {
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = rewriteBlockARM64
		c.lowerValue = rewriteValueARM64
		c.lateLowerBlock = rewriteBlockARM64latelower
		c.lateLowerValue = rewriteValueARM64latelower
		c.registers = registersARM64[:]
		c.gpRegMask = gpRegMaskARM64
		c.fpRegMask = fpRegMaskARM64
		c.simdRegMask = simdRegMaskARM64
		c.intParamRegs = paramIntRegARM64
		c.floatParamRegs = paramFloatRegARM64
		c.FPReg = framepointerRegARM64
		c.LinkReg = linkRegARM64
		c.hasGReg = true
		c.unalignedOK = true
		c.haveBswap64 = true
		c.haveBswap32 = true
		c.haveBswap16 = true
		c.haveCondSelect = true
	})
}
