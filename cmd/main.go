package main

import (
	"log"
	"os"

	_ "func-excel-to-json"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	os.Setenv("FUNCTION_TARGET", "ConvertHTTP")
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}

// https://cloud.google.com/functions/docs/create-deploy-http-go
