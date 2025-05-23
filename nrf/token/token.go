// Package token provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.1-0.20240123090344-d326c01d279a DO NOT EDIT.
package token

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
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

// AccessTokenRequestParams defines parameters for AccessTokenRequest.
type AccessTokenRequestParams struct {
	// ContentEncoding Content-Encoding, described in IETF RFC 7231
	ContentEncoding *string `json:"Content-Encoding,omitempty"`

	// AcceptEncoding Accept-Encoding, described in IETF RFC 7231
	AcceptEncoding *string `json:"Accept-Encoding,omitempty"`
}

// AccessTokenRequestFormdataRequestBody defines body for AccessTokenRequest for application/x-www-form-urlencoded ContentType.
type AccessTokenRequestFormdataRequestBody = externalRef0.AccessTokenReq

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
	// AccessTokenRequestWithBody request with any body
	AccessTokenRequestWithBody(ctx context.Context, params *AccessTokenRequestParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AccessTokenRequestWithFormdataBody(ctx context.Context, params *AccessTokenRequestParams, body AccessTokenRequestFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) AccessTokenRequestWithBody(ctx context.Context, params *AccessTokenRequestParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAccessTokenRequestRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AccessTokenRequestWithFormdataBody(ctx context.Context, params *AccessTokenRequestParams, body AccessTokenRequestFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAccessTokenRequestRequestWithFormdataBody(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewAccessTokenRequestRequestWithFormdataBody calls the generic AccessTokenRequest builder with application/x-www-form-urlencoded body
func NewAccessTokenRequestRequestWithFormdataBody(server string, params *AccessTokenRequestParams, body AccessTokenRequestFormdataRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	bodyStr, err := runtime.MarshalForm(body, nil)
	if err != nil {
		return nil, err
	}
	bodyReader = strings.NewReader(bodyStr.Encode())
	return NewAccessTokenRequestRequestWithBody(server, params, "application/x-www-form-urlencoded", bodyReader)
}

// NewAccessTokenRequestRequestWithBody generates requests for AccessTokenRequest with any type of body
func NewAccessTokenRequestRequestWithBody(server string, params *AccessTokenRequestParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/oauth2/token")
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

	if params != nil {

		if params.ContentEncoding != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithLocation("simple", false, "Content-Encoding", runtime.ParamLocationHeader, *params.ContentEncoding)
			if err != nil {
				return nil, err
			}

			req.Header.Set("Content-Encoding", headerParam0)
		}

		if params.AcceptEncoding != nil {
			var headerParam1 string

			headerParam1, err = runtime.StyleParamWithLocation("simple", false, "Accept-Encoding", runtime.ParamLocationHeader, *params.AcceptEncoding)
			if err != nil {
				return nil, err
			}

			req.Header.Set("Accept-Encoding", headerParam1)
		}

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
	// AccessTokenRequestWithBodyWithResponse request with any body
	AccessTokenRequestWithBodyWithResponse(ctx context.Context, params *AccessTokenRequestParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AccessTokenRequestResponse, error)

	AccessTokenRequestWithFormdataBodyWithResponse(ctx context.Context, params *AccessTokenRequestParams, body AccessTokenRequestFormdataRequestBody, reqEditors ...RequestEditorFn) (*AccessTokenRequestResponse, error)
}

type AccessTokenRequestResponse struct {
	Body                          []byte
	HTTPResponse                  *http.Response
	JSON200                       *externalRef0.AccessTokenRsp
	JSON307                       *externalRef0.RedirectResponse
	JSON308                       *externalRef0.RedirectResponse
	JSON400                       *externalRef0.AccessTokenErr
	ApplicationproblemJSON400     *externalRef0.ProblemDetails
	ApplicationproblemJSON401     *externalRef0.N401
	ApplicationproblemJSON403     *externalRef0.N403
	ApplicationproblemJSON404     *externalRef0.N404
	ApplicationproblemJSON411     *externalRef0.N411
	ApplicationproblemJSON413     *externalRef0.N413
	ApplicationproblemJSON415     *externalRef0.N415
	ApplicationproblemJSON429     *externalRef0.N429
	ApplicationproblemJSON500     *externalRef0.N500
	ApplicationproblemJSON501     *externalRef0.N501
	ApplicationproblemJSON503     *externalRef0.N503
	ApplicationproblemJSONDefault *externalRef0.Default
}

// Status returns HTTPResponse.Status
func (r AccessTokenRequestResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AccessTokenRequestResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// AccessTokenRequestWithBodyWithResponse request with arbitrary body returning *AccessTokenRequestResponse
func (c *ClientWithResponses) AccessTokenRequestWithBodyWithResponse(ctx context.Context, params *AccessTokenRequestParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AccessTokenRequestResponse, error) {
	rsp, err := c.AccessTokenRequestWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAccessTokenRequestResponse(rsp)
}

func (c *ClientWithResponses) AccessTokenRequestWithFormdataBodyWithResponse(ctx context.Context, params *AccessTokenRequestParams, body AccessTokenRequestFormdataRequestBody, reqEditors ...RequestEditorFn) (*AccessTokenRequestResponse, error) {
	rsp, err := c.AccessTokenRequestWithFormdataBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAccessTokenRequestResponse(rsp)
}

// ParseAccessTokenRequestResponse parses an HTTP response from a AccessTokenRequestWithResponse call
func ParseAccessTokenRequestResponse(rsp *http.Response) (*AccessTokenRequestResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AccessTokenRequestResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest externalRef0.AccessTokenRsp
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

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

	case rsp.Header.Get("Content-Type") == "application/json" && rsp.StatusCode == 400:
		var dest externalRef0.AccessTokenErr
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case rsp.Header.Get("Content-Type") == "application/problem+json" && rsp.StatusCode == 400:
		var dest externalRef0.ProblemDetails
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef0.N401
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON401 = &dest

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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 411:
		var dest externalRef0.N411
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON411 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest externalRef0.N413
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest externalRef0.N415
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest externalRef0.N429
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON429 = &dest

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
	// Access Token Request
	// (POST /oauth2/token)
	AccessTokenRequest(c *gin.Context, params AccessTokenRequestParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// AccessTokenRequest operation middleware
func (siw *ServerInterfaceWrapper) AccessTokenRequest(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params AccessTokenRequestParams

	headers := c.Request.Header

	// ------------- Optional header parameter "Content-Encoding" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Content-Encoding")]; found {
		var ContentEncoding string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for Content-Encoding, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "Content-Encoding", valueList[0], &ContentEncoding, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter Content-Encoding: %w", err), http.StatusBadRequest)
			return
		}

		params.ContentEncoding = &ContentEncoding

	}

	// ------------- Optional header parameter "Accept-Encoding" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Accept-Encoding")]; found {
		var AcceptEncoding string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for Accept-Encoding, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "Accept-Encoding", valueList[0], &AcceptEncoding, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter Accept-Encoding: %w", err), http.StatusBadRequest)
			return
		}

		params.AcceptEncoding = &AcceptEncoding

	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AccessTokenRequest(c, params)
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

	router.POST(options.BaseURL+"/oauth2/token", wrapper.AccessTokenRequest)
}

type AccessTokenRequestRequestObject struct {
	Params AccessTokenRequestParams
	Body   *AccessTokenRequestFormdataRequestBody
}

type AccessTokenRequestResponseObject interface {
	VisitAccessTokenRequestResponse(w http.ResponseWriter) error
}

type AccessTokenRequest200ResponseHeaders struct {
	AcceptEncoding  *string
	CacheControl    string
	ContentEncoding *string
	Pragma          string
}

type AccessTokenRequest200JSONResponse struct {
	Body    externalRef0.AccessTokenRsp
	Headers AccessTokenRequest200ResponseHeaders
}

func (response AccessTokenRequest200JSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	if response.Headers.AcceptEncoding != nil {
		w.Header().Set("Accept-Encoding", fmt.Sprint(*response.Headers.AcceptEncoding))
	}
	w.Header().Set("Cache-Control", fmt.Sprint(response.Headers.CacheControl))
	if response.Headers.ContentEncoding != nil {
		w.Header().Set("Content-Encoding", fmt.Sprint(*response.Headers.ContentEncoding))
	}
	w.Header().Set("Pragma", fmt.Sprint(response.Headers.Pragma))
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type AccessTokenRequest307ResponseHeaders struct {
	Location string
}

type AccessTokenRequest307JSONResponse struct {
	Body    externalRef0.RedirectResponse
	Headers AccessTokenRequest307ResponseHeaders
}

func (response AccessTokenRequest307JSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", fmt.Sprint(response.Headers.Location))
	w.WriteHeader(307)

	return json.NewEncoder(w).Encode(response.Body)
}

type AccessTokenRequest308ResponseHeaders struct {
	Location string
}

type AccessTokenRequest308JSONResponse struct {
	Body    externalRef0.RedirectResponse
	Headers AccessTokenRequest308ResponseHeaders
}

func (response AccessTokenRequest308JSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", fmt.Sprint(response.Headers.Location))
	w.WriteHeader(308)

	return json.NewEncoder(w).Encode(response.Body)
}

type AccessTokenRequest400ResponseHeaders struct {
	CacheControl string
	Pragma       string
}

type AccessTokenRequest400JSONResponse struct {
	Body    externalRef0.AccessTokenErr
	Headers AccessTokenRequest400ResponseHeaders
}

func (response AccessTokenRequest400JSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", fmt.Sprint(response.Headers.CacheControl))
	w.Header().Set("Pragma", fmt.Sprint(response.Headers.Pragma))
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response.Body)
}

type AccessTokenRequest400ApplicationProblemPlusJSONResponse struct {
	Body    externalRef0.ProblemDetails
	Headers AccessTokenRequest400ResponseHeaders
}

func (response AccessTokenRequest400ApplicationProblemPlusJSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.Header().Set("Cache-Control", fmt.Sprint(response.Headers.CacheControl))
	w.Header().Set("Pragma", fmt.Sprint(response.Headers.Pragma))
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response.Body)
}

type AccessTokenRequest401ApplicationProblemPlusJSONResponse struct {
	externalRef0.N401ApplicationProblemPlusJSONResponse
}

func (response AccessTokenRequest401ApplicationProblemPlusJSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N401ApplicationProblemPlusJSONResponse))
}

type AccessTokenRequest403ApplicationProblemPlusJSONResponse struct {
	externalRef0.N403ApplicationProblemPlusJSONResponse
}

func (response AccessTokenRequest403ApplicationProblemPlusJSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N403ApplicationProblemPlusJSONResponse))
}

