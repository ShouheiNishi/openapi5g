// Package idmap provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.1-0.20240123090344-d326c01d279a DO NOT EDIT.
package idmap

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

const (
	OAuth2ClientCredentialsScopes = "oAuth2ClientCredentials.Scopes"
)

// GetNfGroupIDsParams defines parameters for GetNfGroupIDs.
type GetNfGroupIDsParams struct {
	// NfType Type of NF
	NfType []externalRef0.NFType `form:"nf-type" json:"nf-type"`

	// SubscriberId Identifier of the subscriber
	SubscriberId externalRef0.SubscriberId `form:"subscriberId" json:"subscriberId"`
}

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
	// GetNfGroupIDs request
	GetNfGroupIDs(ctx context.Context, params *GetNfGroupIDsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetNfGroupIDs(ctx context.Context, params *GetNfGroupIDsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetNfGroupIDsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetNfGroupIDsRequest generates requests for GetNfGroupIDs
func NewGetNfGroupIDsRequest(server string, params *GetNfGroupIDsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/nf-group-ids")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", false, "nf-type", runtime.ParamLocationQuery, params.NfType); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "subscriberId", runtime.ParamLocationQuery, params.SubscriberId); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
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
	// GetNfGroupIDsWithResponse request
	GetNfGroupIDsWithResponse(ctx context.Context, params *GetNfGroupIDsParams, reqEditors ...RequestEditorFn) (*GetNfGroupIDsResponse, error)
}

type GetNfGroupIDsResponse struct {
	Body                          []byte
	HTTPResponse                  *http.Response
	JSON200                       *externalRef0.NfGroupIdMapResult
	ApplicationproblemJSON404     *externalRef0.N404
	ApplicationproblemJSONDefault *externalRef0.ProblemDetails
}

// Status returns HTTPResponse.Status
func (r GetNfGroupIDsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetNfGroupIDsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetNfGroupIDsWithResponse request returning *GetNfGroupIDsResponse
func (c *ClientWithResponses) GetNfGroupIDsWithResponse(ctx context.Context, params *GetNfGroupIDsParams, reqEditors ...RequestEditorFn) (*GetNfGroupIDsResponse, error) {
	rsp, err := c.GetNfGroupIDs(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetNfGroupIDsResponse(rsp)
}

// ParseGetNfGroupIDsResponse parses an HTTP response from a GetNfGroupIDsWithResponse call
func ParseGetNfGroupIDsResponse(rsp *http.Response) (*GetNfGroupIDsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetNfGroupIDsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest externalRef0.NfGroupIdMapResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest externalRef0.N404
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON404 = &dest

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
	// Retrieves NF-Group IDs for provided Subscriber and NF types
	// (GET /nf-group-ids)
	GetNfGroupIDs(c *gin.Context, params GetNfGroupIDsParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetNfGroupIDs operation middleware
func (siw *ServerInterfaceWrapper) GetNfGroupIDs(c *gin.Context) {

	var err error

	c.Set(OAuth2ClientCredentialsScopes, []string{"nudr-group-id-map"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNfGroupIDsParams

	// ------------- Required query parameter "nf-type" -------------

	if paramValue := c.Query("nf-type"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument nf-type is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", false, true, "nf-type", c.Request.URL.Query(), &params.NfType)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter nf-type: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "subscriberId" -------------

	if paramValue := c.Query("subscriberId"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument subscriberId is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "subscriberId", c.Request.URL.Query(), &params.SubscriberId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter subscriberId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetNfGroupIDs(c, params)
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

	router.GET(options.BaseURL+"/nf-group-ids", wrapper.GetNfGroupIDs)
}

type GetNfGroupIDsRequestObject struct {
	Params GetNfGroupIDsParams
}

type GetNfGroupIDsResponseObject interface {
	VisitGetNfGroupIDsResponse(w http.ResponseWriter) error
}

type GetNfGroupIDs200JSONResponse externalRef0.NfGroupIdMapResult

func (response GetNfGroupIDs200JSONResponse) VisitGetNfGroupIDsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(externalRef0.NfGroupIdMapResult(response))
}

type GetNfGroupIDs404ApplicationProblemPlusJSONResponse struct {
	externalRef0.N404ApplicationProblemPlusJSONResponse
}

func (response GetNfGroupIDs404ApplicationProblemPlusJSONResponse) VisitGetNfGroupIDsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(externalRef0.ProblemDetails(response.N404ApplicationProblemPlusJSONResponse))
}

type GetNfGroupIDsdefaultApplicationProblemPlusJSONResponse struct {
	Body       externalRef0.ProblemDetails
	StatusCode int
}

func (response GetNfGroupIDsdefaultApplicationProblemPlusJSONResponse) VisitGetNfGroupIDsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Retrieves NF-Group IDs for provided Subscriber and NF types
	// (GET /nf-group-ids)
	GetNfGroupIDs(ctx context.Context, request GetNfGroupIDsRequestObject) (GetNfGroupIDsResponseObject, error)
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

// GetNfGroupIDs operation middleware
func (sh *strictHandler) GetNfGroupIDs(ctx *gin.Context, params GetNfGroupIDsParams) {
	var request GetNfGroupIDsRequestObject

	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetNfGroupIDs(ctx, request.(GetNfGroupIDsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetNfGroupIDs")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetNfGroupIDsResponseObject); ok {
		if err := validResponse.VisitGetNfGroupIDsResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
