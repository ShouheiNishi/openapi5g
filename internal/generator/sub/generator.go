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

package main

import (
	"reflect"

	"github.com/ShouheiNishi/openapi5g/internal/generator/writer"
	"github.com/getkin/kin-openapi/openapi3"
)

func main() {
	err := OpenApi3WalkerGenerate(&OpenApi3WalkerGeneratorConfig{
		FileName:      "fix_local_refs.go",
		PkgName:       "generator",
		GeneratorName: "github.com/ShouheiNishi/openapi5g/internal/generator/sub",
		Imports: writer.ImportSpecs{
			{ImportPath: "sort"},
			{ImportPath: "strings"},
			{},
			{ImportPath: "github.com/getkin/kin-openapi/openapi3"},
		},

		RootFuncName:      "fixLocalRefInRemoteRef",
		ExtraRootArgs:     ",rootUri string",
		ExtraInit:         "s.rootUri = rootUri\n",
		ExtraWalkArgsInit: ",rootUri",

		StateType:  "fixLocalRefIType",
		ExtraState: "rootUri string\n",

		ExtraWalkArgs:     ",curUri string",
		ExtraWalkArgsCall: ",curUri",
		WalkPreHook: func(t reflect.Type) string {
			if _, exist := t.FieldByName("Ref"); exist {
				return `if v.Ref != "" {
					split := strings.SplitN(v.Ref, "#", 2)
					if len(split) == 2 {
						if split[0] != "" {
							curUri = split[0]
						}
						if curUri == s.rootUri {
							v.Ref = "#" + split[1]
						} else {
							v.Ref = curUri + "#" + split[1]
						}
					} else {
						curUri = split[0]
					}
				}
`
			}
			return ""
		},
	})

	if err != nil {
		panic(err)
	}

	err = OpenApi3WalkerGenerate(&OpenApi3WalkerGeneratorConfig{
		FileName:      "scan_refs.go",
		PkgName:       "generator",
		GeneratorName: "github.com/ShouheiNishi/openapi5g/internal/generator/sub",
		Imports: writer.ImportSpecs{
			{ImportPath: "sort"},
			{ImportPath: "strings"},
			{},
			{ImportPath: "github.com/getkin/kin-openapi/openapi3"},
		},

		RootFuncName:  "scanRef",
		ExtraRootArgs: ", refs map[string]struct{}, cutRefs map[string]struct{}",
		ExtraInit:     "s.refs = refs\ns.cutRefs = cutRefs\n",

		StateType:  "scanRefType",
		ExtraState: "refs map[string]struct{}\ncutRefs map[string]struct{}\n",

		WalkPreHook: func(t reflect.Type) string {
			if t != reflect.TypeOf(openapi3.PathItem{}) {
				if _, exist := t.FieldByName("Ref"); exist {
					cutOp := ""
					if t == reflect.TypeOf(openapi3.SchemaRef{}) {
						cutOp = "if err := fixCutSchemaRef(v) ; err != nil{return err}\nreturn nil\n"
					}
					return `	if v.Ref != "" {
						split := strings.SplitN(v.Ref, "#", 2)
						if len(split) == 2 && split[0] != "" {
							if _, exist := s.cutRefs[split[0]]; exist {
` + cutOp +
						`	} else {
								s.refs[split[0]] = struct{}{}
								return nil
							}
						}
					}
`
				}
			}
			return ""
		},

		WalkPostHook: func(t reflect.Type) string {
			if t == reflect.TypeOf(openapi3.Schema{}) {
				return "if err := fixSkipOptionalPointer(v) ; err != nil{return err}\n" +
					"\nif err := fixIntegerFormat(v) ; err != nil{return err}\n" +
					"\nif err := fixAllOfEnum(v) ; err != nil{return err}\n" +
					"\nif err := fixImplicitArray(v) ; err != nil{return err}\n" +
					"\nif err := fixEliminateCheckerUnion(v) ; err != nil{return err}\n" +
					"\nif err := fixAdditionalProperties(v) ; err != nil{return err}\n"
			}
			return ""
		},

		// vRef := vObj.FieldByName("Ref")
		// vValue := vObj.FieldByName("Value")
		// if vRef.IsValid() && vValue.IsValid() {
		// 	ref := vRef.String()
		// 	if ref != "" {
		// 		split := strings.SplitN(ref, "#", 2)
		// 		if len(split) == 2 && split[0] != "" {
		// 			if _, exist := cutRefs[split[0]]; !exist {
		// 				refs[split[0]] = struct{}{}
		// 				return
		// 			}
		// 		}
		// 	}
		// }
	})

	if err != nil {
		panic(err)
	}
}
