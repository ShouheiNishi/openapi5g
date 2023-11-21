// Package sor provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package sor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	externalRef0 "github.com/ShouheiNishi/openapi5g/commondata"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

const (
	OAuth2ClientCredentialsScopes = "oAuth2ClientCredentials.Scopes"
)

// Defines values for AccessTech.
const (
	CDMA1xRTT                   AccessTech = "CDMA_1xRTT"
	CDMAHRPD                    AccessTech = "CDMA_HRPD"
	ECGSMIoTONLY                AccessTech = "ECGSM_IoT_ONLY"
	EUTRANINNBS1MODEONLY        AccessTech = "EUTRAN_IN_NBS1_MODE_ONLY"
	EUTRANINWBS1MODEANDNBS1MODE AccessTech = "EUTRAN_IN_WBS1_MODE_AND_NBS1_MODE"
	EUTRANINWBS1MODEONLY        AccessTech = "EUTRAN_IN_WBS1_MODE_ONLY"
	GSMANDECGSMIoT              AccessTech = "GSM_AND_ECGSM_IoT"
	GSMCOMPACT                  AccessTech = "GSM_COMPACT"
	GSMWITHOUTECGSMIoT          AccessTech = "GSM_WITHOUT_ECGSM_IoT"
	NR                          AccessTech = "NR"
	UTRAN                       AccessTech = "UTRAN"
)

// AccessTech defines model for AccessTech.
type AccessTech string

// AckInd defines model for AckInd.
type AckInd = bool

// CounterSor defines model for CounterSor.
type CounterSor = string

// SecuredPacket defines model for SecuredPacket.
type SecuredPacket = string

// SorInfo defines model for SorInfo.
type SorInfo struct {
	AckInd               AckInd                          `json:"ackInd"`
	SteeringContainer    *SteeringContainer              `json:"steeringContainer,omitempty"`
	SupportedFeatures    *externalRef0.SupportedFeatures `json:"supportedFeatures,omitempty"`
	AdditionalProperties map[string]interface{}          `json:"-"`
}

// SorMac defines model for SorMac.
type SorMac = string

