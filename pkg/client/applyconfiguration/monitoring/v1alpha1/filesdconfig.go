// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	monitoringv1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
)

// FileSDConfigApplyConfiguration represents a declarative configuration of the FileSDConfig type for use
// with apply.
type FileSDConfigApplyConfiguration struct {
	Files           []monitoringv1alpha1.SDFile `json:"files,omitempty"`
	RefreshInterval *v1.Duration                `json:"refreshInterval,omitempty"`
}

// FileSDConfigApplyConfiguration constructs a declarative configuration of the FileSDConfig type for use with
// apply.
func FileSDConfig() *FileSDConfigApplyConfiguration {
	return &FileSDConfigApplyConfiguration{}
}

// WithFiles adds the given value to the Files field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Files field.
func (b *FileSDConfigApplyConfiguration) WithFiles(values ...monitoringv1alpha1.SDFile) *FileSDConfigApplyConfiguration {
	for i := range values {
		b.Files = append(b.Files, values[i])
	}
	return b
}

// WithRefreshInterval sets the RefreshInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RefreshInterval field is set to the value of the last call.
func (b *FileSDConfigApplyConfiguration) WithRefreshInterval(value v1.Duration) *FileSDConfigApplyConfiguration {
	b.RefreshInterval = &value
	return b
}
