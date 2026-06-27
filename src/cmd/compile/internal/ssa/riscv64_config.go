// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !wasm

package ssa

func init() {
	registerArch("riscv64", func(c *Config) {
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = rewriteBlockRISCV64
		c.lowerValue = rewriteValueRISCV64
		c.lateLowerBlock = rewriteBlockRISCV64latelower
		c.lateLowerValue = rewriteValueRISCV64latelower
		c.registers = registersRISCV64[:]
		c.gpRegMask = gpRegMaskRISCV64
		c.fpRegMask = fpRegMaskRISCV64
		c.intParamRegs = paramIntRegRISCV64
		c.floatParamRegs = paramFloatRegRISCV64
		c.FPReg = framepointerRegRISCV64
		c.hasGReg = true
	})
}
