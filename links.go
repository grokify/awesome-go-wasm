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
			Tags:       []string{"host-call"},
			URL:        "https://programming.vip/docs/istio-service-grid-wasm-filter.html",
			Title:      "Istio service grid Wasm filter",
			AccessDate: "2022-06-14",
			Excerpt: `Host call API

			The host call API is a series of methods provided by proxy wasm to interact with network plug-ins. For example, GetHttpRequestHeaders API can be called in HttpContext to obtain Http request header data, and the LogInfo API can be used to add print information to the log.
			
			All available API s are available through [hostcall.go](https://github.com/tetratelabs/proxy-wasm-go-sdk/blob/v0.14.0/proxywasm/hostcall.go) View details.`,
		},
	}
}
