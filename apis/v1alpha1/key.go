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

// KeySpec defines the desired state of Key.
type KeySpec struct {
	// A flag to indicate whether to bypass the key policy lockout safety check.
	//
	// Setting this value to true increases the risk that the CMK becomes unmanageable.
	// Do not set this value to true indiscriminately.
	//
	// For more information, refer to the scenario in the Default Key Policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam)
	// section in the AWS Key Management Service Developer Guide .
	//
	// Use this parameter only when you include a policy in the request and you
	// intend to prevent the principal that is making the request from making a
	// subsequent PutKeyPolicy request on the CMK.
	//
	// The default value is false.
	BypassPolicyLockoutSafetyCheck *bool `json:"bypassPolicyLockoutSafetyCheck,omitempty"`
	// Creates the CMK in the specified custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
	// and the key material in its associated AWS CloudHSM cluster. To create a
	// CMK in a custom key store, you must also specify the Origin parameter with
	// a value of AWS_CLOUDHSM. The AWS CloudHSM cluster that is associated with
	// the custom key store must have at least two active HSMs, each in a different
	// Availability Zone in the Region.
	//
	// This parameter is valid only for symmetric CMKs. You cannot create an asymmetric
	// CMK in a custom key store.
	//
	// To find the ID of a custom key store, use the DescribeCustomKeyStores operation.
	//
	// The response includes the custom key store ID and the ID of the AWS CloudHSM
	// cluster.
	//
	// This operation is part of the Custom Key Store feature (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
	// feature in AWS KMS, which combines the convenience and extensive integration
	// of AWS KMS with the isolation and control of a single-tenant key store.
	CustomKeyStoreID *string `json:"customKeyStoreID,omitempty"`
	// Specifies the type of CMK to create. The default value, SYMMETRIC_DEFAULT,
	// creates a CMK with a 256-bit symmetric key for encryption and decryption.
	// For help choosing a key spec for your CMK, see How to Choose Your CMK Configuration
	// (https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html)
	// in the AWS Key Management Service Developer Guide.
	//
	// The CustomerMasterKeySpec determines whether the CMK contains a symmetric
	// key or an asymmetric key pair. It also determines the encryption algorithms
	// or signing algorithms that the CMK supports. You can't change the CustomerMasterKeySpec
	// after the CMK is created. To further restrict the algorithms that can be
	// used with the CMK, use a condition key in its key policy or IAM policy. For
	// more information, see kms:EncryptionAlgorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-encryption-algorithm)
	// or kms:Signing Algorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-signing-algorithm)
	// in the AWS Key Management Service Developer Guide.
	//
	// AWS services that are integrated with AWS KMS (http://aws.amazon.com/kms/features/#AWS_Service_Integration)
	// use symmetric CMKs to protect your data. These services do not support asymmetric
	// CMKs. For help determining whether a CMK is symmetric or asymmetric, see
	// Identifying Symmetric and Asymmetric CMKs (https://docs.aws.amazon.com/kms/latest/developerguide/find-symm-asymm.html)
	// in the AWS Key Management Service Developer Guide.
	//
	// AWS KMS supports the following key specs for CMKs:
	//
	//    * Symmetric key (default) SYMMETRIC_DEFAULT (AES-256-GCM)
	//
	//    * Asymmetric RSA key pairs RSA_2048 RSA_3072 RSA_4096
	//
	//    * Asymmetric NIST-recommended elliptic curve key pairs ECC_NIST_P256 (secp256r1)
	//    ECC_NIST_P384 (secp384r1) ECC_NIST_P521 (secp521r1)
	//
	//    * Other asymmetric elliptic curve key pairs ECC_SECG_P256K1 (secp256k1),
	//    commonly used for cryptocurrencies.
	CustomerMasterKeySpec *string `json:"customerMasterKeySpec,omitempty"`
	// A description of the CMK.
	//
	// Use a description that helps you decide whether the CMK is appropriate for
	// a task.
	Description *string `json:"description,omitempty"`
	// Determines the cryptographic operations (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
	// for which you can use the CMK. The default value is ENCRYPT_DECRYPT. This
	// parameter is required only for asymmetric CMKs. You can't change the KeyUsage
	// value after the CMK is created.
	//
	// Select only one valid value.
	//
	//    * For symmetric CMKs, omit the parameter or specify ENCRYPT_DECRYPT.
	//
	//    * For asymmetric CMKs with RSA key material, specify ENCRYPT_DECRYPT or
	//    SIGN_VERIFY.
	//
	//    * For asymmetric CMKs with ECC key material, specify SIGN_VERIFY.
	KeyUsage *string `json:"keyUsage,omitempty"`
	// The source of the key material for the CMK. You cannot change the origin
	// after you create the CMK. The default is AWS_KMS, which means AWS KMS creates
	// the key material.
	//
	// When the parameter value is EXTERNAL, AWS KMS creates a CMK without key material
	// so that you can import key material from your existing key management infrastructure.
	// For more information about importing key material into AWS KMS, see Importing
	// Key Material (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html)
	// in the AWS Key Management Service Developer Guide. This value is valid only
	// for symmetric CMKs.
	//
	// When the parameter value is AWS_CLOUDHSM, AWS KMS creates the CMK in an AWS
	// KMS custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
	// and creates its key material in the associated AWS CloudHSM cluster. You
	// must also use the CustomKeyStoreId parameter to identify the custom key store.
	// This value is valid only for symmetric CMKs.
	Origin *string `json:"origin,omitempty"`
	// The key policy to attach to the CMK.
	//
	// If you provide a key policy, it must meet the following criteria:
	//
	//    * If you don't set BypassPolicyLockoutSafetyCheck to true, the key policy
	//    must allow the principal that is making the CreateKey request to make
	//    a subsequent PutKeyPolicy request on the CMK. This reduces the risk that
	//    the CMK becomes unmanageable. For more information, refer to the scenario
	//    in the Default Key Policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam)
	//    section of the AWS Key Management Service Developer Guide .
	//
	//    * Each statement in the key policy must contain one or more principals.
	//    The principals in the key policy must exist and be visible to AWS KMS.
	//    When you create a new AWS principal (for example, an IAM user or role),
	//    you might need to enforce a delay before including the new principal in
	//    a key policy because the new principal might not be immediately visible
	//    to AWS KMS. For more information, see Changes that I make are not always
	//    immediately visible (https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency)
	//    in the AWS Identity and Access Management User Guide.
	//
	// If you do not provide a key policy, AWS KMS attaches a default key policy
	// to the CMK. For more information, see Default Key Policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default)
	// in the AWS Key Management Service Developer Guide.
	//
	// The key policy size quota is 32 kilobytes (32768 bytes).
	//
	// For help writing and formatting a JSON policy document, see the IAM JSON
	// Policy Reference (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html)
	// in the IAM User Guide .
	Policy *string `json:"policy,omitempty"`
	// One or more tags. Each tag consists of a tag key and a tag value. Both the
	// tag key and the tag value are required, but the tag value can be an empty
	// (null) string.
	//
	// When you add tags to an AWS resource, AWS generates a cost allocation report
	// with usage and costs aggregated by tags. For information about adding, changing,
	// deleting and listing tags for CMKs, see Tagging Keys (https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html).
	//
	// Use this parameter to tag the CMK when it is created. To add tags to an existing
	// CMK, use the TagResource operation.
	//
	// To use this parameter, you must have kms:TagResource (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
	// permission in an IAM policy.
	Tags []*Tag `json:"tags,omitempty"`
}

