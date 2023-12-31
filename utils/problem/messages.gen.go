// Code generated by github.com/ShouheiNishi/openapi5g/internal/generator/problem, DO NOT EDIT.

package problem

import (
	"net/http"

	"github.com/samber/lo"

	"github.com/ShouheiNishi/openapi5g/commondata"
)

const (
	CauseAccessTokenDenied            = "ACCESS_TOKEN_DENIED"
	CauseCCAVerificationFailure       = "CCA_VERIFICATION_FAILURE"
	CauseIncorrectLength              = "INCORRECT_LENGTH"
	CauseInsufficientResources        = "INSUFFICIENT_RESOURCES"
	CauseInvalidAPI                   = "INVALID_API"
	CauseInvalidDiscoveryParam        = "INVALID_DISCOVERY_PARAM"
	CauseInvalidMsgFormat             = "INVALID_MSG_FORMAT"
	CauseInvalidQueryParam            = "INVALID_QUERY_PARAM"
	CauseMandatoryIEIncorrect         = "MANDATORY_IE_INCORRECT"
	CauseMandatoryIEMissing           = "MANDATORY_IE_MISSING"
	CauseMandatoryQueryParamIncorrect = "MANDATORY_QUERY_PARAM_INCORRECT"
	CauseMandatoryQueryParamMissing   = "MANDATORY_QUERY_PARAM_MISSING"
	CauseMissingAccessTokenInfo       = "MISSING_ACCESS_TOKEN_INFO"
	CauseModificationNotAllowed       = "MODIFICATION_NOT_ALLOWED"
	CauseNFCongestion                 = "NF_CONGESTION"
	CauseNFCongestionRisk             = "NF_CONGESTION_RISK"
	CauseNFDiscoveryFailure           = "NF_DISCOVERY_FAILURE"
	CauseNFFailover                   = "NF_FAILOVER"
	CauseNFServiceFailover            = "NF_SERVICE_FAILOVER"
	CauseOptionalIEIncorrect          = "OPTIONAL_IE_INCORRECT"
	CauseOptionalQueryParamIncorrect  = "OPTIONAL_QUERY_PARAM_INCORRECT"
	CauseResourceContextNotFound      = "RESOURCE_CONTEXT_NOT_FOUND"
	CauseResourceURIStructureNotFound = "RESOURCE_URI_STRUCTURE_NOT_FOUND"
	CauseSubscriptionNotFound         = "SUBSCRIPTION_NOT_FOUND"
	CauseSystemFailure                = "SYSTEM_FAILURE"
	CauseTargetNFNotReachable         = "TARGET_NF_NOT_REACHABLE"
	CauseTimedOutRequest              = "TIMED_OUT_REQUEST"
	CauseTokenCCAMismatch             = "TOKEN_CCA_MISMATCH"
	CauseUnspecifiedMsgFailure        = "UNSPECIFIED_MSG_FAILURE"
	CauseUnspecifiedNFFailure         = "UNSPECIFIED_NF_FAILURE"
)

