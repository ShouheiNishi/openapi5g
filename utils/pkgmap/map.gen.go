// Code generated by github.com/ShouheiNishi/openapi5g/internal/generator, DO NOT EDIT.

package pkgmap

import (
	"github.com/getkin/kin-openapi/openapi3"

	amf_communication "github.com/ShouheiNishi/openapi5g/amf/communication"
	amf_event "github.com/ShouheiNishi/openapi5g/amf/event"
	amf_location "github.com/ShouheiNishi/openapi5g/amf/location"
	amf_mt "github.com/ShouheiNishi/openapi5g/amf/mt"
	ausf_authentication "github.com/ShouheiNishi/openapi5g/ausf/authentication"
	ausf_sor "github.com/ShouheiNishi/openapi5g/ausf/sor"
	ausf_upu "github.com/ShouheiNishi/openapi5g/ausf/upu"
	bsf_management "github.com/ShouheiNishi/openapi5g/bsf/management"
	influence "github.com/ShouheiNishi/openapi5g/influence"
	models "github.com/ShouheiNishi/openapi5g/models"
	nef_management "github.com/ShouheiNishi/openapi5g/nef/management"
	northbound_commondata "github.com/ShouheiNishi/openapi5g/northbound/commondata"
	nrf_bootstrapping "github.com/ShouheiNishi/openapi5g/nrf/bootstrapping"
	nrf_discovery "github.com/ShouheiNishi/openapi5g/nrf/discovery"
	nrf_management "github.com/ShouheiNishi/openapi5g/nrf/management"
	nrf_token "github.com/ShouheiNishi/openapi5g/nrf/token"
	nssf_availability "github.com/ShouheiNishi/openapi5g/nssf/availability"
	nssf_selection "github.com/ShouheiNishi/openapi5g/nssf/selection"
	pcf_AMpolicy "github.com/ShouheiNishi/openapi5g/pcf/AMpolicy"
	pcf_BDTpolicy "github.com/ShouheiNishi/openapi5g/pcf/BDTpolicy"
	pcf_SMpolicy "github.com/ShouheiNishi/openapi5g/pcf/SMpolicy"
	pcf_UEpolicy "github.com/ShouheiNishi/openapi5g/pcf/UEpolicy"
	pcf_authorization "github.com/ShouheiNishi/openapi5g/pcf/authorization"
	pcf_event "github.com/ShouheiNishi/openapi5g/pcf/event"
	pfd_management "github.com/ShouheiNishi/openapi5g/pfd/management"
	smf_event "github.com/ShouheiNishi/openapi5g/smf/event"
	smf_session "github.com/ShouheiNishi/openapi5g/smf/session"
	udm_ee "github.com/ShouheiNishi/openapi5g/udm/ee"
	udm_mt "github.com/ShouheiNishi/openapi5g/udm/mt"
	udm_niddau "github.com/ShouheiNishi/openapi5g/udm/niddau"
	udm_pp "github.com/ShouheiNishi/openapi5g/udm/pp"
	udm_sdm "github.com/ShouheiNishi/openapi5g/udm/sdm"
	udm_ueau "github.com/ShouheiNishi/openapi5g/udm/ueau"
	udm_uecm "github.com/ShouheiNishi/openapi5g/udm/uecm"
	udr_application "github.com/ShouheiNishi/openapi5g/udr/application"
	udr_dr "github.com/ShouheiNishi/openapi5g/udr/dr"
	udr_exposure "github.com/ShouheiNishi/openapi5g/udr/exposure"
	udr_idmap "github.com/ShouheiNishi/openapi5g/udr/idmap"
	udr_policy "github.com/ShouheiNishi/openapi5g/udr/policy"
	udr_subscription "github.com/ShouheiNishi/openapi5g/udr/subscription"
)

