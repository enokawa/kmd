package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Get Nodes
	nodesClient := clientset.CoreV1().Nodes()
	nodeList, err := nodesClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, node := range nodeList.Items {
		fmt.Printf("Name: %s \n", node.Name)
		for _, condition := range node.Status.Conditions {
			if condition.Type != v1.NodeReady {
				continue
			}
			if condition.Status == v1.ConditionTrue {
				fmt.Printf("Status: %s \n", condition.Type)
				break
			}
			fmt.Println("Status: NotReady")
			break
		}

		nodeLabels := node.GetLabels()
		if err != nil {
			panic(err)
		}

		var roles []string
		for key, _ := range nodeLabels {
			if strings.Contains(key, "node-role.kubernetes.io") {
				roleLabel := strings.Split(key, "/")
				roles = append(roles, roleLabel[1])
			}
		}

		if len(roles) > 0 {
			fmt.Printf("Roles: %s \n", strings.Join(roles, ","))
		} else if roles == nil {
			fmt.Println("Roles: <none>")
		}
	}
}
