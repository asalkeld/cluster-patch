package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	metal3v1alpha3 "github.com/metal3-io/cluster-api-provider-metal3/api/v1alpha3"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	var clusterName string
	flag.StringVar(&clusterName, "clusterName", "capi-ocp-aws", "cluster name, defaults to 'capi-ocp-aws'")

	var namespace string
	flag.StringVar(&namespace, "namespace", "ocp-cluster-api", "namespace name, defaults to 'ocp-cluster-api'")

	flag.Parse()

	fmt.Printf("cluster name %s, namespace %s\n", clusterName, namespace)

	scheme := runtime.NewScheme()
	_ = metal3v1alpha3.AddToScheme(scheme)

	cl, err := client.New(config.GetConfigOrDie(), client.Options{Scheme: scheme})
	if err != nil {
		fmt.Println("failed to create client:", err)
		os.Exit(1)
	}
	m3Cluster := &metal3v1alpha3.Metal3Cluster{}
	err = cl.Get(context.Background(), client.ObjectKey{Namespace: namespace, Name: clusterName}, m3Cluster)
	if err != nil {
		fmt.Println("failed to get cluster:", err)
		os.Exit(1)
	}

	fmt.Println("Patching cluster status")

	m3ClusterPatch := client.Patch(client.MergeFrom(m3Cluster.DeepCopy()))

	m3Cluster.Status.Ready = true

	err = cl.Status().Patch(context.Background(), m3Cluster, m3ClusterPatch)
	if err != nil {
		fmt.Println("failed to patch cluster status:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully patched cluster status")
}
