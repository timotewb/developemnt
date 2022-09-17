package main

// Import key modules.
import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// Define key global variables.
var (
	subscriptionId = "af272ccb-a7af-4d0e-a513-d3f64afc02b1"
	url            = "https://cpu0blob0test.blob.core.windows.net/" //replace <StorageAccountName> with your Azure storage account name
)

func randomString() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r.Int())
}

// Define the function to create a resource group.

func main() {
	ctx := context.Background()
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Authentication failure: %+v", err)
	}

	// Azure SDK Azure Resource Management clients accept the credential as a parameter
	client, _ := armresources.NewClient(subscriptionId, cred, nil)

	log.Printf("Authenticated to subscription", client)

	serviceClient, err := azblob.NewServiceClient(url, cred, nil)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}
	// Create the container
	containerName := fmt.Sprintf("quickstart-%s", randomString())
	fmt.Printf("Creating a container named %s\n", containerName)

	containerClient, _ := serviceClient.NewContainerClient(containerName)
	_, err = containerClient.Create(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}
}
