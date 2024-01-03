/*
Copyright (C) 2022-2024 Traefik Labs

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
	v1alpha1crd "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1/crd"
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

func checkValidationTestCases(t *testing.T, tests []validationTestCase) {
	t.Helper()

	crds, err := crd.GetCRDs(v1alpha1crd.CRDs)
	require.NoError(t, err)

	validator := validation.NewValidator()
	for _, crd := range crds {
		if err = validator.Register(crd); err != nil {
			require.NoError(t, err)
		}
	}

	decoder, err := crd.NewHubDecoder()
	require.NoError(t, err)

	for _, test := range tests {
		test := test

		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			var object unstructured.Unstructured
			_, decoderErr := decoder.Decode(test.manifest, &object)
			require.NoError(t, decoderErr)

			gotErrs := validator.Validate(&object)
			assert.Equal(t, test.wantErrs, gotErrs)
		})
	}
}
