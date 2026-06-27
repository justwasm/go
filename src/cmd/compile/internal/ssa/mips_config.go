// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !wasm

package ssa

func init() {
	setup := func(c *Config) {
		c.PtrSize = 4
		c.RegSize = 4
		c.lowerBlock = rewriteBlockMIPS
		c.lowerValue = rewriteValueMIPS
		c.registers = registersMIPS[:]
		c.gpRegMask = gpRegMaskMIPS
		c.fpRegMask = fpRegMaskMIPS
		c.specialRegMask = specialRegMaskMIPS
		c.FPReg = framepointerRegMIPS
		c.LinkReg = linkRegMIPS
		c.hasGReg = true
	}
	registerArch("mips", func(c *Config) {
		c.BigEndian = true
		setup(c)
	})
	registerArch("mipsle", setup)
}
