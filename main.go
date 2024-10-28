package main

import (
	"sample-k8s-controller/controllers/sample_controller"
	"sample-k8s-controller/k8sclient"
)

func main(){
	client := k8sclient.NewDynamicClient()
	sample_controller.New(client)
	sample_controller.Run()
}