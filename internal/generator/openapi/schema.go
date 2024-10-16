// Copyright 2024 APRESIA Systems LTD.
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

// Base schema is https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/schemas/v3.0/schema.json
// Base code generated by github.com/atombender/go-jsonschema

package openapi

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type GoTypeImport struct {
	Name string `yaml:"name,omitempty"`
	Path string `yaml:"path,omitempty"`
}

type Schema struct {
	AdditionalProperties AdditionalProperties   `yaml:"additionalProperties,omitempty"`
	AllOf                []Ref[Schema]          `yaml:"allOf,omitempty"`
	AnyOf                []Ref[Schema]          `yaml:"anyOf,omitempty"`
	Default              yaml.Node              `yaml:"default,omitempty"`
	Deprecated           bool                   `yaml:"deprecated,omitempty"`
	Description          string                 `yaml:"description,omitempty"`
	Discriminator        *Discriminator         `yaml:"discriminator,omitempty"`
	Enum                 []yaml.Node            `yaml:"enum,omitempty"`
	Example              yaml.Node              `yaml:"example,omitempty"`
	ExclusiveMaximum     bool                   `yaml:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum     bool                   `yaml:"exclusiveMinimum,omitempty"`
	ExternalDocs         *ExternalDocumentation `yaml:"externalDocs,omitempty"`
	Format               string                 `yaml:"format,omitempty"`
	Items                Ref[Schema]            `yaml:"items,omitempty"`
	MaxItems             int                    `yaml:"maxItems,omitempty"`
	MaxLength            int                    `yaml:"maxLength,omitempty"`
	MaxProperties        int                    `yaml:"maxProperties,omitempty"`
	Maximum              any                    `yaml:"maximum,omitempty"`
	MinItems             *int                   `yaml:"minItems,omitempty"`
	MinLength            int                    `yaml:"minLength,omitempty"`
	MinProperties        *int                   `yaml:"minProperties,omitempty"`
	Minimum              any                    `yaml:"minimum,omitempty"`
	MultipleOf           *float64               `yaml:"multipleOf,omitempty"`
	Not                  Ref[Schema]            `yaml:"not,omitempty"`
	Nullable             *bool                  `yaml:"nullable,omitempty"`
	OneOf                []Ref[Schema]          `yaml:"oneOf,omitempty"`
	Pattern              string                 `yaml:"pattern,omitempty"`
	Properties           SchemaProperties       `yaml:"properties,omitempty"`
	ReadOnly             bool                   `yaml:"readOnly,omitempty"`
	Required             []string               `yaml:"required,omitempty"`
	Title                string                 `yaml:"title,omitempty"`
	Type                 *SchemaType            `yaml:"type,omitempty"`
	UniqueItems          bool                   `yaml:"uniqueItems,omitempty"`
	WriteOnly            bool                   `yaml:"writeOnly,omitempty"`
	Xml                  *XML                   `yaml:"xml,omitempty"`

	GoType                    string        `yaml:"x-go-type,omitempty"`
	GoTypeImport              *GoTypeImport `yaml:"x-go-type-import,omitempty"`
	GoTypeSkipOptionalPointer bool          `yaml:"x-go-type-skip-optional-pointer,omitempty"`
}

var _ yaml.Marshaler = &Schema{}

func yamlNumberFix(number any) (yaml.Node, error) {
	switch n := number.(type) {
	case int:
		return yaml.Node{Kind: yaml.ScalarNode, Value: strconv.FormatInt(int64(n), 10)}, nil
	case int64:
		return yaml.Node{Kind: yaml.ScalarNode, Value: strconv.FormatInt(n, 10)}, nil
	case uint64:
		return yaml.Node{Kind: yaml.ScalarNode, Value: strconv.FormatUint(n, 10)}, nil
	case float64:
		s := strconv.FormatFloat(n, 'g', -1, 64)
		if !strings.ContainsAny(s, ".e") {
			s += ".0"
		}
		return yaml.Node{Kind: yaml.ScalarNode, Value: s}, nil
	case nil:
		return yaml.Node{}, nil
	default:
		return yaml.Node{}, fmt.Errorf("invalid type %T", number)
	}
}

func (s *Schema) MarshalYAML() (interface{}, error) {
	var ret struct {
		AdditionalProperties AdditionalProperties   `yaml:"additionalProperties,omitempty"`
		AllOf                []Ref[Schema]          `yaml:"allOf,omitempty"`
		AnyOf                []Ref[Schema]          `yaml:"anyOf,omitempty"`
		Default              yaml.Node              `yaml:"default,omitempty"`
		Deprecated           bool                   `yaml:"deprecated,omitempty"`
		Description          string                 `yaml:"description,omitempty"`
		Discriminator        *Discriminator         `yaml:"discriminator,omitempty"`
		Enum                 []yaml.Node            `yaml:"enum,omitempty"`
		Example              yaml.Node              `yaml:"example,omitempty"`
		ExclusiveMaximum     bool                   `yaml:"exclusiveMaximum,omitempty"`
		ExclusiveMinimum     bool                   `yaml:"exclusiveMinimum,omitempty"`
		ExternalDocs         *ExternalDocumentation `yaml:"externalDocs,omitempty"`
		Format               string                 `yaml:"format,omitempty"`
		Items                Ref[Schema]            `yaml:"items,omitempty"`
		MaxItems             int                    `yaml:"maxItems,omitempty"`
		MaxLength            int                    `yaml:"maxLength,omitempty"`
		MaxProperties        int                    `yaml:"maxProperties,omitempty"`
		Maximum              yaml.Node              `yaml:"maximum,omitempty"`
		MinItems             *int                   `yaml:"minItems,omitempty"`
		MinLength            int                    `yaml:"minLength,omitempty"`
		MinProperties        *int                   `yaml:"minProperties,omitempty"`
		Minimum              yaml.Node              `yaml:"minimum,omitempty"`
		MultipleOf           *float64               `yaml:"multipleOf,omitempty"`
		Not                  Ref[Schema]            `yaml:"not,omitempty"`
		Nullable             *bool                  `yaml:"nullable,omitempty"`
		OneOf                []Ref[Schema]          `yaml:"oneOf,omitempty"`
		Pattern              string                 `yaml:"pattern,omitempty"`
		Properties           SchemaProperties       `yaml:"properties,omitempty"`
		ReadOnly             bool                   `yaml:"readOnly,omitempty"`
		Required             []string               `yaml:"required,omitempty"`
		Title                string                 `yaml:"title,omitempty"`
		Type                 *SchemaType            `yaml:"type,omitempty"`
		UniqueItems          bool                   `yaml:"uniqueItems,omitempty"`
		WriteOnly            bool                   `yaml:"writeOnly,omitempty"`
		Xml                  *XML                   `yaml:"xml,omitempty"`

		GoType                    string        `yaml:"x-go-type,omitempty"`
		GoTypeImport              *GoTypeImport `yaml:"x-go-type-import,omitempty"`
		GoTypeSkipOptionalPointer bool          `yaml:"x-go-type-skip-optional-pointer,omitempty"`
	}

	ret.AdditionalProperties = s.AdditionalProperties
	ret.AllOf = s.AllOf
	ret.AnyOf = s.AnyOf
	ret.Default = s.Default
	ret.Deprecated = s.Deprecated
	ret.Description = s.Description
	ret.Discriminator = s.Discriminator
	ret.Enum = s.Enum
	ret.Example = s.Example
	ret.ExclusiveMaximum = s.ExclusiveMaximum
	ret.ExclusiveMinimum = s.ExclusiveMinimum
	ret.ExternalDocs = s.ExternalDocs
	ret.Format = s.Format
	ret.Items = s.Items
	ret.MaxItems = s.MaxItems
	ret.MaxLength = s.MaxLength
	ret.MaxProperties = s.MaxProperties
	if v, err := yamlNumberFix(s.Maximum); err != nil {
		return nil, fmt.Errorf("bad Maximum: %w", err)
	} else {
		ret.Maximum = v
	}
	ret.MinItems = s.MinItems
	ret.MinLength = s.MinLength
	ret.MinProperties = s.MinProperties
	if v, err := yamlNumberFix(s.Minimum); err != nil {
		return nil, fmt.Errorf("bad Minimum: %w", err)
	} else {
		ret.Minimum = v
	}
	ret.MultipleOf = s.MultipleOf
	ret.Not = s.Not
	ret.Nullable = s.Nullable
	ret.OneOf = s.OneOf
	ret.Pattern = s.Pattern
	ret.Properties = s.Properties
	ret.ReadOnly = s.ReadOnly
	ret.Required = s.Required
	ret.Title = s.Title
	ret.Type = s.Type
	ret.UniqueItems = s.UniqueItems
	ret.WriteOnly = s.WriteOnly
	ret.Xml = s.Xml
	ret.GoType = s.GoType
	ret.GoTypeImport = s.GoTypeImport
	ret.GoTypeSkipOptionalPointer = s.GoTypeSkipOptionalPointer

	return ret, nil
}
