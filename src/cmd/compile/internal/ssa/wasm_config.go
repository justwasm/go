// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

func init() {
	registerArch("wasm", func(c *Config) {
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = rewriteBlockWasm
		c.lowerValue = rewriteValueWasm
		c.registers = registersWasm[:]
		c.gpRegMask = gpRegMaskWasm
		c.fpRegMask = fpRegMaskWasm
		c.fp32RegMask = fp32RegMaskWasm
		c.fp64RegMask = fp64RegMaskWasm
		c.simdRegMask = simdRegMaskWasm
		c.FPReg = framepointerRegWasm
		c.LinkReg = linkRegWasm
		c.hasGReg = true
		c.unalignedOK = true
		c.haveCondSelect = true
	})
}
