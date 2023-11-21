// Code generated by github.com/ShouheiNishi/openapi5g/internal/generator, DO NOT EDIT.

package management

import (
	"fmt"
	"net/url"
	"path"
	"sync"

	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29122_CommonData"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29122_CpProvisioning"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29122_ResourceManagementOfBdt"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29502_Nsmf_PDUSession"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29503_Nudm_EE"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29503_Nudm_NIDDAU"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29503_Nudm_PP"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29503_Nudm_SDM"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29503_Nudm_UEAU"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29503_Nudm_UECM"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29505_Subscription_Data"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29507_Npcf_AMPolicyControl"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29508_Nsmf_EventExposure"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29509_Nausf_SoRProtection"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29509_Nausf_UPUProtection"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29510_Nnrf_AccessToken"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29510_Nnrf_NFManagement"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29512_Npcf_SMPolicyControl"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29514_Npcf_PolicyAuthorization"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29517_Naf_EventExposure"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29518_Namf_Communication"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29518_Namf_EventExposure"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29519_Policy_Data"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29520_Nnwdaf_AnalyticsInfo"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29520_Nnwdaf_EventsSubscription"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29522_TrafficInfluence"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29523_Npcf_EventExposure"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29531_Nnssf_NSSelection"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29544_Nspaf_SecuredPacket"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29554_Npcf_BDTPolicyControl"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29571_CommonData"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS29572_Nlmf_Location"
	"github.com/ShouheiNishi/openapi5g/internal/embed/TS32291_Nchf_ConvergedCharging"
	"github.com/getkin/kin-openapi/openapi3"
)

var specTable map[string][]byte = map[string][]byte{
	"TS29122_CommonData.yaml":                TS29122_CommonData.SpecYaml,
	"TS29122_CpProvisioning.yaml":            TS29122_CpProvisioning.SpecYaml,
	"TS29122_ResourceManagementOfBdt.yaml":   TS29122_ResourceManagementOfBdt.SpecYaml,
	"TS29502_Nsmf_PDUSession.yaml":           TS29502_Nsmf_PDUSession.SpecYaml,
	"TS29503_Nudm_EE.yaml":                   TS29503_Nudm_EE.SpecYaml,
	"TS29503_Nudm_NIDDAU.yaml":               TS29503_Nudm_NIDDAU.SpecYaml,
	"TS29503_Nudm_PP.yaml":                   TS29503_Nudm_PP.SpecYaml,
	"TS29503_Nudm_SDM.yaml":                  TS29503_Nudm_SDM.SpecYaml,
	"TS29503_Nudm_UEAU.yaml":                 TS29503_Nudm_UEAU.SpecYaml,
	"TS29503_Nudm_UECM.yaml":                 TS29503_Nudm_UECM.SpecYaml,
	"TS29505_Subscription_Data.yaml":         TS29505_Subscription_Data.SpecYaml,
	"TS29507_Npcf_AMPolicyControl.yaml":      TS29507_Npcf_AMPolicyControl.SpecYaml,
	"TS29508_Nsmf_EventExposure.yaml":        TS29508_Nsmf_EventExposure.SpecYaml,
	"TS29509_Nausf_SoRProtection.yaml":       TS29509_Nausf_SoRProtection.SpecYaml,
	"TS29509_Nausf_UPUProtection.yaml":       TS29509_Nausf_UPUProtection.SpecYaml,
	"TS29510_Nnrf_AccessToken.yaml":          TS29510_Nnrf_AccessToken.SpecYaml,
	"TS29510_Nnrf_NFManagement.yaml":         TS29510_Nnrf_NFManagement.SpecYaml,
	"TS29512_Npcf_SMPolicyControl.yaml":      TS29512_Npcf_SMPolicyControl.SpecYaml,
	"TS29514_Npcf_PolicyAuthorization.yaml":  TS29514_Npcf_PolicyAuthorization.SpecYaml,
	"TS29517_Naf_EventExposure.yaml":         TS29517_Naf_EventExposure.SpecYaml,
	"TS29518_Namf_Communication.yaml":        TS29518_Namf_Communication.SpecYaml,
	"TS29518_Namf_EventExposure.yaml":        TS29518_Namf_EventExposure.SpecYaml,
	"TS29519_Policy_Data.yaml":               TS29519_Policy_Data.SpecYaml,
	"TS29520_Nnwdaf_AnalyticsInfo.yaml":      TS29520_Nnwdaf_AnalyticsInfo.SpecYaml,
	"TS29520_Nnwdaf_EventsSubscription.yaml": TS29520_Nnwdaf_EventsSubscription.SpecYaml,
	"TS29522_TrafficInfluence.yaml":          TS29522_TrafficInfluence.SpecYaml,
	"TS29523_Npcf_EventExposure.yaml":        TS29523_Npcf_EventExposure.SpecYaml,
	"TS29531_Nnssf_NSSelection.yaml":         TS29531_Nnssf_NSSelection.SpecYaml,
	"TS29544_Nspaf_SecuredPacket.yaml":       TS29544_Nspaf_SecuredPacket.SpecYaml,
	"TS29554_Npcf_BDTPolicyControl.yaml":     TS29554_Npcf_BDTPolicyControl.SpecYaml,
	"TS29571_CommonData.yaml":                TS29571_CommonData.SpecYaml,
	"TS29572_Nlmf_Location.yaml":             TS29572_Nlmf_Location.SpecYaml,
	"TS32291_Nchf_ConvergedCharging.yaml":    TS32291_Nchf_ConvergedCharging.SpecYaml,
}

var kinOpenApi3Doc *openapi3.T
var kinOpenApi3Err error
var kinOpenApi3Once sync.Once

func GetKinOpenApi3Document() (*openapi3.T, error) {
	kinOpenApi3Once.Do(func() {
		loader := openapi3.NewLoader()
		loader.IsExternalRefsAllowed = true
		loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
			spec := specTable[path.Base(url.Path)]
			if spec == nil {
				return nil, fmt.Errorf("%s is missing", url.String())
			}
			return spec, nil
		}
		kinOpenApi3Doc, kinOpenApi3Err = loader.LoadFromFile("TS29510_Nnrf_NFManagement.yaml")
	})
	return kinOpenApi3Doc, kinOpenApi3Err
}

func GetKinOpenApi3DocumentMust() *openapi3.T {
	doc, err := GetKinOpenApi3Document()
	if err != nil {
		panic(err)
	}
	return doc
}
