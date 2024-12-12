/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KuberhealthyStateSpec defines the desired state of KuberhealthyState
type KuberhealthyStateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of KuberhealthyState. Edit kuberhealthystate_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// KuberhealthyStateStatus defines the observed state of KuberhealthyState
type KuberhealthyStateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KuberhealthyState is the Schema for the kuberhealthystates API
type KuberhealthyState struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KuberhealthyStateSpec   `json:"spec,omitempty"`
	Status KuberhealthyStateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KuberhealthyStateList contains a list of KuberhealthyState
type KuberhealthyStateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KuberhealthyState `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KuberhealthyState{}, &KuberhealthyStateList{})
}
