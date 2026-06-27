// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// archSetup sets arch-specific fields on a Config.
type archSetup func(c *Config)

var archConfigs map[string]archSetup

func registerArch(arch string, fn archSetup) {
	if archConfigs == nil {
		archConfigs = make(map[string]archSetup)
	}
	archConfigs[arch] = fn
}
