// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !wasm

package ssa

func init() {
	registerArch("arm", func(c *Config) {
		c.PtrSize = 4
		c.RegSize = 4
		c.lowerBlock = rewriteBlockARM
		c.lowerValue = rewriteValueARM
		c.registers = registersARM[:]
		c.gpRegMask = gpRegMaskARM
		c.fpRegMask = fpRegMaskARM
		c.FPReg = framepointerRegARM
		c.LinkReg = linkRegARM
		c.hasGReg = true
	})
}
