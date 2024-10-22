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

import (
	"fmt"
	"math"
	"math/big"
	"regexp"
	"sort"
	"strings"

	"github.com/mohae/deepcopy"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"

	"github.com/ShouheiNishi/openapi5g/internal/generator/openapi"
)

const nameForModels = "f5gcModels"
const importForModels = "github.com/free5gc/openapi/models"

func (s *GeneratorState) RewriteYaml(spec string) error {
	doc := s.Specs[spec]
	if doc == nil {
		return fmt.Errorf("spec %s is not exist", spec)
	}

	refs := map[string]struct{}{}
	cutRefs := map[string]struct{}{}
	for _, s := range pkgList[spec].cutRefs {
		cutRefs[s] = struct{}{}
	}

	if err := scanRef(doc, refs, cutRefs); err != nil {
		return err
	}

	if spec == "TS29571_CommonData.yaml" {
		schema := doc.Components.Schemas["ProblemDetails"].Value.Properties["status"].Value
		schema.GoTypeSkipOptionalPointer = true
	}

	if spec == "TS29571_CommonData.yaml" {
		for statusCode, res := range doc.Components.Responses {
			if statusCode == "default" {
				res.Value.Content = map[string]*openapi.MediaType{
					"application/problem+json": {
						Schema: openapi.Ref[openapi.Schema]{
							RefFile:    "TS29571_CommonData.yaml",
							RefPointer: "/components/schemas/ProblemDetails",
						},
					},
				}
			}
		}
	} else {
		for _, pathItem := range doc.Paths {
			if pathItem.HasRef() {
				continue
			}
			for _, op := range pathItem.Value.Operations() {
				if op.Responses == nil {
					op.Responses = make(map[string]*openapi.Ref[openapi.Response])
				}
				if op.Responses["default"] == nil {
					op.Responses["default"] = &openapi.Ref[openapi.Response]{Value: &openapi.Response{}}
				}
				if op.Responses["default"].Ref() != "TS29571_CommonData.yaml#/components/responses/default" {
					if op.Responses["default"].Value.Content == nil {
						op.Responses["default"].Value.Content = make(map[string]*openapi.MediaType)
					}
					if op.Responses["default"].Value.Content["application/problem+json"] == nil {
						resNew := deepcopy.Copy(op.Responses["default"]).(*openapi.Ref[openapi.Response])
						resNew.RefFile = ""
						resNew.RefPointer = ""
						resNew.Value.Content["application/problem+json"] = &openapi.MediaType{
							Schema: openapi.Ref[openapi.Schema]{
								RefFile:    "TS29571_CommonData.yaml",
								RefPointer: "/components/schemas/ProblemDetails",
							},
						}
						op.Responses["default"] = resNew
					}
				}
			}
		}
	}

	if spec == "TS29503_Nudm_SDM.yaml" {
		for _, parameterRef := range doc.Paths["/shared-data"].Value.Get.Parameters {
			parameter := parameterRef.Value
			if parameter.Name == "supportedFeatures" {
				parameter.GoName = "SupportedFeaturesShouldNotBeUsed"
			}
		}
	}

	if spec == "TS29571_CommonData.yaml" {
		schema := doc.Components.Schemas["ExtSnssai"].Value
		for i := range schema.AllOf {
			if strings.HasSuffix(schema.AllOf[i].RefPointer, "/Snssai") {
				newSchema := *schema.AllOf[i].Value
				schema.AllOf[i] = openapi.Ref[openapi.Schema]{
					Value: &newSchema,
				}
			}
		}

		schema = doc.Components.Schemas["Snssai"].Value
		schema.GoType = nameForModels + ".Snssai"
		schema.GoTypeImport = &openapi.GoTypeImport{
			Name: nameForModels,
			Path: importForModels,
		}

		schema = doc.Components.Schemas["PlmnId"].Value
		schema.GoType = nameForModels + ".PlmnId"
		schema.GoTypeImport = &openapi.GoTypeImport{
			Name: nameForModels,
			Path: importForModels,
		}

		schema = doc.Components.Schemas["Guami"].Value
		schema.GoType = nameForModels + ".Guami"
		schema.GoTypeImport = &openapi.GoTypeImport{
			Name: nameForModels,
			Path: importForModels,
		}

		for _, t := range []string{"Int32", "Int64", "Uint16", "Uint32", "Uint64"} {
			doc.Components.Schemas[t].Value.Format = strings.ToLower(t)
			doc.Components.Schemas[t].Value.Minimum = nil
			doc.Components.Schemas[t].Value.ExclusiveMinimum = false
			doc.Components.Schemas[t].Value.Maximum = nil
			doc.Components.Schemas[t].Value.ExclusiveMaximum = false
			doc.Components.Schemas[t+"Rm"].Value.Format = strings.ToLower(t)
			doc.Components.Schemas[t+"Rm"].Value.Minimum = nil
			doc.Components.Schemas[t+"Rm"].Value.ExclusiveMinimum = false
			doc.Components.Schemas[t+"Rm"].Value.Maximum = nil
			doc.Components.Schemas[t+"Rm"].Value.ExclusiveMaximum = false
		}
		doc.Components.Schemas["Uinteger"].Value.Format = "uint"
		doc.Components.Schemas["Uinteger"].Value.Minimum = nil
		doc.Components.Schemas["Uinteger"].Value.ExclusiveMinimum = false
		doc.Components.Schemas["Uinteger"].Value.Maximum = nil
		doc.Components.Schemas["Uinteger"].Value.ExclusiveMaximum = false
		doc.Components.Schemas["UintegerRm"].Value.Format = "uint"
		doc.Components.Schemas["UintegerRm"].Value.Minimum = nil
		doc.Components.Schemas["UintegerRm"].Value.ExclusiveMinimum = false
		doc.Components.Schemas["UintegerRm"].Value.Maximum = nil
		doc.Components.Schemas["UintegerRm"].Value.ExclusiveMaximum = false
	}

	deps := make([]string, 0, len(refs))
	for r := range refs {
		if r == spec {
			continue
		}
		if _, exist := pkgList[r]; !exist {
			if _, exist := cutRefs[r]; !exist {
				panic(fmt.Sprintf("%s is not defined.", r))
			}
		}
		deps = append(deps, r)
	}
	sort.Strings(deps)

	s.DepsForImport[spec] = deps

	return nil
}

