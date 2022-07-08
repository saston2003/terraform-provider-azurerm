package client

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2021-08-27/attacheddatabaseconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2021-08-27/clusterprincipalassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2021-08-27/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2021-08-27/databaseprincipalassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2021-08-27/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2021-08-27/dataconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2021-08-27/scripts"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	AttachedDatabaseConfigurationsClient *attacheddatabaseconfigurations.AttachedDatabaseConfigurationsClient
	ClustersClient                       *clusters.ClustersClient
	ClusterPrincipalAssignmentsClient    *clusterprincipalassignments.ClusterPrincipalAssignmentsClient
	DatabasesClient                      *databases.DatabasesClient
	DataConnectionsClient                *dataconnections.DataConnectionsClient
	DatabasePrincipalAssignmentsClient   *databaseprincipalassignments.DatabasePrincipalAssignmentsClient
	ScriptsClient                        *scripts.ScriptsClient
}

func NewClient(o *common.ClientOptions) *Client {
	ClustersClient := clusters.NewClustersClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&ClustersClient.Client, o.ResourceManagerAuthorizer)

	ClusterPrincipalAssignmentsClient := clusterprincipalassignments.NewClusterPrincipalAssignmentsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&ClusterPrincipalAssignmentsClient.Client, o.ResourceManagerAuthorizer)

	DatabasesClient := databases.NewDatabasesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&DatabasesClient.Client, o.ResourceManagerAuthorizer)

	DatabasePrincipalAssignmentsClient := databaseprincipalassignments.NewDatabasePrincipalAssignmentsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&DatabasePrincipalAssignmentsClient.Client, o.ResourceManagerAuthorizer)

	DataConnectionsClient := dataconnections.NewDataConnectionsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&DataConnectionsClient.Client, o.ResourceManagerAuthorizer)

	AttachedDatabaseConfigurationsClient := attacheddatabaseconfigurations.NewAttachedDatabaseConfigurationsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&AttachedDatabaseConfigurationsClient.Client, o.ResourceManagerAuthorizer)

	ScriptsClient := scripts.NewScriptsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&ScriptsClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		AttachedDatabaseConfigurationsClient: &AttachedDatabaseConfigurationsClient,
		ClustersClient:                       &ClustersClient,
		ClusterPrincipalAssignmentsClient:    &ClusterPrincipalAssignmentsClient,
		DatabasesClient:                      &DatabasesClient,
		DataConnectionsClient:                &DataConnectionsClient,
		DatabasePrincipalAssignmentsClient:   &DatabasePrincipalAssignmentsClient,
		ScriptsClient:                        &ScriptsClient,
	}
}
