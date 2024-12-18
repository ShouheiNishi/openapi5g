// Package bootstrapping provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.1-0.20240123090344-d326c01d279a DO NOT EDIT.
package bootstrapping

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	externalRef0 "github.com/ShouheiNishi/openapi5g/models"
	"github.com/gin-gonic/gin"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// BootstrappingInfoRequest request
	BootstrappingInfoRequest(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) BootstrappingInfoRequest(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBootstrappingInfoRequestRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewBootstrappingInfoRequestRequest generates requests for BootstrappingInfoRequest
func NewBootstrappingInfoRequestRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/bootstrapping")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// BootstrappingInfoRequestWithResponse request
	BootstrappingInfoRequestWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*BootstrappingInfoRequestResponse, error)
}

type BootstrappingInfoRequestResponse struct {
	Body                          []byte
	HTTPResponse                  *http.Response
	Application3gppHalJSON200     *externalRef0.BootstrappingInfo
	JSON307                       *externalRef0.RedirectResponse
	JSON308                       *externalRef0.RedirectResponse
	ApplicationproblemJSON400     *externalRef0.N400
	ApplicationproblemJSON500     *externalRef0.N500
	ApplicationproblemJSONDefault *externalRef0.Default
}

// Status returns HTTPResponse.Status
func (r BootstrappingInfoRequestResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r BootstrappingInfoRequestResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// BootstrappingInfoRequestWithResponse request returning *BootstrappingInfoRequestResponse
func (c *ClientWithResponses) BootstrappingInfoRequestWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*BootstrappingInfoRequestResponse, error) {
	rsp, err := c.BootstrappingInfoRequest(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBootstrappingInfoRequestResponse(rsp)
}

// ParseBootstrappingInfoRequestResponse parses an HTTP response from a BootstrappingInfoRequestWithResponse call
func ParseBootstrappingInfoRequestResponse(rsp *http.Response) (*BootstrappingInfoRequestResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &BootstrappingInfoRequestResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest externalRef0.BootstrappingInfo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.Application3gppHalJSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 307:
		var dest externalRef0.RedirectResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON307 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 308:
		var dest externalRef0.RedirectResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON308 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef0.N400
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.N500
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest externalRef0.Default
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSONDefault = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Bootstrapping Info Request
	// (GET /bootstrapping)
	BootstrappingInfoRequest(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// BootstrappingInfoRequest operation middleware
func (siw *ServerInterfaceWrapper) BootstrappingInfoRequest(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.BootstrappingInfoRequest(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/bootstrapping", wrapper.BootstrappingInfoRequest)
}

type BootstrappingInfoRequestRequestObject struct {
}

type BootstrappingInfoRequestResponseObject interface {
	VisitBootstrappingInfoRequestResponse(w http.ResponseWriter) error
}

type BootstrappingInfoRequest200Application3gppHalPlusJSONResponse externalRef0.BootstrappingInfo

func (response BootstrappingInfoRequest200Application3gppHalPlusJSONResponse) VisitBootstrappingInfoRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/3gppHal+json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(externalRef0.BootstrappingInfo(response))
}

type BootstrappingInfoRequest307ResponseHeaders struct {
	Location string
}

type BootstrappingInfoRequest307JSONResponse struct {
	Body    externalRef0.RedirectResponse
	Headers BootstrappingInfoRequest307ResponseHeaders
}

func (response BootstrappingInfoRequest307JSONResponse) VisitBootstrappingInfoRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", fmt.Sprint(response.Headers.Location))
	w.WriteHeader(307)

	return json.NewEncoder(w).Encode(response.Body)
}

type BootstrappingInfoRequest308ResponseHeaders struct {
	Location string
}

type BootstrappingInfoRequest308JSONResponse struct {
	Body    externalRef0.RedirectResponse
	Headers BootstrappingInfoRequest308ResponseHeaders
}

func (response BootstrappingInfoRequest308JSONResponse) VisitBootstrappingInfoRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", fmt.Sprint(response.Headers.Location))
	w.WriteHeader(308)

	return json.NewEncoder(w).Encode(response.Body)
}

type BootstrappingInfoRequest400ApplicationProblemPlusJSONResponse struct {
	externalRef0.N400ApplicationProblemPlusJSONResponse
}

func (response BootstrappingInfoRequest400ApplicationProblemPlusJSONResponse) VisitBootstrappingInfoRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N400ApplicationProblemPlusJSONResponse))
}

type BootstrappingInfoRequest500ApplicationProblemPlusJSONResponse struct {
	externalRef0.N500ApplicationProblemPlusJSONResponse
}

func (response BootstrappingInfoRequest500ApplicationProblemPlusJSONResponse) VisitBootstrappingInfoRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N500ApplicationProblemPlusJSONResponse))
}

type BootstrappingInfoRequestdefaultApplicationProblemPlusJSONResponse struct {
	Body       externalRef0.ProblemDetails
	StatusCode int
}

func (response BootstrappingInfoRequestdefaultApplicationProblemPlusJSONResponse) VisitBootstrappingInfoRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Bootstrapping Info Request
	// (GET /bootstrapping)
	BootstrappingInfoRequest(ctx context.Context, request BootstrappingInfoRequestRequestObject) (BootstrappingInfoRequestResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// BootstrappingInfoRequest operation middleware
func (sh *strictHandler) BootstrappingInfoRequest(ctx *gin.Context) {
	var request BootstrappingInfoRequestRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.BootstrappingInfoRequest(ctx, request.(BootstrappingInfoRequestRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "BootstrappingInfoRequest")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(BootstrappingInfoRequestResponseObject); ok {
		if err := validResponse.VisitBootstrappingInfoRequestResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