func getSchemaType(schema *openapi.Schema) openapi.SchemaType {
	if schema.Type != nil {
		return *schema.Type
	}
	return ""
}

func fixAnyOfEnum(v *openapi.Schema) error {
	if len(v.AnyOf) == 2 && !v.AnyOf[0].HasRef() && !v.AnyOf[1].HasRef() {
		v0 := v.AnyOf[0].Value
		v1 := v.AnyOf[1].Value
		if getSchemaType(v0) == getSchemaType(v1) && v0.Format == v1.Format &&
			(getSchemaType(v0) == openapi.SchemaTypeInteger || getSchemaType(v1) == openapi.SchemaTypeString) {
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
				v.GoTypeSkipOptionalPointer = v0.GoTypeSkipOptionalPointer && v1.GoTypeSkipOptionalPointer
			}
		}
	}
	return nil
}

func fixAnyOfString(v *openapi.Schema) error {
	if len(v.AnyOf) > 0 {
		for _, vRef := range v.AnyOf {
			if vRef.Value.Type == nil || *vRef.Value.Type != openapi.SchemaTypeString {
				return nil
			}
		}
		newDescription := []string{"Merged type of"}
		newSkipOptionalPointer := true
		for _, vRef := range v.AnyOf {
			if !vRef.HasRef() {
				if vRef.Value.Description == "" {
					newDescription = append(newDescription, "  Anonymous string")
				} else {
					newDescription = append(newDescription, "  "+vRef.Value.Description)
				}
			} else {
				if vRef.Value.Description == "" {
					newDescription = append(newDescription, "  string in "+vRef.Ref())
				} else {
					newDescription = append(newDescription, "  "+vRef.Value.Description+" in "+vRef.Ref())
				}
			}
			if !vRef.Value.GoTypeSkipOptionalPointer {
				newSkipOptionalPointer = false
			}
		}
		*v = openapi.Schema{
			Type:                      lo.ToPtr(openapi.SchemaTypeString),
			Description:               strings.Join(newDescription, "\n"),
			GoTypeSkipOptionalPointer: newSkipOptionalPointer,
		}
	}
	return nil
}

func fixImplicitArray(v *openapi.Schema) error {
	if getSchemaType(v) == "" && !v.Items.IsZero() {
		v.Type = lo.ToPtr(openapi.SchemaTypeArray)
	}
	return nil
}

