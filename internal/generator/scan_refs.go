// Code generated by github.com/ShouheiNishi/openapi5g/internal/generator/sub, DO NOT EDIT.

package generator

import (
	"sort"

	"gopkg.in/yaml.v3"

	"github.com/ShouheiNishi/openapi5g/internal/generator/openapi"
)

type scanRefType struct {
	visited map[interface{}]struct{}
	doc     *openapi.Document
	refs    map[string]struct{}
	cutRefs map[string]struct{}
}

func scanRef(d *openapi.Document, refs map[string]struct{}, cutRefs map[string]struct{}) error {
	s := &scanRefType{
		visited: make(map[interface{}]struct{}),
		doc:     d,
	}

	s.refs = refs
	s.cutRefs = cutRefs

	return s.walkDocument(s.doc)
}

func (s *scanRefType) walkAdditionalProperties(v *openapi.AdditionalProperties) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	if v.SchemaRef != nil {
		if err := s.walkSchemaRef(v.SchemaRef); err != nil {
			return err
		}
	}

	return nil
}

func (s *scanRefType) walkCallbackRef(v *openapi.Ref[openapi.Callback]) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	if v.HasRef() {
		if _, exist := s.cutRefs[v.RefFile]; exist {
		} else {
			s.refs[v.RefFile] = struct{}{}
			return nil
		}
	}

	if v.Value != nil {
		k1s5 := make([]string, 0, len(*v.Value))
		for k1 := range *v.Value {
			k1s5 = append(k1s5, k1)
		}
		sort.Strings(k1s5)
		for _, k1 := range k1s5 {
			if (*v.Value)[k1] != nil {
				if err := s.walkPathItemBaseRef((*v.Value)[k1]); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (s *scanRefType) walkComponents(v *openapi.Components) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	k0s0 := make([]string, 0, len(v.Callbacks))
	for k0 := range v.Callbacks {
		k0s0 = append(k0s0, k0)
	}
	sort.Strings(k0s0)
	for _, k0 := range k0s0 {
		if v.Callbacks[k0] != nil {
			if err := s.walkCallbackRef(v.Callbacks[k0]); err != nil {
				return err
			}
		}
	}

	k0s1 := make([]string, 0, len(v.Examples))
	for k0 := range v.Examples {
		k0s1 = append(k0s1, k0)
	}
	sort.Strings(k0s1)
	for _, k0 := range k0s1 {
		if v.Examples[k0] != nil {
			if err := s.walkExampleRef(v.Examples[k0]); err != nil {
				return err
			}
		}
	}

	k0s2 := make([]string, 0, len(v.Headers))
	for k0 := range v.Headers {
		k0s2 = append(k0s2, k0)
	}
	sort.Strings(k0s2)
	for _, k0 := range k0s2 {
		if v.Headers[k0] != nil {
			if err := s.walkHeaderRef(v.Headers[k0]); err != nil {
				return err
			}
		}
	}

	k0s3 := make([]string, 0, len(v.Links))
	for k0 := range v.Links {
		k0s3 = append(k0s3, k0)
	}
	sort.Strings(k0s3)
	for _, k0 := range k0s3 {
		if v.Links[k0] != nil {
			if err := s.walkLinkRef(v.Links[k0]); err != nil {
				return err
			}
		}
	}

	k0s4 := make([]string, 0, len(v.Parameters))
	for k0 := range v.Parameters {
		k0s4 = append(k0s4, k0)
	}
	sort.Strings(k0s4)
	for _, k0 := range k0s4 {
		if v.Parameters[k0] != nil {
			if err := s.walkParameterRef(v.Parameters[k0]); err != nil {
				return err
			}
		}
	}

	k0s5 := make([]string, 0, len(v.RequestBodies))
	for k0 := range v.RequestBodies {
		k0s5 = append(k0s5, k0)
	}
	sort.Strings(k0s5)
	for _, k0 := range k0s5 {
		if v.RequestBodies[k0] != nil {
			if err := s.walkRequestBodyRef(v.RequestBodies[k0]); err != nil {
				return err
			}
		}
	}

	k0s6 := make([]string, 0, len(v.Responses))
	for k0 := range v.Responses {
		k0s6 = append(k0s6, k0)
	}
	sort.Strings(k0s6)
	for _, k0 := range k0s6 {
		if v.Responses[k0] != nil {
			if err := s.walkResponseRef(v.Responses[k0]); err != nil {
				return err
			}
		}
	}

	k0s7 := make([]string, 0, len(v.Schemas))
	for k0 := range v.Schemas {
		k0s7 = append(k0s7, k0)
	}
	sort.Strings(k0s7)
	for _, k0 := range k0s7 {
		if v.Schemas[k0] != nil {
			if err := s.walkSchemaRef(v.Schemas[k0]); err != nil {
				return err
			}
		}
	}

	k0s8 := make([]string, 0, len(v.SecuritySchemes))
	for k0 := range v.SecuritySchemes {
		k0s8 = append(k0s8, k0)
	}
	sort.Strings(k0s8)
	for _, k0 := range k0s8 {
		if v.SecuritySchemes[k0] != nil {
			if err := s.walkNodeRef(v.SecuritySchemes[k0]); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *scanRefType) walkDocument(v *openapi.Document) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	if v.Components != nil {
		if err := s.walkComponents(v.Components); err != nil {
			return err
		}
	}

	k0s4 := make([]string, 0, len(v.Paths))
	for k0 := range v.Paths {
		k0s4 = append(k0s4, k0)
	}
	sort.Strings(k0s4)
	for _, k0 := range k0s4 {
		if v.Paths[k0] != nil {
			if err := s.walkPathItemBaseRef(v.Paths[k0]); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *scanRefType) walkEncoding(v *openapi.Encoding) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	k0s3 := make([]string, 0, len(v.Headers))
	for k0 := range v.Headers {
		k0s3 = append(k0s3, k0)
	}
	sort.Strings(k0s3)
	for _, k0 := range k0s3 {
		if v.Headers[k0] != nil {
			if err := s.walkHeaderRef(v.Headers[k0]); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *scanRefType) walkExampleRef(v *openapi.Ref[openapi.Example]) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	if v.HasRef() {
		if _, exist := s.cutRefs[v.RefFile]; exist {
		} else {
			s.refs[v.RefFile] = struct{}{}
			return nil
		}
	}

	return nil
}

func (s *scanRefType) walkHeader(v *openapi.Header) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	k0s2 := make([]string, 0, len(v.Content))
	for k0 := range v.Content {
		k0s2 = append(k0s2, k0)
	}
	sort.Strings(k0s2)
	for _, k0 := range k0s2 {
		if v.Content[k0] != nil {
			if err := s.walkMediaType(v.Content[k0]); err != nil {
				return err
			}
		}
	}

	k0s6 := make([]string, 0, len(v.Examples))
	for k0 := range v.Examples {
		k0s6 = append(k0s6, k0)
	}
	sort.Strings(k0s6)
	for _, k0 := range k0s6 {
		if v.Examples[k0] != nil {
			if err := s.walkExampleRef(v.Examples[k0]); err != nil {
				return err
			}
		}
	}

	if err := s.walkSchemaRef(&v.Schema); err != nil {
		return err
	}

	return nil
}

func (s *scanRefType) walkHeaderRef(v *openapi.Ref[openapi.Header]) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	if v.HasRef() {
		if _, exist := s.cutRefs[v.RefFile]; exist {
		} else {
			s.refs[v.RefFile] = struct{}{}
			return nil
		}
	}

	if v.Value != nil {
		if err := s.walkHeader(v.Value); err != nil {
			return err
		}
	}

	return nil
}

func (s *scanRefType) walkLinkRef(v *openapi.Ref[openapi.Link]) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	if v.HasRef() {
		if _, exist := s.cutRefs[v.RefFile]; exist {
		} else {
			s.refs[v.RefFile] = struct{}{}
			return nil
		}
	}

	return nil
}

func (s *scanRefType) walkMediaType(v *openapi.MediaType) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	k0s0 := make([]string, 0, len(v.Encoding))
	for k0 := range v.Encoding {
		k0s0 = append(k0s0, k0)
	}
	sort.Strings(k0s0)
	for _, k0 := range k0s0 {
		if v.Encoding[k0] != nil {
			if err := s.walkEncoding(v.Encoding[k0]); err != nil {
				return err
			}
		}
	}

	k0s2 := make([]string, 0, len(v.Examples))
	for k0 := range v.Examples {
		k0s2 = append(k0s2, k0)
	}
	sort.Strings(k0s2)
	for _, k0 := range k0s2 {
		if v.Examples[k0] != nil {
			if err := s.walkExampleRef(v.Examples[k0]); err != nil {
				return err
			}
		}
	}

	if err := s.walkSchemaRef(&v.Schema); err != nil {
		return err
	}

	return nil
}

func (s *scanRefType) walkNodeRef(v *openapi.Ref[yaml.Node]) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	if v.HasRef() {
		if _, exist := s.cutRefs[v.RefFile]; exist {
		} else {
			s.refs[v.RefFile] = struct{}{}
			return nil
		}
	}

	return nil
}