var s2p = map[string]string{
	"TS29122_CommonData.yaml":               "github.com/ShouheiNishi/openapi5g/northbound/commondata",
	"TS29122_PfdManagement.yaml":            "github.com/ShouheiNishi/openapi5g/pfd/management",
	"TS29502_Nsmf_PDUSession.yaml":          "github.com/ShouheiNishi/openapi5g/smf/session",
	"TS29503_Nudm_EE.yaml":                  "github.com/ShouheiNishi/openapi5g/udm/ee",
	"TS29503_Nudm_MT.yaml":                  "github.com/ShouheiNishi/openapi5g/udm/mt",
	"TS29503_Nudm_NIDDAU.yaml":              "github.com/ShouheiNishi/openapi5g/udm/niddau",
	"TS29503_Nudm_PP.yaml":                  "github.com/ShouheiNishi/openapi5g/udm/pp",
	"TS29503_Nudm_SDM.yaml":                 "github.com/ShouheiNishi/openapi5g/udm/sdm",
	"TS29503_Nudm_UEAU.yaml":                "github.com/ShouheiNishi/openapi5g/udm/ueau",
	"TS29503_Nudm_UECM.yaml":                "github.com/ShouheiNishi/openapi5g/udm/uecm",
	"TS29504_Nudr_DR.yaml":                  "github.com/ShouheiNishi/openapi5g/udr/dr",
	"TS29504_Nudr_GroupIDmap.yaml":          "github.com/ShouheiNishi/openapi5g/udr/idmap",
	"TS29505_Subscription_Data.yaml":        "github.com/ShouheiNishi/openapi5g/udr/subscription",
	"TS29507_Npcf_AMPolicyControl.yaml":     "github.com/ShouheiNishi/openapi5g/pcf/AMpolicy",
	"TS29508_Nsmf_EventExposure.yaml":       "github.com/ShouheiNishi/openapi5g/smf/event",
	"TS29509_Nausf_SoRProtection.yaml":      "github.com/ShouheiNishi/openapi5g/ausf/sor",
	"TS29509_Nausf_UEAuthentication.yaml":   "github.com/ShouheiNishi/openapi5g/ausf/authentication",
	"TS29509_Nausf_UPUProtection.yaml":      "github.com/ShouheiNishi/openapi5g/ausf/upu",
	"TS29510_Nnrf_AccessToken.yaml":         "github.com/ShouheiNishi/openapi5g/nrf/token",
	"TS29510_Nnrf_Bootstrapping.yaml":       "github.com/ShouheiNishi/openapi5g/nrf/bootstrapping",
	"TS29510_Nnrf_NFDiscovery.yaml":         "github.com/ShouheiNishi/openapi5g/nrf/discovery",
	"TS29510_Nnrf_NFManagement.yaml":        "github.com/ShouheiNishi/openapi5g/nrf/management",
	"TS29512_Npcf_SMPolicyControl.yaml":     "github.com/ShouheiNishi/openapi5g/pcf/SMpolicy",
	"TS29514_Npcf_PolicyAuthorization.yaml": "github.com/ShouheiNishi/openapi5g/pcf/authorization",
	"TS29518_Namf_Communication.yaml":       "github.com/ShouheiNishi/openapi5g/amf/communication",
	"TS29518_Namf_EventExposure.yaml":       "github.com/ShouheiNishi/openapi5g/amf/event",
	"TS29518_Namf_Location.yaml":            "github.com/ShouheiNishi/openapi5g/amf/location",
	"TS29518_Namf_MT.yaml":                  "github.com/ShouheiNishi/openapi5g/amf/mt",
	"TS29519_Application_Data.yaml":         "github.com/ShouheiNishi/openapi5g/udr/application",
	"TS29519_Exposure_Data.yaml":            "github.com/ShouheiNishi/openapi5g/udr/exposure",
	"TS29519_Policy_Data.yaml":              "github.com/ShouheiNishi/openapi5g/udr/policy",
	"TS29521_Nbsf_Management.yaml":          "github.com/ShouheiNishi/openapi5g/bsf/management",
	"TS29522_TrafficInfluence.yaml":         "github.com/ShouheiNishi/openapi5g/influence",
	"TS29523_Npcf_EventExposure.yaml":       "github.com/ShouheiNishi/openapi5g/pcf/event",
	"TS29525_Npcf_UEPolicyControl.yaml":     "github.com/ShouheiNishi/openapi5g/pcf/UEpolicy",
	"TS29531_Nnssf_NSSAIAvailability.yaml":  "github.com/ShouheiNishi/openapi5g/nssf/availability",
	"TS29531_Nnssf_NSSelection.yaml":        "github.com/ShouheiNishi/openapi5g/nssf/selection",
	"TS29551_Nnef_PFDmanagement.yaml":       "github.com/ShouheiNishi/openapi5g/nef/management",
	"TS29554_Npcf_BDTPolicyControl.yaml":    "github.com/ShouheiNishi/openapi5g/pcf/BDTpolicy",
	"TS29571_CommonData.yaml":               "github.com/ShouheiNishi/openapi5g/models",
}

