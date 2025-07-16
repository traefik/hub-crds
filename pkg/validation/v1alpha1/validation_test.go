/*
Copyright (C) 2022-2025 Traefik Labs

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package v1alpha1_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	hubcrd "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1/crd"
	"github.com/traefik/hub-crds/pkg/crd"
	"github.com/traefik/hub-crds/pkg/validation"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

type validationTestCase struct {
	desc     string
	manifest []byte
	wantErrs field.ErrorList
}

func checkValidation(t *testing.T, test validationTestCase) {
	t.Helper()

	crds, err := crd.GetCRDs(hubcrd.CRDs)
	require.NoError(t, err)

	validator := validation.NewValidator()
	for _, definition := range crds {
		require.NoError(t, validator.Register(definition))
	}

	decoder, err := crd.NewHubDecoder()
	require.NoError(t, err)

	var object unstructured.Unstructured

	decoderErr := decoder.Decode(test.manifest, &object)
	require.NoError(t, decoderErr)

	gotErrs := validator.Validate(&object)
	assert.Equal(t, test.wantErrs, gotErrs)
}
