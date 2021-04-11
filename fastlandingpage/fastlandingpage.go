package main

import (
	"github.com/valyala/fasthttp"
	"log"
)

type MyHandler struct {
	foobar string
}

func main() {
	fs := &fasthttp.FS{
		// Path to directory to serve.
		Root: "./",

		// Generate index pages if client requests directory contents.
		GenerateIndexPages: false,

		IndexNames: []string{"index.html", "index.htm"},

		// Enable transparent compression to save network traffic.
		Compress: true,
	}

	// Create request handler for serving static files.
	h := fs.NewRequestHandler()

	// Start the server.
	if err := fasthttp.ListenAndServe(":8080", h); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}



	// request handler in net/http style, i.e. method bound to MyHandler struct.


	// request handler in fasthttp style, i.e. just plain function.

	// pass bound struct method to fasthttp
	//myHandler := &MyHandler{
	//	foobar: "foobar",
	//}
	//// pass plain function to fasthttp
	//fasthttp.ListenAndServe(":8081", fastHTTPHandler)
	//
	//fasthttp.ListenAndServe(":8080", myHandler.HandleFastHTTP)
}

//func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
//	// notice that we may access MyHandler properties here - see h.foobar.
//	fmt.Fprintf(ctx, "Hello, world! Requested path is %q. Foobar is %q",
//		ctx.Path(), h.foobar)
//}
//
//func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
//	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
//}