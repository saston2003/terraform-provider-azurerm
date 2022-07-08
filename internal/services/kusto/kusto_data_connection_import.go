package kusto

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2021-08-27/dataconnections"

	"github.com/Azure/azure-sdk-for-go/services/kusto/mgmt/2022-02-01/kusto"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

func importDataConnection(kind dataconnections.DataConnectionKind) pluginsdk.ImporterFunc {
	return func(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) (data []*pluginsdk.ResourceData, err error) {
		id, err := dataconnections.ParseDataConnectionID(d.Id())
		if err != nil {
			return []*pluginsdk.ResourceData{}, err
		}

		client := meta.(*clients.Client).Kusto.DataConnectionsClient
		dataConnection, err := client.Get(ctx, *id)
		if err != nil {
			return []*pluginsdk.ResourceData{}, fmt.Errorf("retrieving %s: %+v", id, err)
		}

		if test, ok := *dataConnection.Model.(dataconnections.EventHubDataConnection); ok {

		}
		if _, ok := ; ok && kind != dataconnections.DataConnectionKindEventHub {
			return nil, fmt.Errorf(`kusto data connection "kind" mismatch, expected "%s", got "%s"`, kind, kusto.KindBasicDataConnectionKindEventHub)
		}

		if _, ok := dataConnection.Value.AsIotHubDataConnection(); ok && kind != kusto.KindBasicDataConnectionKindIotHub {
			return nil, fmt.Errorf(`kusto data connection "kind" mismatch, expected "%s", got "%s"`, kind, kusto.KindBasicDataConnectionKindIotHub)
		}

		if _, ok := dataConnection.Value.AsEventGridDataConnection(); ok && kind != kusto.KindBasicDataConnectionKindEventGrid {
			return nil, fmt.Errorf(`kusto data connection "kind" mismatch, expected "%s", got "%s"`, kind, kusto.KindBasicDataConnectionKindEventGrid)
		}

		return []*pluginsdk.ResourceData{d}, nil
	}
}
