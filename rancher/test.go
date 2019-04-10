package main

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
)

func gethostip(node *corev1.Node) {
	for _, address := range node.Status.Addresses {
		if address.Type == corev1.NodeInternalIP {
			fmt.Println("==========", address.Address)
		}
	}
}
