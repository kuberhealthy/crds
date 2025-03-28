/*
Copyright 2025 Kuberhealthy Authors.

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

package v2

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KuberhealthyCheckSpec defines the desired state of KuberhealthyCheck
type KuberhealthyCheckSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of KuberhealthyCheck. Edit kuberhealthycheck_types.go to remove/update
	Name      string                  `json:"name"`
	SingleRun bool                    `json:"singleRunOnly,omitempty"`
	PodSpec   v1.PodSpec              `json:"podSpec"`
	Status    KuberhealthyCheckStatus `json:"status"`
}

// KuberhealthyCheckStatus defines the observed state of KuberhealthyCheck
type KuberhealthyCheckStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	OK                 bool     `json:"ok"`
	Errors             []string `json:"errors"`
	CurrentMaster      string   `json:"currentMaster"`
	RunDuration        string   `json:"runDuration"`
	Namespace          string   `json:"namespace"`
	CurrentUUID        string   `json:"currentUUID"`
	LastRunUnix        int64    `json:"lastRunUnix"`
	AdditionalMetadata string   `json:"additionalMetadata"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KuberhealthyCheck is the Schema for the kuberhealthychecks API
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
