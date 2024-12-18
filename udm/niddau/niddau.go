// Package niddau provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.1-0.20240123090344-d326c01d279a DO NOT EDIT.
package niddau

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

// AuthorizeNiddDataJSONRequestBody defines body for AuthorizeNiddData for application/json ContentType.
type AuthorizeNiddDataJSONRequestBody = externalRef0.AuthorizationInfo

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
	// AuthorizeNiddDataWithBody request with any body
	AuthorizeNiddDataWithBody(ctx context.Context, ueIdentity string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AuthorizeNiddData(ctx context.Context, ueIdentity string, body AuthorizeNiddDataJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) AuthorizeNiddDataWithBody(ctx context.Context, ueIdentity string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAuthorizeNiddDataRequestWithBody(c.Server, ueIdentity, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AuthorizeNiddData(ctx context.Context, ueIdentity string, body AuthorizeNiddDataJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAuthorizeNiddDataRequest(c.Server, ueIdentity, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewAuthorizeNiddDataRequest calls the generic AuthorizeNiddData builder with application/json body
func NewAuthorizeNiddDataRequest(server string, ueIdentity string, body AuthorizeNiddDataJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAuthorizeNiddDataRequestWithBody(server, ueIdentity, "application/json", bodyReader)
}

// NewAuthorizeNiddDataRequestWithBody generates requests for AuthorizeNiddData with any type of body
func NewAuthorizeNiddDataRequestWithBody(server string, ueIdentity string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "ueIdentity", runtime.ParamLocationPath, ueIdentity)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/%s/authorize", pathParam0)
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
	// AuthorizeNiddDataWithBodyWithResponse request with any body
	AuthorizeNiddDataWithBodyWithResponse(ctx context.Context, ueIdentity string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AuthorizeNiddDataResponse, error)

	AuthorizeNiddDataWithResponse(ctx context.Context, ueIdentity string, body AuthorizeNiddDataJSONRequestBody, reqEditors ...RequestEditorFn) (*AuthorizeNiddDataResponse, error)
}

type AuthorizeNiddDataResponse struct {
	Body                          []byte
	HTTPResponse                  *http.Response
	JSON200                       *externalRef0.AuthorizationData
	ApplicationproblemJSON400     *externalRef0.N400
	ApplicationproblemJSON403     *externalRef0.N403
	ApplicationproblemJSON404     *externalRef0.N404
	ApplicationproblemJSON500     *externalRef0.N500
	ApplicationproblemJSON501     *externalRef0.N501
	ApplicationproblemJSON503     *externalRef0.N503
	ApplicationproblemJSONDefault *externalRef0.ProblemDetails
}

// Status returns HTTPResponse.Status
func (r AuthorizeNiddDataResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AuthorizeNiddDataResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// AuthorizeNiddDataWithBodyWithResponse request with arbitrary body returning *AuthorizeNiddDataResponse
func (c *ClientWithResponses) AuthorizeNiddDataWithBodyWithResponse(ctx context.Context, ueIdentity string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AuthorizeNiddDataResponse, error) {
	rsp, err := c.AuthorizeNiddDataWithBody(ctx, ueIdentity, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAuthorizeNiddDataResponse(rsp)
}

func (c *ClientWithResponses) AuthorizeNiddDataWithResponse(ctx context.Context, ueIdentity string, body AuthorizeNiddDataJSONRequestBody, reqEditors ...RequestEditorFn) (*AuthorizeNiddDataResponse, error) {
	rsp, err := c.AuthorizeNiddData(ctx, ueIdentity, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAuthorizeNiddDataResponse(rsp)
}

// ParseAuthorizeNiddDataResponse parses an HTTP response from a AuthorizeNiddDataWithResponse call
func ParseAuthorizeNiddDataResponse(rsp *http.Response) (*AuthorizeNiddDataResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AuthorizeNiddDataResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest externalRef0.AuthorizationData
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef0.N400
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest externalRef0.N403
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest externalRef0.N404
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.N500
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 501:
		var dest externalRef0.N501
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON501 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest externalRef0.N503
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
	// Authorize the NIDD configuration request.
	// (POST /{ueIdentity}/authorize)
	AuthorizeNiddData(c *gin.Context, ueIdentity string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// AuthorizeNiddData operation middleware
func (siw *ServerInterfaceWrapper) AuthorizeNiddData(c *gin.Context) {

	var err error

	// ------------- Path parameter "ueIdentity" -------------
	var ueIdentity string

	err = runtime.BindStyledParameterWithOptions("simple", "ueIdentity", c.Param("ueIdentity"), &ueIdentity, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ueIdentity: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(OAuth2ClientCredentialsScopes, []string{"nudm-niddau"})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AuthorizeNiddData(c, ueIdentity)
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

	router.POST(options.BaseURL+"/:ueIdentity/authorize", wrapper.AuthorizeNiddData)
}

type AuthorizeNiddDataRequestObject struct {
	UeIdentity string `json:"ueIdentity"`
	Body       *AuthorizeNiddDataJSONRequestBody
}

type AuthorizeNiddDataResponseObject interface {
	VisitAuthorizeNiddDataResponse(w http.ResponseWriter) error
}

type AuthorizeNiddData200JSONResponse externalRef0.AuthorizationData

func (response AuthorizeNiddData200JSONResponse) VisitAuthorizeNiddDataResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(externalRef0.AuthorizationData(response))
}

type AuthorizeNiddData400ApplicationProblemPlusJSONResponse struct {
	externalRef0.N400ApplicationProblemPlusJSONResponse
}

func (response AuthorizeNiddData400ApplicationProblemPlusJSONResponse) VisitAuthorizeNiddDataResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N400ApplicationProblemPlusJSONResponse))
}

type AuthorizeNiddData403ApplicationProblemPlusJSONResponse struct {
	externalRef0.N403ApplicationProblemPlusJSONResponse
}

func (response AuthorizeNiddData403ApplicationProblemPlusJSONResponse) VisitAuthorizeNiddDataResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N403ApplicationProblemPlusJSONResponse))
}

type AuthorizeNiddData404ApplicationProblemPlusJSONResponse struct {
	externalRef0.N404ApplicationProblemPlusJSONResponse
}

func (response AuthorizeNiddData404ApplicationProblemPlusJSONResponse) VisitAuthorizeNiddDataResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N404ApplicationProblemPlusJSONResponse))
}

