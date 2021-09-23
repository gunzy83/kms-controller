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

package key

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.BypassPolicyLockoutSafetyCheck, b.ko.Spec.BypassPolicyLockoutSafetyCheck) {
		delta.Add("Spec.BypassPolicyLockoutSafetyCheck", a.ko.Spec.BypassPolicyLockoutSafetyCheck, b.ko.Spec.BypassPolicyLockoutSafetyCheck)
	} else if a.ko.Spec.BypassPolicyLockoutSafetyCheck != nil && b.ko.Spec.BypassPolicyLockoutSafetyCheck != nil {
		if *a.ko.Spec.BypassPolicyLockoutSafetyCheck != *b.ko.Spec.BypassPolicyLockoutSafetyCheck {
			delta.Add("Spec.BypassPolicyLockoutSafetyCheck", a.ko.Spec.BypassPolicyLockoutSafetyCheck, b.ko.Spec.BypassPolicyLockoutSafetyCheck)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CustomKeyStoreID, b.ko.Spec.CustomKeyStoreID) {
		delta.Add("Spec.CustomKeyStoreID", a.ko.Spec.CustomKeyStoreID, b.ko.Spec.CustomKeyStoreID)
	} else if a.ko.Spec.CustomKeyStoreID != nil && b.ko.Spec.CustomKeyStoreID != nil {
		if *a.ko.Spec.CustomKeyStoreID != *b.ko.Spec.CustomKeyStoreID {
			delta.Add("Spec.CustomKeyStoreID", a.ko.Spec.CustomKeyStoreID, b.ko.Spec.CustomKeyStoreID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CustomerMasterKeySpec, b.ko.Spec.CustomerMasterKeySpec) {
		delta.Add("Spec.CustomerMasterKeySpec", a.ko.Spec.CustomerMasterKeySpec, b.ko.Spec.CustomerMasterKeySpec)
	} else if a.ko.Spec.CustomerMasterKeySpec != nil && b.ko.Spec.CustomerMasterKeySpec != nil {
		if *a.ko.Spec.CustomerMasterKeySpec != *b.ko.Spec.CustomerMasterKeySpec {
			delta.Add("Spec.CustomerMasterKeySpec", a.ko.Spec.CustomerMasterKeySpec, b.ko.Spec.CustomerMasterKeySpec)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Description, b.ko.Spec.Description) {
		delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
	} else if a.ko.Spec.Description != nil && b.ko.Spec.Description != nil {
		if *a.ko.Spec.Description != *b.ko.Spec.Description {
			delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.KeyUsage, b.ko.Spec.KeyUsage) {
		delta.Add("Spec.KeyUsage", a.ko.Spec.KeyUsage, b.ko.Spec.KeyUsage)
	} else if a.ko.Spec.KeyUsage != nil && b.ko.Spec.KeyUsage != nil {
		if *a.ko.Spec.KeyUsage != *b.ko.Spec.KeyUsage {
			delta.Add("Spec.KeyUsage", a.ko.Spec.KeyUsage, b.ko.Spec.KeyUsage)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Origin, b.ko.Spec.Origin) {
		delta.Add("Spec.Origin", a.ko.Spec.Origin, b.ko.Spec.Origin)
	} else if a.ko.Spec.Origin != nil && b.ko.Spec.Origin != nil {
		if *a.ko.Spec.Origin != *b.ko.Spec.Origin {
			delta.Add("Spec.Origin", a.ko.Spec.Origin, b.ko.Spec.Origin)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Policy, b.ko.Spec.Policy) {
		delta.Add("Spec.Policy", a.ko.Spec.Policy, b.ko.Spec.Policy)
	} else if a.ko.Spec.Policy != nil && b.ko.Spec.Policy != nil {
		if *a.ko.Spec.Policy != *b.ko.Spec.Policy {
			delta.Add("Spec.Policy", a.ko.Spec.Policy, b.ko.Spec.Policy)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.Tags, b.ko.Spec.Tags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}

	return delta
}
