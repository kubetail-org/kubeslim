package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/kubetail-org/kubeslim"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type namespaceList struct {
	Items []struct {
		Metadata struct {
			Name string `json:"name"`
		} `json:"metadata"`
	} `json:"items"`
}

type podList struct {
	Items []struct {
		Metadata struct {
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
		} `json:"metadata"`
	} `json:"items"`
}

type pod struct {
	Metadata struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
}

func main() {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	// Load configuration from kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("failed to load kubeconfig: %v", err)
	}

	// Create the kubeslim client
	client, err := kubeslim.NewForConfig(config)
	if err != nil {
		log.Fatalf("failed to create kubeslim client: %v", err)
	}

	// Get namespaces
	namespaceGVR := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "namespaces"}

	namespaces, err := kubeslim.List[namespaceList](context.TODO(), client, namespaceGVR)
	if err != nil {
		log.Fatalf("failed to list namespaces: %v", err)
	}

	if len(namespaces.Items) == 0 {
		fmt.Println("No namespaces found.")
		return
	}

	for _, ns := range namespaces.Items {
		fmt.Println(ns.Metadata.Name)
	}

	// Get pods
	podGVR := schema.GroupVersionResource{Group: "core", Version: "v1", Resource: "pods"}

	pods, err := kubeslim.List[podList](context.TODO(), client, podGVR)
	if err != nil {
		log.Fatalf("failed to list pods: %v", err)
	}

	if len(pods.Items) == 0 {
		fmt.Println("No pods found.")
		return
	}

	for _, pod := range pods.Items {
		fmt.Printf("%s/%s\n", pod.Metadata.Namespace, pod.Metadata.Name)
	}

	fmt.Println("\n--- Testing Get Method ---")
	firstPod := pods.Items[0]
	fmt.Printf("Fetching pod %s in namespace %s...\n", firstPod.Metadata.Name, firstPod.Metadata.Namespace)

	singlePod, err := kubeslim.Get[pod](context.TODO(), client, podGVR, firstPod.Metadata.Namespace, firstPod.Metadata.Name)
	if err != nil {
		log.Fatalf("failed to get pod: %v", err)
	}
	fmt.Printf("Successfully retrieved single pod: %s/%s\n", singlePod.Metadata.Namespace, singlePod.Metadata.Name)
}
