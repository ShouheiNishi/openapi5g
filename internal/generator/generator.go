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
	"net/url"
	"path"
	"path/filepath"
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

func Generate(rootDir string) error {
	outLists, deps, err := RewriteYamlAndGenerateConfig(rootDir)
	if err != nil {
		return fmt.Errorf("GenerateOapiCodeGenConfig: %w", err)
	}

	var outs []string
	outs, err = GenerateEmbed(rootDir, deps)
	if err != nil {
		return fmt.Errorf("GenerateEmbed: %w", err)
	}
	outLists = append(outLists, outs...)

	outs, err = GenerateLoaderTest(rootDir)
	if err != nil {
		return fmt.Errorf("GenerateLoaderTest: %w", err)
	}
	outLists = append(outLists, outs...)

	outs, err = GeneratePkgMap(rootDir)
	if err != nil {
		return fmt.Errorf("GeneratePkgMap: %w", err)
	}
	outLists = append(outLists, outs...)

	err = RemoveOldFiles(rootDir, outLists)
	if err != nil {
		return fmt.Errorf("RemoveOldFiles: %w", err)
	}

	return nil
}

func RewriteYamlAndGenerateConfig(rootDir string) (outLists []string, depsAll []string, err error) {
	dep := make(map[string]struct{})

	for spec := range pkgList {
		doc, depsOne, err := LoadOpenApi3Spec(rootDir, spec)
		if err != nil {
			return nil, nil, fmt.Errorf("LoadOpenApi3Spec: %w", err)
		}
		for _, d := range depsOne {
			dep[d] = struct{}{}
		}

		var outs []string
		var depsRewrite []string
		outs, depsRewrite, err = RewriteYaml(rootDir, spec, doc)
		if err != nil {
			return nil, nil, fmt.Errorf("RewriteYaml: %w", err)
		}
		outLists = append(outLists, outs...)

		outs, err = GenerateConfig(rootDir, spec, doc, depsRewrite)
		if err != nil {
			return nil, nil, fmt.Errorf("GenerateConfig: %w", err)
		}
		outLists = append(outLists, outs...)

		outs, err = GenerateLoader(rootDir, spec, depsOne)
		if err != nil {
			return nil, nil, fmt.Errorf("GenerateLoader(: %w", err)
		}
		outLists = append(outLists, outs...)
	}

	for d := range dep {
		depsAll = append(depsAll, d)
	}
	sort.Strings(depsAll)

	return outLists, depsAll, nil
}

func LoadOpenApi3Spec(rootDir string, spec string) (doc *openapi3.T, deps []string, err error) {
	d := make(map[string]struct{})
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		d[path.Base(url.Path)] = struct{}{}
		return openapi3.DefaultReadFromURI(loader, url)
	}
	doc, err = loader.LoadFromFile(filepath.Join(rootDir, "specs", spec))
	if err != nil {
		return nil, nil, err
	}

	for s := range d {
		deps = append(deps, s)
	}
	sort.Strings(deps)

	return doc, deps, nil
}
