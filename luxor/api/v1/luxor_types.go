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

// PodSpec defines the specification for a pod
type PodSpec struct {
	// PodSpec fields
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec               PodTemplateSpec `json:"spec,omitempty"`
}

// PodTemplateSpec defines the template for creating pods
type PodTemplateSpec struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec               PodSpec `json:"spec,omitempty"`
}

// LuxorSpec defines the desired state of Luxor
type LuxorSpec struct {
	// Pod defines the pod template
	Pod PodSpec `json:"pod,omitempty"`
}

// LuxorStatus defines the observed state of Luxor
type LuxorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Luxor is the Schema for the luxors API
type Luxor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LuxorSpec   `json:"spec,omitempty"`
	Status LuxorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LuxorList contains a list of Luxor
type LuxorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Luxor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Luxor{}, &LuxorList{})
}
