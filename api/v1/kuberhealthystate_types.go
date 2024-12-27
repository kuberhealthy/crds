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
// Important: Run "make" to regenerate code after modifying this file
type KuberhealthyStateSpec struct {
	OK          bool     `json:"OK" yaml:"OK"`                   // true or false status of the khWorkload, whether or not it completed successfully
	Errors      []string `json:"Errors" yaml:"Errors"`           // the list of errors reported from the khWorkload run
	RunDuration string   `json:"RunDuration" yaml:"RunDuration"` // the time it took for the khWorkload to complete
	Namespace   string   `json:"Namespace" yaml:"Namespace"`     // the namespace the khWorkload was run in
	Node        string   `json:"Node" yaml:"Node"`               // the node the khWorkload ran on
	// +nullable
	LastRun          *metav1.Time `json:"LastRun,omitempty" yaml:"LastRun,omitempty"` // the time the khWorkload was last run
	AuthoritativePod string       `json:"AuthoritativePod" yaml:"AuthoritativePod"`   // the main kuberhealthy pod creating and updating the khstate
	CurrentUUID      string       `json:"uuid" yaml:"uuid"`                           // the UUID that is authorized to report statuses into the kuberhealthy endpoint
	// +nullable
	KHWorkload *KHWorkload `json:"khWorkload,omitempty" yaml:"khWorkload,omitempty"`
}

// KHWorkload is used to describe the different types of kuberhealthy workloads: KhCheck or KHJob
type KHWorkload string

// Two types of KHWorkloads are available: Kuberhealthy Check or Kuberhealthy Job
// KHChecks run on a scheduled run interval
// KHJobs run once
const (
	KHCheck KHWorkload = "KHCheck"
	KHJob   KHWorkload = "KHJob"
)

// KuberhealthyStateStatus defines the observed state of KuberhealthyState
type KuberhealthyStateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KuberhealthyState is the Schema for the kuberhealthystates API
// +k8s:openapi-gen=true
// +kubebuilder:printcolumn:name="OK",type=string,JSONPath=`.spec.OK`,description="OK status"
// +kubebuilder:printcolumn:name="Age LastRun",type=date,JSONPath=`.spec.LastRun`,description="Last Run"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`,description="Age"
// +kubebuilder:resource:path="khstates"
// +kubebuilder:resource:singular="khstate"
// +kubebuilder:resource:shortName="khs"
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
