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
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/ShouheiNishi/openapi5g/internal/generator/writer"
)

type OpenApi3WalkerGeneratorConfig struct {
	FileName      string
	PkgName       string
	GeneratorName string
	Imports       writer.ImportSpecs

	RootFuncName      string
	ExtraRootArgs     string
	ExtraInit         string
	ExtraWalkArgsInit string

	StateType  string
	ExtraState string

	ExtraWalkArgs     string
	ExtraWalkArgsCall string
	WalkPreHook       func(t reflect.Type) string
	WalkPostHook      func(t reflect.Type) string
}

type openApi3WalkerGeneratorState struct {
	*OpenApi3WalkerGeneratorConfig
	out   io.WriteCloser
	funcs map[string]string
}

const generatedFunctionPrefix = "walk"

var dummyForAnyReflectType interface{}
var anyReflectType = reflect.TypeOf(&dummyForAnyReflectType).Elem()

func typeName(t reflect.Type) string {
	n := t.Name()
	if n == "" {
		panic("empty type name")
	}
	p := t.PkgPath()
	if p == "" {
		return n
	}
	sp := strings.Split(p, "/")
	return sp[len(sp)-1] + "." + n
}

func (s *openApi3WalkerGeneratorState) generateValueOp(n string, t reflect.Type, depth int, id int) string {
	switch t.Kind() {
	case reflect.Struct:
		if !strings.HasSuffix(t.PkgPath(), "/openapi3") {
			return ""
		}
		s.generateFuncs(t)
		if s.funcs[generatedFunctionPrefix+t.Name()] == "*" {
			return ""
		}
		n2 := "&" + n
		if strings.HasPrefix(n, "*") {
			n2 = strings.TrimPrefix(n, "*")
		}
		return fmt.Sprintf("if err := s.%s(%s %s) ; err != nil {\nreturn err\n}", generatedFunctionPrefix+t.Name(), n2, s.ExtraWalkArgsCall)
	case reflect.Bool:
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Uintptr:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Complex64:
	case reflect.Complex128:
	case reflect.String:
		return ""
	case reflect.Interface:
		return fmt.Sprintf("if %s != nil {panic(\"Non nil interface\")}", n)
	case reflect.Pointer:
		str := s.generateValueOp("*"+n, t.Elem(), depth+1, id)
		if str == "" {
			return ""
		}
		return fmt.Sprintf("if %s != nil {\n%s\n}", n, str)
	case reflect.Slice:
		n2 := n
		if strings.HasPrefix(n, "*") {
			n2 = "(" + n + ")"
		}
		str := s.generateValueOp(fmt.Sprintf("%s[i%d]", n2, depth), t.Elem(), depth+1, id)
		if str == "" {
			return ""
		}
		return fmt.Sprintf("for i%d := range %s {\n%s\n}", depth, n, str)
	case reflect.Map:
		n2 := n
		if strings.HasPrefix(n, "*") {
			n2 = "(" + n + ")"
		}
		str := s.generateValueOp(fmt.Sprintf("%s[k%d]", n2, depth), t.Elem(), depth+1, id)
		if str == "" {
			return ""
		}
		if t.Key() == reflect.TypeOf(string("")) {
			return fmt.Sprintf("k%ds%d := make([]string, 0, len(%s))\nfor k%d := range %s {\nk%ds%d = append(k%ds%d, k%d)\n}\n"+
				"sort.Strings(k%ds%d)\nfor _, k%d := range k%ds%d {\n%s\n}",
				depth, id, n, depth, n, depth, id, depth, id, depth,
				depth, id, depth, depth, id, str)
		} else {
			return fmt.Sprintf("for k%d := range %s {\n%s\n}", depth, n, str)
		}
	default:
		panic(fmt.Sprintf("%s is unsupported type %s", n, t))
	}
	return ""
}

func (s *openApi3WalkerGeneratorState) generateFuncs(t reflect.Type) {
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("%s is not struct", t))
	}

	fname := "walk" + t.Name()

	if _, exist := s.funcs[fname]; exist {
		return
	}
	s.funcs[fname] = ""

	str := fmt.Sprintf("func (s *%s) %s(v *%s %s) error {\n", s.StateType, fname, typeName(t), s.ExtraWalkArgs)
	str += "if _, exist := s.visited[v] ; exist {\nreturn nil\n}\ns.visited[v] = struct{}{}\n"

	var elems []string
