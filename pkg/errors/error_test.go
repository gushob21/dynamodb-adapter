// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	e := New("E0001")
	assert.Error(t, e)
	code, err := HTTPResponse(e, nil)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, code)
}

func TestHTTPResponse(t *testing.T) {
	e := New("E0001")
	assert.Error(t, e)
	code, err := e.HTTPResponse(nil)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, code)
}

func TestNewForExceptionalError(t *testing.T) {
	e := New("E1001")
	assert.Error(t, e)
}

func TestNewForSystemErrors(t *testing.T) {
	code, e := HTTPResponse(errors.New("Test"), nil)
	assert.NotNil(t, e)
	assert.Equal(t, http.StatusInternalServerError, code)

}
