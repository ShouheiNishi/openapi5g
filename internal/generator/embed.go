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
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ShouheiNishi/openapi5g/internal/generator/writer"
)

func GenerateEmbed(rootDir string, deps []string) (outLists []string, err error) {
	for _, d := range deps {
		base := strings.TrimSuffix(d, ".yaml")
		dir := path.Join(rootDir, "internal/embed", base)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, err
		}

		name := path.Join(dir, "embed.go")
		out := writer.NewOutputFile(name, base, generatorName, writer.ImportSpecs{
			{PackageName: "_", ImportPath: "embed"},
		})
		outLists = append(outLists, name)
		fmt.Fprintf(out, "//go:embed %s\n", d)
		fmt.Fprintln(out, "var SpecYaml []byte")
		if err := out.Close(); err != nil {
			return nil, err
		}

		if fIn, err := os.Open(filepath.Join(rootDir, "specs", d)); err != nil {
			return nil, err
		} else {
			defer fIn.Close()
			name := path.Join(dir, d)
			outLists = append(outLists, name)
			if fOut, err := os.Create(name); err != nil {
				return nil, err
			} else {
				defer fOut.Close()
				if _, err := io.Copy(fOut, fIn); err != nil {
					return nil, err
				}
			}
		}
	}

	return outLists, nil
}
