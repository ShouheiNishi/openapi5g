package problem

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/ShouheiNishi/openapi5g/commondata"
)

func ExtractBodyAndHttpResponse(v any) (body []byte, httpResponse *http.Response, err error) {
	defer func() {
		if p := recover(); p != nil {
			if e, ok := p.(error); ok {
				err = fmt.Errorf("ExtractBodyAndHttpResponse: panic occured %w", e)
			} else {
				err = fmt.Errorf("ExtractBodyAndHttpResponse: panic occured %+v", p)
			}
		}
	}()
	return extractBodyAndHttpResponseMain(v)
}

func extractBodyAndHttpResponseMain(v any) (body []byte, httpResponse *http.Response, err error) {
	vValue := reflect.ValueOf(v)
pointerLoop:
	for {
		if !vValue.IsValid() {
			return nil, nil, errors.New("invalid value")
		}
		switch vValue.Kind() {
		case reflect.Pointer, reflect.Interface:
			vValue = vValue.Elem()
			continue
		case reflect.Struct:
			break pointerLoop
		default:
			return nil, nil, fmt.Errorf("invalid type %T", v)
		}
	}

	vBody := vValue.FieldByName("Body")
	if !vBody.IsValid() {
		return nil, nil, errors.New("body is not exist")
	}
	if !vBody.Type().AssignableTo(reflect.TypeOf(body)) {
		return nil, nil, errors.New("body is invalid type")
	}
	reflect.ValueOf(&body).Elem().Set(vBody)

	vHTTPResponse := vValue.FieldByName("HTTPResponse")
	if !vHTTPResponse.IsValid() {
		return nil, nil, errors.New("HTTPResponse is not exist")
	}
	if !vHTTPResponse.Type().AssignableTo(reflect.TypeOf(httpResponse)) {
		return nil, nil, errors.New("HTTPResponse is invalid type")
	}
	reflect.ValueOf(&httpResponse).Elem().Set(vHTTPResponse)

	return
}

func ExtractStatusCodeAndProblemDetails(v any) (statusCode int, problemDetails *commondata.ProblemDetails, err error) {
	body, httpResponse, err := ExtractBodyAndHttpResponse(v)
	if err != nil {
		return 0, nil, err
	}

	if httpResponse == nil {
		return 0, nil, errors.New("HTTPResponse is nil")
	}

	statusCode = httpResponse.StatusCode

	if !strings.Contains(httpResponse.Header.Get("Content-Type"), "application/problem+json") {
		return statusCode, nil, errors.New("Content-Type mismatch")
	}

	if len(body) == 0 {
		return statusCode, nil, errors.New("empty body")
	}

	var pd commondata.ProblemDetails
	if err := json.Unmarshal(body, &pd); err != nil {
		return statusCode, nil, err
	}

	return statusCode, &pd, nil
}
