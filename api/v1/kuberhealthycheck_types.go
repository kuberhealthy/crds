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
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KuberhealthyCheckSpec defines the desired state of KuberhealthyCheck
// Important: Run "make" to regenerate code after modifying this file
type KuberhealthyCheckSpec struct {
	RunInterval string        `json:"runInterval" yaml:"runInterval"` // the interval at which the check runs
	Timeout     string        `json:"timeout" yaml:"timeout"`         // the maximum time the pod is allowed to run before a failure is assumed
	PodSpec     apiv1.PodSpec `json:"podSpec" yaml:"podSpec"`         // a spec for the external checker
	// +optional
	ExtraAnnotations map[string]string `json:"extraAnnotations" yaml:"extraAnnotations"` // a map of extra annotations that will be applied to the pod
	// +optional
	ExtraLabels map[string]string `json:"extraLabels" yaml:"extraLabels"` // a map of extra labels that will be applied to the pod
}

// KuberhealthyCheckStatus defines the observed state of KuberhealthyCheck
type KuberhealthyCheckStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KuberhealthyCheck is the Schema for the kuberhealthychecks API
// +k8s:openapi-gen=true
// +kubebuilder:resource:path="khchecks"
// +kubebuilder:resource:singular="khcheck"
// +kubebuilder:resource:shortName="khc"
type KuberhealthyCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KuberhealthyCheckSpec   `json:"spec,omitempty"`
	Status KuberhealthyCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KuberhealthyCheckList contains a list of KuberhealthyCheck
type KuberhealthyCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KuberhealthyCheck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KuberhealthyCheck{}, &KuberhealthyCheckList{})
}
