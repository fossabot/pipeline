/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.3.0
 * Contact: info@banzaicloud.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

type ClusterProfileAksAks struct {
	KubernetesVersion string                        `json:"kubernetesVersion,omitempty"`
	NodePools         ClusterProfileAksAksNodePools `json:"nodePools,omitempty"`
}