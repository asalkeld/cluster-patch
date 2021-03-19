package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	infrav1alpha3 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
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
	_ = infrav1alpha3.AddToScheme(scheme)

	cl, err := client.New(config.GetConfigOrDie(), client.Options{Scheme: scheme})
	if err != nil {
		fmt.Println("failed to create client:", err)
		os.Exit(1)
	}

	awsCluster := &infrav1alpha3.AWSCluster{}
	err = cl.Get(context.Background(), client.ObjectKey{Namespace: namespace, Name: clusterName}, awsCluster)
	if err != nil {
		fmt.Println("failed to get cluster:", err)
		os.Exit(1)
	}

	fmt.Println("Patching cluster status")

	awsClusterPatch := client.Patch(client.MergeFrom(awsCluster.DeepCopy()))

	awsCluster.Status.Ready = true

	err = cl.Status().Patch(context.Background(), awsCluster, awsClusterPatch)
	if err != nil {
		fmt.Println("failed to patch cluster status:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully patched cluster status")
}
