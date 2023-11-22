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

package writer

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/tools/imports"
)

func init() {
	imports.LocalPrefix = "github.com/ShouheiNishi/openapi5g"
}

// Writer and formatter of golang source file
type outputFile struct {
	*bytes.Buffer
	fileName string
}

type ImportSpec struct {
	PackageName string
	ImportPath  string
}

type ImportSpecs []ImportSpec

func NewOutputFile(fileName string, pkgName string, generatorName string, imports ImportSpecs) *outputFile {
	o := outputFile{
		Buffer:   new(bytes.Buffer),
		fileName: fileName,
	}

	fmt.Fprintf(o, "// Code generated by %s, DO NOT EDIT.\n", generatorName)
	fmt.Fprintln(o, "")
	fmt.Fprintf(o, "package %s\n", pkgName)
	fmt.Fprintln(o, "")
	if len(imports) > 0 {
		fmt.Fprintln(o, "import (")
		for _, s := range imports {
			if s.ImportPath == "" {
				fmt.Fprintln(o, "")
			} else {
				fmt.Fprintf(o, "%s \"%s\"\n", s.PackageName, s.ImportPath)
			}
		}
		fmt.Fprintln(o, ")")
		fmt.Fprintln(o, "")
	}

	return &o
}

func (o *outputFile) Close() (err error) {
	// Output to file
	out, err := imports.Process(o.fileName, o.Bytes(), nil)
	if err != nil {
		return fmt.Errorf("format error: %w\n%s\n", err, o.String())
	}
	fWrite, err := os.Create(o.fileName)
	if err != nil {
		return err
	}
	defer func() {
		errClose := fWrite.Close()
		if errClose != nil && err == nil {
			err = errClose
		}
	}()
	_, err = fWrite.Write(out)
	if err != nil {
		return err
	}
	return nil
}
