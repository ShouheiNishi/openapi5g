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

package pkgmap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ShouheiNishi/openapi5g/utils/pkgmap"
)

func TestPkgMap(t *testing.T) {
	p, exist := pkgmap.SpecName2PkgName("TS29571_CommonData.yaml")
	assert.True(t, exist)
	assert.Equal(t, "github.com/ShouheiNishi/openapi5g/commondata", p)

	s, exist := pkgmap.PkgName2specName("github.com/ShouheiNishi/openapi5g/commondata")
	assert.True(t, exist)
	assert.Equal(t, "TS29571_CommonData.yaml", s)

	doc, err := pkgmap.SpecName2Doc("TS29571_CommonData.yaml")
	assert.NoError(t, err)
	assert.Equal(t, "Common Data Types", doc.Info.Title)

	doc, err = pkgmap.PkgName2Doc("github.com/ShouheiNishi/openapi5g/commondata")
	assert.NoError(t, err)
	assert.Equal(t, "Common Data Types", doc.Info.Title)

	_, exist = pkgmap.SpecName2PkgName("TSXXXXX_NotExist.yaml")
	assert.False(t, exist)

	_, exist = pkgmap.PkgName2specName("github.com/ShouheiNishi/openapi5g/notexist")
	assert.False(t, exist)

	_, err = pkgmap.SpecName2Doc("TSXXXXX_NotExist.yaml")
	assert.Error(t, err)

	_, err = pkgmap.PkgName2Doc("github.com/ShouheiNishi/openapi5g/notexist")
	assert.Error(t, err)
}
