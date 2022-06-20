package wasmlinks

type Resource struct {
	URL        string
	Title      string
	Excerpt    string
	AccessDate string // RFC-3339 `full-date`
	Tags       []string
}

func Resources() []Resource {
	return []Resource{
		{
			Tags:       []string{"Go Runtimes"},
			URL:        "https://github.com/second-state/WasmEdge-go",
			Title:      "WasmEdge WebAssembly runtime for Go",
			AccessDate: "2022-06-20",
			Excerpt: `The [WasmEdge](https://github.com/WasmEdge/WasmEdge) is a high performance WebAssembly runtime optimized for server side applications. This project provides a [golang](https://go.dev/) package for accessing to WasmEdge.`,
		},
		{
			Tags:       []string{"Go Runtimes"},
			URL:        "https://github.com/wasmerio/wasmer-go",
			Title:      "Wasmer WebAssembly runtime for Go",
			AccessDate: "2022-06-20",
			Excerpt: `A complete and mature WebAssembly runtime for Go based on [Wasmer](https://github.com/wasmerio/wasmer) - The leading WebAssembly Runtime supporting WASI and Emscripten.`,
		},
		{
			Tags:       []string{"Go Runtimes"},
			URL:        "github.com/bytecodealliance/wasmtime-go",
			Title:      "Wasmtime WebAssembly runtime for Go",
			AccessDate: "2022-06-20",
			Excerpt: `This Go library uses CGO to consume the C API of the [Wasmtime](https://github.com/bytecodealliance/wasmtime) project which is written in Rust. Precompiled binaries of Wasmtime are checked into this repository on tagged releases so you won't have to install Wasmtime locally, but it means that this project only works on Linux x86_64, macOS x86_64 , and Windows x86_64 currently. Building on other platforms will need to arrange to build Wasmtime and use CGO_* env vars to compile correctly. Wasmtime is a standalone JIT-style runtime for WebAssembly, using Cranelift.`,
		},
	}
}
