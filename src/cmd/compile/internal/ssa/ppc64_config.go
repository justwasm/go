// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !wasm

package ssa

func init() {
	setup := func(c *Config) {
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = rewriteBlockPPC64
		c.lowerValue = rewriteValuePPC64
		c.lateLowerBlock = rewriteBlockPPC64latelower
		c.lateLowerValue = rewriteValuePPC64latelower
		c.registers = registersPPC64[:]
		c.gpRegMask = gpRegMaskPPC64
		c.fpRegMask = fpRegMaskPPC64
		c.specialRegMask = specialRegMaskPPC64
		c.intParamRegs = paramIntRegPPC64
		c.floatParamRegs = paramFloatRegPPC64
		c.FPReg = framepointerRegPPC64
		c.LinkReg = linkRegPPC64
		c.hasGReg = true
		c.unalignedOK = true
		// Note: ppc64 has register bswap ops only when GOPPC64>=10.
		// But it has bswap+load and bswap+store ops for all ppc64 variants.
		// That is the sense we're using them here - they are only used
		// in contexts where they can be merged with a load or store.
		c.haveBswap64 = true
		c.haveBswap32 = true
		c.haveBswap16 = true
		c.haveCondSelect = true
	}
	registerArch("ppc64", func(c *Config) {
		c.BigEndian = true
		setup(c)
	})
	registerArch("ppc64le", setup)
}
