// Copyright 2023 APRESIA Systems LTD.
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

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/invopop/yaml"
	"github.com/mohae/deepcopy"
)

const nameForModels = "f5gcModels"
const importForModels = "github.com/free5gc/openapi/models"

const maxSafeInteger = (1<<53 - 1)
const minSafeInteger = -maxSafeInteger

func RewriteYaml(rootDir string, spec string, doc *openapi3.T) (outLists []string, deps []string, err error) {
	if err := fixLocalRefInRemoteRef(doc, spec); err != nil {
		return nil, nil, err
	}

	refs := map[string]struct{}{}
	cutRefs := map[string]struct{}{}
	for _, s := range pkgList[spec].cutRefs {
		cutRefs[s] = struct{}{}
	}
	if err := scanRef(doc, spec, refs, cutRefs); err != nil {
		return nil, nil, err
	}

	if spec == "TS29571_CommonData.yaml" {
		schema := doc.Components.Schemas["ProblemDetails"].Value.Properties["status"].Value
		if schema.Extensions == nil {
			schema.Extensions = make(map[string]interface{})
		}
		schema.Extensions["x-go-type-skip-optional-pointer"] = true
	}

	if spec == "TS29571_CommonData.yaml" {
		for statusCode, res := range doc.Components.Responses {
			if statusCode == "default" {
				res.Value.Content = openapi3.NewContent()
				res.Value.Content["application/problem+json"] = openapi3.NewMediaType().WithSchemaRef(
					openapi3.NewSchemaRef("#/components/schemas/ProblemDetails", nil),
				)
			}
		}
	} else {
		for _, pathItem := range doc.Paths {
			if pathItem.Ref != "" {
				continue
			}
			for _, op := range pathItem.Operations() {
				if op.Responses == nil {
					op.Responses = make(openapi3.Responses)
				}
				if op.Responses["default"] == nil {
					op.Responses["default"] = &openapi3.ResponseRef{Value: openapi3.NewResponse()}
				}
				if op.Responses["default"].Ref != "TS29571_CommonData.yaml#/components/responses/default" {
					if op.Responses["default"].Value.Content == nil {
						op.Responses["default"].Value.Content = openapi3.NewContent()
					}
					if op.Responses["default"].Value.Content["application/problem+json"] == nil {
						resNew := deepcopy.Copy(op.Responses["default"]).(*openapi3.ResponseRef)
						resNew.Ref = ""
						resNew.Value.Content["application/problem+json"] = openapi3.NewMediaType().WithSchemaRef(
							openapi3.NewSchemaRef("TS29571_CommonData.yaml#/components/schemas/ProblemDetails", nil),
						)
						op.Responses["default"] = resNew
					}
				}
			}
		}
	}

	if spec == "TS29503_Nudm_SDM.yaml" {
		for _, parameterRef := range doc.Paths["/shared-data"].Get.Parameters {
			parameter := parameterRef.Value
			if parameter.Name == "supportedFeatures" {
				if parameter.Extensions == nil {
					parameter.Extensions = make(map[string]interface{})
				}
				parameter.Extensions["x-go-name"] = "SupportedFeaturesShouldNotBeUsed"
			}
		}
	}

	if spec == "TS29571_CommonData.yaml" {
		schema := doc.Components.Schemas["ExtSnssai"].Value
		for i := range schema.AllOf {
			if strings.HasSuffix(schema.AllOf[i].Ref, "/Snssai") {
				newSchema := *schema.AllOf[i].Value
				if newSchema.Extensions != nil {
					newSchema.Extensions = deepcopy.Copy(newSchema.Extensions).(map[string]interface{})
				}
				schema.AllOf[i] = &openapi3.SchemaRef{
					Value: &newSchema,
				}
			}
		}

		schema = doc.Components.Schemas["Snssai"].Value
		if schema.Extensions == nil {
			schema.Extensions = make(map[string]interface{})
		}
		schema.Extensions["x-go-type"] = nameForModels + ".Snssai"
		schema.Extensions["x-go-type-import"] = map[string]interface{}{
			"name": nameForModels,
			"path": importForModels,
		}

		schema = doc.Components.Schemas["PlmnId"].Value
		if schema.Extensions == nil {
			schema.Extensions = make(map[string]interface{})
		}
		schema.Extensions["x-go-type"] = nameForModels + ".PlmnId"
		schema.Extensions["x-go-type-import"] = map[string]interface{}{
			"name": nameForModels,
			"path": importForModels,
		}

		schema = doc.Components.Schemas["Guami"].Value
		if schema.Extensions == nil {
			schema.Extensions = make(map[string]interface{})
		}
		schema.Extensions["x-go-type"] = nameForModels + ".Guami"
		schema.Extensions["x-go-type-import"] = map[string]interface{}{
			"name": nameForModels,
			"path": importForModels,
		}

		for _, t := range []string{"Int32", "Int64", "Uint16", "Uint32", "Uint64"} {
			doc.Components.Schemas[t].Value.Format = strings.ToLower(t)
			doc.Components.Schemas[t+"Rm"].Value.Format = strings.ToLower(t)
		}
		doc.Components.Schemas["Uinteger"].Value.Format = "uint"
		doc.Components.Schemas["UintegerRm"].Value.Format = "uint"
	}

	if err := os.MkdirAll(filepath.Join(rootDir, "modSpecs"), 0o755); err != nil {
		return nil, nil, err
	}
	if f, err := os.Create(filepath.Join(rootDir, "modSpecs", spec)); err != nil {
		return nil, nil, err
	} else {
		if buf, err := yaml.Marshal(doc); err != nil {
			return nil, nil, err
		} else {
			if _, err := f.Write(buf); err != nil {
				return nil, nil, err
			}
			if err := f.Close(); err != nil {
				return nil, nil, err
			}
		}
	}
	outLists = append(outLists, filepath.Join(rootDir, "modSpecs", spec))

	deps = make([]string, 0, len(refs))
	for r := range refs {
		if _, exist := pkgList[r]; !exist {
			if _, exist := cutRefs[r]; !exist {
				panic(fmt.Sprintf("%s is not defined.", r))
			}
		}
		deps = append(deps, r)
	}
	sort.Strings(deps)

	return outLists, deps, nil
}

