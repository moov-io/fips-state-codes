// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package fips_state_codes_test

import (
	"fmt"
	fips_state_codes "github.com/moov-io/fips-state-codes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFIPSCodeFromStateCode(t *testing.T) {
	result, err := fips_state_codes.FIPSCodeFromStateCode("IL")
	require.NoError(t, err)
	require.Equal(t, "17", result)
}

func TestFIPSCodeFromStateCode_ErrorsOnUnknownStateCode(t *testing.T) {
	inputCode := "ZZ"
	result, err := fips_state_codes.FIPSCodeFromStateCode(inputCode)
	require.Equal(t, "", result)
	require.Error(t, err)
	require.Equal(t, fmt.Sprintf("unknown state code: %s", inputCode), err.Error())
}
