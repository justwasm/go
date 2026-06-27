# js/wasm Compiler Optimization: Arch-Specific Code Exclusion

## Principle

On js/wasm, the Go compiler only ever targets wasm. Every non-wasm architecture's
SSA rewrite rules, SIMD intrinsic tables, register definitions, and config setup
are dead weight — they consume memory in the wasm binary and during compilation
but are never executed.

The optimization: **exclude non-wasm arch-specific code at compile time** using
build tags, and **sever direct symbol references** from shared code so the linker
can dead-code eliminate the data as well.

## The Pattern

For each arch-specific file (or set of files) that should be excluded on wasm:

### 1. Add a build constraint

```go
//go:build !wasm
```

Place this at the top of every non-wasm-arch-specific source file. On wasm hosts
(GOARCH=wasm), these files are never compiled or linked. On all other hosts,
they compile as before.

### 2. Replace direct symbol references with a registry

If the excluded file exports symbols (functions, variables) that are referenced
from shared code (e.g., `config.go`'s `rewriteBlockAMD64`), those references
create linker roots that prevent dead-code elimination. Fix this with a registry:

**Infrastructure** (always compiled, e.g. `config_init.go`):
```go
type archSetup func(c *Config)
var archConfigs map[string]archSetup

func registerArch(arch string, fn archSetup) {
    if archConfigs == nil {
        archConfigs = make(map[string]archSetup)
    }
    archConfigs[arch] = fn
}
```

**Registration** (in the build-tagged file):
```go
//go:build !wasm
package ssa

func init() {
    registerArch("amd64", func(c *Config) {
        c.lowerBlock = rewriteBlockAMD64
        c.registers = registersAMD64[:]
        // ... all arch-specific config
    })
}
```

**Lookup** (in shared code):
```go
// Before (direct ref — linker keeps all arch symbols):
switch arch {
case "amd64":
    c.lowerBlock = rewriteBlockAMD64
case "wasm":
    c.lowerBlock = rewriteBlockWasm
}

// After (registry lookup — linker sees no ref to excluded symbols):
if fn := archConfigs[arch]; fn != nil {
    fn(c)
} else {
    ctxt.Diag("arch %s not implemented", arch)
}
```

### 3. File naming trap

**Never name files `*_GOARCH.go` or `*_GOOS.go`** — Go's build system
automatically applies build constraints to files matching these patterns.

Bad:  `config_amd64.go`  → implicit `//go:build amd64`
Good: `amd64_config.go`  → no implicit constraint

This is why the config registration files use the `ARCH_config.go` naming
convention instead of `config_ARCH.go`.

## What Was Excluded

| Layer | Files | Source size | Technique |
|-------|-------|-------------|-----------|
| SIMD intrinsics (ssagen) | `simdAMD64intrinsics.go`, `simdARM64intrinsics.go` | 266 KB | `//go:build` + function variable + nil check |
| SSA rewrite rules (ssa) | 18 rewrite*.go files | ~5.8 MB | `//go:build !wasm` + config registry + linker DCE |

## Testing

- Full `go test cmd/compile/internal/ssa` passes on non-wasm hosts
- `go test cmd/compile/internal/ssagen` passes
- The `generate_test.go` was updated to tolerate leading build constraint
  lines before generated-file headers when comparing file contents

## Checklist for Future Optimizations

When adding more arch-specific code that should be excluded on js/wasm:

1. Can a `//go:build !wasm` tag be added? (Is the file purely arch-specific?)
2. If it exports symbols referenced from shared code, create a registration
   pattern instead of direct references.
3. Avoid `*_GOARCH.go` / `*_GOOS.go` filename patterns.
4. Run the full test suite for the affected package.
5. Update any "generated file is up to date" tests to handle build tags.
