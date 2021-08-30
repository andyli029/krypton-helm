/*
Copyright 2021 RadonDB.

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

package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func AddFinalizer(objMeta metav1.ObjectMeta, finalizer string) {
	if !HasFinalizer(objMeta, finalizer) {
		objMeta.Finalizers = append(objMeta.Finalizers, finalizer)
	}
}

func HasFinalizer(objMeta metav1.ObjectMeta, finalizer string) bool {
	for _, metaFinalizer := range objMeta.Finalizers {
		if metaFinalizer == finalizer {
			return true
		}
	}
	return false
}
