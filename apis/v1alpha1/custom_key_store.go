// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CustomKeyStoreSpec defines the desired state of CustomKeyStore.
type CustomKeyStoreSpec struct {
	// Identifies the AWS CloudHSM cluster for the custom key store. Enter the cluster
	// ID of any active AWS CloudHSM cluster that is not already associated with
	// a custom key store. To find the cluster ID, use the DescribeClusters (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_DescribeClusters.html)
	// operation.
	// +kubebuilder:validation:Required
	CloudHsmClusterID *string `json:"cloudHsmClusterID"`
	// Specifies a friendly name for the custom key store. The name must be unique
	// in your AWS account.
	// +kubebuilder:validation:Required
	CustomKeyStoreName *string `json:"customKeyStoreName"`
	// Enter the password of the kmsuser crypto user (CU) account (https://docs.aws.amazon.com/kms/latest/developerguide/key-store-concepts.html#concept-kmsuser)
	// in the specified AWS CloudHSM cluster. AWS KMS logs into the cluster as this
	// user to manage key material on your behalf.
	//
	// The password must be a string of 7 to 32 characters. Its value is case sensitive.
	//
	// This parameter tells AWS KMS the kmsuser account password; it does not change
	// the password in the AWS CloudHSM cluster.
	// +kubebuilder:validation:Required
	KeyStorePassword *string `json:"keyStorePassword"`
	// Enter the content of the trust anchor certificate for the cluster. This is
	// the content of the customerCA.crt file that you created when you initialized
	// the cluster (https://docs.aws.amazon.com/cloudhsm/latest/userguide/initialize-cluster.html).
	// +kubebuilder:validation:Required
	TrustAnchorCertificate *string `json:"trustAnchorCertificate"`
}

// CustomKeyStoreStatus defines the observed state of CustomKeyStore
type CustomKeyStoreStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// A unique identifier for the new custom key store.
	// +kubebuilder:validation:Optional
	CustomKeyStoreID *string `json:"customKeyStoreID,omitempty"`
}

// CustomKeyStore is the Schema for the CustomKeyStores API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type CustomKeyStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomKeyStoreSpec   `json:"spec,omitempty"`
	Status            CustomKeyStoreStatus `json:"status,omitempty"`
}

// CustomKeyStoreList contains a list of CustomKeyStore
// +kubebuilder:object:root=true
type CustomKeyStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomKeyStore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CustomKeyStore{}, &CustomKeyStoreList{})
}
