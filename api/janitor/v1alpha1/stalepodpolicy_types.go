package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// desired state of StalePodPolicy.
type StalePodPolicySpec struct {
	// TTL is how long a pod may remain in a terminal phase before it is deleted.
	TTL metav1.Duration `json:"ttl"`
}

// StalePodPolicyStatus defines the observed state of StalePodPolicy.
type StalePodPolicyStatus struct {
	// DeletedCount is the number of pods deleted under this policy so far.
	DeletedCount int64 `json:"deletedCount,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type StalePodPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StalePodPolicySpec   `json:"spec,omitempty"`
	Status StalePodPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type StalePodPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StalePodPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StalePodPolicy{}, &StalePodPolicyList{})
}