func fixAnyOfEnum(v *openapi3.Schema) error {
	if len(v.AnyOf) == 2 && v.AnyOf[0].Ref == "" && v.AnyOf[1].Ref == "" {
		v0 := v.AnyOf[0].Value
		v1 := v.AnyOf[1].Value
		// v0.
		if v0.Type == v1.Type && v0.Format == v1.Format &&
			(v0.Type == openapi3.TypeInteger || v0.Type == openapi3.TypeString) {
			newDescription := v.Description
			if newDescription == "" {
				newDescription = v0.Description
			}
			if newDescription == "" {
				newDescription = v1.Description
			}
			merged := false
			if v0.Enum == nil && len(v1.Enum) > 0 {
				*v = *v1
				v.Description = newDescription
				merged = true
			} else if v1.Enum == nil && len(v0.Enum) > 0 {
				*v = *v0
				v.Description = newDescription
				merged = true
			}
			if merged {
				newSkipOptionalPointer := false
				if b, ok := v0.Extensions["x-go-type-skip-optional-pointer"].(bool); b && ok {
					if b, ok := v1.Extensions["x-go-type-skip-optional-pointer"].(bool); b && ok {
						newSkipOptionalPointer = true
					}
				}
				if newSkipOptionalPointer {
					if v.Extensions == nil {
						v.Extensions = make(map[string]interface{})
					}
					v.Extensions["x-go-type-skip-optional-pointer"] = true
				} else {
					delete(v.Extensions, "x-go-type-skip-optional-pointer")
				}
			}
		}
	}
	return nil
}

func fixAnyOfString(v *openapi3.Schema) error {
	if len(v.AnyOf) > 0 {
		for _, vRef := range v.AnyOf {
			if vRef.Value.Type != openapi3.TypeString {
				return nil
			}
		}
		newDescription := []string{"Merged type of"}
		newSkipOptionalPointer := true
		for _, vRef := range v.AnyOf {
			if vRef.Ref == "" {
				if vRef.Value.Description == "" {
					newDescription = append(newDescription, "  Anonymous string")
				} else {
					newDescription = append(newDescription, "  "+vRef.Value.Description)
				}
			} else {
				if vRef.Value.Description == "" {
					newDescription = append(newDescription, "  string in "+vRef.Ref)
				} else {
					newDescription = append(newDescription, "  "+vRef.Value.Description+" in "+vRef.Ref)
				}
			}
			if b, ok := vRef.Value.Extensions["x-go-type-skip-optional-pointer"].(bool); !ok || !b {
				newSkipOptionalPointer = false
			}
		}
		*v = openapi3.Schema{
			Type:        openapi3.TypeString,
			Description: strings.Join(newDescription, "\n"),
		}
		if newSkipOptionalPointer {
			v.Extensions = map[string]interface{}{
				"x-go-type-skip-optional-pointer": true,
			}
		}
	}
	return nil
}

