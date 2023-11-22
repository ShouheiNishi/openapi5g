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
	"io/fs"
	"os"
	"path/filepath"
)

func RemoveOldFiles(rootDir string, lists []string) error {
	e := make(map[string]struct{})
	for _, p := range lists {
		e[p] = struct{}{}
	}

	return filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path == filepath.Join(rootDir, ".git") {
			if d.IsDir() {
				return filepath.SkipDir
			} else {
				return nil
			}
		}

		if path == filepath.Join(rootDir, ".github") {
			return filepath.SkipDir
		}

		if path == filepath.Join(rootDir, "specs") {
			return filepath.SkipDir
		}

		if path == filepath.Join(rootDir, "utils") {
			return filepath.SkipDir
		}

		if path == filepath.Join(rootDir, "internal/generator") {
			return filepath.SkipDir
		}

		if path == filepath.Join(rootDir, "internal/generate.go") {
			return nil
		}

		if path == filepath.Join(rootDir, ".gitmodules") {
			return nil
		}

		if path == filepath.Join(rootDir, ".gitignore") {
			return nil
		}

		if path == filepath.Join(rootDir, ".golangci.yml") {
			return nil
		}

		if path == filepath.Join(rootDir, "LICENSE") {
			return nil
		}

		if path == filepath.Join(rootDir, "NOTICE") {
			return nil
		}

		if path == filepath.Join(rootDir, "README.md") {
			return nil
		}

		if path == filepath.Join(rootDir, "go.mod") {
			return nil
		}

		if path == filepath.Join(rootDir, "go.sum") {
			return nil
		}

		if path == filepath.Join(rootDir, "internal/tools.go") {
			return nil
		}

		if _, exist := e[path]; exist {
			return nil
		}

		if d.Type().IsRegular() {
			os.Remove(path)
		}

		return nil
	})
}
