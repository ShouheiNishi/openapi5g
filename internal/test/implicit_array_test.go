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

package test_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ShouheiNishi/openapi5g/udr/dr"
)

func TestFixImplicitArray(t *testing.T) {
	ty := reflect.TypeOf(dr.QueryEeGroupSubscriptionResponse{}.JSON200)
	require.Equal(t, reflect.Pointer, ty.Kind())
	assert.Equal(t, reflect.Slice, ty.Elem().Kind())
}
