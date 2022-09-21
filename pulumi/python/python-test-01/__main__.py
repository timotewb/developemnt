"""An Azure RM Python Pulumi program"""

import pulumi
from pulumi_azure_native import resources
import pulumi_azure_native as azure_native

# Create an Azure Resource Group
rg = resources.ResourceGroup('pulumi-py-rg',
    location="australiaeast",
    resource_group_name="pulumi-py-rg"
)
rgdbr = resources.ResourceGroup('pulumi-py-rgdbr',
    location="australiaeast",
    resource_group_name="pulumi-py-rgdbr"
)

# Create role assignment on managed identity
ra = azure_native.authorization.RoleAssignment("pulumi-py-ra",
    principal_id="64765f4d-06b5-4f56-8cc4-1c068f624992",
    principal_type="ServicePrincipal",
    role_assignment_name="5a53e7cc-3e62-4357-a85d-6ac4af0d6c18",
    role_definition_id="/subscriptions/5f3d7f2f-1189-427d-aaa3-5c220e2b3e9a/providers/Microsoft.Authorization/roleDefinitions/8e3af657-a8ff-443c-a75c-2fe8c4bcb635",
    scope=rgdbr.id
)

# create managed identity
midbr = azure_native.managedidentity.UserAssignedIdentity("pulumi-py-midbr",
    location="australiaeast",
    resource_group_name=rgdbr.name,
    resource_name_="pulumi-py-midbr",
    tags={
        "applicaiton": "databricks",
        "databricks-environment": "true",
    }
)


# create network security group
network_security_group = azure_native.network.NetworkSecurityGroup("pulumi-sgdbr",
    location="australiaeast",
    network_security_group_name="pulumi-sgdbr",
    resource_group_name=rgdbr.name,
    security_rules=[azure_native.network.SecurityRuleArgs(
        access="Allow",
        direction="Inbound",
        protocol="*",
        description="Required for Databricks control plane management of worker nodes.",
        destination_address_prefix="*",
        destination_port_range="22",
        name="databricks-control-plane-ssh",
        priority=100,
        source_address_prefix="20.37.156.208/32,23.101.152.95/32",
        source_port_range="*",
    )]
)