fieldLoop:
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		switch f.Name {
		case "Default":
			if t == reflect.TypeOf(openapi3.Schema{}) && f.Type == anyReflectType {
				continue fieldLoop
			}
		case "Enum":
			if t == reflect.TypeOf(openapi3.Schema{}) && f.Type == reflect.TypeOf([]interface{}{}) {
				continue fieldLoop
			}
		case "Example":
			if (t == reflect.TypeOf(openapi3.MediaType{}) ||
				t == reflect.TypeOf(openapi3.Parameter{}) ||
				t == reflect.TypeOf(openapi3.Schema{})) &&
				f.Type == anyReflectType {
				continue fieldLoop
			}
		case "Extensions":
			if f.Type == reflect.TypeOf(map[string]interface{}{}) {
				continue fieldLoop
			}
		case "Parameters":
			if t == reflect.TypeOf(openapi3.Link{}) && f.Type == reflect.TypeOf(map[string]interface{}{}) {
				continue fieldLoop
			}
		case "RequestBody":
			if t == reflect.TypeOf(openapi3.Link{}) && f.Type == anyReflectType {
				continue fieldLoop
			}
		case "Value":
			if t == reflect.TypeOf(openapi3.Example{}) && f.Type == anyReflectType {
				continue fieldLoop
			}
		}
		elems = append(elems, f.Name)
	}

	sort.Strings(elems)
	empty := true

	if s.WalkPreHook != nil {
		if hs := s.WalkPreHook(t); hs != "" {
			str += "\n" + hs
			empty = false
		}
	}

	for i, n := range elems {
		f, _ := t.FieldByName(n)
		if s2 := s.generateValueOp("v."+n, f.Type, 0, i); s2 != "" {
			str += "\n" + s2 + "\n"
			empty = false
		}
	}

	if s.WalkPostHook != nil {
		if hs := s.WalkPostHook(t); hs != "" {
			str += "\n" + hs
			empty = false
		}
	}

	str += "\nreturn nil\n}\n"

	if empty {
		s.funcs[fname] = "*"
		return
	}
	s.funcs[fname] = str
}

func OpenApi3WalkerGenerate(c *OpenApi3WalkerGeneratorConfig) error {
	s := openApi3WalkerGeneratorState{
		OpenApi3WalkerGeneratorConfig: c,
		out:                           writer.NewOutputFile(c.FileName, c.PkgName, c.GeneratorName, c.Imports),
		funcs:                         make(map[string]string),
	}

	fmt.Fprintf(s.out, "type %s struct{", c.StateType)
	fmt.Fprintln(s.out, "visited map[interface{}]struct{}")
	fmt.Fprintln(s.out, "doc *openapi3.T")
	fmt.Fprint(s.out, s.ExtraState)
	fmt.Fprintln(s.out, "}")
	fmt.Fprintln(s.out, "")

	fmt.Fprintf(s.out, "func %s(d *openapi3.T %s) error {\n", c.RootFuncName, c.ExtraRootArgs)
	fmt.Fprintf(s.out, "s := &%s {\n", c.StateType)
	fmt.Fprintln(s.out, "visited: make(map[interface{}]struct{}),")
	fmt.Fprintln(s.out, "doc: d,")
	fmt.Fprintln(s.out, "}")
	if c.ExtraInit != "" {
		fmt.Fprintf(s.out, "\n%s", c.ExtraInit)
	}
	fmt.Fprintln(s.out, "")
	fmt.Fprintf(s.out, "return s.%sT(s.doc %s)\n", generatedFunctionPrefix, s.ExtraWalkArgsInit)
	fmt.Fprintln(s.out, "}")
	fmt.Fprintln(s.out, "")

	oldEmptyCount := 0
	for {
		s.generateFuncs(reflect.TypeOf(openapi3.T{}))
		emptyCount := 0
		for n := range s.funcs {
			if s.funcs[n] == "*" {
				emptyCount++
			} else {
				delete(s.funcs, n)
			}
		}
		if emptyCount == oldEmptyCount {
			break
		}
		oldEmptyCount = emptyCount
	}
	s.generateFuncs(reflect.TypeOf(openapi3.T{}))
	fn := make([]string, 0, len(s.funcs))
	for n := range s.funcs {
		fn = append(fn, n)
	}
	sort.Strings(fn)
	for _, n := range fn {
		if s.funcs[n] != "*" {
			fmt.Fprint(s.out, s.funcs[n])
			fmt.Fprint(s.out, "\n\n")
		}
	}

	return s.out.Close()
}
