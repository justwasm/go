// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !wasm

package ssa

func init() {
	registerArch("s390x", func(c *Config) {
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = rewriteBlockS390X
		c.lowerValue = rewriteValueS390X
		c.registers = registersS390X[:]
		c.gpRegMask = gpRegMaskS390X
		c.fpRegMask = fpRegMaskS390X
		c.intParamRegs = paramIntRegS390X
		c.floatParamRegs = paramFloatRegS390X
		c.FPReg = framepointerRegS390X
		c.LinkReg = linkRegS390X
		c.hasGReg = true
		c.BigEndian = true
		c.unalignedOK = true
		c.haveBswap64 = true
		c.haveBswap32 = true
		c.haveBswap16 = true // only for loads&stores, see ppc64 comment
	})
}