// SorSecurityInfo defines model for SorSecurityInfo.
type SorSecurityInfo struct {
	CounterSor           CounterSor             `json:"counterSor"`
	SorMacIausf          SorMac                 `json:"sorMacIausf"`
	SorXmacIue           *SorMac                `json:"sorXmacIue,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// SteeringContainer defines model for SteeringContainer.
type SteeringContainer struct {
	union json.RawMessage
}

// SteeringContainer0 defines model for .
type SteeringContainer0 = []SteeringInfo

// SteeringInfo defines model for SteeringInfo.
type SteeringInfo struct {
	AccessTechList       *[]AccessTech          `json:"accessTechList,omitempty"`
	PlmnId               externalRef0.PlmnId    `json:"plmnId"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// PostSupiUeSorJSONRequestBody defines body for PostSupiUeSor for application/json ContentType.
type PostSupiUeSorJSONRequestBody = SorInfo

// Getter for additional properties for SorInfo. Returns the specified
// element and whether it was found
func (a SorInfo) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for SorInfo
func (a *SorInfo) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for SorInfo to handle AdditionalProperties
func (a *SorInfo) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["ackInd"]; found {
		err = json.Unmarshal(raw, &a.AckInd)
		if err != nil {
			return fmt.Errorf("error reading 'ackInd': %w", err)
		}
		delete(object, "ackInd")
	}

	if raw, found := object["steeringContainer"]; found {
		err = json.Unmarshal(raw, &a.SteeringContainer)
		if err != nil {
			return fmt.Errorf("error reading 'steeringContainer': %w", err)
		}
		delete(object, "steeringContainer")
	}

	if raw, found := object["supportedFeatures"]; found {
		err = json.Unmarshal(raw, &a.SupportedFeatures)
		if err != nil {
			return fmt.Errorf("error reading 'supportedFeatures': %w", err)
		}
		delete(object, "supportedFeatures")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for SorInfo to handle AdditionalProperties
func (a SorInfo) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["ackInd"], err = json.Marshal(a.AckInd)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'ackInd': %w", err)
	}

	if a.SteeringContainer != nil {
		object["steeringContainer"], err = json.Marshal(a.SteeringContainer)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'steeringContainer': %w", err)
		}
	}

	if a.SupportedFeatures != nil {
		object["supportedFeatures"], err = json.Marshal(a.SupportedFeatures)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'supportedFeatures': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for SorSecurityInfo. Returns the specified
// element and whether it was found
func (a SorSecurityInfo) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for SorSecurityInfo
func (a *SorSecurityInfo) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for SorSecurityInfo to handle AdditionalProperties
func (a *SorSecurityInfo) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["counterSor"]; found {
		err = json.Unmarshal(raw, &a.CounterSor)
		if err != nil {
			return fmt.Errorf("error reading 'counterSor': %w", err)
		}
		delete(object, "counterSor")
	}

	if raw, found := object["sorMacIausf"]; found {
		err = json.Unmarshal(raw, &a.SorMacIausf)
		if err != nil {
			return fmt.Errorf("error reading 'sorMacIausf': %w", err)
		}
		delete(object, "sorMacIausf")
	}

	if raw, found := object["sorXmacIue"]; found {
		err = json.Unmarshal(raw, &a.SorXmacIue)
		if err != nil {
			return fmt.Errorf("error reading 'sorXmacIue': %w", err)
		}
		delete(object, "sorXmacIue")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for SorSecurityInfo to handle AdditionalProperties
func (a SorSecurityInfo) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["counterSor"], err = json.Marshal(a.CounterSor)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'counterSor': %w", err)
	}

	object["sorMacIausf"], err = json.Marshal(a.SorMacIausf)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'sorMacIausf': %w", err)
	}

	if a.SorXmacIue != nil {
		object["sorXmacIue"], err = json.Marshal(a.SorXmacIue)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'sorXmacIue': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for SteeringInfo. Returns the specified
// element and whether it was found
func (a SteeringInfo) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for SteeringInfo
func (a *SteeringInfo) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for SteeringInfo to handle AdditionalProperties
func (a *SteeringInfo) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["accessTechList"]; found {
		err = json.Unmarshal(raw, &a.AccessTechList)
		if err != nil {
			return fmt.Errorf("error reading 'accessTechList': %w", err)
		}
		delete(object, "accessTechList")
	}

	if raw, found := object["plmnId"]; found {
		err = json.Unmarshal(raw, &a.PlmnId)
		if err != nil {
			return fmt.Errorf("error reading 'plmnId': %w", err)
		}
		delete(object, "plmnId")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for SteeringInfo to handle AdditionalProperties
func (a SteeringInfo) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.AccessTechList != nil {
		object["accessTechList"], err = json.Marshal(a.AccessTechList)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'accessTechList': %w", err)
		}
	}

	object["plmnId"], err = json.Marshal(a.PlmnId)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'plmnId': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// AsSteeringContainer0 returns the union data inside the SteeringContainer as a SteeringContainer0
func (t SteeringContainer) AsSteeringContainer0() (SteeringContainer0, error) {
	var body SteeringContainer0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSteeringContainer0 overwrites any union data inside the SteeringContainer as the provided SteeringContainer0
func (t *SteeringContainer) FromSteeringContainer0(v SteeringContainer0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSteeringContainer0 performs a merge with any union data inside the SteeringContainer, using the provided SteeringContainer0
func (t *SteeringContainer) MergeSteeringContainer0(v SteeringContainer0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsSecuredPacket returns the union data inside the SteeringContainer as a SecuredPacket
func (t SteeringContainer) AsSecuredPacket() (SecuredPacket, error) {
	var body SecuredPacket
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSecuredPacket overwrites any union data inside the SteeringContainer as the provided SecuredPacket
func (t *SteeringContainer) FromSecuredPacket(v SecuredPacket) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSecuredPacket performs a merge with any union data inside the SteeringContainer, using the provided SecuredPacket
func (t *SteeringContainer) MergeSecuredPacket(v SecuredPacket) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t SteeringContainer) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *SteeringContainer) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
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
	// PostSupiUeSorWithBody request with any body
	PostSupiUeSorWithBody(ctx context.Context, supi externalRef0.Supi, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostSupiUeSor(ctx context.Context, supi externalRef0.Supi, body PostSupiUeSorJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) PostSupiUeSorWithBody(ctx context.Context, supi externalRef0.Supi, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostSupiUeSorRequestWithBody(c.Server, supi, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostSupiUeSor(ctx context.Context, supi externalRef0.Supi, body PostSupiUeSorJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostSupiUeSorRequest(c.Server, supi, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewPostSupiUeSorRequest calls the generic PostSupiUeSor builder with application/json body
func NewPostSupiUeSorRequest(server string, supi externalRef0.Supi, body PostSupiUeSorJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostSupiUeSorRequestWithBody(server, supi, "application/json", bodyReader)
}

// NewPostSupiUeSorRequestWithBody generates requests for PostSupiUeSor with any type of body
func NewPostSupiUeSorRequestWithBody(server string, supi externalRef0.Supi, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/%s/ue-sor", pathParam0)
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
	// PostSupiUeSorWithBodyWithResponse request with any body
	PostSupiUeSorWithBodyWithResponse(ctx context.Context, supi externalRef0.Supi, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostSupiUeSorResponse, error)

	PostSupiUeSorWithResponse(ctx context.Context, supi externalRef0.Supi, body PostSupiUeSorJSONRequestBody, reqEditors ...RequestEditorFn) (*PostSupiUeSorResponse, error)
}

type PostSupiUeSorResponse struct {
	Body                      []byte
	HTTPResponse              *http.Response
	JSON200                   *SorSecurityInfo
	JSON307                   *externalRef0.N307
	JSON308                   *externalRef0.N308
	ApplicationproblemJSON503 *externalRef0.ProblemDetails
}

// Status returns HTTPResponse.Status
func (r PostSupiUeSorResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostSupiUeSorResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostSupiUeSorWithBodyWithResponse request with arbitrary body returning *PostSupiUeSorResponse
func (c *ClientWithResponses) PostSupiUeSorWithBodyWithResponse(ctx context.Context, supi externalRef0.Supi, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostSupiUeSorResponse, error) {
	rsp, err := c.PostSupiUeSorWithBody(ctx, supi, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostSupiUeSorResponse(rsp)
}

func (c *ClientWithResponses) PostSupiUeSorWithResponse(ctx context.Context, supi externalRef0.Supi, body PostSupiUeSorJSONRequestBody, reqEditors ...RequestEditorFn) (*PostSupiUeSorResponse, error) {
	rsp, err := c.PostSupiUeSor(ctx, supi, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostSupiUeSorResponse(rsp)
}

// ParsePostSupiUeSorResponse parses an HTTP response from a PostSupiUeSorWithResponse call
func ParsePostSupiUeSorResponse(rsp *http.Response) (*PostSupiUeSorResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostSupiUeSorResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SorSecurityInfo
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
		var dest externalRef0.N308
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

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /{supi}/ue-sor)
	PostSupiUeSor(c *gin.Context, supi externalRef0.Supi)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// PostSupiUeSor operation middleware
func (siw *ServerInterfaceWrapper) PostSupiUeSor(c *gin.Context) {

	var err error

	// ------------- Path parameter "supi" -------------
	var supi externalRef0.Supi

	err = runtime.BindStyledParameter("simple", false, "supi", c.Param("supi"), &supi)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter supi: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(OAuth2ClientCredentialsScopes, []string{"nausf-sorprotection"})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostSupiUeSor(c, supi)
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

	router.POST(options.BaseURL+"/:supi/ue-sor", wrapper.PostSupiUeSor)
}

type PostSupiUeSorRequestObject struct {
	Supi externalRef0.Supi `json:"supi"`
	Body *PostSupiUeSorJSONRequestBody
}

type PostSupiUeSorResponseObject interface {
	VisitPostSupiUeSorResponse(w http.ResponseWriter) error
}

type PostSupiUeSor200JSONResponse SorSecurityInfo

func (response PostSupiUeSor200JSONResponse) VisitPostSupiUeSorResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostSupiUeSor307JSONResponse struct{ externalRef0.N307JSONResponse }

func (response PostSupiUeSor307JSONResponse) VisitPostSupiUeSorResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("3gpp-Sbi-Target-Nf-Id", fmt.Sprint(response.Headers.N3gppSbiTargetNfId))
	w.Header().Set("Location", fmt.Sprint(response.Headers.Location))
	w.WriteHeader(307)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostSupiUeSor308JSONResponse struct{ externalRef0.N308JSONResponse }

func (response PostSupiUeSor308JSONResponse) VisitPostSupiUeSorResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("3gpp-Sbi-Target-Nf-Id", fmt.Sprint(response.Headers.N3gppSbiTargetNfId))
	w.Header().Set("Location", fmt.Sprint(response.Headers.Location))
	w.WriteHeader(308)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostSupiUeSor503ApplicationProblemPlusJSONResponse externalRef0.ProblemDetails

func (response PostSupiUeSor503ApplicationProblemPlusJSONResponse) VisitPostSupiUeSorResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(503)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (POST /{supi}/ue-sor)
	PostSupiUeSor(ctx context.Context, request PostSupiUeSorRequestObject) (PostSupiUeSorResponseObject, error)
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

// PostSupiUeSor operation middleware
func (sh *strictHandler) PostSupiUeSor(ctx *gin.Context, supi externalRef0.Supi) {
	var request PostSupiUeSorRequestObject

	request.Supi = supi

	var body PostSupiUeSorJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostSupiUeSor(ctx, request.(PostSupiUeSorRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostSupiUeSor")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostSupiUeSorResponseObject); ok {
		if err := validResponse.VisitPostSupiUeSorResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
