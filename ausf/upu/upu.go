// Package upu provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.1-0.20240123090344-d326c01d279a DO NOT EDIT.
package upu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	externalRef0 "github.com/ShouheiNishi/openapi5g/models"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

const (
	OAuth2ClientCredentialsScopes = "oAuth2ClientCredentials.Scopes"
)

// PostSupiUeUpuJSONRequestBody defines body for PostSupiUeUpu for application/json ContentType.
type PostSupiUeUpuJSONRequestBody = externalRef0.AusfUpuInfo

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
	// PostSupiUeUpuWithBody request with any body
	PostSupiUeUpuWithBody(ctx context.Context, supi externalRef0.Supi, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostSupiUeUpu(ctx context.Context, supi externalRef0.Supi, body PostSupiUeUpuJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) PostSupiUeUpuWithBody(ctx context.Context, supi externalRef0.Supi, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostSupiUeUpuRequestWithBody(c.Server, supi, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostSupiUeUpu(ctx context.Context, supi externalRef0.Supi, body PostSupiUeUpuJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostSupiUeUpuRequest(c.Server, supi, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewPostSupiUeUpuRequest calls the generic PostSupiUeUpu builder with application/json body
func NewPostSupiUeUpuRequest(server string, supi externalRef0.Supi, body PostSupiUeUpuJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostSupiUeUpuRequestWithBody(server, supi, "application/json", bodyReader)
}

// NewPostSupiUeUpuRequestWithBody generates requests for PostSupiUeUpu with any type of body
func NewPostSupiUeUpuRequestWithBody(server string, supi externalRef0.Supi, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "supi", runtime.ParamLocationPath, supi)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/%s/ue-upu", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

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
	// PostSupiUeUpuWithBodyWithResponse request with any body
	PostSupiUeUpuWithBodyWithResponse(ctx context.Context, supi externalRef0.Supi, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostSupiUeUpuResponse, error)

	PostSupiUeUpuWithResponse(ctx context.Context, supi externalRef0.Supi, body PostSupiUeUpuJSONRequestBody, reqEditors ...RequestEditorFn) (*PostSupiUeUpuResponse, error)
}

type PostSupiUeUpuResponse struct {
	Body                          []byte
	HTTPResponse                  *http.Response
	JSON200                       *externalRef0.UpuSecurityInfo
	JSON307                       *externalRef0.N307
	JSON308                       *externalRef0.N307
	ApplicationproblemJSON503     *externalRef0.ProblemDetails
	ApplicationproblemJSONDefault *externalRef0.ProblemDetails
}

// Status returns HTTPResponse.Status
func (r PostSupiUeUpuResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostSupiUeUpuResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostSupiUeUpuWithBodyWithResponse request with arbitrary body returning *PostSupiUeUpuResponse
func (c *ClientWithResponses) PostSupiUeUpuWithBodyWithResponse(ctx context.Context, supi externalRef0.Supi, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostSupiUeUpuResponse, error) {
	rsp, err := c.PostSupiUeUpuWithBody(ctx, supi, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostSupiUeUpuResponse(rsp)
}

func (c *ClientWithResponses) PostSupiUeUpuWithResponse(ctx context.Context, supi externalRef0.Supi, body PostSupiUeUpuJSONRequestBody, reqEditors ...RequestEditorFn) (*PostSupiUeUpuResponse, error) {
	rsp, err := c.PostSupiUeUpu(ctx, supi, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostSupiUeUpuResponse(rsp)
}

// ParsePostSupiUeUpuResponse parses an HTTP response from a PostSupiUeUpuWithResponse call
func ParsePostSupiUeUpuResponse(rsp *http.Response) (*PostSupiUeUpuResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostSupiUeUpuResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest externalRef0.UpuSecurityInfo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 307:
		var dest externalRef0.N307
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON307 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 308:
		var dest externalRef0.N307
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON308 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest externalRef0.ProblemDetails
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON503 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest externalRef0.ProblemDetails
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSONDefault = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /{supi}/ue-upu)
	PostSupiUeUpu(c *gin.Context, supi externalRef0.Supi)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// PostSupiUeUpu operation middleware
func (siw *ServerInterfaceWrapper) PostSupiUeUpu(c *gin.Context) {

	var err error

	// ------------- Path parameter "supi" -------------
	var supi externalRef0.Supi

	err = runtime.BindStyledParameterWithOptions("simple", "supi", c.Param("supi"), &supi, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter supi: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(OAuth2ClientCredentialsScopes, []string{"nausf-upuprotection"})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostSupiUeUpu(c, supi)
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

	router.POST(options.BaseURL+"/:supi/ue-upu", wrapper.PostSupiUeUpu)
}

type PostSupiUeUpuRequestObject struct {
	Supi externalRef0.Supi `json:"supi"`
	Body *PostSupiUeUpuJSONRequestBody
}

type PostSupiUeUpuResponseObject interface {
	VisitPostSupiUeUpuResponse(w http.ResponseWriter) error
}

type PostSupiUeUpu200JSONResponse externalRef0.UpuSecurityInfo

func (response PostSupiUeUpu200JSONResponse) VisitPostSupiUeUpuResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(externalRef0.UpuSecurityInfo(response))
}

type PostSupiUeUpu307JSONResponse struct{ externalRef0.N307JSONResponse }

func (response PostSupiUeUpu307JSONResponse) VisitPostSupiUeUpuResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	if response.Headers.N3gppSbiTargetNfId != nil {
		w.Header().Set("3gpp-Sbi-Target-Nf-Id", fmt.Sprint(*response.Headers.N3gppSbiTargetNfId))
	}
	w.Header().Set("Location", fmt.Sprint(response.Headers.Location))
	w.WriteHeader(307)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostSupiUeUpu308ResponseHeaders struct {
	N3gppSbiTargetNfId *string
	Location           string
}

type PostSupiUeUpu308JSONResponse struct {
	Body    externalRef0.RedirectResponse
	Headers PostSupiUeUpu308ResponseHeaders
}

func (response PostSupiUeUpu308JSONResponse) VisitPostSupiUeUpuResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	if response.Headers.N3gppSbiTargetNfId != nil {
		w.Header().Set("3gpp-Sbi-Target-Nf-Id", fmt.Sprint(*response.Headers.N3gppSbiTargetNfId))
	}
	w.Header().Set("Location", fmt.Sprint(response.Headers.Location))
	w.WriteHeader(308)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostSupiUeUpu503ApplicationProblemPlusJSONResponse externalRef0.ProblemDetails

func (response PostSupiUeUpu503ApplicationProblemPlusJSONResponse) VisitPostSupiUeUpuResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(503)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response))
}

type PostSupiUeUpudefaultApplicationProblemPlusJSONResponse struct {
	Body       externalRef0.ProblemDetails
	StatusCode int
}

func (response PostSupiUeUpudefaultApplicationProblemPlusJSONResponse) VisitPostSupiUeUpuResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (POST /{supi}/ue-upu)
	PostSupiUeUpu(ctx context.Context, request PostSupiUeUpuRequestObject) (PostSupiUeUpuResponseObject, error)
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

// PostSupiUeUpu operation middleware
func (sh *strictHandler) PostSupiUeUpu(ctx *gin.Context, supi externalRef0.Supi) {
	var request PostSupiUeUpuRequestObject

	request.Supi = supi

	var body PostSupiUeUpuJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostSupiUeUpu(ctx, request.(PostSupiUeUpuRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostSupiUeUpu")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostSupiUeUpuResponseObject); ok {
		if err := validResponse.VisitPostSupiUeUpuResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
