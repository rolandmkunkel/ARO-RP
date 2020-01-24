package e2e

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/ARO-RP/test/util/ready"
)

var _ = Describe("Check the node count is correct [CheckNodeCount][EveryPR]", func() {
	It("should be possible to list nodes and confirm they are as expected", func() {
		By("Verifying that the expected number of nodes exist and are ready")
		ctx := context.Background()
		cs, err := Clients.openshiftclusters.Get(ctx, os.Getenv("RESOURCEGROUP"), os.Getenv("CLUSTER"))
		Expect(err).NotTo(HaveOccurred())
		var expectedNodeCount int32 = 3 // for masters
		for _, wp := range *cs.WorkerProfiles {
			expectedNodeCount += *wp.Count
		}

		nodes, err := Clients.kubernetes.CoreV1().Nodes().List(metav1.ListOptions{})
		Expect(err).NotTo(HaveOccurred())
		var nodeCount int64
		for _, node := range nodes.Items {
			if ready.NodeIsReady(&node) {
				nodeCount++
			} else {
				for _, c := range node.Status.Conditions {
					Log.Warnf("node %s status %s", node.Name, c.String())
				}
			}
		}
		Expect(nodeCount).To(Equal(expectedNodeCount))
	})
})
