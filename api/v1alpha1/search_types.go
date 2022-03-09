/*
Copyright 2022.

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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SearchSpec defines the desired state of Search
type SearchSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	// Kubernetes secret name containing user provided db secret
	// Secret contains connection url, certificates
	CustomDbConfig string `json:"customDbConfig,omitempty"`

	// +optional
	// Storgeclass to be used by default db
	DBStorage StorageSpec `json:"dbStorage,omitempty"`

	// +optional
	//configmap name contains parameters to override default db parameters
	DbConfig string `json:"dbConfig,omitempty"`

	// +optional
	// Customize search deployments
	Deployments SearchDeployments `json:"deployments,omitempty"`

	// +optional
	// flag to turn on/off High Availability for search components
	EnableHA bool `json:"enableHA,omitempty"`

	// +optional
	// Control list of Kubernetes resources indexed by search-collector
	AllowDenyResources FilterSpec `json:"allowDenyResources,omitempty"`
}

type SearchDeployments struct {
	// +optional
	// Configuration for DB Deployment
	Database DeploymentConfig `json:"database,omitempty"`

	// +optional
	// Configuration for indexer Deployment
	Indexer DeploymentConfig `json:"indexer,omitempty"`

	// +optional
	// Configuration for collector Deployment
	Collector DeploymentConfig `json:"collector,omitempty"`

	// +optional
	// Configuration for api Deployment
	API DeploymentConfig `json:"api,omitempty"`

	// +optional
	// Configuration for addon installed collector Deployment
	RemoteCollector DeploymentConfig `json:"remoteCollector,omitempty"`
}

type DeploymentConfig struct {
	// +optional
	// Number of pod instances
	ReplicaCount int `json:"replicaCount,omitempty"`
	// +optional
	// Compute Resources required by deployment
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
	//Image_override
	ImageOverride string `json:"imageOverride,omitempty"`

	// +optional
	//ImagePullSecret
	ImagePullSecret string `json:"imagePullSecret,omitempty"`

	//ImagePullPolicy
	// +optional
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`

	// NodeSelector to schedule on nodes with matching labels
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	//Proxy config , if remote collectors need to override
	// +optional
	ProxyConfig map[string]string `json:"proxyConfig,omitempty"`
}

type StorageSpec struct {
	// +optional
	// name of the storage class
	StorageClassName string `json:"storageClassName,omitempty"`
	// +optional
	// storage capacity
	Size *resource.Quantity `json:"size,omitempty"`
}

type FilterSpec struct {
	// +optional
	// Allowed resources from collector
	AllowedResources map[string][]string `json:"allowedResources,omitempty"`

	// +optional
	// Denied resources from collector
	DeniedResources map[string][]string `json:"deniedResources,omitempty"`
}

// SearchStatus defines the observed state of Search
type SearchStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Human readable health state
	SearchHealth string `json:"searchHealth,omitempty"`
	// Database used by search
	DBInUse string `json:"dbInUse,omitempty"`
	// Storage used by database
	StorageInUse string `json:"storageInUse,omitempty"`
	// +optional
	Conditions SearchConditions `json:"conditions,omitempty"`
}

type SearchCondition struct {
	Type   SearchConditionType    `json:"type"`
	Status corev1.ConditionStatus `json:"status"`
	// Last time the condition transitioned
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	// +optional
	// Reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// +optional
	// Human readable message
	Message string `json:"message,omitempty" `
}
type SearchConditionType string
type SearchConditions []SearchCondition

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Search is the Schema for the searches API
type Search struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SearchSpec   `json:"spec,omitempty"`
	Status SearchStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SearchList contains a list of Search
type SearchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Search `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Search{}, &SearchList{})
}
