package k8sclient

import "k8s.io/client-go/dynamic"

func NewDynamicClient() *dynamic.Interface {
	config := buildKubeConfig()
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return &client
}