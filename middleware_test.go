package http_middleware

import (
	"fmt"
	"net/http"
	"testing"
)

func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello World!")
}

func MiddlewareA(writer http.ResponseWriter, request *http.Request, handleFunc MiddlewareHandlerFunc) {
	fmt.Fprintln(writer, "StartMiddleware A")
	handleFunc(writer, request)
	fmt.Fprintln(writer, "EndMiddleware A")
}

func MiddlewareB(writer http.ResponseWriter, request *http.Request, handleFunc MiddlewareHandlerFunc) {
	fmt.Fprintln(writer, "StartMiddleware B")
	handleFunc(writer, request)
	fmt.Fprintln(writer, "EndMiddleware B")
}

func TestMiddlewareChain_Handler(t *testing.T) {
	middlewareChain := MiddlewareChain{MiddlewareA, MiddlewareB}
	helloWorld := middlewareChain.Handler(HelloWorld)

	http.ListenAndServe("127.0.0.1:8088", helloWorld)
}
