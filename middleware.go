package http_middleware

import "net/http"

type MiddlewareHandlerFunc http.HandlerFunc

type MiddlewareInterceptor func(http.ResponseWriter, *http.Request, MiddlewareHandlerFunc)

func (cont MiddlewareHandlerFunc) Intercept(mw MiddlewareInterceptor) MiddlewareHandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		mw(writer, request, cont)
	}
}

type MiddlewareChain []MiddlewareInterceptor

func (chain MiddlewareChain) Handler(handler http.HandlerFunc) http.Handler {
	curr := MiddlewareHandlerFunc(handler)
	for i := len(chain) - 1; i >= 0; i-- {
		mw := chain[i]
		curr = curr.Intercept(mw)
	}

	return http.HandlerFunc(curr)
}
