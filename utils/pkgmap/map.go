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

package pkgmap

import (
	"errors"

	"github.com/getkin/kin-openapi/openapi3"
)

func SpecName2Doc(specName string) (doc *openapi3.T, err error) {
	if p, exist := SpecName2PkgName(specName); !exist {
		return nil, errors.New("spec is not extist")
	} else {
		return PkgName2Doc(p)
	}
}

func SpecName2PkgName(specName string) (pkgName string, exist bool) {
	pkgName, exist = s2p[specName]
	return
}

func PkgName2Doc(pkgName string) (doc *openapi3.T, err error) {
	if l, exist := p2l[pkgName]; !exist {
		return nil, errors.New("packge is not extist")
	} else {
		return l()
	}
}

func PkgName2specName(pkgName string) (specName string, exist bool) {
	specName, exist = p2s[pkgName]
	return
}
