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

package main

import (
	"go/types"

	"github.com/ShouheiNishi/openapi5g/internal/generator/writer"
)

func main() {
	if err := LoadPackage("../.."); err != nil {
		panic(err)
	}

	if err := OpenApi3WalkerGenerate(&OpenApi3WalkerGeneratorConfig{
		FileName:      "setup_refs.go",
		PkgName:       "generator",
		GeneratorName: "github.com/ShouheiNishi/openapi5g/internal/generator/sub",
		Imports: writer.ImportSpecs{
			{ImportPath: "github.com/ShouheiNishi/openapi5g/internal/generator/openapi"},
			{ImportPath: "gopkg.in/yaml.v3"},
		},

		RootFuncName:  "setupRefs",
		ExtraRootArgs: ", curFile string, deps map[string]struct{}",
		ExtraInit:     "s.curFile = curFile\ns.deps = deps\n",

		StateType:  "setupRefsType",
		ExtraState: "curFile string\ndeps map[string]struct{}\n",

		WalkPreHook: func(t *types.Named) string {
			if t.Obj().Name() == "Ref" {
				return `if v.Ref != "" {
					split := strings.SplitN(v.Ref, "#", 2)
					if len(split) == 2 {
						if split[0] == "" {
							v.Ref = s.curFile + "#" + split[1]
						} else {
							s.deps[split[0]] = struct{}{}
						}
					}
					return nil
				}
`
			}
			return ""
		},
	}); err != nil {
		panic(err)
	}

	if err := OpenApi3WalkerGenerate(&OpenApi3WalkerGeneratorConfig{
		FileName:      "resolve_refs.go",
		PkgName:       "generator",
		GeneratorName: "github.com/ShouheiNishi/openapi5g/internal/generator/sub",
		Imports: writer.ImportSpecs{
			{ImportPath: "github.com/ShouheiNishi/openapi5g/internal/generator/openapi"},
			{ImportPath: "gopkg.in/yaml.v3"},
		},

		RootFuncName:  "resolveRefs",
		ExtraRootArgs: ",gs *GeneratorState",
		ExtraInit:     "s.generatorState = gs\n",

		StateType:  "resolveRefsType",
		ExtraState: "generatorState *GeneratorState\n",

		WalkPreHook: func(t *types.Named) string {
			if t.Obj().Name() == "Ref" {
				return `if v.Ref != "" {
					if vRes, err := ResolveRef[` + typeName(t.TypeArgs().At(0).(*types.Named)) + `](s.generatorState, v.Ref); err != nil {
						return fmt.Errorf("ResolveRef(%s): %w", v.Ref, err)
					} else {
						v.Value = vRes
						return nil
					}
				}
`
			}
			return ""
		},
	}); err != nil {
		panic(err)
	}

	if err := OpenApi3WalkerGenerate(&OpenApi3WalkerGeneratorConfig{
		FileName:      "post_refs.go",
		PkgName:       "generator",
		GeneratorName: "github.com/ShouheiNishi/openapi5g/internal/generator/sub",
		Imports: writer.ImportSpecs{
			{ImportPath: "github.com/ShouheiNishi/openapi5g/internal/generator/openapi"},
			{ImportPath: "gopkg.in/yaml.v3"},
		},

		RootFuncName:  "postRefs",
		ExtraRootArgs: ", curFile *string",
		ExtraInit:     "s.curFile = curFile\n",

		StateType:  "postRefsType",
		ExtraState: "curFile *string\n",

		WalkPreHook: func(t *types.Named) string {
			if t.Obj().Name() == "Ref" {
				return `v.CurFile = s.curFile
				if v.Ref != "" {
					return nil
				}
`
			}
			return ""
		},
	}); err != nil {
		panic(err)
	}

	if err := OpenApi3WalkerGenerate(&OpenApi3WalkerGeneratorConfig{
		FileName:      "scan_refs.go",
		PkgName:       "generator",
		GeneratorName: "github.com/ShouheiNishi/openapi5g/internal/generator/sub",
		Imports: writer.ImportSpecs{
			{ImportPath: "github.com/ShouheiNishi/openapi5g/internal/generator/openapi"},
			{ImportPath: "gopkg.in/yaml.v3"},
		},

		RootFuncName:  "scanRef",
		ExtraRootArgs: ", refs map[string]struct{}, cutRefs map[string]struct{}",
		ExtraInit:     "s.refs = refs\ns.cutRefs = cutRefs\n",

		StateType:  "scanRefType",
		ExtraState: "refs map[string]struct{}\ncutRefs map[string]struct{}\n",

		WalkPreHook: func(t *types.Named) string {
			if t.Obj().Name() == "Ref" {
				if t.TypeArgs().At(0).(*types.Named).Obj().Name() != "PathItemBase" {
					cutOp := ""
					if t.TypeArgs().At(0).(*types.Named).Obj().Name() == "Schema" {
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

		WalkPostHook: func(t *types.Named) string {
			if t.Obj().Name() == "Schema" {
				return "if err := fixSkipOptionalPointer(v) ; err != nil{return err}\n" +
					"\nif err := fixIntegerFormat(v) ; err != nil{return err}\n" +
					"\nif err := fixAnyOfEnum(v) ; err != nil{return err}\n" +
					"\nif err := fixAnyOfString(v) ; err != nil{return err}\n" +
					"\nif err := fixNullable(v) ; err != nil{return err}\n" +
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
	}); err != nil {
		panic(err)
	}
}
