# CGO

CGO in Go is a feature that allows Go code to call C code directly. It enables interoperability between Go and C, which is useful when you need to use:

- existing C libraries,
- system-level APIs not exposed to Go,
- or performance-critical native code.

## CGO generates C bindings so that your Go code can:

- call C functions,
- use C types (C.int, C.char\*, etc.),
- and link against C libraries.

## ✅ When to Use CGO?

- Calling existing C libraries (like libxml, OpenSSL, libpng)
- Interfacing with system calls or OS-level APIs not in the Go standard library
- Performance tuning where low-level control is needed

## ⚠️ Downsides of CGO

- Slower build time (uses a C compiler under the hood)
- Less portable (not pure Go anymore)
- Cross-compilation is harder
- Potential for memory issues (manual memory management if mixing with malloc)
