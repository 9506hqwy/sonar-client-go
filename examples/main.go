package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func basicAuth(ctx context.Context, req *http.Request) error {
	req.SetBasicAuth("admin", "admin")
	return nil
}

func main() {
	hc := http.Client{}

	c, err := client.NewClientWithResponses("http://127.0.0.1:9000", client.WithHTTPClient(&hc))
	if err != nil {
		log.Fatal(err)
	}

	params := client.ApiProjectsSearchParams{}
	resp, err := c.ApiProjectsSearchWithResponse(context.TODO(), &params, basicAuth)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("Expected HTTP 200 but received %d", resp.StatusCode())
	}

	for _, project := range *resp.JSON2XX.Components {
		fmt.Printf("Project Key: %s, Name: %s\n", *project.Key, *project.Name)
	}
}
