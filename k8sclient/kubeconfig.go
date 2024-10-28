package k8sclient

import (
	"fmt"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func buildKubeConfig() *rest.Config {
	var kubeconfig string
	home := homedir.HomeDir()
	
	if home == "" {
		panic("Home directory path not found.")
	}

	kubeconfig = filepath.Join(home, ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		fmt.Println("Kubeconfig not found, building from in-cluster config")
		config, err = rest.InClusterConfig()

		if err != nil {
			panic(err.Error())
		}
	}
	return config
}