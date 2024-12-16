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

// KuberhealthyJobSpec defines the desired state of KuberhealthyJob
// Important: Run "make" to regenerate code after modifying this file
type KuberhealthyJobSpec struct {
	// +optional
	Phase   JobPhase      `json:"phase" yaml:"phase"`     // the state or phase of the job
	Timeout string        `json:"timeout" yaml:"timeout"` // the maximum time the pod is allowed to run before a failure is assumed
	PodSpec apiv1.PodSpec `json:"podSpec" yaml:"podSpec"` // a spec for the external job
	// +optional
	ExtraAnnotations map[string]string `json:"extraAnnotations" yaml:"extraAnnotations"` // a map of extra annotations that will be applied to the pod
	// +optional
	ExtraLabels map[string]string `json:"extraLabels" yaml:"extraLabels"` // a map of extra labels that will be applied to the pod
}

// JobPhase is a label for the condition of the job at the current time.
type JobPhase string

// These are the valid phases of jobs.
const (
	JobRunning   JobPhase = "Running"
	JobCompleted JobPhase = "Completed"
)


// KuberhealthyJobStatus defines the observed state of KuberhealthyJob
type KuberhealthyJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KuberhealthyJob is the Schema for the kuberhealthyjobs API
// +k8s:openapi-gen=true
// +kubebuilder:resource:path="khjobs"
// +kubebuilder:resource:singular="khjob"
// +kubebuilder:resource:shortName="khj"
type KuberhealthyJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KuberhealthyJobSpec   `json:"spec,omitempty"`
	Status KuberhealthyJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KuberhealthyJobList contains a list of KuberhealthyJob
type KuberhealthyJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KuberhealthyJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KuberhealthyJob{}, &KuberhealthyJobList{})
}