func AccessTokenDenied(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusForbidden,
		Cause:  lo.ToPtr(CauseAccessTokenDenied),
		Title:  lo.ToPtr("Access token denied"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func CCAVerificationFailure(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusForbidden,
		Cause:  lo.ToPtr(CauseCCAVerificationFailure),
		Title:  lo.ToPtr("CCA verification failure"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func IncorrectLength(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusLengthRequired,
		Cause:  lo.ToPtr(CauseIncorrectLength),
		Title:  lo.ToPtr("Incorrect length"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func InsufficientResources(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusInternalServerError,
		Cause:  lo.ToPtr(CauseInsufficientResources),
		Title:  lo.ToPtr("Insufficient resources"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func InvalidAPI(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusBadRequest,
		Cause:  lo.ToPtr(CauseInvalidAPI),
		Title:  lo.ToPtr("Invalid API"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func InvalidDiscoveryParam(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusBadRequest,
		Cause:  lo.ToPtr(CauseInvalidDiscoveryParam),
		Title:  lo.ToPtr("Invalid discovery param"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func InvalidMsgFormat(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusBadRequest,
		Cause:  lo.ToPtr(CauseInvalidMsgFormat),
		Title:  lo.ToPtr("Invalid msg format"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func InvalidQueryParam(invalidParams []commondata.InvalidParam, detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status:        http.StatusBadRequest,
		Cause:         lo.ToPtr(CauseInvalidQueryParam),
		Title:         lo.ToPtr("Invalid query param"),
		InvalidParams: invalidParams,
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func MandatoryIEIncorrect(invalidParams []commondata.InvalidParam, detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status:        http.StatusBadRequest,
		Cause:         lo.ToPtr(CauseMandatoryIEIncorrect),
		Title:         lo.ToPtr("Mandatory IE incorrect"),
		InvalidParams: invalidParams,
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func MandatoryIEMissing(invalidParams []commondata.InvalidParam, detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status:        http.StatusBadRequest,
		Cause:         lo.ToPtr(CauseMandatoryIEMissing),
		Title:         lo.ToPtr("Mandatory IE missing"),
		InvalidParams: invalidParams,
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func MandatoryQueryParamIncorrect(invalidParams []commondata.InvalidParam, detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status:        http.StatusBadRequest,
		Cause:         lo.ToPtr(CauseMandatoryQueryParamIncorrect),
		Title:         lo.ToPtr("Mandatory query param incorrect"),
		InvalidParams: invalidParams,
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func MandatoryQueryParamMissing(invalidParams []commondata.InvalidParam, detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status:        http.StatusBadRequest,
		Cause:         lo.ToPtr(CauseMandatoryQueryParamMissing),
		Title:         lo.ToPtr("Mandatory query param missing"),
		InvalidParams: invalidParams,
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func MissingAccessTokenInfo(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusBadRequest,
		Cause:  lo.ToPtr(CauseMissingAccessTokenInfo),
		Title:  lo.ToPtr("Missing access token info"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func ModificationNotAllowed(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusForbidden,
		Cause:  lo.ToPtr(CauseModificationNotAllowed),
		Title:  lo.ToPtr("Modification not allowed"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func NFCongestion(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusServiceUnavailable,
		Cause:  lo.ToPtr(CauseNFCongestion),
		Title:  lo.ToPtr("NF congestion"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func NFCongestionRisk(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusTooManyRequests,
		Cause:  lo.ToPtr(CauseNFCongestionRisk),
		Title:  lo.ToPtr("NF congestion risk"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func NFDiscoveryFailure(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusBadRequest,
		Cause:  lo.ToPtr(CauseNFDiscoveryFailure),
		Title:  lo.ToPtr("NF discovery failure"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func NFFailover(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusInternalServerError,
		Cause:  lo.ToPtr(CauseNFFailover),
		Title:  lo.ToPtr("NF failover"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func NFServiceFailover(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusInternalServerError,
		Cause:  lo.ToPtr(CauseNFServiceFailover),
		Title:  lo.ToPtr("NF service failover"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func OptionalIEIncorrect(invalidParams []commondata.InvalidParam, detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status:        http.StatusBadRequest,
		Cause:         lo.ToPtr(CauseOptionalIEIncorrect),
		Title:         lo.ToPtr("Optional IE incorrect"),
		InvalidParams: invalidParams,
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func OptionalQueryParamIncorrect(invalidParams []commondata.InvalidParam, detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status:        http.StatusBadRequest,
		Cause:         lo.ToPtr(CauseOptionalQueryParamIncorrect),
		Title:         lo.ToPtr("Optional query param incorrect"),
		InvalidParams: invalidParams,
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func ResourceContextNotFound(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusBadRequest,
		Cause:  lo.ToPtr(CauseResourceContextNotFound),
		Title:  lo.ToPtr("Resource context not found"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func ResourceURIStructureNotFound(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusNotFound,
		Cause:  lo.ToPtr(CauseResourceURIStructureNotFound),
		Title:  lo.ToPtr("Resource URI structure not found"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func SubscriptionNotFound(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusNotFound,
		Cause:  lo.ToPtr(CauseSubscriptionNotFound),
		Title:  lo.ToPtr("Subscription not found"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func SystemFailure(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusInternalServerError,
		Cause:  lo.ToPtr(CauseSystemFailure),
		Title:  lo.ToPtr("System failure"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func TargetNFNotReachable(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusGatewayTimeout,
		Cause:  lo.ToPtr(CauseTargetNFNotReachable),
		Title:  lo.ToPtr("Target NF not reachable"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func TimedOutRequest(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusGatewayTimeout,
		Cause:  lo.ToPtr(CauseTimedOutRequest),
		Title:  lo.ToPtr("Timed out request"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func TokenCCAMismatch(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusForbidden,
		Cause:  lo.ToPtr(CauseTokenCCAMismatch),
		Title:  lo.ToPtr("Token CCA mismatch"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func UnspecifiedMsgFailure(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusBadRequest,
		Cause:  lo.ToPtr(CauseUnspecifiedMsgFailure),
		Title:  lo.ToPtr("Unspecified msg failure"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}

func UnspecifiedNFFailure(detail string) commondata.ProblemDetails {
	pd := commondata.ProblemDetails{
		Status: http.StatusInternalServerError,
		Cause:  lo.ToPtr(CauseUnspecifiedNFFailure),
		Title:  lo.ToPtr("Unspecified NF failure"),
	}
	if detail != "" {
		pd.Detail = &detail
	}
	return pd
}
