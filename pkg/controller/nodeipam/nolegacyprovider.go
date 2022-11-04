//go:build providerless
// +build providerless

/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nodeipam

import (
	"net"

	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog/v2"
)

type fakeController struct {
}

func (f *fakeController) Run(stopCh <-chan struct{}) {
	<-stopCh
}

func createLegacyIPAM(
	ic *Controller,
	nodeInformer coreinformers.NodeInformer,
	cloud cloudprovider.Interface,
	kubeClient clientset.Interface,
	clusterCIDRs []*net.IPNet,
	serviceCIDR *net.IPNet,
	nodeCIDRMaskSizes []int,
) *fakeController {
	klog.Fatal("Error trying to Init(): legacy cloud provider support disabled at build time")
	return &fakeController{}
}
