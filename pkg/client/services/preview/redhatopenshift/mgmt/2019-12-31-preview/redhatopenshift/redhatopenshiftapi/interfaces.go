package redhatopenshiftapi

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"

	"github.com/jim-minter/rp/pkg/client/services/preview/redhatopenshift/mgmt/2019-12-31-preview/redhatopenshift"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result redhatopenshift.OperationList, err error)
}

var _ OperationsClientAPI = (*redhatopenshift.OperationsClient)(nil)

// OpenShiftClustersClientAPI contains the set of methods on the OpenShiftClustersClient type.
type OpenShiftClustersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, resourceName string, parameters redhatopenshift.OpenShiftCluster) (result redhatopenshift.OpenShiftClustersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.OpenShiftClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.OpenShiftCluster, err error)
	GetCredentials(ctx context.Context, resourceGroupName string, resourceName string) (result redhatopenshift.OpenShiftClusterCredentials, err error)
	List(ctx context.Context) (result redhatopenshift.OpenShiftClusterList, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result redhatopenshift.OpenShiftClusterList, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, parameters redhatopenshift.OpenShiftCluster) (result redhatopenshift.OpenShiftClustersUpdateFuture, err error)
}

var _ OpenShiftClustersClientAPI = (*redhatopenshift.OpenShiftClustersClient)(nil)
