package main

import (
	"embed"
	algolia "github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/joho/godotenv"
	"github.com/tylermmorton/torque"
	"github.com/tylermmorton/torque/.www/docsite/routes/docs"
	"github.com/tylermmorton/torque/.www/docsite/routes/landing"
	"github.com/tylermmorton/torque/.www/docsite/routes/search"
	"github.com/tylermmorton/torque/.www/docsite/services/content"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:generate tmpl bind ./... --outfile=tmpl.gen.go

//go:embed .build/static/*
var staticAssets embed.FS

//go:embed content/docs/*
var embeddedContent embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("failed to load env: %+v", err)
	}

	algoliaAppId, ok := os.LookupEnv("ALGOLIA_APP_ID")
	if !ok {
		log.Fatalf("ALGOLIA_APP_ID not set in environment")
	}

	algoliaApiKey, ok := os.LookupEnv("ALGOLIA_API_KEY")
	if !ok {
		log.Fatalf("ALGOLIA_API_KEY not set in environment")
	}

	algoliaClient := algolia.NewClient(algoliaAppId, algoliaApiKey)

	contentSvc, err := content.New(embeddedContent, algoliaClient)
	if err != nil {
		log.Fatalf("failed to create content service: %+v", err)
	}

	staticAssets, err := fs.Sub(staticAssets, ".build/static")
	if err != nil {
		log.Fatalf("failed to create static assets filesystem: %+v", err)
	}

	r := torque.NewRouter(

		torque.WithFileSystemServer("/s", staticAssets),

		torque.WithRouteModule("/", &landing.RouteModule{}),

		torque.WithRouteModule("/docs/{pageName}", &docs.RouteModule{
			ContentSvc: contentSvc,
		}),

		torque.WithRouteModule("/search", &search.RouteModule{
			ContentSvc: contentSvc,
		}),
	)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("failed to start server: %+v", err)
	}
}