func (s *scanRefType) walkOperation(v *openapi.Operation) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	k0s0 := make([]string, 0, len(v.Callbacks))
	for k0 := range v.Callbacks {
		k0s0 = append(k0s0, k0)
	}
	sort.Strings(k0s0)
	for _, k0 := range k0s0 {
		if v.Callbacks[k0] != nil {
			if err := s.walkCallbackRef(v.Callbacks[k0]); err != nil {
				return err
			}
		}
	}

	for i0 := range v.Parameters {
		if err := s.walkParameterRef(&v.Parameters[i0]); err != nil {
			return err
		}
	}

	if err := s.walkRequestBodyRef(&v.RequestBody); err != nil {
		return err
	}

	k0s7 := make([]string, 0, len(v.Responses))
	for k0 := range v.Responses {
		k0s7 = append(k0s7, k0)
	}
	sort.Strings(k0s7)
	for _, k0 := range k0s7 {
		if v.Responses[k0] != nil {
			if err := s.walkResponseRef(v.Responses[k0]); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *scanRefType) walkParameter(v *openapi.Parameter) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	k0s2 := make([]string, 0, len(v.Content))
	for k0 := range v.Content {
		k0s2 = append(k0s2, k0)
	}
	sort.Strings(k0s2)
	for _, k0 := range k0s2 {
		if v.Content[k0] != nil {
			if err := s.walkMediaType(v.Content[k0]); err != nil {
				return err
			}
		}
	}

	k0s6 := make([]string, 0, len(v.Examples))
	for k0 := range v.Examples {
		k0s6 = append(k0s6, k0)
	}
	sort.Strings(k0s6)
	for _, k0 := range k0s6 {
		if v.Examples[k0] != nil {
			if err := s.walkExampleRef(v.Examples[k0]); err != nil {
				return err
			}
		}
	}

	if err := s.walkSchemaRef(&v.Schema); err != nil {
		return err
	}

	return nil
}