type AuthorizeNiddData500ApplicationProblemPlusJSONResponse struct {
	externalRef0.N500ApplicationProblemPlusJSONResponse
}

func (response AuthorizeNiddData500ApplicationProblemPlusJSONResponse) VisitAuthorizeNiddDataResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N500ApplicationProblemPlusJSONResponse))
}

type AuthorizeNiddData501ApplicationProblemPlusJSONResponse struct {
	externalRef0.N501ApplicationProblemPlusJSONResponse
}

func (response AuthorizeNiddData501ApplicationProblemPlusJSONResponse) VisitAuthorizeNiddDataResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(501)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N501ApplicationProblemPlusJSONResponse))
}

type AuthorizeNiddData503ApplicationProblemPlusJSONResponse struct {
	externalRef0.N503ApplicationProblemPlusJSONResponse
}

func (response AuthorizeNiddData503ApplicationProblemPlusJSONResponse) VisitAuthorizeNiddDataResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(503)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N503ApplicationProblemPlusJSONResponse))
}

type AuthorizeNiddDatadefaultApplicationProblemPlusJSONResponse struct {
	Body       externalRef0.ProblemDetails
	StatusCode int
}

func (response AuthorizeNiddDatadefaultApplicationProblemPlusJSONResponse) VisitAuthorizeNiddDataResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Authorize the NIDD configuration request.
	// (POST /{ueIdentity}/authorize)
	AuthorizeNiddData(ctx context.Context, request AuthorizeNiddDataRequestObject) (AuthorizeNiddDataResponseObject, error)
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

// AuthorizeNiddData operation middleware
func (sh *strictHandler) AuthorizeNiddData(ctx *gin.Context, ueIdentity string) {
	var request AuthorizeNiddDataRequestObject

	request.UeIdentity = ueIdentity

	var body AuthorizeNiddDataJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AuthorizeNiddData(ctx, request.(AuthorizeNiddDataRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AuthorizeNiddData")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AuthorizeNiddDataResponseObject); ok {
		if err := validResponse.VisitAuthorizeNiddDataResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
