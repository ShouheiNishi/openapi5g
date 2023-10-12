package error

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/ShouheiNishi/openapi5g/commondata"
	"github.com/ShouheiNishi/openapi5g/utils/problem"
)

type WrappedOpenAPIError struct {
	Method    string
	BaseError error
}

func (e *WrappedOpenAPIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Method, e.BaseError.Error())
}

func (e *WrappedOpenAPIError) Unwrap() error {
	return e.BaseError
}

type ProblemDetailsError struct {
	Message        string
	ProblemDetails commondata.ProblemDetails
}

func (e *ProblemDetailsError) Error() string {
	return fmt.Sprintf("%s %v", e.Message, e.ProblemDetails)
}

func ExtractAndWrapOpenAPIError(method string, res any, errIn error) (err error) {
	defer func() {
		if err == nil {
			err = errors.New("unknown error")
		}
		err = &WrappedOpenAPIError{
			Method:    method,
			BaseError: err,
		}
	}()

	if errIn != nil {
		return errIn
	}

	body, httpResponse, err := problem.ExtractBodyAndHttpResponse(res)
	if err != nil {
		return fmt.Errorf("problem.ExtractBodyAndHttpResponse: %w", err)
	}

	if httpResponse == nil {
		return errors.New("no HTTP response")
	}

	if !strings.Contains(httpResponse.Header.Get("Content-Type"), "application/problem+json") || len(body) == 0 {
		return fmt.Errorf("no problemDetails, status code = %d, content-type = %s", httpResponse.StatusCode, httpResponse.Header.Get("Content-Type"))
	}

	var pd commondata.ProblemDetails
	if err := json.Unmarshal(body, &pd); err != nil {
		return fmt.Errorf("json.Unmarshal: %w, status code = %d", err, httpResponse.StatusCode)
	}

	return &ProblemDetailsError{
		Message:        "problemDetails received",
		ProblemDetails: pd,
	}
}

func ErrorToProblemDetails(err error) commondata.ProblemDetails {
	var pe *ProblemDetailsError
	if errors.As(err, &pe) {
		return pe.ProblemDetails
	}

	return problem.SystemFailure(err.Error())
}
