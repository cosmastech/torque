package torque

import "net/http"

// Action is executed during an http POST request. Actions perform
// data mutations such as creating or updating resources and are
// usually triggered by a form submission in the browser.
type Action interface {
	Action(wr http.ResponseWriter, req *http.Request) error
}

// Loader is executed during an http GET request and provides
// data to the Renderer
// It can parse URL values, attach session data, etc.
type Loader interface {
	Load(req *http.Request) (any, error)
}

// Renderer is a response to an http GET that renders a template
type Renderer interface {
	Render(wr http.ResponseWriter, req *http.Request, loaderData any) error
}

// ErrorBoundary handles all errors returned by read and write operations in a .
type ErrorBoundary interface {
	ErrorBoundary(wr http.ResponseWriter, req *http.Request, err error) http.HandlerFunc
}

// PanicBoundary is a panic recovery handler. It catches any unhandled errors.
//
// If a handler is not returned to redirect the request, a stack trace is printed
// to the server logs.
type PanicBoundary interface {
	PanicBoundary(wr http.ResponseWriter, req *http.Request, err error) http.HandlerFunc
}

// SubmoduleProvider is executed when the torque app is initialized. It can
// return a list of RouterConfigs to be registered as children of the current
// RouteModule. The parent RouteModule's path will be prefixed to any provided
// paths in the RouterConfig.
type SubmoduleProvider interface {
	Submodules() []Route
}
