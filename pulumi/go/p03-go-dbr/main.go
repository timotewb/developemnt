package main

import (
	"fmt"

	authorization "github.com/pulumi/pulumi-azure-native/sdk/go/azure/authorization"
	databricks "github.com/pulumi/pulumi-azure-native/sdk/go/azure/databricks"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/network"
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

		// create role assignment on managed identity
		ra, err := authorization.NewRoleAssignment(ctx, "pulumi-ra", &authorization.RoleAssignmentArgs{
			PrincipalId:        pulumi.String("64765f4d-06b5-4f56-8cc4-1c068f624992"),                                                                                                       //databricks resource provider objID
			PrincipalType:      pulumi.String("ServicePrincipal"),                                                                                                                           // type
			RoleAssignmentName: pulumi.String("5a53e7cc-3e62-4357-a85d-6ac4af0d6c18"),                                                                                                       //midbr clientID
			RoleDefinitionId:   pulumi.String("/subscriptions/5f3d7f2f-1189-427d-aaa3-5c220e2b3e9a/providers/Microsoft.Authorization/roleDefinitions/8e3af657-a8ff-443c-a75c-2fe8c4bcb635"), //id for owner role form web
			Scope:              rg.ID().ToStringOutput(),
		})
		if err != nil {
			return err
		}

		// create security group
		// sg, err := network.NewNetworkSecurityGroup(ctx, "pulumi-sg", &network.NetworkSecurityGroupArgs{
		// 	Location:                 pulumi.String(location),
		// 	NetworkSecurityGroupName: pulumi.String("pulumi-sg"),
		// 	ResourceGroupName:        rg.Name,
		// 	Tags: pulumi.StringMap{
		// 		"applicaiton":            pulumi.String("databricks"),
		// 		"databricks-environment": pulumi.String("true"),
		// 	},
		// })
		// if err != nil {
		// 	return err
		// }
		// manually create security groups 'sgdbr_template.json'

		sg, err := network.NewNetworkSecurityGroup(ctx, "pulumi-sg", &network.NetworkSecurityGroupArgs{
			Location:          pulumi.String(location),
			ResourceGroupName: rg.Name,
			Name:              pulumi.String("pulumi-sg"),
			SecurityRules: network.NetworkSecurityGroupSecurityRuleArray{
				&network.NetworkSecurityGroupSecurityRuleArgs{
					Name:                     pulumi.String("Microsoft.Databricks-workspaces_UseOnly_databricks-worker-to-worker-inbound"),
					Description:              pulumi.String("Required for worker nodes communication within a cluster."),
					Protocol:                 pulumi.String("*"),
					SourcePortRange:          pulumi.String("*"),
					DestinationPortRange:     pulumi.String("*"),
					SourceAddressPrefix:      pulumi.String("VirtualNetwork"),
					DestinationAddressPrefix: pulumi.String("VirtualNetwork"),
					Access:                   pulumi.String("Allow"),
					Priority:                 pulumi.Int(100),
					Direction:                pulumi.String("Inbound"),
				},
				&network.NetworkSecurityGroupSecurityRuleArgs{
					Name:                     pulumi.String("Microsoft.Databricks-workspaces_UseOnly_databricks-worker-to-databricks-webapp"),
					Description:              pulumi.String("Required for workers communication with Databricks Webapp."),
					Protocol:                 pulumi.String("TCP"),
					SourcePortRange:          pulumi.String("*"),
					DestinationPortRange:     pulumi.String("443"),
					SourceAddressPrefix:      pulumi.String("VirtualNetwork"),
					DestinationAddressPrefix: pulumi.String("AzureDatabricks"),
					Access:                   pulumi.String("Allow"),
					Priority:                 pulumi.Int(100),
					Direction:                pulumi.String("Outbound"),
				},
				&network.NetworkSecurityGroupSecurityRuleArgs{
					Name:                     pulumi.String("Microsoft.Databricks-workspaces_UseOnly_databricks-worker-to-sql"),
					Description:              pulumi.String("Required for workers communication with Azure SQL services."),
					Protocol:                 pulumi.String("TCP"),
					SourcePortRange:          pulumi.String("*"),
					DestinationPortRange:     pulumi.String("3306"),
					SourceAddressPrefix:      pulumi.String("VirtualNetwork"),
					DestinationAddressPrefix: pulumi.String("Sql"),
					Access:                   pulumi.String("Allow"),
					Priority:                 pulumi.Int(101),
					Direction:                pulumi.String("Outbound"),
				},
				&network.NetworkSecurityGroupSecurityRuleArgs{
					Name:                     pulumi.String("Microsoft.Databricks-workspaces_UseOnly_databricks-worker-to-storage"),
					Description:              pulumi.String("Required for workers communication with Azure Storage services."),
					Protocol:                 pulumi.String("TCP"),
					SourcePortRange:          pulumi.String("*"),
					DestinationPortRange:     pulumi.String("443"),
					SourceAddressPrefix:      pulumi.String("VirtualNetwork"),
					DestinationAddressPrefix: pulumi.String("Storage"),
					Access:                   pulumi.String("Allow"),
					Priority:                 pulumi.Int(102),
					Direction:                pulumi.String("Outbound"),
				},
				&network.NetworkSecurityGroupSecurityRuleArgs{
					Name:                     pulumi.String("Microsoft.Databricks-workspaces_UseOnly_databricks-worker-to-worker-outbound"),
					Description:              pulumi.String("Required for worker nodes communication within a cluster."),
					Protocol:                 pulumi.String("TCP"),
					SourcePortRange:          pulumi.String("*"),
					DestinationPortRange:     pulumi.String("*"),
					SourceAddressPrefix:      pulumi.String("VirtualNetwork"),
					DestinationAddressPrefix: pulumi.String("VirtualNetwork"),
					Access:                   pulumi.String("Allow"),
					Priority:                 pulumi.Int(103),
					Direction:                pulumi.String("Outbound"),
				},
				&network.NetworkSecurityGroupSecurityRuleArgs{
					Name:                     pulumi.String("Microsoft.Databricks-workspaces_UseOnly_databricks-worker-to-eventhub"),
					Description:              pulumi.String("Required for worker communication with Azure Eventhub services."),
					Protocol:                 pulumi.String("TCP"),
					SourcePortRange:          pulumi.String("*"),
					DestinationPortRange:     pulumi.String("8080"),
					SourceAddressPrefix:      pulumi.String("VirtualNetwork"),
					DestinationAddressPrefix: pulumi.String("EventHub"),
					Access:                   pulumi.String("Allow"),
					Priority:                 pulumi.Int(104),
					Direction:                pulumi.String("Outbound"),
				},
			},
			Tags: pulumi.StringMap{
				"applicaiton":            pulumi.String("databricks"),
				"databricks-environment": pulumi.String("true"),
			},
		})
		if err != nil {
			return err
		}

		// Create virtual network
		vn, err := network.NewVirtualNetwork(ctx, "pulumi-vn", &network.VirtualNetworkArgs{
			Location:          pulumi.String(location),
			ResourceGroupName: rg.Name,
			Name:              pulumi.String("pulumi-vn"),
			AddressSpaces: pulumi.StringArray{
				pulumi.String("10.139.0.0/16"),
			},
			Subnets: network.VirtualNetworkSubnetArray{
				&network.VirtualNetworkSubnetArgs{
					Name:          pulumi.String("public-subnet"),
					AddressPrefix: pulumi.String("10.139.0.0/24"),
					SecurityGroup: sg.ID(),
				},
				&network.VirtualNetworkSubnetArgs{
					Name:          pulumi.String("private-subnet"),
					AddressPrefix: pulumi.String("10.139.1.0/24"),
					SecurityGroup: sg.ID(),
				},
			},
			Tags: pulumi.StringMap{
				"applicaiton":            pulumi.String("databricks"),
				"databricks-environment": pulumi.String("true"),
			},
		})
		if err != nil {
			return err
		}

		// add delegations to databricks

		// create databricks workspace
		dbrws, err := databricks.NewWorkspace(ctx, "pulumi-dbrws", &databricks.WorkspaceArgs{
			Location: pulumi.String(location),
			// ManagedResourceGroupId: rgdbr.ID().ToStringOutput(),
			ManagedResourceGroupId: pulumi.String("/subscriptions/5f3d7f2f-1189-427d-aaa3-5c220e2b3e9a/resourceGroups/pulumi-rgdbr-auto"),
			Parameters: &databricks.WorkspaceCustomParametersArgs{
				CustomVirtualNetworkId: &databricks.WorkspaceCustomStringParameterArgs{
					Value: vn.ID(),
				},
				CustomPrivateSubnetName: &databricks.WorkspaceCustomStringParameterArgs{
					Value: pulumi.String("private-subnet"),
				},
				CustomPublicSubnetName: &databricks.WorkspaceCustomStringParameterArgs{
					Value: pulumi.String("public-subnet"),
				},
				RequireInfrastructureEncryption: &databricks.WorkspaceCustomBooleanParameter{
					Value: true,
				},
			},
			ResourceGroupName: rg.Name,
			WorkspaceName:     pulumi.String("pulumi-dbrws"),
		})
		if err != nil {
			return err
		}

		fmt.Println(ra.Name.ToStringOutput())
		fmt.Println(sg.ID())
		fmt.Println(vn.ID())
		fmt.Println(dbrws.Name.ToStringOutput())

		return nil
	})
}