func fixImplicitArray(v *openapi3.Schema) error {
	if v.Type == "" && v.Items != nil {
		v.Type = openapi3.TypeArray
	}
	return nil
}

func fixEliminateCheckerUnion(v *openapi3.Schema) error {
	var newOneOf openapi3.SchemaRefs
	for _, ref := range v.OneOf {
		if !(ref.Ref == "" &&
			ref.Value.Type == "" &&
			ref.Value.Description == "" &&
			len(ref.Value.Properties) == 0 &&
			ref.Value.OneOf == nil &&
			ref.Value.AnyOf == nil &&
			ref.Value.AllOf == nil) {
			newOneOf = append(newOneOf, ref)
		}
	}
	v.OneOf = newOneOf

	var newAnyOf openapi3.SchemaRefs
	for _, ref := range v.AnyOf {
		if !(ref.Ref == "" &&
			ref.Value.Type == "" &&
			ref.Value.Description == "" &&
			len(ref.Value.Properties) == 0 &&
			ref.Value.OneOf == nil &&
			ref.Value.AnyOf == nil &&
			ref.Value.AllOf == nil) {
			newAnyOf = append(newAnyOf, ref)
		}
	}
	v.AnyOf = newAnyOf

	var newAllOf openapi3.SchemaRefs
	for _, ref := range v.AllOf {
		if !(ref.Ref == "" &&
			ref.Value.Type == "" &&
			ref.Value.Description == "" &&
			len(ref.Value.Properties) == 0 &&
			ref.Value.OneOf == nil &&
			ref.Value.AnyOf == nil &&
			ref.Value.AllOf == nil) {
			newAllOf = append(newAllOf, ref)
		}
	}
	v.AllOf = newAllOf

	return nil
}

func fixAdditionalProperties(v *openapi3.Schema) error {
	if (v.Type == openapi3.TypeObject || v.Type == "") && len(v.Properties) > 0 &&
		v.AdditionalProperties.Has == nil && v.AdditionalProperties.Schema == nil {
		v.WithAnyAdditionalProperties()
	}
	return nil
}

func fixSkipOptionalPointer(v *openapi3.Schema) error {
	skipOptionalPointer := false

	if v.Nullable {
		return nil
	}

	switch v.Type {
	case openapi3.TypeString:
		// TODO format check
		// Check whether allow empty string
		if v.MinLength > 0 {
			skipOptionalPointer = true
			break
		}
		if r, err := regexp.Compile(v.Pattern); r != nil && err == nil {
			if !r.MatchString("") {
				skipOptionalPointer = true
				break
			}
		}
		if len(v.Enum) != 0 {
			existEmptyMember := false
			for _, m := range v.Enum {
				if s, ok := m.(string); ok {
					if s == "" {
						existEmptyMember = true
						break
					}
				}
			}
			if !existEmptyMember {
				skipOptionalPointer = true
				break
			}
		}

	case openapi3.TypeArray:
		if v.MinItems > 0 {
			skipOptionalPointer = true
			break
		}

	case openapi3.TypeInteger, openapi3.TypeNumber:
		if v.Min != nil {
			if v.ExclusiveMin {
				if *v.Min >= 0 {
					skipOptionalPointer = true
					break
				}
			} else {
				if *v.Min > 0 {
					skipOptionalPointer = true
					break
				}
			}
		}
		if v.Max != nil {
			if v.ExclusiveMax {
				if *v.Max <= 0 {
					skipOptionalPointer = true
					break
				}
			} else {
				if *v.Max < 0 {
					skipOptionalPointer = true
					break
				}
			}
		}
	}

	if skipOptionalPointer {
		if v.Extensions == nil {
			v.Extensions = make(map[string]interface{})
		}
		v.Extensions["x-go-type-skip-optional-pointer"] = true
	}

	return nil
}