func fixEliminateCheckerUnion(v *openapi.Schema) error {
	var newOneOf []openapi.Ref[openapi.Schema]
	for _, ref := range v.OneOf {
		if !(!ref.HasRef() &&
			getSchemaType(ref.Value) == "" &&
			ref.Value.Description == "" &&
			len(ref.Value.Properties) == 0 &&
			ref.Value.OneOf == nil &&
			ref.Value.AnyOf == nil &&
			ref.Value.AllOf == nil) {
			newOneOf = append(newOneOf, ref)
		}
	}
	v.OneOf = newOneOf

	var newAnyOf []openapi.Ref[openapi.Schema]
	for _, ref := range v.AnyOf {
		if !(!ref.HasRef() &&
			getSchemaType(ref.Value) == "" &&
			ref.Value.Description == "" &&
			len(ref.Value.Properties) == 0 &&
			ref.Value.OneOf == nil &&
			ref.Value.AnyOf == nil &&
			ref.Value.AllOf == nil) {
			newAnyOf = append(newAnyOf, ref)
		}
	}
	v.AnyOf = newAnyOf

	var newAllOf []openapi.Ref[openapi.Schema]
	for _, ref := range v.AllOf {
		if !(!ref.HasRef() &&
			getSchemaType(ref.Value) == "" &&
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

func fixAdditionalProperties(v *openapi.Schema) error {
	if (getSchemaType(v) == openapi.SchemaTypeObject || getSchemaType(v) == "") && len(v.Properties) > 0 &&
		v.AdditionalProperties.Bool == nil && v.AdditionalProperties.SchemaRef == nil {
		v.AdditionalProperties.Bool = lo.ToPtr(true)
	}
	return nil
}

func maySkipOptionalPointerByMin[T int | int64 | uint64 | float64](v T, exclusive bool) bool {
	if exclusive {
		if v >= T(0) {
			return true
		}
	} else {
		if v > T(0) {
			return true
		}
	}
	return false
}

func maySkipOptionalPointerByMax[T int | int64 | uint64 | float64](v T, exclusive bool) bool {
	if exclusive {
		if v <= T(0) {
			return true
		}
	} else {
		if v < T(0) {
			return true
		}
	}
	return false
}

func fixSkipOptionalPointer(v *openapi.Schema) error {
	skipOptionalPointer := false

	if v.Nullable != nil && *v.Nullable {
		return nil
	}

	switch getSchemaType(v) {
	case openapi.SchemaTypeString:
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
				if m.Kind == yaml.ScalarNode && m.Value == "" {
					existEmptyMember = true
					break
				}
			}
			if !existEmptyMember {
				skipOptionalPointer = true
				break
			}
		}

	case openapi.SchemaTypeArray:
		if v.MinItems != nil && *v.MinItems > 0 {
			skipOptionalPointer = true
			break
		}

	case openapi.SchemaTypeInteger, openapi.SchemaTypeNumber:
		switch m := v.Minimum.(type) {
		case nil:
		case int:
			if maySkipOptionalPointerByMin(m, v.ExclusiveMinimum) {
				skipOptionalPointer = true
			}
		case int64:
			if maySkipOptionalPointerByMin(m, v.ExclusiveMinimum) {
				skipOptionalPointer = true
			}
		case uint64:
			if maySkipOptionalPointerByMin(m, v.ExclusiveMinimum) {
				skipOptionalPointer = true
			}
		case float64:
			if maySkipOptionalPointerByMin(m, v.ExclusiveMinimum) {
				skipOptionalPointer = true
			}
		default:
			return fmt.Errorf("unknown minimum type %T", m)
		}
		switch m := v.Maximum.(type) {
		case nil:
		case int:
			if maySkipOptionalPointerByMax(m, v.ExclusiveMaximum) {
				skipOptionalPointer = true
			}
		case int64:
			if maySkipOptionalPointerByMax(m, v.ExclusiveMaximum) {
				skipOptionalPointer = true
			}
		case uint64:
			if maySkipOptionalPointerByMax(m, v.ExclusiveMaximum) {
				skipOptionalPointer = true
			}
		case float64:
			if maySkipOptionalPointerByMax(m, v.ExclusiveMaximum) {
				skipOptionalPointer = true
			}
		default:
			return fmt.Errorf("unknown maximum type %T", m)
		}
	}

	if skipOptionalPointer {
		v.GoTypeSkipOptionalPointer = true
	}

	return nil
}

func fixCutSchemaRef(v *openapi.Ref[openapi.Schema]) error {
	origSchema := v.Value

	newDescription := v.Value.Description
	if newDescription == "" {
		newDescription = fmt.Sprintf("Original reference %s", v.Ref())
	} else {
		newDescription = fmt.Sprintf("%s (Original reference %s)", v.Value.Description, v.Ref())
	}

	v.RefFile = ""
	v.RefPointer = ""
	v.Value = &openapi.Schema{
		Description: newDescription,
	}

	skipPointer := true

	switch getSchemaType(origSchema) {
	case "":
		if len(origSchema.AnyOf) == 2 &&
			origSchema.AnyOf[0].Value.Type != nil &&
			*origSchema.AnyOf[0].Value.Type == openapi.SchemaTypeString &&
			origSchema.AnyOf[1].Value.Type != nil &&
			*origSchema.AnyOf[1].Value.Type == openapi.SchemaTypeString {
			v.Value.Type = lo.ToPtr(openapi.SchemaTypeString)
			skipPointer = false
		}
	case openapi.SchemaTypeBoolean, openapi.SchemaTypeString:
		v.Value.Type = origSchema.Type
		skipPointer = false
	}

	if skipPointer {
		v.Value.GoTypeSkipOptionalPointer = true
	}

	return nil
}

func fixNullable(v *openapi.Schema) error {
	nullValueRef := "TS29571_CommonData.yaml#/components/schemas/NullValue"

	if len(v.AnyOf) == 2 {
		if v.AnyOf[0].Ref() == nullValueRef {
			*v = *(deepcopy.Copy(v.AnyOf[1].Value).(*openapi.Schema))
			// kin-openapi don`t allow this
			// v.Nullable = true
			v.Nullable = nil
		} else if v.AnyOf[1].Ref() == nullValueRef {
			*v = *(deepcopy.Copy(v.AnyOf[0].Value).(*openapi.Schema))
			// kin-openapi don`t allow this
			// v.Nullable = true
			v.Nullable = nil
		}
	}
	return nil
}

func getRangeForGeneratedType(v *openapi.Schema) (minValue *big.Int, maxValue *big.Int) {
	switch v.Format {
	case "int64":
		return big.NewInt(math.MinInt64), big.NewInt(math.MaxInt64)
	case "int32":
		return big.NewInt(math.MinInt32), big.NewInt(math.MaxInt32)
	case "int16":
		return big.NewInt(math.MinInt16), big.NewInt(math.MaxInt16)
	case "int8":
		return big.NewInt(math.MinInt8), big.NewInt(math.MaxInt8)
	case "int":
		// support for 32bit arch
		return big.NewInt(math.MinInt32), big.NewInt(math.MaxInt32)
	case "uint64":
		return big.NewInt(0), new(big.Int).SetUint64(math.MaxUint64)
	case "uint32":
		return big.NewInt(0), big.NewInt(math.MaxUint32)
	case "uint16":
		return big.NewInt(0), big.NewInt(math.MaxUint16)
	case "uint8":
		return big.NewInt(0), big.NewInt(math.MaxUint8)
	case "uint":
		// support for 32bit arch
		return big.NewInt(0), big.NewInt(math.MaxUint32)
	default:
		// use int type
		// support for 32bit arch
		return big.NewInt(math.MinInt32), big.NewInt(math.MaxInt32)
	}
}

func isFitRange(v *openapi.Schema, min *big.Int, max *big.Int) bool {
	minValue, maxValue := getRangeForGeneratedType(v)
	if min != nil {
		if min.Cmp(minValue) < 0 || min.Cmp(maxValue) > 0 {
			return false
		}
	}
	if max != nil {
		if max.Cmp(minValue) < 0 || max.Cmp(maxValue) > 0 {
			return false
		}
	}
	return true
}

func dumpAndPanicSchema(v *openapi.Schema) {
	panic(fmt.Sprintf("%+v %+v %+v", *v, v.Minimum, v.Maximum))
}

var bigOne = big.NewInt(1)

func setupMinMax(in any, exclusive bool, max bool) (r *big.Int, err error) {
	var minMax string
	if max {
		minMax = "max"
	} else {
		minMax = "min"
	}
	switch in := in.(type) {
	case nil:
		return nil, nil
	case int:
		r = big.NewInt(int64(in))
	case int64:
		r = big.NewInt(in)
	case uint64:
		r = new(big.Int).SetUint64(in)
	case float64:
		if i, a := new(big.Float).SetFloat64(in).Int(nil); a != big.Exact {
			return nil, fmt.Errorf("%s value is not integer", minMax)
		} else {
			r = i
		}
	default:
		return nil, fmt.Errorf("unknown %s type %T", minMax, in)
	}

	if exclusive {
		if max {
			r = r.Sub(r, bigOne)
		} else {
			r = r.Add(r, bigOne)
		}
	}

	return r, nil
}

func fixIntegerFormat(v *openapi.Schema) error {
	if v.Type == nil || *v.Type != openapi.SchemaTypeInteger {
		return nil
	}

	var min, max *big.Int
	if m, err := setupMinMax(v.Minimum, v.ExclusiveMinimum, false); err != nil {
		return err
	} else {
		min = m
	}
	if m, err := setupMinMax(v.Maximum, v.ExclusiveMaximum, true); err != nil {
		return err
	} else {
		max = m
	}

	// Check whether min/max value fit to generated type
	if isFitRange(v, min, max) {
		return nil
	}

	v.Format = "" // assume int type
	if isFitRange(v, min, max) {
		return nil
	}
	v.Format = "int64"
	if isFitRange(v, min, max) {
		return nil
	}
	v.Format = "uint64"

	// Check whether min/max value fit to generated type
	if !isFitRange(v, min, max) {
		dumpAndPanicSchema(v)
	}

	return nil
}