type AccessTokenRequest404ApplicationProblemPlusJSONResponse struct {
	externalRef0.N404ApplicationProblemPlusJSONResponse
}

func (response AccessTokenRequest404ApplicationProblemPlusJSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N404ApplicationProblemPlusJSONResponse))
}

type AccessTokenRequest411ApplicationProblemPlusJSONResponse struct {
	externalRef0.N411ApplicationProblemPlusJSONResponse
}

func (response AccessTokenRequest411ApplicationProblemPlusJSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(411)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N411ApplicationProblemPlusJSONResponse))
}

type AccessTokenRequest413ApplicationProblemPlusJSONResponse struct {
	externalRef0.N413ApplicationProblemPlusJSONResponse
}

func (response AccessTokenRequest413ApplicationProblemPlusJSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(413)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N413ApplicationProblemPlusJSONResponse))
}

type AccessTokenRequest415ApplicationProblemPlusJSONResponse struct {
	externalRef0.N415ApplicationProblemPlusJSONResponse
}

func (response AccessTokenRequest415ApplicationProblemPlusJSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(415)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N415ApplicationProblemPlusJSONResponse))
}

type AccessTokenRequest429ApplicationProblemPlusJSONResponse struct {
	externalRef0.N429ApplicationProblemPlusJSONResponse
}

