# Awesome Go Wasm [![Awesome](https://awesome.re/badge.svg)](https://awesome.re)

## Contents

- [Awesome Go WASM](#awesome-go-wasm)
  - [Contents](#contents)
  - [Go Guest Examples](#go-guests)
  - [TinyGo Guest Examples](#tinygo-guests)
  - [Go Libraries](#go-libraries)
  - [Go Runtimes](#go-runtimes)
  - [waPC - WebAssembly Procedure Calls](wapc-webAssembly-procedure-calls)

## Go Guest Examples

_WebAssembly examples in Go_

- [Understanding WebAssembly and interoperability with Go and Javascript](https://github.com/gptankit/go-wasm) - This is a gentle introduction to the world of WebAssembly from the lens of Go and Javascript. The purpose of this primer is to introduce WebAssembly to people who are already familiar with Go and want to use their understanding to build fast programs for the web and other environments outside the world of Go. If you are already substantially familiar with WebAssembly, you may want to directly jump to the experiments.
- [go-wasm-examples](https://github.com/danieljoos/go-wasm-examples) — Some small examples of using Go and WebAssembly
- [Go WebAssembly Tutorial - Building a Calculator Tutorial](https://tutorialedge.net/golang/go-webassembly-tutorial/)

**[⬆ back to top](#contents)**

## TinyGo Guest Examples

- [Official TinyGo Browser Example](https://tinygo.org/docs/guides/webassembly/) - How to call WebAssembly from JavaScript in a browser. You can call a JavaScript function from Go and call a Go function from WebAssembly.
- [WASI Hello World Example in Go from Wasm by Example](https://wasmbyexample.dev/examples/wasi-hello-world/wasi-hello-world.go.en-us.html) - In this example, We will be writing "Hello world!" to both the console (`stdout`), and a newly created file `helloworld.txt`. We highly reccomended that you have read the [WASI Introduction](https://wasmbyexample.dev/examples/wasi-introduction/wasi-introduction.all.en-us.html) before procedding with this example.

**[⬆ back to top](#contents)**

## Go Libraries

- [Go-app](https://github.com/maxence-charriere/go-app) [![](https://pkg.go.dev/badge/github.com/maxence-charriere/go-app)](https://pkg.go.dev/github.com/maxence-charriere/go-app/pkg/app) — Go-app is a package for building progressive web apps (PWA) with the Go programming language (Golang) and WebAssembly (Wasm).
- [Wat2Wasm function](https://github.com/bytecodealliance/wasmtime-go/blob/main/wat2wasm.go) [![](https://pkg.go.dev/badge/github.com/bytecodealliance/wasmtime-go)](https://pkg.go.dev/github.com/bytecodealliance/wasmtime-go#Wat2Wasm) - function from Wasmtime to convert Wat/Witx text format of WebAssembly to the binary format.
- [Vugu](https://github.com/vugu/vugu) [![](https://pkg.go.dev/badge/github.com/vugu/vugu)](https://pkg.go.dev/github.com/vugu/vugu) — Vugu is an experimental library for web UIs written in Go and targeting webassembly. Guide and docs at https://www.vugu.org. Godoc at https://godoc.org/github.com/vugu/vugu.
- [WASM-FETCH](https://github.com/marwan-at-work/wasm-fetch) [![](https://pkg.go.dev/badge/github.com/marwan-at-work/wasm-fetch)](https://pkg.go.dev/marwan.io/wasm-fetch) — A go-wasm library that wraps the Fetch API. This is useful since TinyGo does not support `net/http`.

**[⬆ back to top](#contents)**

## Go Runtimes

_Go Runtimes for WASM runtimes_

- [WasmEdge WebAssembly runtime for Go](https://github.com/second-state/WasmEdge-go) - The [WasmEdge](https://github.com/WasmEdge/WasmEdge) is a high performance WebAssembly runtime optimized for server side applications. This project provides a [golang](https://go.dev/) package for accessing to WasmEdge.
- [Wasmer WebAssembly runtime for Go](https://github.com/wasmerio/wasmer-go) - A complete and mature WebAssembly runtime for Go based on [Wasmer](https://github.com/wasmerio/wasmer) - The leading WebAssembly Runtime supporting WASI and Emscripten.
- [Wasmtime WebAssembly runtime for Go](https://github.com/bytecodealliance/wasmtime-go) [![](https://pkg.go.dev/badge/github.com/bytecodealliance/wasmtime-go)](https://pkg.go.dev/github.com/bytecodealliance/wasmtime-go) — This Go library uses CGO to consume the C API of the [Wasmtime](https://pkg.go.dev/github.com/bytecodealliance/wasmtime) project which is written in Rust. Precompiled binaries of Wasmtime are checked into this repository on tagged releases so you won't have to install Wasmtime locally, but it means that this project only works on Linux x86_64, macOS x86_64 , and Windows x86_64 currently. Building on other platforms will need to arrange to build Wasmtime and use `CGO_*` env vars to compile correctly. Wasmtime is a standalone JIT-style runtime for WebAssembly, using Cranelift.
- [Wasmtime host example](https://docs.wasmtime.dev/lang-go.html) - Example of calling a GCD function (provided as witx). Func `(func $gcd (param i32 i32) (result i32) ...)` is called as `val, err := gcd.Call(store, 6, 27)`.
- [wazero WebAssembly runtime for Go with zero dependencies](https://github.com/tetratelabs/wazero) [![](https://pkg.go.dev/badge/github.com/tetratelabs/wazero)](https://pkg.go.dev/github.com/tetratelabs/wazero) — wazero is a WebAssembly Core Specification [1.0](https://www.w3.org/TR/2019/REC-wasm-core-1-20191205/) and [2.0](https://www.w3.org/TR/2022/WD-wasm-core-2-20220419/) compliant runtime written in Go. It has *zero dependencies*, and doesn't rely on CGO. This means you can run applications in other languages and still keep cross compilation.
- [fastlike](https://github.com/avidal/fastlike) [![](https://pkg.go.dev/badge/github.com/avidal/fastlike)](https://pkg.go.dev/fastlike.dev) — fastlike is a Go project that implements the Fastly Compute@Edge ABI using `wasmtime` and exposes a `http.Handler` for you to use.

**[⬆ back to top](#contents)**

## waPC - WebAssembly Procedure Calls

- [waPC Host for Go](https://github.com/wapc/wapc-go) [![](https://pkg.go.dev/badge/github.com/wapc/wapc-go)](https://pkg.go.dev/github.com/wapc/wapc-go) [![Gitter](https://badges.gitter.im/wapc/community.svg)](https://gitter.im/wapc/community) — Golang-based WebAssembly Host Runtime for waPC-compliant modules
- [waPC Guest Library for TinyGo](https://github.com/wapc/wapc-guest-tinygo) [![](https://pkg.go.dev/badge/github.com/wapc/wapc-guest-tinygo)](https://pkg.go.dev/github.com/wapc/wapc-guest-tinygo) [![Gitter](https://badges.gitter.im/wapc/community.svg)](https://gitter.im/wapc/community) — SDK for creating waPC WebAssembly Guest Modules in TinyGo

**[⬆ back to top](#contents)**
