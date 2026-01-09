/*
Copyright (C) 2022-2026 Traefik Labs

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
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	hubcrd "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1/crd"
	"github.com/traefik/hub-crds/pkg/crd"
	"github.com/traefik/hub-crds/pkg/validation"
	apiextensionsvalidation "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/validation"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// TestCRDsAreValid validates that CRD definitions are valid according to Kubernetes rules,
// including CEL cost budgets. This catches issues before release rather than after.
func TestCRDsAreValid(t *testing.T) {
	t.Parallel()

	crds, err := crd.GetCRDs(hubcrd.CRDs)
	require.NoError(t, err)

	for _, definition := range crds {
		t.Run(definition.Name, func(t *testing.T) {
			t.Parallel()

			// Kubernetes populates this field when the CRD is applied, not in the source YAML files.
			definition.Status.StoredVersions = []string{"v1alpha1"}

			errs := apiextensionsvalidation.ValidateCustomResourceDefinition(context.Background(), definition)
			assert.Empty(t, errs)
		})
	}
}

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
