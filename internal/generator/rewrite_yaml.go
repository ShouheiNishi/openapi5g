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
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/invopop/yaml"
	"github.com/mohae/deepcopy"
)

const nameForModels = "f5gcModels"
const importForModels = "github.com/free5gc/openapi/models"

func RewriteYaml(rootDir string, spec string, doc *openapi3.T) (outLists []string, deps []string, err error) {
	if err := fixLocalRefInRemoteRef(doc, spec); err != nil {
		return nil, nil, err
	}

	refs := map[string]struct{}{}
	cutRefs := map[string]struct{}{}
	for _, s := range pkgList[spec].cutRefs {
		cutRefs[s] = struct{}{}
	}
	if err := scanRef(doc, refs, cutRefs); err != nil {
		return nil, nil, err
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

func fixAllOfEnum(v *openapi3.Schema) error {
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
			if v0.Enum == nil && len(v1.Enum) > 0 {
				*v = *v1
				v.Description = newDescription
			} else if v1.Enum == nil && len(v0.Enum) > 0 {
				*v = *v0
				v.Description = newDescription
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

func fixCutSchemaRef(v *openapi3.SchemaRef) error {
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

	return nil
}
