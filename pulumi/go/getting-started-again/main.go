package main

import (
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		conf := config.New(ctx, "")
		location := conf.Require("location")

		// Create an Azure Resource Group
		resourceGroup, err := resources.NewResourceGroup(ctx, "pulumi-rg", &resources.ResourceGroupArgs{
			Location:          pulumi.String(location),
			ResourceGroupName: pulumi.String("pulumi-rg"),
		})
		if err != nil {
			return err
		}

		// Create an Azure resource (Storage Account)
		account, err := storage.NewStorageAccount(ctx, "pulumi0sa", &storage.StorageAccountArgs{
			ResourceGroupName: resourceGroup.Name,
			Sku: &storage.SkuArgs{
				Name: pulumi.String("Standard_LRS"),
			},
			Kind: pulumi.String("StorageV2"),
		})
		if err != nil {
			return err
		}

		// create databricks workspace
		// resourceGroupDBr, err := resources.NewResourceGroup(ctx, "pulumi-rgdbr", &resources.ResourceGroupArgs{
		// 	Location:          pulumi.String(location),
		// 	ResourceGroupName: pulumi.String("pulumi-rgdbr"),
		// })
		// if err != nil {
		// 	return err
		// }
		// databricks.NewWorkspace(ctx, "pulumi-dbr", &databricks.WorkspaceArgs{
		// 	Location:               pulumi.String(location),
		// 	ManagedResourceGroupId: resourceGroupDBr.ID().ToStringOutput(),
		// 	ResourceGroupName:      resourceGroup.Name,
		// 	WorkspaceName:          pulumi.String("pulumi-dbrws"),
		// })
		// if err != nil {
		// 	return err
		// }

		// Export the primary key of the Storage Account
		ctx.Export("primaryStorageKey", pulumi.All(resourceGroup.Name, account.Name).ApplyT(
			func(args []interface{}) (string, error) {
				resourceGroupName := args[0].(string)
				accountName := args[1].(string)
				accountKeys, err := storage.ListStorageAccountKeys(ctx, &storage.ListStorageAccountKeysArgs{
					ResourceGroupName: resourceGroupName,
					AccountName:       accountName,
				})
				if err != nil {
					return "", err
				}

				return accountKeys.Keys[0].Value, nil
			},
		))

		return nil
	})
}