func (response AccessTokenRequest429ApplicationProblemPlusJSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(429)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N429ApplicationProblemPlusJSONResponse))
}

type AccessTokenRequest500ApplicationProblemPlusJSONResponse struct {
	externalRef0.N500ApplicationProblemPlusJSONResponse
}

func (response AccessTokenRequest500ApplicationProblemPlusJSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N500ApplicationProblemPlusJSONResponse))
}

type AccessTokenRequest501ApplicationProblemPlusJSONResponse struct {
	externalRef0.N501ApplicationProblemPlusJSONResponse
}

func (response AccessTokenRequest501ApplicationProblemPlusJSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(501)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N501ApplicationProblemPlusJSONResponse))
}

type AccessTokenRequest503ApplicationProblemPlusJSONResponse struct {
	externalRef0.N503ApplicationProblemPlusJSONResponse
}

func (response AccessTokenRequest503ApplicationProblemPlusJSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(503)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N503ApplicationProblemPlusJSONResponse))
}

type AccessTokenRequestdefaultApplicationProblemPlusJSONResponse struct {
	Body       externalRef0.ProblemDetails
	StatusCode int
}

func (response AccessTokenRequestdefaultApplicationProblemPlusJSONResponse) VisitAccessTokenRequestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Access Token Request
	// (POST /oauth2/token)
	AccessTokenRequest(ctx context.Context, request AccessTokenRequestRequestObject) (AccessTokenRequestResponseObject, error)
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

// AccessTokenRequest operation middleware
func (sh *strictHandler) AccessTokenRequest(ctx *gin.Context, params AccessTokenRequestParams) {
	var request AccessTokenRequestRequestObject

	request.Params = params

	if err := ctx.Request.ParseForm(); err != nil {
		ctx.Error(err)
		return
	}
	var body AccessTokenRequestFormdataRequestBody
	if err := runtime.BindForm(&body, ctx.Request.Form, nil, nil); err != nil {
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AccessTokenRequest(ctx, request.(AccessTokenRequestRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AccessTokenRequest")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AccessTokenRequestResponseObject); ok {
		if err := validResponse.VisitAccessTokenRequestResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
