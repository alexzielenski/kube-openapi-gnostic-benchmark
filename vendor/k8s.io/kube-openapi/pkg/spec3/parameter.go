/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package spec3

import (
	"encoding/json"

	"github.com/go-openapi/swag"
	"gopkg.in/yaml.v3"
	"k8s.io/kube-openapi/pkg/validation/spec"
)

// Parameter a struct that describes a single operation parameter, more at https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#parameterObject
//
// Note that this struct is actually a thin wrapper around ParameterProps to make it referable and extensible
type Parameter struct {
	spec.Refable
	ParameterProps
	spec.VendorExtensible
}

// MarshalJSON is a custom marshal function that knows how to encode Parameter as JSON
func (p *Parameter) MarshalJSON() ([]byte, error) {
	b1, err := json.Marshal(p.Refable)
	if err != nil {
		return nil, err
	}
	b2, err := json.Marshal(p.ParameterProps)
	if err != nil {
		return nil, err
	}
	b3, err := json.Marshal(p.VendorExtensible)
	if err != nil {
		return nil, err
	}
	return swag.ConcatJSON(b1, b2, b3), nil
}

func (p *Parameter) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &p.Refable); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &p.ParameterProps); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &p.VendorExtensible); err != nil {
		return err
	}

	return nil
}

func (p *Parameter) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&p.Refable); err != nil {
		return err
	}
	if err := value.Decode(&p.ParameterProps); err != nil {
		return err
	}
	if err := p.VendorExtensible.UnmarshalYAML(value); err != nil {
		return err
	}
	return nil
}

// ParameterProps a struct that describes a single operation parameter, more at https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#parameterObject
type ParameterProps struct {
	// Name holds the name of the parameter
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// In holds the location of the parameter
	In string `json:"in,omitempty" yaml:"in,omitempty"`
	// Description holds a brief description of the parameter
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Required determines whether this parameter is mandatory
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`
	// Deprecated declares this operation to be deprecated
	Deprecated bool `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	// AllowEmptyValue sets the ability to pass empty-valued parameters
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	// Style describes how the parameter value will be serialized depending on the type of the parameter value
	Style string `json:"style,omitempty" yaml:"style,omitempty"`
	// Explode when true, parameter values of type array or object generate separate parameters for each value of the array or key-value pair of the map
	Explode bool `json:"explode,omitempty" yaml:"explode,omitempty"`
	// AllowReserved determines whether the parameter value SHOULD allow reserved characters, as defined by RFC3986
	AllowReserved bool `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	// Schema holds the schema defining the type used for the parameter
	Schema *spec.Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
	// Content holds a map containing the representations for the parameter
	Content map[string]*MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	// Example of the parameter's potential value
	Example interface{} `json:"example,omitempty" yaml:"example,omitempty"`
	// Examples of the parameter's potential value. Each example SHOULD contain a value in the correct format as specified in the parameter encoding
	Examples map[string]*Example `json:"examples,omitempty" yaml:"examples,omitempty"`
}
