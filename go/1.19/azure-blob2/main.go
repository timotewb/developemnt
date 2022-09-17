package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func randomString() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r.Int())
}

func main() {
	// define context
	ctx := context.Background()
	url := "https://test0blob0sa.blob.core.windows.net/"

	// authenticate
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Authentication failure: %+v", err)
	}

	// define client to access blob
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
