package handler

import (
	"net/http"
	"runtime/debug"

	"clean_arch/domain/model"
	"clean_arch/endpoint/api"
	"clean_arch/endpoint/api/respond"
	"clean_arch/registry"
)

// NewErrorHandler -
func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{
		rsp: respond.NewRespond(registry.Cfg.Serializer.Code),
	}
}

// ErrorHandler -
type ErrorHandler struct {
	rsp api.Respond
}

// RouteNotFound -
func (e *ErrorHandler) RouteNotFound(w http.ResponseWriter, h *http.Request) {
	e.rsp.Error(w, model.ErrRouteNotFound)
}

// MethodNotAllowed -
func (e *ErrorHandler) MethodNotAllowed(w http.ResponseWriter, h *http.Request) {
	e.rsp.Error(w, model.ErrMethodNotAllowed)
}

// Recoverer -
func (e *ErrorHandler) Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {

				logEntry := registry.Log
				if logEntry != nil {
					logEntry.Error(rvr, debug.Stack())
				}
				//				fmt.Fprintf(os.Stderr, "Panic: %+v\n", rvr)
				//				debug.PrintStack()

				e.rsp.Error(w, model.ErrInternalServerError)
				// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
