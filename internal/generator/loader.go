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
	"path"
	"sort"
	"strings"

	"github.com/ShouheiNishi/openapi5g/internal/generator/writer"
)

func (s *GeneratorState) MakeDepsForLoader() error {
	s.DepsForLoader = make(map[string][]string)
	for spec := range pkgList {
		depsLoaderMap := map[string]struct{}{spec: struct{}{}}
		for {
			populated := false
			for spec2 := range depsLoaderMap {
				for spec3 := range s.DepsBase[spec2] {
					if _, exist := depsLoaderMap[spec3]; !exist {
						depsLoaderMap[spec3] = struct{}{}
						populated = true
					}
				}
			}
			if !populated {
				break
			}
		}

		depsLoader := make([]string, 0, len(depsLoaderMap))
		for spec2 := range depsLoaderMap {
			depsLoader = append(depsLoader, spec2)
		}
		sort.Strings(depsLoader)
		s.DepsForLoader[spec] = depsLoader
	}

	return nil
}

func (s *GeneratorState) GenerateLoader(spec string) (err error) {
	depsOne := s.DepsForLoader[spec]
	imports := writer.ImportSpecs{
		{ImportPath: "fmt"},
		{ImportPath: "net/url"},
		{ImportPath: "path"},
		{ImportPath: "sync"},
		{},
		{ImportPath: "github.com/getkin/kin-openapi/openapi3"},
	}

	for _, d := range depsOne {
		imports = append(imports, writer.ImportSpec{
			ImportPath: modBase + "/internal/embed/" + strings.TrimSuffix(d, ".yaml"),
		})
	}

	name, err := s.CreateFileName(pkgList[spec].path, "loader.go")
	if err != nil {
		return fmt.Errorf("CreateFileName: %w", err)
	}
	out := writer.NewOutputFile(name, path.Base(pkgList[spec].path), generatorName, imports)

	fmt.Fprintln(out, "var specTable map[string][]byte=map[string][]byte{")
	for _, d := range depsOne {
		fmt.Fprintf(out, "\"%s\": %s.SpecYaml,\n", d, strings.TrimSuffix(d, ".yaml"))
	}
	fmt.Fprintln(out, "}")

	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "var kinOpenApi3Doc *openapi3.T")
	fmt.Fprintln(out, "var kinOpenApi3Err error")
	fmt.Fprintln(out, "var kinOpenApi3Once sync.Once")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "func GetKinOpenApi3Document() (*openapi3.T, error) {")
	fmt.Fprintln(out, "kinOpenApi3Once.Do(func() {")
	fmt.Fprintln(out, "loader := openapi3.NewLoader()")
	fmt.Fprintln(out, "loader.IsExternalRefsAllowed = true")
	fmt.Fprintln(out, "loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {")
	fmt.Fprintln(out, "spec := specTable[path.Base(url.Path)]")
	fmt.Fprintln(out, "if spec == nil {")
	fmt.Fprintf(out, "return nil, fmt.Errorf(\"%%s is missing\", url.String())\n")
	fmt.Fprintln(out, "}")
	fmt.Fprintln(out, "return spec, nil")
	fmt.Fprintln(out, "}")
	fmt.Fprintf(out, "kinOpenApi3Doc, kinOpenApi3Err = loader.LoadFromFile(\"%s\")\n", spec)
	fmt.Fprintln(out, "})")
	fmt.Fprintln(out, "return kinOpenApi3Doc, kinOpenApi3Err")
	fmt.Fprintln(out, "}")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "func GetKinOpenApi3DocumentMust() *openapi3.T {")
	fmt.Fprintln(out, "doc, err := GetKinOpenApi3Document()")
	fmt.Fprintln(out, "if err != nil {")
	fmt.Fprintln(out, "panic(err)")
	fmt.Fprintln(out, "}")
	fmt.Fprintln(out, "return doc")
	fmt.Fprintln(out, "}")

	if err := out.Close(); err != nil {
		return err
	}

	return nil
}

func (s *GeneratorState) GenerateLoaderTest() error {
	l := make([]string, 0, len(pkgList))
	for spec := range pkgList {
		l = append(l, spec)
	}
	sort.Strings(l)

	imp := writer.ImportSpecs{
		{ImportPath: "context"},
		{ImportPath: "testing"},
		{},
		{ImportPath: "github.com/getkin/kin-openapi/openapi3"},
		{ImportPath: "github.com/stretchr/testify/assert"},
	}
	for _, spec := range l {
		imp = append(imp, writer.ImportSpec{
			PackageName: strings.TrimSuffix(spec, ".yaml"),
			ImportPath:  modBase + "/" + pkgList[spec].path,
		})
	}
	name, err := s.CreateFileName("internal/test/loader_gen_test.go")
	if err != nil {
		return fmt.Errorf("CreateFileName: %w", err)
	}
	f := writer.NewOutputFile(name, "test_test", generatorName, imp)

	fmt.Fprintf(f, "func TestLoader(t *testing.T) {\n")
	fmt.Fprintf(f, "var doc *openapi3.T\n")
	fmt.Fprintf(f, "var err error\n")
	for _, spec := range l {
		sbase := strings.TrimSuffix(spec, ".yaml")
		fmt.Fprintf(f, "\n")
		fmt.Fprintf(f, "doc, err = %s.GetKinOpenApi3Document()\n", sbase)
		fmt.Fprintf(f, "assert.NoError(t, err)\n")
		fmt.Fprintf(f, "err = doc.Validate(context.Background())\n")
		fmt.Fprintf(f, "assert.NoError(t, err)\n")
	}
	fmt.Fprintf(f, "}\n")
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
