package main

import (
	"flag"
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"log"
	"os"
	"time"
)

var APIClient *blizzard_api.WoWClient

func main() {
	appKey := flag.String("key", "", "Battle net APP key")
	appSecret := flag.String("secret", "", "Battle net APP secret")
	region := flag.String("region", "us", "API region, defaults to us")
	output := flag.String("output", "./dump", "Output directory")
	concurrency := flag.Int("concurrency", 100, "Concurrency")

	flag.Parse()

	if appKey == nil || *appKey == "" {
		log.Fatal("Missing argument: key")
	}

	if appSecret == nil || *appSecret == "" {
		log.Fatal("Missing argument: secret")
	}

	if _, err := os.Stat(*output); os.IsNotExist(err) {
		err := os.MkdirAll(*output, os.ModePerm)
		if err != nil {
			log.Fatal("Could not create the output directory")
		}
	}

	if *concurrency > 100 {
		*concurrency = 100
	}

	apiRegion := blizzard_api.Region(*region)

	APIClient = blizzard_api.NewWoWClient(apiRegion, nil, nil, false)
	APIClient.CreateAccessToken(*appKey, *appSecret, apiRegion)

	ahDump := &AHDump{
		OutputDir: *output,
	}
	start := time.Now()
	ahDump.Run(*concurrency)
	elapsed := time.Since(start)
	log.Printf("Completed in %s", elapsed)
}
