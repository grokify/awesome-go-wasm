# Awesome Go Wasm [![Awesome](https://awesome.re/badge.svg)](https://awesome.re)

## Contents

- [Awesome Go WASM](#awesome-go-wasm)
  - [Contents](#contents)
  - [Go Guest Examples](#go-guest-examples)
  - [TinyGo Guest Examples](#tinygo-guest-examples)
  - [Go Runtimes](#go-runtimes)
  - [Go Runtimes for Fastly](#go-runtimes-for-fastly)

## Go Guest Examples

_WebAssembly examples in Go_

- [Go WebAssembly Tutorial - Building a Calculator Tutorial](https://tutorialedge.net/golang/go-webassembly-tutorial/)

**[⬆ back to top](#contents)**

## TinyGo Guest Examples

- [Official TinyGo Browser Example](https://tinygo.org/docs/guides/webassembly/) - How to call WebAssembly from JavaScript in a browser. You can call a JavaScript function from Go and call a Go function from WebAssembly.
- [WASI Hello World Example in Go from Wasm by Example](https://wasmbyexample.dev/examples/wasi-hello-world/wasi-hello-world.go.en-us.html) - In this example, We will be writing "Hello world!" to both the console (`stdout`), and a newly created file `helloworld.txt`. We highly reccomended that you have read the [WASI Introduction](https://wasmbyexample.dev/examples/wasi-introduction/wasi-introduction.all.en-us.html) before procedding with this example.

**[⬆ back to top](#contents)**

## Go Runtimes

_Go Runtimes for WASM runtimes_

- [WasmEdge WebAssembly runtime for Go](https://github.com/second-state/WasmEdge-go) - The [WasmEdge](https://github.com/WasmEdge/WasmEdge) is a high performance WebAssembly runtime optimized for server side applications. This project provides a [golang](https://go.dev/) package for accessing to WasmEdge.
- [Wasmer WebAssembly runtime for Go](https://github.com/wasmerio/wasmer-go) - A complete and mature WebAssembly runtime for Go based on [Wasmer](https://github.com/wasmerio/wasmer) - The leading WebAssembly Runtime supporting WASI and Emscripten.
- [Wasmtime WebAssembly runtime for Go](https://github.com/bytecodealliance/wasmtime-go) [![](https://pkg.go.dev/badge/github.com/bytecodealliance/wasmtime-go)](https://pkg.go.dev/github.com/bytecodealliance/wasmtime-g) — This Go library uses CGO to consume the C API of the [Wasmtime](https://pkg.go.dev/github.com/bytecodealliance/wasmtime) project which is written in Rust. Precompiled binaries of Wasmtime are checked into this repository on tagged releases so you won't have to install Wasmtime locally, but it means that this project only works on Linux x86_64, macOS x86_64 , and Windows x86_64 currently. Building on other platforms will need to arrange to build Wasmtime and use CGO_* env vars to compile correctly. Wasmtime is a standalone JIT-style runtime for WebAssembly, using Cranelift.
- [Wasmtime host example](https://docs.wasmtime.dev/lang-go.html) - Example of calling a GCD function (provided as witx). Func `(func $gcd (param i32 i32) (result i32) ...)` is called as `val, err := gcd.Call(store, 6, 27)`.

**[⬆ back to top](#contents)**

## waPC - WebAssembly Procedure Calls

- [waPC Host for Go](https://github.com/wapc/wapc-go) [![](https://pkg.go.dev/badge/github.com/wapc/wapc-go)](https://pkg.go.dev/github.com/wapc/wapc-go) [![Gitter](https://badges.gitter.im/wapc/community.svg)](https://gitter.im/wapc/community) — Golang-based WebAssembly Host Runtime for waPC-compliant modules
- [waPC Guest Library for TinyGo](https://github.com/wapc/wapc-guest-tinygo) [![](https://pkg.go.dev/badge/github.com/wapc/wapc-guest-tinygo)](https://pkg.go.dev/github.com/wapc/wapc-guest-tinygo) [![Gitter](https://badges.gitter.im/wapc/community.svg)](https://gitter.im/wapc/community) — SDK for creating waPC WebAssembly Guest Modules in TinyGo

**[⬆ back to top](#contents)**

## Go Runtimes for Fastly

- [fastlike](https://github.com/avidal/fastlike) — fastlike is a Go project that implements the Fastly Compute@Edge ABI using `wasmtime` and exposes a `http.Handler` for you to use.

**[⬆ back to top](#contents)**


