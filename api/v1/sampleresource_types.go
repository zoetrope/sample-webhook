package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SampleResourceSpec defines the desired state of SampleResource
type SampleResourceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of SampleResource. Edit sampleresource_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// SampleResourceStatus defines the observed state of SampleResource
type SampleResourceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SampleResource is the Schema for the sampleresources API
type SampleResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SampleResourceSpec   `json:"spec,omitempty"`
	Status SampleResourceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SampleResourceList contains a list of SampleResource
type SampleResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SampleResource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SampleResource{}, &SampleResourceList{})
}
