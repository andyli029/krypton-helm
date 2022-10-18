/*
Copyright 2022 RadonDB.

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

package framework

import (
	"fmt"
	"time"

	. "github.com/onsi/gomega"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/retry"
)

var (
	SuperUserTemplate = `
apiVersion: mysql.radondb.com/v1alpha1
kind: MysqlUser
metadata:
  name: super-user
spec:
  user: super_user
  withGrantOption: true
  tlsOptions:
    type: NONE
  hosts:
  - '%%'
  permissions:
  - database: '*'
    tables:
    - '*'
    privileges:
    - ALL
  userOwner:
    clusterName: %s
    nameSpace: %s
  secretSelector:
    secretName: sample-user-password
    secretKey: superUser
`
	NormalUserTemplate = `
apiVersion: mysql.radondb.com/v1alpha1
kind: MysqlUser
metadata:
  name: normal-user
spec:
  user: normal_user
  withGrantOption: true
  tlsOptions: 
    type: NONE
  hosts: 
  - "%%"
  permissions:
  - database: "*"
    tables:
    - "*"
    privileges:
    - USAGE
  userOwner:
    clusterName: %s
    nameSpace: %s
  secretSelector:
    secretName: sample-user-passwordj
    secretKey: normalUser
`
	UserSecretTemplate = `
apiVersion: v1
kind: Secret
metadata:
  name: sample-user-password
data:
  superUser: UmFkb25EQkAxMjM=
  normalUser: UmFkb25EQkAxMjM=
`
	UserAsset = `https://github.com/radondb/radondb-mysql-kubernetes/releases/latest/download/mysql_v1alpha1_mysqluser.yaml`
)

func (f *Framework) CreateUserSecret() {
	k8s.KubectlApplyFromString(f.t, f.kubectlOptions, UserSecretTemplate)
}

func (f *Framework) CreateNormalUser() {
	user := fmt.Sprintf(NormalUserTemplate, TestContext.ClusterReleaseName, f.kubectlOptions.Namespace)
	k8s.KubectlApplyFromString(f.t, f.kubectlOptions, user)
}

func (f *Framework) CreateSuperUser() {
	user := fmt.Sprintf(SuperUserTemplate, TestContext.ClusterReleaseName, f.kubectlOptions.Namespace)
	k8s.KubectlApplyFromString(f.t, f.kubectlOptions, user)
}

func (f *Framework) CreateUserUsingAsset() {
	k8s.KubectlApply(f.t, f.kubectlOptions, UserAsset)
}

func (f *Framework) CleanUpUser() {
	IgnoreNotFound(k8s.KubectlDeleteFromStringE(f.t, f.kubectlOptions, UserSecretTemplate))
	k8s.RunKubectl(f.t, f.kubectlOptions, "delete", "mysqluser", "--all")
}

func (f *Framework) CheckGantsForUser(user string, withGrant bool) {
	podName := fmt.Sprintf("%s-mysql-0", TestContext.ClusterReleaseName)
	grants := retry.DoWithRetry(f.t, fmt.Sprintf("check grants for %s", user), 12, 10*time.Second, func() (string, error) {
		grants, err := k8s.RunKubectlAndGetOutputE(f.t, f.kubectlOptions, "exec", "-it", podName, "-c", "mysql", "--", "mysql", "-u", "root", "-e", "show grants for "+user)
		if err != nil {
			return "", err
		}
		return grants, nil
	})
	if withGrant {
		Expect(grants).Should(ContainSubstring("WITH GRANT OPTION"))
	}
}

func (f *Framework) CheckLogIn(user, pass string) {
	podName := fmt.Sprintf("%s-mysql-0", TestContext.ClusterReleaseName)
	var err error
	if pass != "" {
		_, err = k8s.RunKubectlAndGetOutputE(f.t, f.kubectlOptions, "exec", "-it", podName, "-c", "mysql", "--", "mysql", "-u", user, "-p"+pass, "-e", "select 1")
	} else {
		_, err = k8s.RunKubectlAndGetOutputE(f.t, f.kubectlOptions, "exec", "-it", podName, "-c", "mysql", "--", "mysql", "-u", user, "-e", "select 1")
	}
	Expect(err).To(BeNil())
}
