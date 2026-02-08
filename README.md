# kubeslim

_An alternative to K8s client-go optimized for binary size, memory consumption and speed_

<img width="239" height="349" alt="gopher-slim" src="https://github.com/user-attachments/assets/76a56ea2-652e-41d8-b60d-c768ba35d520" />

## Introduction

**kubeslim** is a Go library designed as a lightweight, performant alternative to Kubernetes [client-go](https://github.com/kubernetes/client-go). While building the [`kubetail`](https://github.com/kubetail-org/kubetail) CLI tool, we noticed that final binary sizes for tools in the K8s ecosystem (e.g. `kubectl`, `helm`) were disproportionately large compared to other unix utilities. We traced the root cause to the [Kubernetes client-go library](https://github.com/amorey/size-matters), which typically adds 20MB+ to final binaries. We also found that `client-go` informers consume significant amounts of memory at runtime. This library addresses both issues with a design optimized for binary size, memory consumption, and speed.

| Client | Binary Size¹ |
|---|---|
| `client-go.Clientset` | 38 MB |
| `client-go.DynamicClient` | 16 MB |
| **`kubeslim`** | **10 MB** |

¹ Measured using a minimal Go executable performing a simple API call. See [size-matters](https://github.com/amorey/size-matters) for methodology.

This library is a work in progress and we welcome contributions!

## Installation

```console
go get github.com/kubetail-org/kubeslim
```

## Basic usage

```go
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
      Name string `json:"name"`
    } `json:"metadata"`
  } `json:"items"`
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
		fmt.Println(pod.Metadata.Name)
	}
}
```

## Get Involved

This library is very much a work in progress! If you have ideas on how to improve it or if you'd like help using it you can:

* Create a [GitHub Issue](https://github.com/kubetail-org/kubeslim/issues)
* Send us an email ([hello@kubetail.com](hello@kubetail.com))
* Join our [Discord server](https://discord.gg/CmsmWAVkvX) or [Slack channel](https://join.slack.com/t/kubetail/shared_invite/zt-2cq01cbm8-e1kbLT3EmcLPpHSeoFYm1w)
