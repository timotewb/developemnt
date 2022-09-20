package main

import (
	"fmt"

	authorization "github.com/pulumi/pulumi-azure-native/sdk/go/azure/authorization"
	managedidentity "github.com/pulumi/pulumi-azure-native/sdk/go/azure/managedidentity"
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

		// create role assignment on managed identity
		ra, err := authorization.NewRoleAssignment(ctx, "ra", &authorization.RoleAssignmentArgs{
			PrincipalId:        pulumi.String("64765f4d-06b5-4f56-8cc4-1c068f624992"), //databricks resource provider objID
			PrincipalType:      pulumi.String("ServicePrincipal"),                     // type
			RoleAssignmentName: pulumi.String("5a53e7cc-3e62-4357-a85d-6ac4af0d6c18"), //midbr clientID
			RoleDefinitionId:   pulumi.String("/subscriptions/5f3d7f2f-1189-427d-aaa3-5c220e2b3e9a/providers/Microsoft.Authorization/roleDefinitions/8e3af657-a8ff-443c-a75c-2fe8c4bcb635"),
			Scope:              rgdbr.ID().ToStringOutput(),
		})
		if err != nil {
			return err
		}

		fmt.Println(rg.Name)
		fmt.Println(rgdbr.Name)
		fmt.Println(midbr.Name)
		fmt.Println(ra.Name)
		return nil
	})
}
