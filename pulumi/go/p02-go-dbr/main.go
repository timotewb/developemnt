package main

import (
	"fmt"

	authorization "github.com/pulumi/pulumi-azure-native/sdk/go/azure/authorization"
	managedidentity "github.com/pulumi/pulumi-azure-native/sdk/go/azure/managedidentity"
	network "github.com/pulumi/pulumi-azure-native/sdk/go/azure/network"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		conf := config.New(ctx, "")
		location := conf.Require("location")

		// Create an Azure Resource Group
		rg, err := resources.NewResourceGroup(ctx, "pulumi-rg", &resources.ResourceGroupArgs{
			Location:          pulumi.String(location),
			ResourceGroupName: pulumi.String("pulumi-rg"),
		})
		if err != nil {
			return err
		}
		rgdbr, err := resources.NewResourceGroup(ctx, "pulumi-rgdbr", &resources.ResourceGroupArgs{
			Location:          pulumi.String(location),
			ResourceGroupName: pulumi.String("pulumi-rgdbr"),
		})
		if err != nil {
			return err
		}

		// create role assignment on managed identity
		radbr, err := authorization.NewRoleAssignment(ctx, "pulumi-radbr", &authorization.RoleAssignmentArgs{
			PrincipalId:        pulumi.String("64765f4d-06b5-4f56-8cc4-1c068f624992"),                                                                                                       //databricks resource provider objID
			PrincipalType:      pulumi.String("ServicePrincipal"),                                                                                                                           // type
			RoleAssignmentName: pulumi.String("5a53e7cc-3e62-4357-a85d-6ac4af0d6c18"),                                                                                                       //midbr clientID
			RoleDefinitionId:   pulumi.String("/subscriptions/5f3d7f2f-1189-427d-aaa3-5c220e2b3e9a/providers/Microsoft.Authorization/roleDefinitions/8e3af657-a8ff-443c-a75c-2fe8c4bcb635"), //id for owner role form web
			Scope:              rgdbr.ID().ToStringOutput(),
		})
		if err != nil {
			return err
		}

		// create managed identity in rgdbr
		midbr, err := managedidentity.NewUserAssignedIdentity(ctx, "pulumi-midbr", &managedidentity.UserAssignedIdentityArgs{
			Location:          pulumi.String(location),
			ResourceGroupName: rgdbr.Name,
			ResourceName:      pulumi.String("pulumi-midbr"),
			Tags: pulumi.StringMap{
				"applicaiton":            pulumi.String("databricks"),
				"databricks-environment": pulumi.String("true"),
			},
		})
		if err != nil {
			return err
		}

		// create network security group
		// need to create security rules for valid ip range
		sgdbr, err := network.NewNetworkSecurityGroup(ctx, "pulumi-sgdbr", &network.NetworkSecurityGroupArgs{
			Location:                 pulumi.String(location),
			NetworkSecurityGroupName: pulumi.String("pulumi-sgdbr"),
			ResourceGroupName:        rgdbr.Name,
			Tags: pulumi.StringMap{
				"applicaiton":            pulumi.String("databricks"),
				"databricks-environment": pulumi.String("true"),
			},
		})
		if err != nil {
			return err
		}

		// create virtual network
		// need to manaully create private and public subnets
		vndbr, err := network.NewVirtualNetwork(ctx, "pulumi-vndbr", &network.VirtualNetworkArgs{
			AddressSpace: &network.AddressSpaceArgs{
				AddressPrefixes: pulumi.StringArray{
					pulumi.String("10.139.0.0/16"),
				},
			},
			Location:           pulumi.String(location),
			ResourceGroupName:  rgdbr.Name,
			VirtualNetworkName: pulumi.String("pulumi-vndbr"),
			Tags: pulumi.StringMap{
				"applicaiton":            pulumi.String("databricks"),
				"databricks-environment": pulumi.String("true"),
			},
		})
		if err != nil {
			return err
		}

		// create databricks workspace
		// not you cannot run the below without the subnets created!
		// wsdbr, err := databricks.NewWorkspace(ctx, "pulumi-wsdbr", &databricks.WorkspaceArgs{
		// 	Location:               pulumi.String(location),
		// 	ManagedResourceGroupId: rgdbr.ID().ToStringOutput(),
		// 	Parameters: &databricks.WorkspaceCustomParametersArgs{
<<<<<<< Updated upstream
		// 		CustomVirtualNetworkId: &databricks.WorkspaceCustomStringParameterArgs{
		// 			Value: vndbr.ID().ToStringOutput(),
		// 		},
		// 		CustomPrivateSubnetName: &databricks.WorkspaceCustomStringParameterArgs{
		// 			Value: pulumi.String("private-subnet-t"),
		// 		},
		// 		CustomPublicSubnetName: &databricks.WorkspaceCustomStringParameterArgs{
		// 			Value: pulumi.String("public-subnet-t"),
=======
		// 		CustomPrivateSubnetName: &databricks.WorkspaceCustomStringParameterArgs{
		// 			Value: pulumi.String("private-subnet"),
		// 		},
		// 		CustomPublicSubnetName: &databricks.WorkspaceCustomStringParameterArgs{
		// 			Value: pulumi.String("public-subnet"),
		// 		},
		// 		CustomVirtualNetworkId: &databricks.WorkspaceCustomStringParameterArgs{
		// 			Value: vndbr.ID().ToStringOutput(),
>>>>>>> Stashed changes
		// 		},
		// 	},
		// 	ResourceGroupName: rg.Name,
		// 	WorkspaceName:     pulumi.String("pulumi-wsdbr"),
		// })
		// if err != nil {
		// 	return err
		// }

		fmt.Println(rg.Name)
		fmt.Println(rgdbr.Name)
		fmt.Println(midbr.Name)
		fmt.Println(radbr.Name)
		fmt.Println(sgdbr.Name)
		fmt.Println(vndbr.Name)
		// fmt.Println(wsdbr.Name)

		return nil
	})
}
