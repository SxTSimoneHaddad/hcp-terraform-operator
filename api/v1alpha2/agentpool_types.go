package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Agent Token is a secret token that a Terraform Cloud Agent is used to connect to the Terraform Cloud Agent Pool.
// More infromation:
//   - https://developer.hashicorp.com/terraform/cloud-docs/agents
type AgentToken struct {
	// Agent Token name.
	Name string `json:"name"`
	// Agent Token ID.
	// +optional
	ID string `json:"id,omitempty"`
	// Timestamp of when the agent token was created.
	// +optional
	CreatedAt int64 `json:"createdAt,omitempty"`
	// Timestamp of when the agent token was last used.
	// +optional
	LastUsedAt int64 `json:"lastUsedAt,omitempty"`
}

// AgentPoolSpec defines the desired state of AgentPool.
type AgentPoolSpec struct {
	// Agent Pool name:
	// More information:
	//   - https://developer.hashicorp.com/terraform/cloud-docs/agents/agent-pools
	Name string `json:"name"`
	// API Token to be used for API calls.
	Token Token `json:"token"`
	// Organization name where the Workspace will be created.
	// More information:
	//  - https://developer.hashicorp.com/terraform/cloud-docs/users-teams-organizations/organizations
	Organization string `json:"organization"`

	// List of the agent tokens to generate.
	// +kubebuilder:validation:MinItems:=1
	// +optional
	AgentTokens []*AgentToken `json:"agentTokens,omitempty"`
}

// AgentPoolStatus defines the observed state of AgentPool
type AgentPoolStatus struct {
	// Real world state generation.
	ObservedGeneration int64 `json:"observedGeneration"`
	// Agent Pool ID that is managed by the controller.
	AgentPoolID string `json:"agentPoolID"`

	// List of the agent tokens generated by the controller.
	// +optional
	AgentTokens []*AgentToken `json:"agentTokens,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AgentPool is the Schema for the agentpools API
type AgentPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AgentPoolSpec   `json:"spec"`
	Status AgentPoolStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AgentPoolList contains a list of AgentPool
type AgentPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgentPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AgentPool{}, &AgentPoolList{})
}
