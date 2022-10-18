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

package v1alpha1

import (
	"fmt"
	"net"
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var mysqlclusterlog = logf.Log.WithName("mysqlcluster-resource")

func (r *MysqlCluster) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-mysql-radondb-com-v1alpha1-mysqlcluster,mutating=false,failurePolicy=fail,sideEffects=None,groups=mysql.radondb.com,resources=mysqlclusters,verbs=create;update,versions=v1alpha1,name=vmysqlcluster.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &MysqlCluster{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *MysqlCluster) ValidateCreate() error {
	mysqlclusterlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	if err := r.validateNFSServerAddress(r); err != nil {
		return err
	}
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *MysqlCluster) ValidateUpdate(old runtime.Object) error {
	mysqlclusterlog.Info("validate update", "name", r.Name)

	oldCluster, ok := old.(*MysqlCluster)
	if !ok {
		return apierrors.NewBadRequest(fmt.Sprintf("expected an MysqlCluster but got a %T", old))
	}
	if err := r.validateVolumeSize(oldCluster); err != nil {
		return err
	}
	if err := r.validateLowTableCase(oldCluster); err != nil {
		return err
	}
	if err := r.validateMysqlVersionAndImage(); err != nil {
		return err
	}
	if err := r.validateNFSServerAddress(oldCluster); err != nil {
		return err
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *MysqlCluster) ValidateDelete() error {
	mysqlclusterlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

// TODO: Add NFSServerAddress webhook & backup schedule.
func (r *MysqlCluster) validateNFSServerAddress(oldCluster *MysqlCluster) error {
	isIP := net.ParseIP(r.Spec.NFSServerAddress) != nil
	if len(r.Spec.NFSServerAddress) != 0 && !isIP {
		return apierrors.NewForbidden(schema.GroupResource{}, "", fmt.Errorf("nfsServerAddress should be set as IP"))
	}
	if len(r.Spec.BackupSchedule) != 0 && len(r.Spec.BackupSecretName) == 0 && !isIP {
		return apierrors.NewForbidden(schema.GroupResource{}, "", fmt.Errorf("backupSchedule is set without any backupSecretName or nfsServerAddress"))
	}
	return nil
}

// Validate volume size, forbidden shrink storage size.
func (r *MysqlCluster) validateVolumeSize(oldCluster *MysqlCluster) error {
	oldStorageSize, err := resource.ParseQuantity(oldCluster.Spec.Persistence.Size)
	if err != nil {
		return err
	}
	newStorageSize, err := resource.ParseQuantity(r.Spec.Persistence.Size)
	if err != nil {
		return err
	}
	// =1 means that old storage size is greater than new.
	if oldStorageSize.Cmp(newStorageSize) == 1 {
		return apierrors.NewForbidden(schema.GroupResource{}, "", fmt.Errorf("volesize can not be decreased"))
	}
	return nil
}

// Validate low table case for mysqlcluster.
func (r *MysqlCluster) validateLowTableCase(oldCluster *MysqlCluster) error {
	oldmyconf := oldCluster.Spec.MysqlOpts.MysqlConf
	newmyconf := r.Spec.MysqlOpts.MysqlConf
	if strings.Contains(r.Spec.MysqlOpts.Image, "8.0") &&
		oldmyconf["lower_case_table_names"] != newmyconf["lower_case_table_names"] {
		return apierrors.NewForbidden(schema.GroupResource{}, "", fmt.Errorf("lower_case_table_names cannot be changed in MySQL8.0+"))
	}
	return nil
}

// Validate MysqlVersion and spec.MysqlOpts.image are conflict.
func (r *MysqlCluster) validateMysqlVersionAndImage() error {
	if r.Spec.MysqlOpts.Image != "" && r.Spec.MysqlVersion != "" {
		if !strings.Contains(r.Spec.MysqlOpts.Image, r.Spec.MysqlVersion) {
			return apierrors.NewForbidden(schema.GroupResource{}, "", fmt.Errorf("spec.MysqlOpts.Image and spec.MysqlVersion are conflict"))
		}
	}
	return nil
}