var p2s = map[string]string{
	"github.com/ShouheiNishi/openapi5g/amf/communication":     "TS29518_Namf_Communication.yaml",
	"github.com/ShouheiNishi/openapi5g/amf/event":             "TS29518_Namf_EventExposure.yaml",
	"github.com/ShouheiNishi/openapi5g/amf/location":          "TS29518_Namf_Location.yaml",
	"github.com/ShouheiNishi/openapi5g/amf/mt":                "TS29518_Namf_MT.yaml",
	"github.com/ShouheiNishi/openapi5g/ausf/authentication":   "TS29509_Nausf_UEAuthentication.yaml",
	"github.com/ShouheiNishi/openapi5g/ausf/sor":              "TS29509_Nausf_SoRProtection.yaml",
	"github.com/ShouheiNishi/openapi5g/ausf/upu":              "TS29509_Nausf_UPUProtection.yaml",
	"github.com/ShouheiNishi/openapi5g/bsf/management":        "TS29521_Nbsf_Management.yaml",
	"github.com/ShouheiNishi/openapi5g/influence":             "TS29522_TrafficInfluence.yaml",
	"github.com/ShouheiNishi/openapi5g/models":                "TS29571_CommonData.yaml",
	"github.com/ShouheiNishi/openapi5g/nef/management":        "TS29551_Nnef_PFDmanagement.yaml",
	"github.com/ShouheiNishi/openapi5g/northbound/commondata": "TS29122_CommonData.yaml",
	"github.com/ShouheiNishi/openapi5g/nrf/bootstrapping":     "TS29510_Nnrf_Bootstrapping.yaml",
	"github.com/ShouheiNishi/openapi5g/nrf/discovery":         "TS29510_Nnrf_NFDiscovery.yaml",
	"github.com/ShouheiNishi/openapi5g/nrf/management":        "TS29510_Nnrf_NFManagement.yaml",
	"github.com/ShouheiNishi/openapi5g/nrf/token":             "TS29510_Nnrf_AccessToken.yaml",
	"github.com/ShouheiNishi/openapi5g/nssf/availability":     "TS29531_Nnssf_NSSAIAvailability.yaml",
	"github.com/ShouheiNishi/openapi5g/nssf/selection":        "TS29531_Nnssf_NSSelection.yaml",
	"github.com/ShouheiNishi/openapi5g/pcf/AMpolicy":          "TS29507_Npcf_AMPolicyControl.yaml",
	"github.com/ShouheiNishi/openapi5g/pcf/BDTpolicy":         "TS29554_Npcf_BDTPolicyControl.yaml",
	"github.com/ShouheiNishi/openapi5g/pcf/SMpolicy":          "TS29512_Npcf_SMPolicyControl.yaml",
	"github.com/ShouheiNishi/openapi5g/pcf/UEpolicy":          "TS29525_Npcf_UEPolicyControl.yaml",
	"github.com/ShouheiNishi/openapi5g/pcf/authorization":     "TS29514_Npcf_PolicyAuthorization.yaml",
	"github.com/ShouheiNishi/openapi5g/pcf/event":             "TS29523_Npcf_EventExposure.yaml",
	"github.com/ShouheiNishi/openapi5g/pfd/management":        "TS29122_PfdManagement.yaml",
	"github.com/ShouheiNishi/openapi5g/smf/event":             "TS29508_Nsmf_EventExposure.yaml",
	"github.com/ShouheiNishi/openapi5g/smf/session":           "TS29502_Nsmf_PDUSession.yaml",
	"github.com/ShouheiNishi/openapi5g/udm/ee":                "TS29503_Nudm_EE.yaml",
	"github.com/ShouheiNishi/openapi5g/udm/mt":                "TS29503_Nudm_MT.yaml",
	"github.com/ShouheiNishi/openapi5g/udm/niddau":            "TS29503_Nudm_NIDDAU.yaml",
	"github.com/ShouheiNishi/openapi5g/udm/pp":                "TS29503_Nudm_PP.yaml",
	"github.com/ShouheiNishi/openapi5g/udm/sdm":               "TS29503_Nudm_SDM.yaml",
	"github.com/ShouheiNishi/openapi5g/udm/ueau":              "TS29503_Nudm_UEAU.yaml",
	"github.com/ShouheiNishi/openapi5g/udm/uecm":              "TS29503_Nudm_UECM.yaml",
	"github.com/ShouheiNishi/openapi5g/udr/application":       "TS29519_Application_Data.yaml",
	"github.com/ShouheiNishi/openapi5g/udr/dr":                "TS29504_Nudr_DR.yaml",
	"github.com/ShouheiNishi/openapi5g/udr/exposure":          "TS29519_Exposure_Data.yaml",
	"github.com/ShouheiNishi/openapi5g/udr/idmap":             "TS29504_Nudr_GroupIDmap.yaml",
	"github.com/ShouheiNishi/openapi5g/udr/policy":            "TS29519_Policy_Data.yaml",
	"github.com/ShouheiNishi/openapi5g/udr/subscription":      "TS29505_Subscription_Data.yaml",
}

