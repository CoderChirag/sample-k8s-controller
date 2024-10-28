package sample_controller

import "k8s.io/apimachinery/pkg/runtime/schema"

func getSampleCrd() *schema.GroupVersionResource {
	sampleCrd := schema.GroupVersionResource{
		Group: "myk8s.io",
		Version: "v1",
		Resource: "samples",
	}
	return &sampleCrd
}