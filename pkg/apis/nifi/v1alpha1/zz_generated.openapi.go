// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/nifi/v1alpha1.NifiUser":       schema_pkg_apis_nifi_v1alpha1_NifiUser(ref),
		"./pkg/apis/nifi/v1alpha1.NifiUserSpec":   schema_pkg_apis_nifi_v1alpha1_NifiUserSpec(ref),
		"./pkg/apis/nifi/v1alpha1.NifiUserStatus": schema_pkg_apis_nifi_v1alpha1_NifiUserStatus(ref),
	}
}

func schema_pkg_apis_nifi_v1alpha1_NifiUser(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Nifi User is the Schema for the nifi users API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/nifi/v1alpha1.NifiUserSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/nifi/v1alpha1.NifiUserStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/nifi/v1alpha1.NifiUserSpec", "./pkg/apis/nifi/v1alpha1.NifiUserStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_nifi_v1alpha1_NifiUserSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NifiUserSpec defines the desired state of NifiUser",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"secretName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"clusterRef": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/nifi/v1alpha1.ClusterReference"),
						},
					},
					"dnsNames": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"includeJKS": {
						SchemaProps: spec.SchemaProps{
							Description: "TopicGrants []UserTopicGrant `json:\"topicGrants,omitempty\"`",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"secretName", "clusterRef"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/nifi/v1alpha1.ClusterReference"},
	}
}

func schema_pkg_apis_nifi_v1alpha1_NifiUserStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NifiUserStatus defines the observed state of NifiUser",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"state": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"acls": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"state"},
			},
		},
	}
}
