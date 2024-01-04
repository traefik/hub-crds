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
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strings"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apimachinery/pkg/util/yaml"
)

// GetCRDs returns CRDs.
func GetCRDs(filesystem fs.FS) ([]*apiextensions.CustomResourceDefinition, error) {
	decoder, err := NewDecoder()
	if err != nil {
		return nil, fmt.Errorf("creating CRD decoder: %w", err)
	}

	manifests, err := loadManifests(filesystem)
	if err != nil {
		return nil, fmt.Errorf("loading CRD documents: %w", err)
	}

	crds := make([]*apiextensions.CustomResourceDefinition, 0, len(manifests))

	for _, m := range manifests {
		crd, decodeErr := decoder.Decode(m.Data)
		if decodeErr != nil {
			return nil, fmt.Errorf("decoding manifest %s: %w", m.Path, decodeErr)
		}

		crds = append(crds, crd)
	}

	return crds, nil
}

type manifest struct {
	Path string
	Data []byte
}

func loadManifests(filesystem fs.FS) ([]manifest, error) {
	var manifests []manifest

	err := fs.WalkDir(filesystem, ".", func(path string, entry fs.DirEntry, fileErr error) error {
		if fileErr != nil {
			return fileErr
		}

		if filepath.Base(path) == ".git" {
			return filepath.SkipDir
		}

		if !entry.Type().IsRegular() || !isYAMLOrJSON(path) {
			return nil
		}

		reader, err := filesystem.Open(path)
		if err != nil {
			return fmt.Errorf("opening file: %w", err)
		}
		defer func() { _ = reader.Close() }()

		r := yaml.NewYAMLReader(bufio.NewReader(reader))

		for {
			data, err := r.Read()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}

				return fmt.Errorf("reading file content: %w", err)
			}

			manifests = append(manifests, manifest{
				Data: data,
				Path: path,
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return manifests, nil
}

func isYAMLOrJSON(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))

	return ext == ".yaml" || ext == ".yml" || ext == ".json"
}
