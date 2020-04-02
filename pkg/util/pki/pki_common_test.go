// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pki

import (
	"fmt"
	"reflect"
	"testing"
	"github.com/erdrix/nifikop/pkg/apis/nifi/v1alpha1"
	"github.com/erdrix/nifikop/pkg/resources/templates"
	"github.com/erdrix/nifikop/pkg/util"
	certutil "github.com/erdrix/nifikop/pkg/util/cert"
)

func testCluster(t *testing.T) *v1alpha1.NifiCluster {
	t.Helper()
	cluster := &v1alpha1.NifiCluster{}
	cluster.Name = "test-cluster"
	cluster.Namespace = "test-namespace"
	cluster.Spec = v1alpha1.NifiClusterSpec{}

	cluster.Spec.Nodes = []v1alpha1.Node{
		{Id: 0},
		{Id: 1},
		{Id: 2},
	}
	return cluster
}

func TestDN(t *testing.T) {
	cert, _, expected, err := certutil.GenerateTestCert()
	if err != nil {
		t.Fatal("failed to generate certificate for testing:", err)
	}
	userCert := &UserCertificate{
		Certificate: cert,
	}
	dn := userCert.DN()
	if dn != expected {
		t.Error("Expected:", expected, "got:", dn)
	}
}

func TestGetCommonName(t *testing.T) {
	cluster := &v1alpha1.NifiCluster{}
	cluster.Name = "test-cluster"
	cluster.Namespace = "test-namespace"

	cluster.Spec = v1alpha1.NifiClusterSpec{HeadlessServiceEnabled: true}
	headlessCN := GetCommonName(cluster)
	expected := "test-cluster-headless.test-namespace.svc.cluster.local"
	if headlessCN != expected {
		t.Error("Expected:", expected, "Got:", headlessCN)
	}

	cluster.Spec = v1alpha1.NifiClusterSpec{HeadlessServiceEnabled: false}
	allNodeCN := GetCommonName(cluster)
	expected = "test-cluster-all-node.test-namespace.svc.cluster.local"
	if allNodeCN != expected {
		t.Error("Expected:", expected, "Got:", allNodeCN)
	}
}

func TestLabelsForNifiPKI(t *testing.T) {
	expected := map[string]string{
		"app":          "nifi",
		"nifi_issuer": fmt.Sprintf(NodeIssuerTemplate, "test"),
	}
	got := LabelsForNifiPKI("test")
	if !reflect.DeepEqual(got, expected) {
		t.Error("Expected:", expected, "got:", got)
	}
}

func TestGetInternalDNSNames(t *testing.T) {
	cluster := testCluster(t)

	for _, node := range cluster.Spec.Nodes {
		cluster.Spec.HeadlessServiceEnabled = true
		headlessNames := GetInternalDNSNames(cluster, node.Id)
		expected := []string{
			fmt.Sprintf("test-cluster-%d-node.test-cluster-headless.test-namespace.svc.cluster.local", node.Id),
			fmt.Sprintf("test-cluster-%d-node.test-cluster-headless.test-namespace.svc", node.Id),
			fmt.Sprintf("test-cluster-%d-node.test-cluster-headless.test-namespace", node.Id),
			fmt.Sprintf("test-cluster-%d-node.test-cluster-headless", node.Id),
			fmt.Sprintf(templates.NodeNameTemplate, cluster.Name, node.Id),
		}
		if !reflect.DeepEqual(expected, headlessNames) {
			t.Error("Expected:", expected, "got:", headlessNames)
		}

		cluster.Spec.HeadlessServiceEnabled = false
		allNodeNames := GetInternalDNSNames(cluster, node.Id)
		expected = []string{
			fmt.Sprintf("test-cluster-%d-node.test-cluster-all-node.test-namespace.svc.cluster.local", node.Id),
			fmt.Sprintf("test-cluster-%d-node.test-cluster-all-node.test-namespace.svc", node.Id),
			fmt.Sprintf("test-cluster-%d-node.test-cluster-all-node.test-namespace", node.Id),
			fmt.Sprintf("test-cluster-%d-node.test-cluster-all-node", node.Id),
			fmt.Sprintf(templates.NodeNameTemplate, cluster.Name, node.Id),

		}
		if !reflect.DeepEqual(expected, allNodeNames) {
			t.Error("Expected:", expected, "got:", allNodeNames)
		}
	}
}

func TestNodeUsersForCluster(t *testing.T) {
	cluster := testCluster(t)
	users := NodeUsersForCluster(cluster, []string{})

	for _, node := range cluster.Spec.Nodes {
		expected := &v1alpha1.NifiUser{
			ObjectMeta: templates.ObjectMeta(GetNodeUserName(cluster, node.Id), LabelsForNifiPKI(cluster.Name), cluster),
			Spec: v1alpha1.NifiUserSpec{
				SecretName: fmt.Sprintf(NodeServerCertTemplate, cluster.Name, node.Id),
				DNSNames:   GetInternalDNSNames(cluster, node.Id),
				IncludeJKS: true,
				ClusterRef: v1alpha1.ClusterReference{
					Name:      cluster.Name,
					Namespace: cluster.Namespace,
				},
			},
		}
		if !util.NifiUserSliceContains(users, expected) {
			t.Errorf("Expected %+v\ninto %+v", expected, users)
		}
	}
}

func TestControllerUserForCluster(t *testing.T) {
	cluster := testCluster(t)
	user := ControllerUserForCluster(cluster)

	expected := &v1alpha1.NifiUser{
		ObjectMeta: templates.ObjectMeta(
			fmt.Sprintf(NodeControllerFQDNTemplate, fmt.Sprintf(NodeControllerTemplate, cluster.Name), cluster.Namespace),
			LabelsForNifiPKI(cluster.Name), cluster,
		),
		Spec: v1alpha1.NifiUserSpec{
			SecretName: fmt.Sprintf(NodeControllerTemplate, cluster.Name),
			IncludeJKS: true,
			ClusterRef: v1alpha1.ClusterReference{
				Name:      cluster.Name,
				Namespace: cluster.Namespace,
			},
		},
	}

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("Expected %+v\nGot %+v", expected, user)
	}
}
