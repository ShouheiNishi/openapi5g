// Copyright 2023-2024 APRESIA Systems LTD.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generator

type pkgEntry struct {
	path    string
	cutRefs []string
}

var pkgList map[string]pkgEntry = map[string]pkgEntry{
	// Common Data Types
	"TS29571_CommonData.yaml": {"models", []string{
		"TS29510_Nnrf_AccessToken.yaml",  // To avoid import cycle
		"TS29510_Nnrf_NFManagement.yaml", // To avoid import cycle
		"TS29122_ResourceManagementOfBdt.yaml",
		"TS29515_Ngmlc_Location.yaml",
		"TS29517_Naf_EventExposure.yaml",
		"TS29520_Nnwdaf_AnalyticsInfo.yaml",
		"TS29520_Nnwdaf_EventsSubscription.yaml",
		"TS29522_IPTVConfiguration.yaml",
		"TS29522_ServiceParameter.yaml",
		"TS29544_Nspaf_SecuredPacket.yaml",
		"TS29572_Nlmf_Location.yaml",
		"TS32291_Nchf_ConvergedCharging.yaml",
	}},

	// // Other dependency
	"TS29122_CommonData.yaml": {"northbound/commondata", []string{
		"TS29554_Npcf_BDTPolicyControl.yaml", // To avoid import cycle
		"TS29572_Nlmf_Location.yaml",
	}},
	"TS29122_PfdManagement.yaml": {"pfd/management", nil},
	// "TS29122_ResourceManagementOfBdt.yaml": {"northbound/bdt", false, nil},
	// "TS29122_CpProvisioning.yaml":          {"northbound/provisioning", false, nil},
	// "TS29515_Ngmlc_Location.yaml": {"gmlc/location", false, []string{
	// 	"TS29518_Namf_Location.yaml",
	// }},
	// "TS29522_IPTVConfiguration.yaml": {"iptv", false, nil},
	// "TS29522_ServiceParameter.yaml":  {"service", false, nil},
	// "TS29517_Naf_EventExposure.yaml": {"af/event", false, []string{
	// 	"TS29503_Nudm_SDM.yaml",
	// 	"TS29508_Nsmf_EventExposure.yaml",
	// 	"TS29512_Npcf_SMPolicyControl.yaml",
	// 	"TS29514_Npcf_PolicyAuthorization.yaml",
	// 	"TS29523_Npcf_EventExposure.yaml",
	// }},
	// "TS29520_Nnwdaf_AnalyticsInfo.yaml": {"nwdaf/analytics", false, nil},
	// "TS29520_Nnwdaf_EventsSubscription.yaml": {"nwdaf/event", false, []string{
	// 	"TS29503_Nudm_PP.yaml",
	// 	"TS29503_Nudm_SDM.yaml",
	// 	"TS29508_Nsmf_EventExposure.yaml",
	// 	"TS29510_Nnrf_NFManagement.yaml",
	// 	"TS29512_Npcf_SMPolicyControl.yaml",
	// 	"TS29514_Npcf_PolicyAuthorization.yaml",
	// 	"TS29517_Naf_EventExposure.yaml",
	// 	"TS29523_Npcf_EventExposure.yaml",
	// 	"TS29531_Nnssf_NSSelection.yaml",
	// }},
	"TS29521_Nbsf_Management.yaml":  {"bsf/management", nil},
	"TS29522_TrafficInfluence.yaml": {"influence", nil},
	// "TS29544_Nspaf_SecuredPacket.yaml": {"spaf/secure", false, []string{
	// 	"TS29503_Nudm_SDM.yaml",
	// }},
	"TS29551_Nnef_PFDmanagement.yaml": {"nef/management", nil},
	// "TS29572_Nlmf_Location.yaml": {"lmf/location", false, []string{
	// 	"TS29518_Namf_EventExposure.yaml",
	// }},
	// "TS32291_Nchf_ConvergedCharging.yaml": {"chf/charging", false, []string{
	// 	"TS29512_Npcf_SMPolicyControl.yaml",
	// }},

	// free5GC modules
	"TS29518_Namf_Communication.yaml": {"amf/communication", []string{
		"TS29502_Nsmf_PDUSession.yaml", // To avoid import cycle
		"TS29572_Nlmf_Location.yaml",
	}},
	"TS29518_Namf_EventExposure.yaml": {"amf/event", []string{
		"TS29503_Nudm_EE.yaml", // To avoid import cycle
	}},
	"TS29518_Namf_Location.yaml": {"amf/location", []string{
		"TS29515_Ngmlc_Location.yaml",
		"TS29572_Nlmf_Location.yaml",
	}},
	"TS29518_Namf_MT.yaml": {"amf/mt", nil},

	"TS29509_Nausf_SoRProtection.yaml":    {"ausf/sor", nil},
	"TS29509_Nausf_UEAuthentication.yaml": {"ausf/authentication", nil},
	"TS29509_Nausf_UPUProtection.yaml": {"ausf/upu", []string{
		"TS29544_Nspaf_SecuredPacket.yaml",
	}},

	"TS29510_Nnrf_AccessToken.yaml":   {"nrf/token", nil},
	"TS29510_Nnrf_Bootstrapping.yaml": {"nrf/bootstrapping", nil},
	"TS29510_Nnrf_NFDiscovery.yaml": {"nrf/discovery", []string{
		"TS29520_Nnwdaf_AnalyticsInfo.yaml",
		"TS29520_Nnwdaf_EventsSubscription.yaml",
		"TS29572_Nlmf_Location.yaml",
	}},
	"TS29510_Nnrf_NFManagement.yaml": {"nrf/management", []string{
		"TS29517_Naf_EventExposure.yaml",
		"TS29518_Namf_Communication.yaml", // To avoid import cycle
		"TS29520_Nnwdaf_AnalyticsInfo.yaml",
		"TS29520_Nnwdaf_EventsSubscription.yaml",
		"TS29572_Nlmf_Location.yaml",
	}},

	"TS29531_Nnssf_NSSAIAvailability.yaml": {"nssf/availability", nil},
	"TS29531_Nnssf_NSSelection.yaml":       {"nssf/selection", nil},

	"TS29507_Npcf_AMPolicyControl.yaml": {"pcf/AMpolicy", nil},
	"TS29512_Npcf_SMPolicyControl.yaml": {"pcf/SMpolicy", []string{
		"TS32291_Nchf_ConvergedCharging.yaml",
	}},
	"TS29514_Npcf_PolicyAuthorization.yaml": {"pcf/authorization", []string{
		"TS29512_Npcf_SMPolicyControl.yaml", // To avoid import cycle
		"TS32291_Nchf_ConvergedCharging.yaml",
	}},
	"TS29523_Npcf_EventExposure.yaml":   {"pcf/event", nil},
	"TS29525_Npcf_UEPolicyControl.yaml": {"pcf/UEpolicy", nil},
	"TS29554_Npcf_BDTPolicyControl.yaml": {"pcf/BDTpolicy", []string{
		"TS29122_ResourceManagementOfBdt.yaml",
	}},

	"TS29502_Nsmf_PDUSession.yaml": {"smf/session", []string{
		"TS29512_Npcf_SMPolicyControl.yaml", // To avoid import cycle
		"TS32291_Nchf_ConvergedCharging.yaml",
	}},
	"TS29508_Nsmf_EventExposure.yaml": {"smf/event", nil},

	"TS29503_Nudm_EE.yaml": {"udm/ee", nil},
	"TS29503_Nudm_MT.yaml": {"udm/mt", []string{
		"TS29572_Nlmf_Location.yaml",
	}},
	"TS29503_Nudm_NIDDAU.yaml": {"udm/niddau", nil},
	"TS29503_Nudm_PP.yaml": {"udm/pp", []string{
		"TS29503_Nudm_SDM.yaml", // To avoid import cycle
		"TS29572_Nlmf_Location.yaml",
	}},
	"TS29503_Nudm_SDM.yaml": {"udm/sdm", []string{
		"TS29503_Nudm_UECM.yaml", // To avoid import cycle
		"TS29572_Nlmf_Location.yaml",
	}},
	"TS29503_Nudm_UEAU.yaml": {"udm/ueau", nil},
	"TS29503_Nudm_UECM.yaml": {"udm/uecm", nil},

	"TS29504_Nudr_DR.yaml": {"udr/dr", []string{
		"TS29522_IPTVConfiguration.yaml",
		"TS29522_ServiceParameter.yaml",
	}},
	"TS29504_Nudr_GroupIDmap.yaml": {"udr/idmap", nil},
	"TS29505_Subscription_Data.yaml": {"udr/subscription", []string{
		"TS29572_Nlmf_Location.yaml",
	}},
	"TS29519_Application_Data.yaml": {"udr/application", []string{
		"TS29522_IPTVConfiguration.yaml",
		"TS29522_ServiceParameter.yaml",
	}},
	"TS29519_Exposure_Data.yaml": {"udr/exposure", nil},
	"TS29519_Policy_Data.yaml": {"udr/policy", []string{
		"TS29122_ResourceManagementOfBdt.yaml",
		"TS29505_Subscription_Data.yaml",    // To avoid import cycle
		"TS29512_Npcf_SMPolicyControl.yaml", // To avoid import cycle
	}},
}