func (s *scanRefType) walkParameterRef(v *openapi.Ref[openapi.Parameter]) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	if v.HasRef() {
		if _, exist := s.cutRefs[v.RefFile]; exist {
		} else {
			s.refs[v.RefFile] = struct{}{}
			return nil
		}
	}

	if v.Value != nil {
		if err := s.walkParameter(v.Value); err != nil {
			return err
		}
	}

	return nil
}

func (s *scanRefType) walkPathItemBase(v *openapi.PathItemBase) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	if v.Delete != nil {
		if err := s.walkOperation(v.Delete); err != nil {
			return err
		}
	}

	if v.Get != nil {
		if err := s.walkOperation(v.Get); err != nil {
			return err
		}
	}

	if v.Head != nil {
		if err := s.walkOperation(v.Head); err != nil {
			return err
		}
	}

	if v.Options != nil {
		if err := s.walkOperation(v.Options); err != nil {
			return err
		}
	}

	for i0 := range v.Parameters {
		if err := s.walkParameterRef(&v.Parameters[i0]); err != nil {
			return err
		}
	}

	if v.Patch != nil {
		if err := s.walkOperation(v.Patch); err != nil {
			return err
		}
	}

	if v.Post != nil {
		if err := s.walkOperation(v.Post); err != nil {
			return err
		}
	}

	if v.Put != nil {
		if err := s.walkOperation(v.Put); err != nil {
			return err
		}
	}

	if v.Trace != nil {
		if err := s.walkOperation(v.Trace); err != nil {
			return err
		}
	}

	return nil
}

func (s *scanRefType) walkPathItemBaseRef(v *openapi.Ref[openapi.PathItemBase]) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	if v.Value != nil {
		if err := s.walkPathItemBase(v.Value); err != nil {
			return err
		}
	}

	return nil
}