func fixCutSchemaRef(v *openapi3.SchemaRef) error {
	origSchema := v.Value

	newDescription := v.Value.Description
	if newDescription == "" {
		newDescription = fmt.Sprintf("Original reference %s", v.Ref)
	} else {
		newDescription = fmt.Sprintf("%s (Original reference %s)", v.Value.Description, v.Ref)
	}

	v.Ref = ""
	v.Value = &openapi3.Schema{
		Description: newDescription,
	}

	skipPointer := true

	switch origSchema.Type {
	case "":
		if len(origSchema.AnyOf) == 2 &&
			origSchema.AnyOf[0].Value.Type == openapi3.TypeString &&
			origSchema.AnyOf[1].Value.Type == openapi3.TypeString {
			v.Value.Type = openapi3.TypeString
			skipPointer = false
		}
	case openapi3.TypeBoolean, openapi3.TypeString:
		v.Value.Type = origSchema.Type
		skipPointer = false
	}

	if skipPointer {
		v.Value.Extensions = map[string]interface{}{
			"x-go-type-skip-optional-pointer": true,
		}
	}

	return nil
}

func fixNullable(v *openapi3.Schema, specName string) error {
	nullValueRef := ""
	if specName == "TS29571_CommonData.yaml" {
		nullValueRef = "#/components/schemas/NullValue"
	} else {
		nullValueRef = "TS29571_CommonData.yaml#/components/schemas/NullValue"
	}

	if len(v.AnyOf) == 2 {
		if v.AnyOf[0].Ref == nullValueRef {
			*v = *(deepcopy.Copy(v.AnyOf[1].Value).(*openapi3.Schema))
			// kin-openapi don`t allow this
			// v.Nullable = true
			v.Nullable = false
		} else if v.AnyOf[1].Ref == nullValueRef {
			*v = *(deepcopy.Copy(v.AnyOf[0].Value).(*openapi3.Schema))
			// kin-openapi don`t allow this
			// v.Nullable = true
			v.Nullable = false
		}
	}
	return nil
}

func getRangeForGeneratedType(v *openapi3.Schema) (minValue float64, maxValue float64) {
	switch v.Format {
	case "int64":
		return math.MinInt64, math.MaxInt64
	case "int32":
		return math.MinInt32, math.MaxInt32
	case "int16":
		return math.MinInt16, math.MaxInt16
	case "int8":
		return math.MinInt8, math.MaxInt8
	case "int":
		// support for 32bit arch
		return math.MinInt32, math.MaxInt32
	case "uint64":
		return 0, math.MaxUint64
	case "uint32":
		return 0, math.MaxUint32
	case "uint16":
		return 0, math.MaxUint16
	case "uint8":
		return 0, math.MaxUint8
	case "uint":
		// support for 32bit arch
		return 0, math.MaxUint32
	default:
		// use int type
		// support for 32bit arch
		return math.MinInt32, math.MaxInt32
	}
}

func isFitRange(v *openapi3.Schema) bool {
	minValue, maxValue := getRangeForGeneratedType(v)
	if v.Min != nil {
		if *v.Min < minValue || *v.Min > maxValue {
			return false
		}
	}
	if v.Max != nil {
		if *v.Max < minValue || *v.Max > maxValue {
			return false
		}
	}
	return true
}

func dumpAndPanicSchema(v *openapi3.Schema) {
	vMin := "<nil>"
	if v.Min != nil {
		vMin = fmt.Sprint(*v.Min)
	}
	vMax := "<nil>"
	if v.Max != nil {
		vMax = fmt.Sprint(*v.Max)
	}
	panic(fmt.Sprintf("%+v %+v %+v", *v, vMin, vMax))
}

func fixIntegerFormat(v *openapi3.Schema) error {
	if v.Type != openapi3.TypeInteger {
		return nil
	}

	if (v.Min != nil && (*v.Min > maxSafeInteger || *v.Min < minSafeInteger)) ||
		(v.Max != nil && (*v.Max > maxSafeInteger || *v.Max < minSafeInteger)) {
		switch {
		case v.Min != nil && *v.Min == 0 && v.Max != nil && *v.Max == 1.8446744073709552e+19:
			// TS29571_CommonData.yaml#/components/schemas/Uint64
			// TS29571_CommonData.yaml#/components/schemas/Uint64Rm
			v.Min = nil
			v.Max = nil
			v.Format = "uint64"
		default:
			dumpAndPanicSchema(v)
		}
	}

	// Check whether min/max value fit to generated type
	if isFitRange(v) {
		return nil
	}

	v.Format = "" // assume int type
	if v.Min != nil {
		if *v.Min < math.MinInt32 || *v.Min > math.MaxInt32 {
			v.Format = "int64"
		}
	}
	if v.Max != nil {
		if *v.Max < math.MinInt32 || *v.Max > math.MaxInt32 {
			v.Format = "int64"
		}
	}

	// Double check
	if !isFitRange(v) {
		dumpAndPanicSchema(v)
	}

	return nil
}
