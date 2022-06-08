package main

import (
	"context"
	"fmt"
	"github.com/kjuulh/stiletto/pkg/client"
	"github.com/kjuulh/stiletto/pkg/featurestores"
)

func main() {

	sClient := client.NewStilettoClient()

	sClient.SetFeatureStore(
		"userIDs",
		featurestores.NewInMemoryFeatureStore([]string{"some-user-id"}))

	featureActive, err := sClient.GetFeature(
		context.Background(),
		"userIDs",
		"some-user-id")
	if err != nil {
		panic(err)
	}

	if featureActive {
		fmt.Println("Feature is active, yay!")
	}
}
