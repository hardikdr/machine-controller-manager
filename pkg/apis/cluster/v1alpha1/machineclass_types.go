/*
Copyright 2018 The Kubernetes Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineClass can be used to templatize and re-use provider configuration
// across multiple Machines / MachineSets / MachineDeployments.
// +k8s:openapi-gen=true
// +resource:path=machineclasses,strategy=MachineClassStrategy
type MachineClass struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The total capacity available on this machine type (cpu/memory/disk).
	//
	// WARNING: It is up to the creator of the MachineClass to ensure that
	// this field is consistent with the underlying machine that will
	// be provisioned when this class is used, to inform higher level
	// automation (e.g. the cluster autoscaler).
	Capacity corev1.ResourceList `json:"capacity"`

	// How much capacity is actually allocatable on this machine.
	// Must be equal to or less than the capacity, and when less
	// indicates the resources reserved for system overhead.
	//
	// WARNING: It is up to the creator of the MachineClass to ensure that
	// this field is consistent with the underlying machine that will
	// be provisioned when this class is used, to inform higher level
	// automation (e.g. the cluster autoscaler).
	Allocatable corev1.ResourceList `json:"allocatable"`

	// Provider-specific configuration to use during node creation.
	ProviderConfig runtime.RawExtension `json:"providerConfig"`

	// TODO: should this use an api.ObjectReference to a 'MachineTemplate' instead?
	// A link to the MachineTemplate that will be used to create provider
	// specific configuration for Machines of this class.
	// MachineTemplate corev1.ObjectReference `json:machineTemplate`

	SecretRef *corev1.SecretReference `json:"secretRef,omitempty"`

	//Cloud-provider name
	Provider string `json:"provider"`
}

// // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// type MachineClassList struct {
// 	// +optional
// 	metav1.TypeMeta `json:",inline"`

// 	// +optional
// 	metav1.ListMeta `json:"metadata,omitempty"`

// 	// +optional
// 	Items []MachineClass `json:"items"`
// }
