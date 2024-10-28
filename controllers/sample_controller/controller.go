package sample_controller

import (
	"context"
	"fmt"
	sampleEvents "sample-k8s-controller/controllers/sample_controller/events"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/cache"
)

var informer *cache.SharedIndexInformer

func New(client *dynamic.Interface){
	sampleCrd := getSampleCrd()
	informer = createSampleInformer(client, sampleCrd)
}

func Run(){
	addEventHandlers()
	run()
}

func createSampleInformer(client *dynamic.Interface, resource *schema.GroupVersionResource) *cache.SharedIndexInformer{
	informer := cache.NewSharedIndexInformer(&cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			return (*client).Resource(*resource).Namespace("").List(context.TODO(), options)
		},
		WatchFunc: func(options metav1.ListOptions)(watch.Interface, error) {
			return (*client).Resource(*resource).Namespace("").Watch(context.TODO(), options)
		},
	}, &unstructured.Unstructured{}, 0, cache.Indexers{})
	return &informer
}

func addEventHandlers(){
	(*informer).AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: sampleEvents.AdditionEvent,
		UpdateFunc: sampleEvents.UpdationEvent,
		DeleteFunc: sampleEvents.DeletionEvent,
	})
}

func run() {
	stop := make(chan struct{})
	defer close(stop)

	go (*informer).Run(stop)

	if !cache.WaitForCacheSync(stop, (*informer).HasSynced) {
		panic("Timeout waiting for cache sync")
	}

	fmt.Println("Sample Resource Controller started successfully")

	<-stop
}