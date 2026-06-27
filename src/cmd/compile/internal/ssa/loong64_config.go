// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !wasm

package ssa

func init() {
	registerArch("loong64", func(c *Config) {
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = rewriteBlockLOONG64
		c.lowerValue = rewriteValueLOONG64
		c.lateLowerBlock = rewriteBlockLOONG64latelower
		c.lateLowerValue = rewriteValueLOONG64latelower
		c.registers = registersLOONG64[:]
		c.gpRegMask = gpRegMaskLOONG64
		c.fpRegMask = fpRegMaskLOONG64
		c.intParamRegs = paramIntRegLOONG64
		c.floatParamRegs = paramFloatRegLOONG64
		c.FPReg = framepointerRegLOONG64
		c.LinkReg = linkRegLOONG64
		c.hasGReg = true
		c.unalignedOK = true
		c.haveBswap64 = true
		c.haveBswap32 = true
		c.haveBswap16 = true
		c.haveCondSelect = true
	})
}
