// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !wasm

package ssa

func init() {
	setup := func(c *Config) {
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = rewriteBlockMIPS64
		c.lowerValue = rewriteValueMIPS64
		c.lateLowerBlock = rewriteBlockMIPS64latelower
		c.lateLowerValue = rewriteValueMIPS64latelower
		c.registers = registersMIPS64[:]
		c.gpRegMask = gpRegMaskMIPS64
		c.fpRegMask = fpRegMaskMIPS64
		c.specialRegMask = specialRegMaskMIPS64
		c.FPReg = framepointerRegMIPS64
		c.LinkReg = linkRegMIPS64
		c.hasGReg = true
	}
	registerArch("mips64", func(c *Config) {
		c.BigEndian = true
		setup(c)
	})
	registerArch("mips64le", setup)
}