var p2l = map[string]func() (*openapi3.T, error){
	"github.com/ShouheiNishi/openapi5g/amf/communication":     amf_communication.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/amf/event":             amf_event.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/amf/location":          amf_location.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/amf/mt":                amf_mt.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/ausf/authentication":   ausf_authentication.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/ausf/sor":              ausf_sor.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/ausf/upu":              ausf_upu.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/bsf/management":        bsf_management.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/influence":             influence.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/models":                models.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/nef/management":        nef_management.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/northbound/commondata": northbound_commondata.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/nrf/bootstrapping":     nrf_bootstrapping.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/nrf/discovery":         nrf_discovery.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/nrf/management":        nrf_management.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/nrf/token":             nrf_token.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/nssf/availability":     nssf_availability.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/nssf/selection":        nssf_selection.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/pcf/AMpolicy":          pcf_AMpolicy.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/pcf/BDTpolicy":         pcf_BDTpolicy.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/pcf/SMpolicy":          pcf_SMpolicy.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/pcf/UEpolicy":          pcf_UEpolicy.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/pcf/authorization":     pcf_authorization.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/pcf/event":             pcf_event.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/pfd/management":        pfd_management.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/smf/event":             smf_event.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/smf/session":           smf_session.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/udm/ee":                udm_ee.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/udm/mt":                udm_mt.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/udm/niddau":            udm_niddau.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/udm/pp":                udm_pp.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/udm/sdm":               udm_sdm.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/udm/ueau":              udm_ueau.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/udm/uecm":              udm_uecm.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/udr/application":       udr_application.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/udr/dr":                udr_dr.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/udr/exposure":          udr_exposure.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/udr/idmap":             udr_idmap.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/udr/policy":            udr_policy.GetKinOpenApi3Document,
	"github.com/ShouheiNishi/openapi5g/udr/subscription":      udr_subscription.GetKinOpenApi3Document,
}
