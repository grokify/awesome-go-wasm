# Awesome Go WASM

## Contents

- [Awesome Go WASM](#awesome-go-wasm)
  - [Contents](#contents)
  - [Examples](#examples)
  - [Go Runtimes](#go-runtimes)

## Go Examples

_WebAssembly examples in Go_

- [Go WebAssembly Tutorial - Building a Calculator Tutorial](https://tutorialedge.net/golang/go-webassembly-tutorial/)

**[⬆ back to top](#contents)**

## Go Runtimes

_Go Runtimes for WASM runtimes_

- [WasmEdge WebAssembly runtime for Go - github.com/second-state/WasmEdge-go](https://github.com/second-state/WasmEdge-go) - The [WasmEdge](https://github.com/WasmEdge/WasmEdge) is a high performance WebAssembly runtime optimized for server side applications. This project provides a [golang](https://go.dev/) package for accessing to WasmEdge.
- [Wasmer WebAssembly runtime for Go](https://github.com/wasmerio/wasmer-go) - A complete and mature WebAssembly runtime for Go based on [Wasmer](https://github.com/wasmerio/wasmer) - The leading WebAssembly Runtime supporting WASI and Emscripten.
- [Wasmtime WebAssembly runtime for Go - github.com/bytecodealliance/wasmtime-go](https://github.com/bytecodealliance/wasmtime-go) - This Go library uses CGO to consume the C API of the [Wasmtime](https://github.com/bytecodealliance/wasmtime) project which is written in Rust. Precompiled binaries of Wasmtime are checked into this repository on tagged releases so you won't have to install Wasmtime locally, but it means that this project only works on Linux x86_64, macOS x86_64 , and Windows x86_64 currently. Building on other platforms will need to arrange to build Wasmtime and use CGO_* env vars to compile correctly. Wasmtime is a standalone JIT-style runtime for WebAssembly, using Cranelift.

**[⬆ back to top](#contents)**