func (s *scanRefType) walkRequestBody(v *openapi.RequestBody) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	k0s0 := make([]string, 0, len(v.Content))
	for k0 := range v.Content {
		k0s0 = append(k0s0, k0)
	}
	sort.Strings(k0s0)
	for _, k0 := range k0s0 {
		if v.Content[k0] != nil {
			if err := s.walkMediaType(v.Content[k0]); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *scanRefType) walkRequestBodyRef(v *openapi.Ref[openapi.RequestBody]) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	if v.HasRef() {
		if _, exist := s.cutRefs[v.RefFile]; exist {
		} else {
			s.refs[v.RefFile] = struct{}{}
			return nil
		}
	}

	if v.Value != nil {
		if err := s.walkRequestBody(v.Value); err != nil {
			return err
		}
	}

	return nil
}

func (s *scanRefType) walkResponse(v *openapi.Response) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	k0s0 := make([]string, 0, len(v.Content))
	for k0 := range v.Content {
		k0s0 = append(k0s0, k0)
	}
	sort.Strings(k0s0)
	for _, k0 := range k0s0 {
		if v.Content[k0] != nil {
			if err := s.walkMediaType(v.Content[k0]); err != nil {
				return err
			}
		}
	}

	k0s2 := make([]string, 0, len(v.Headers))
	for k0 := range v.Headers {
		k0s2 = append(k0s2, k0)
	}
	sort.Strings(k0s2)
	for _, k0 := range k0s2 {
		if v.Headers[k0] != nil {
			if err := s.walkHeaderRef(v.Headers[k0]); err != nil {
				return err
			}
		}
	}

	k0s3 := make([]string, 0, len(v.Links))
	for k0 := range v.Links {
		k0s3 = append(k0s3, k0)
	}
	sort.Strings(k0s3)
	for _, k0 := range k0s3 {
		if v.Links[k0] != nil {
			if err := s.walkLinkRef(v.Links[k0]); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *scanRefType) walkResponseRef(v *openapi.Ref[openapi.Response]) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	if v.HasRef() {
		if _, exist := s.cutRefs[v.RefFile]; exist {
		} else {
			s.refs[v.RefFile] = struct{}{}
			return nil
		}
	}

	if v.Value != nil {
		if err := s.walkResponse(v.Value); err != nil {
			return err
		}
	}

	return nil
}

func (s *scanRefType) walkSchema(v *openapi.Schema) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	if err := s.walkAdditionalProperties(&v.AdditionalProperties); err != nil {
		return err
	}

	for i0 := range v.AllOf {
		if err := s.walkSchemaRef(&v.AllOf[i0]); err != nil {
			return err
		}
	}

	for i0 := range v.AnyOf {
		if err := s.walkSchemaRef(&v.AnyOf[i0]); err != nil {
			return err
		}
	}

	if err := s.walkSchemaRef(&v.Items); err != nil {
		return err
	}

	if err := s.walkSchemaRef(&v.Not); err != nil {
		return err
	}

	for i0 := range v.OneOf {
		if err := s.walkSchemaRef(&v.OneOf[i0]); err != nil {
			return err
		}
	}

	k0s28 := make([]string, 0, len(v.Properties))
	for k0 := range v.Properties {
		k0s28 = append(k0s28, k0)
	}
	sort.Strings(k0s28)
	for _, k0 := range k0s28 {
		if v.Properties[k0] != nil {
			if err := s.walkSchemaRef(v.Properties[k0]); err != nil {
				return err
			}
		}
	}

	if err := fixSkipOptionalPointer(v); err != nil {
		return err
	}

	if err := fixIntegerFormat(v); err != nil {
		return err
	}

	if err := fixAnyOfEnum(v); err != nil {
		return err
	}

	if err := fixAnyOfString(v); err != nil {
		return err
	}

	if err := fixNullable(v); err != nil {
		return err
	}

	if err := fixImplicitArray(v); err != nil {
		return err
	}

	if err := fixEliminateCheckerUnion(v); err != nil {
		return err
	}

	if err := fixAdditionalProperties(v); err != nil {
		return err
	}

	return nil
}

func (s *scanRefType) walkSchemaRef(v *openapi.Ref[openapi.Schema]) error {
	if _, exist := s.visited[v]; exist {
		return nil
	}
	s.visited[v] = struct{}{}

	if v.HasRef() {
		if _, exist := s.cutRefs[v.RefFile]; exist {
			if err := fixCutSchemaRef(v); err != nil {
				return err
			}
			return nil
		} else {
			s.refs[v.RefFile] = struct{}{}
			return nil
		}
	}

	if v.Value != nil {
		if err := s.walkSchema(v.Value); err != nil {
			return err
		}
	}

	return nil
}
