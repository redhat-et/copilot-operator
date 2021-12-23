/*
Copyright 2021.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DataCollectionPolicySpec defines the desired state of DataCollectionPolicy
type DataCollectionPolicySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Master on/off switch for data collection policy.
	// Explicit set to `true` is required to opt-in for the operator
	// to collect data from the cluster and send to a central AI server
	// to fine-train the completion AI model.
	Enabled bool `json:"enabled,omitempty"`
}

// DataCollectionPolicyStatus defines the observed state of DataCollectionPolicy
type DataCollectionPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The latest generation of the resource which was observed by the operator.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// number of yamls collected from the cluster
	CollectedCount int `json:"collectedCount,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Enabled",type="boolean",JSONPath=".spec.enabled"
//+kubebuilder:printcolumn:name="Collected",type="integer",JSONPath=".status.collectedCount"

// DataCollectionPolicy is the Schema for the datacollectionpolicies API
type DataCollectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataCollectionPolicySpec   `json:"spec,omitempty"`
	Status DataCollectionPolicyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DataCollectionPolicyList contains a list of DataCollectionPolicy
type DataCollectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCollectionPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataCollectionPolicy{}, &DataCollectionPolicyList{})
}