// KeyStatus defines the observed state of Key
type KeyStatus struct {
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
	// The twelve-digit account ID of the AWS account that owns the CMK.
	// +kubebuilder:validation:Optional
	AWSAccountID *string `json:"awsAccountID,omitempty"`
	// The cluster ID of the AWS CloudHSM cluster that contains the key material
	// for the CMK. When you create a CMK in a custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html),
	// AWS KMS creates the key material for the CMK in the associated AWS CloudHSM
	// cluster. This value is present only when the CMK is created in a custom key
	// store.
	// +kubebuilder:validation:Optional
	CloudHsmClusterID *string `json:"cloudHsmClusterID,omitempty"`
	// The date and time when the CMK was created.
	// +kubebuilder:validation:Optional
	CreationDate *metav1.Time `json:"creationDate,omitempty"`
	// The date and time after which AWS KMS deletes the CMK. This value is present
	// only when KeyState is PendingDeletion.
	// +kubebuilder:validation:Optional
	DeletionDate *metav1.Time `json:"deletionDate,omitempty"`
	// Specifies whether the CMK is enabled. When KeyState is Enabled this value
	// is true, otherwise it is false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty"`
	// The encryption algorithms that the CMK supports. You cannot use the CMK with
	// other encryption algorithms within AWS KMS.
	//
	// This field appears only when the KeyUsage of the CMK is ENCRYPT_DECRYPT.
	// +kubebuilder:validation:Optional
	EncryptionAlgorithms []*string `json:"encryptionAlgorithms,omitempty"`
	// Specifies whether the CMK's key material expires. This value is present only
	// when Origin is EXTERNAL, otherwise this value is omitted.
	// +kubebuilder:validation:Optional
	ExpirationModel *string `json:"expirationModel,omitempty"`
	// The globally unique identifier for the CMK.
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyID,omitempty"`
	// The manager of the CMK. CMKs in your AWS account are either customer managed
	// or AWS managed. For more information about the difference, see Customer Master
	// Keys (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#master_keys)
	// in the AWS Key Management Service Developer Guide.
	// +kubebuilder:validation:Optional
	KeyManager *string `json:"keyManager,omitempty"`
	// The current status of the CMK.
	//
	// For more information about how key state affects the use of a CMK, see Key
	// state: Effect on your CMK (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
	// in the AWS Key Management Service Developer Guide.
	// +kubebuilder:validation:Optional
	KeyState *string `json:"keyState,omitempty"`
	// The signing algorithms that the CMK supports. You cannot use the CMK with
	// other signing algorithms within AWS KMS.
	//
	// This field appears only when the KeyUsage of the CMK is SIGN_VERIFY.
	// +kubebuilder:validation:Optional
	SigningAlgorithms []*string `json:"signingAlgorithms,omitempty"`
	// The time at which the imported key material expires. When the key material
	// expires, AWS KMS deletes the key material and the CMK becomes unusable. This
	// value is present only for CMKs whose Origin is EXTERNAL and whose ExpirationModel
	// is KEY_MATERIAL_EXPIRES, otherwise this value is omitted.
	// +kubebuilder:validation:Optional
	ValidTo *metav1.Time `json:"validTo,omitempty"`
}

// Key is the Schema for the Keys API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Key struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeySpec   `json:"spec,omitempty"`
	Status            KeyStatus `json:"status,omitempty"`
}

// KeyList contains a list of Key
// +kubebuilder:object:root=true
type KeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Key `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Key{}, &KeyList{})
}
