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

package crd

import (
	"fmt"

	hubv1alpha1 "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

// Decoder decodes CRD objects.
type Decoder struct {
	decoder runtime.Decoder
	scheme  *runtime.Scheme
}

// NewDecoder creates a new Decoder.
func NewDecoder() (*Decoder, error) {
	scheme := runtime.NewScheme()
	if err := apiextensionsv1.AddToScheme(scheme); err != nil {
		return nil, err
	}

	decoder := serializer.NewCodecFactory(scheme).UniversalDeserializer()

	return &Decoder{
		decoder: decoder,
		scheme:  scheme,
	}, nil
}

// Decode decodes the given YAML/JSON manifest into a CustomResourceDefinition.
func (d *Decoder) Decode(document []byte) (*apiextensions.CustomResourceDefinition, error) {
	object, _, err := d.decoder.Decode(document, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("decoding object: %w", err)
	}

	var internalObject apiextensions.CustomResourceDefinition
	if err = d.scheme.Convert(object, &internalObject, nil); err != nil {
		return nil, fmt.Errorf("converting CRD to internal object: %w", err)
	}

	return &internalObject, nil
}

// HubDecoder decodes Traefik Hub Kubernetes objects.
type HubDecoder struct {
	decoder runtime.Decoder
}

// NewHubDecoder creates a new HubDecoder.
func NewHubDecoder() (*HubDecoder, error) {
	scheme := runtime.NewScheme()
	if err := hubv1alpha1.AddToScheme(scheme); err != nil {
		return nil, fmt.Errorf("adding hub.traefik.io/v1alpha1 resources: %w", err)
	}

	decoder := serializer.NewCodecFactory(scheme).UniversalDeserializer()

	return &HubDecoder{decoder: decoder}, nil
}

// Decode decodes the given YAML/JSON manifest into a runtime object.
// If `into` is nil, a new runtime.Object will be created.
// A nil runtime.Object will be returned if it doesn't match a known resource type of is not a Kubernetes manifest.
func (d *HubDecoder) Decode(document []byte, into runtime.Object) (runtime.Object, error) {
	object, _, err := d.decoder.Decode(document, nil, into)
	if err != nil {
		switch {
		case runtime.IsMissingKind(err), runtime.IsNotRegisteredError(err):
			return nil, nil
		default:
			return nil, fmt.Errorf("decoding: %w", err)
		}
	}

	return object, nil
}
