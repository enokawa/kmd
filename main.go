package main

import (
	// "context"
	"flag"
	"fmt"
	"path/filepath"

	// "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	fmt.Printf(config.Host)

	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	panic(err)
	// }

	// Get Nodes
	// nodesClient := clientset.NodeV1alpha1()
	// get := nodesClient.RESTClient().Get().Do(context.TODO())
	// err = get.Error()
	// if err != nil {
	// 	panic(err)
	// }

	// getRaw, err := get.Raw()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%s", getRaw)
}

// func int32Ptr(i int32) *int32 { return &i }
